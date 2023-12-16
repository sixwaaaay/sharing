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
	"net/http"
	"os"
	"os/signal"
	"time"

	pb "codeberg.org/sixwaaaay/sharing-pb"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sixwaaaay/must"
	"github.com/sixwaaaay/token/web"
	"github.com/spf13/viper"
	_ "go.uber.org/automaxprocs"

	"github.com/sixwaaaay/sharing/cmd/userapi/api"
)

type Config struct {
	ListenOn           string
	UserServiceAddress string
	Secret             string
}

var configFile = flag.String("f", "configs/config.yaml", "the config file")

func main() {
	viper.SetConfigFile(*configFile)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	var config Config
	must.RunE(viper.ReadInConfig())
	must.RunE(viper.Unmarshal(&config))

	e := newServer()

	conn := must.Must(pb.Dial(config.UserServiceAddress))
	client := pb.NewUserServiceClient(conn)

	auth := echo.WrapMiddleware(web.Middleware([]byte(config.Secret)))
	// add user api
	userApi := api.NewUserApi(client, auth)
	userApi.Update(e)

	// add follow api
	followApi := api.NewFollowApi(client, auth)
	followApi.Update(e)

	// Start server
	go func() {
		if err := e.Start(config.ListenOn); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
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
