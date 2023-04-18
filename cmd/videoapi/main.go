/*
 * Copyright (c) 2023 sixwaaaay.
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
	"flag"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/minio/minio-go/v7"
	"github.com/sixwaaaay/sharing/pkg/blobstore"
	"github.com/sixwaaaay/sharing/pkg/configs"
	"github.com/sixwaaaay/sharing/pkg/encoder"
	"github.com/sixwaaaay/sharing/pkg/pb"
	"github.com/sixwaaaay/sharing/pkg/rpc"
	"github.com/sixwaaaay/sharing/pkg/sign"
	"mime/multipart"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

type Config struct {
	ListenOn     string
	VideoService rpc.GrpcConfig
	Jwt          sign.JWT
	MinIO        blobstore.MinioConfig
	Secret       string
	ImageBucket  string
	VideoBucket  string
}

var configFile = flag.String("f", "config.yaml", "the config file")

func main() {
	config, err := configs.NewConfig[Config](*configFile)
	log.Printf("%+v", config)
	handleErr(err)
	e := newServer()
	client, err := rpc.NewVideoClient(config.VideoService)
	if err != nil {
		panic(err)
	}
	handleErr(err)
	mc, err := blobstore.NewMinioClient(config.MinIO)
	handleErr(err)
	handler := NewHandler(client, config.Secret, config.ImageBucket, config.VideoBucket, mc)
	handler.Update(e)

	// Start server
	go func() {
		if err := e.Start(config.ListenOn); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
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
	return e
}

type Handler struct {
	cli         pb.VideoServiceClient
	secret      []byte
	ImageBucket string
	VideoBucket string
	mc          *minio.Client
}

func NewHandler(cli pb.VideoServiceClient, secret string, bucket string, videoBucket string, mc *minio.Client) *Handler {
	return &Handler{
		cli:         cli,
		secret:      []byte(secret),
		ImageBucket: bucket,
		VideoBucket: videoBucket,
		mc:          mc,
	}
}

func (h *Handler) Update(e *echo.Echo) {
	e.POST("/videos/liked", h.VideoLiked, echo.WrapMiddleware(sign.Middleware(h.secret, false)))
	e.POST("/videos/recent", h.VideoRecent, echo.WrapMiddleware(sign.Middleware(h.secret, false)))
	e.POST("/videos/user", h.VideoUser, echo.WrapMiddleware(sign.Middleware(h.secret, false)))
	e.POST("/videos", h.VideoCreate, echo.WrapMiddleware(sign.Middleware(h.secret, true)))
	e.GET("/videos/:id", h.VideoGet, echo.WrapMiddleware(sign.Middleware(h.secret, false)))
	e.PATCH("/videos/liked", h.VideoLike, echo.WrapMiddleware(sign.Middleware(h.secret, true)))
}

func (h *Handler) VideoLiked(c echo.Context) error {
	var req = new(pb.GetLikedVideosRequest)
	if err := encoder.Unmarshal(c.Request().Body, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	req.SubjectId, _ = h.subjectId(c)
	resp, err := h.cli.GetLikedVideos(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return encoder.Marshal(c.Response().Writer, resp)
}

func (h *Handler) VideoUser(c echo.Context) error {
	var req = new(pb.GetUserVideosRequest)
	if err := encoder.Unmarshal(c.Request().Body, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	req.SubjectId, _ = h.subjectId(c)
	resp, err := h.cli.GetUserVideos(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return encoder.Marshal(c.Response().Writer, resp)
}

func (h *Handler) subjectId(ctx echo.Context) (int64, error) {
	subjectID, _ := ctx.Request().Context().Value("x-id").(string)
	id, err := strconv.ParseInt(subjectID, 10, 64)
	return id, err
}

func (h *Handler) VideoRecent(c echo.Context) error {
	var req = new(pb.GetRecentVideoReq)
	if err := encoder.Unmarshal(c.Request().Body, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	req.SubjectId, _ = h.subjectId(c)
	resp, err := h.cli.GetRecentVideo(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return encoder.Marshal(c.Response().Writer, resp)
}

func (h *Handler) VideoCreate(ctx echo.Context) error {
	var req = new(pb.CreateVideoRequest)
	var err error
	req.SubjectId, err = h.subjectId(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if video, err := ctx.FormFile("video"); err == nil {
		req.VideoUrl, err = h.uploadFile(ctx.Request().Context(), video, h.VideoBucket)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	if cover, err := ctx.FormFile("cover"); err == nil {
		req.CoverUrl, err = h.uploadFile(ctx.Request().Context(), cover, h.ImageBucket)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	err = echo.FormFieldBinder(ctx).
		String("title", &req.Title).
		String("description", &req.Description).
		String("category", &req.Category).
		//Strings("tags", &req.Tags). // need to refactor
		BindError()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, err := h.cli.CreateVideo(ctx.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return encoder.Marshal(ctx.Response().Writer, resp)
}

func (u *Handler) uploadFile(ctx context.Context, avatar *multipart.FileHeader, bucket string) (string, error) {
	//	gen a uuid
	id := uuid.New().String()
	// open source file
	src, err := avatar.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	_, err = u.mc.PutObject(ctx, bucket, id, src, -1, minio.PutObjectOptions{
		ContentType: avatar.Header.Get("Content-Type"),
	})
	if err != nil {
		return "", err
	}
	return id, nil
}

func (h *Handler) VideoGet(c echo.Context) error {
	var req = new(pb.GetVideoRequest)
	var err error
	req.Id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	req.SubjectId, _ = h.subjectId(c)
	resp, err := h.cli.GetVideo(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return encoder.Marshal(c.Response().Writer, resp)
}

func (h *Handler) VideoLike(c echo.Context) error {
	var req = new(pb.LikeActionRequest)
	var err error
	if err := encoder.Unmarshal(c.Request().Body, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	req.SubjectId, err = h.subjectId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	resp, err := h.cli.LikeAction(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return encoder.Marshal(c.Response().Writer, resp)
}
