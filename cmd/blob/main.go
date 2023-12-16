package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sixwaaaay/must"
	"github.com/sixwaaaay/sharing/pkg/configs"
)

type Request struct {
	Kind string `json:"kind"`
}

func (req *Request) Validate() error {
	if req.Kind != "img" && req.Kind != "avatar" && req.Kind != "video" {
		return errors.New("unsupported kind of asserts")
	}
	return nil
}

type Reply struct {
	Url string `json:"url"`
}

var configFile = flag.String("f", "configs/config.yaml", "the config file")

func main() {
	config := must.Must(configs.NewConfig[Conf](*configFile))

	client := must.Must(NewMinioClient(config.Minio))
	e := newServer()

	e.POST("/assets/new", func(c echo.Context) error {
		var req Request
		if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
			return err
		}

		if err := req.Validate(); err != nil {
			return err
		}

		presignedURL, err := GeneratePresignedURL(c.Request().Context(), client, req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, Reply{presignedURL})
	})

	svr := &http.Server{Addr: config.ListenOn, Handler: e}

	go func() {
		if err := svr.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	must.RunE(svr.Shutdown(ctx))
}

func GeneratePresignedURL(ctx context.Context, client *minio.Client, req Request) (string, error) {
	random, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	s := random.String()
	presignedURL, err := client.PresignedPutObject(ctx, req.Kind, s, time.Hour*24)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}

type Conf struct {
	ListenOn string
	Minio    MinioConfig
}

func NewMinioClient(config MinioConfig) (*minio.Client, error) {
	client, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: config.UseSSL,
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

type MinioConfig struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	UseSSL    bool
}

func newServer() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Logger.SetLevel(log.INFO)
	middleware.RequestID()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		c.Logger().Error(err)
		e.DefaultHTTPErrorHandler(errors.New("server error"), c)
	}
	return e
}
