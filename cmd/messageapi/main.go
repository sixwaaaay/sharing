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
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sixwaaaay/sharing/pkg/configs"
	"github.com/sixwaaaay/sharing/pkg/encoder"
	"github.com/sixwaaaay/sharing/pkg/pb"
	"github.com/sixwaaaay/sharing/pkg/rpc"
	"github.com/sixwaaaay/sharing/pkg/sign"
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
	client, err := rpc.NewMessageClient(config.CommentService)
	handleErr(err)
	handler := NewHandler(client, config.Secret)
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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	return e
}

type Handler struct {
	cli    pb.MessageServiceClient
	secret []byte
}

func NewHandler(cli pb.MessageServiceClient, secret string) *Handler {
	return &Handler{
		cli:    cli,
		secret: []byte(secret),
	}
}

func (h *Handler) Update(e *echo.Echo) {
	e.POST("/messages", h.MessageList, echo.WrapMiddleware(sign.Middleware(h.secret, true)))
	e.PATCH("/messages", h.Message, echo.WrapMiddleware(sign.Middleware(h.secret, true)))
}

func (h *Handler) MessageList(ctx echo.Context) error {
	var req = new(pb.MessageListRequest)
	if err := encoder.Unmarshal(ctx.Request().Body, req); err != nil {
		return err
	}
	id, err := h.subjectId(ctx)
	if err != nil || id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid subject id", err.Error())
	}
	req.UserId = id
	resp, err := h.cli.List(ctx.Request().Context(), req)

	return encoder.Marshal(ctx.Response().Writer, resp)
}

func (h *Handler) Message(ctx echo.Context) error {
	var req = new(pb.MessageActionRequest)
	if err := encoder.Unmarshal(ctx.Request().Body, req); err != nil {
		return err
	}
	id, err := h.subjectId(ctx)
	if err != nil || id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid subject id", err.Error())
	}
	req.UserId = id
	resp, err := h.cli.Put(ctx.Request().Context(), req)

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
