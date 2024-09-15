/*
 * Copyright (c) 2024 sixwaaaay.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"errors"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sixwaaaay/must"
	"github.com/spf13/viper"
)

type Request struct {
	Name string `json:"name"`
}

func (req *Request) Validate() error {
	if req.Name == "" || len(req.Name) > 255 {
		return echo.NewHTTPError(http.StatusBadRequest, "name can not be empty or exceed 255 characters")
	}
	return nil
}

type Reply struct {
	Url string `json:"url"`
}

var configFile = flag.String("f", "configs/config.yaml", "the config file")

func main() {
	viper.SetConfigFile(*configFile)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	var config Conf
	must.RunE(viper.ReadInConfig())
	must.RunE(viper.Unmarshal(&config))

	handler := must.Must(NewHandler(config.Minio))

	e := newServer()

	e.POST("/assert/new", handler.NewAssert)

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

// Handler is a struct that encapsulates a MinIO client and a bucket name.
// It is used to interact with MinIO storage, performing operations on the specified bucket.
//
// Fields:
// - client: A pointer to a MinIO client instance used for communicating with the MinIO server.
// - bucket: A string representing the name of the bucket to perform operations on.
type Handler struct {
	client *minio.Client
	bucket string
}

func NewHandler(config MinioConfig) (*Handler, error) {
	client, err := NewMinioClient(config)
	if err != nil {
		return nil, err
	}
	return &Handler{
		client: client,
	}, nil
}

// GeneratePresignedURL generates a presigned URL for uploading an object to a specified bucket.
// The URL is valid for 24 hours.
//
// Parameters:
//
//	ctx - The context for the request.
//	req - The request containing the name of the object to be uploaded.
//
// Returns:
//
//	A string containing the presigned URL, or an error if the URL could not be generated.
func (h *Handler) GeneratePresignedURL(ctx context.Context, req Request) (string, error) {
	random, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	hex := strings.ReplaceAll(random.String(), "-", "")
	name := hex + "_" + req.Name

	presignedURL, err := h.client.PresignedPutObject(ctx, h.bucket, name, time.Hour*24)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}

// NewAssert handles the creation of a new assertion. It binds the request data,
// validates it, generates a presigned URL, and returns the URL in the response.
//
// Parameters:
//
//	ctx - the Echo context containing the request and response objects.
//
// Returns:
//
//	An error if any step in the process fails, otherwise it returns a JSON response
//	with the presigned URL.
func (h *Handler) NewAssert(ctx echo.Context) error {
	var req Request
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	if err := req.Validate(); err != nil {
		return err
	}

	presignedURL, err := h.GeneratePresignedURL(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, Reply{presignedURL})
}

type Conf struct {
	ListenOn string
	Minio    MinioConfig
}

func NewMinioClient(config MinioConfig) (*minio.Client, error) {
	return minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: config.UseSSL,
	})
}

type MinioConfig struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	UseSSL    bool
	Bucket    string
}

func newServer() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Logger.SetLevel(log.INFO)
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

/*

$env:GIT_COMMITTER_DATE="Fri Apr 12 21:38:14 2024 +0800"
git commit --amend --date "Fri Apr 12 21:38:14 2024 +0800"
*/
