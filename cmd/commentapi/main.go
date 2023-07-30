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
	"errors"
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sixwaaaay/sharing/pkg/configs"
	"github.com/sixwaaaay/sharing/pkg/encoder"
	"github.com/sixwaaaay/sharing/pkg/pb"
	"github.com/sixwaaaay/sharing/pkg/rpc"
	"github.com/sixwaaaay/sharing/pkg/sign"
	_ "go.uber.org/automaxprocs"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

type Config struct {
	ListenOn       string
	CommentService rpc.GrpcConfig
	Jwt            sign.JWT
	Secret         string
}

var configFile = flag.String("f", "configs/config.yaml", "the config file")

func main() {
	config, err := configs.NewConfig[Config](*configFile)
	handleErr(err)
	e := newServer()
	client, err := rpc.NewCommentClient(config.CommentService)
	handleErr(err)
	handler := NewHandler(client, config.Secret)
	handler.Update(e)

	// Start server
	go func() {
		if err := e.Start(config.ListenOn); err != nil && !errors.Is(err, http.ErrServerClosed) {
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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	return e
}

type Handler struct {
	cli    pb.CommentServiceClient
	secret []byte
}

func NewHandler(cli pb.CommentServiceClient, secret string) *Handler {
	return &Handler{
		cli:    cli,
		secret: []byte(secret),
	}
}

func (h *Handler) Update(e *echo.Echo) {
	e.POST("/comments", h.CommentList, echo.WrapMiddleware(sign.Middleware(h.secret, false)))
	e.PATCH("/comments", h.Comment, echo.WrapMiddleware(sign.Middleware(h.secret, true)))
}

func (h *Handler) CommentList(ctx echo.Context) error {
	var req = new(pb.CommentListReq)
	if err := encoder.Unmarshal(ctx.Request().Body, req); err != nil {
		return err
	}
	id, err := h.subjectId(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid subject id")
	}
	req.SubjectId = id
	resp, err := h.cli.CommentList(ctx.Request().Context(), req)

	return encoder.Marshal(ctx.Response().Writer, resp)
}

func (h *Handler) Comment(ctx echo.Context) error {
	var req = new(pb.CommentActionReq)
	if err := encoder.Unmarshal(ctx.Request().Body, req); err != nil {
		return err
	}
	id, err := h.subjectId(ctx)
	if err != nil || id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid subject id")
	}
	req.SubjectId = id
	resp, err := h.cli.CommentAction(ctx.Request().Context(), req)

	if err != nil {
		return err
	}

	return encoder.Marshal(ctx.Response().Writer, resp)
}

func (h *Handler) subjectId(ctx echo.Context) (int64, error) {
	subjectID, _ := ctx.Request().Context().Value("x-id").(string)
	id, err := strconv.ParseInt(subjectID, 10, 64)
	return id, err
}

type CommentReq struct {
	VideoID   int64  `json:"video_id"`
	Action    int32  `json:"action"`
	Comment   string `json:"comment"`
	CommentID int64  `json:"comment_id"`
}

func (c *CommentReq) Validate() error {
	if c.Action != 1 && c.Action != 2 {
		return errors.New("action is required")
	}
	if c.VideoID <= 0 {
		return errors.New("video_id is required")
	}
	if c.Action == 1 && c.Comment == "" {
		return errors.New("comment is required")
	}
	if c.Action == 2 && c.CommentID <= 0 {
		return errors.New("comment_id is required")
	}
	return nil
}

type CommentListReq struct {
	VideoID int64 `json:"video_id"`
	Token   int64 `json:"token"`
	Limit   int32 `json:"limit"`
}

func (c *CommentListReq) Validate() error {
	if c.VideoID <= 0 {
		return errors.New("video_id is required")
	}
	return nil
}

type CommentListReply struct {
	Comments []*pb.Comment `json:"comments"`
}
