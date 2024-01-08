/*
 * Copyright (c) 2023-2024 sixwaaaay.
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

package web

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sixwaaaay/must"
	"github.com/sixwaaaay/token/web"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"

	"github.com/sixwaaaay/shauser/internal/config"
	"github.com/sixwaaaay/shauser/internal/server"
	"github.com/sixwaaaay/shauser/internal/web/api"
)

func M(conf config.Config, client *server.UserServiceServer) *echo.Echo {
	e := newServer()
	auth := echo.WrapMiddleware(web.Middleware([]byte(conf.Secret)))
	userApi := api.NewUserApi(client, auth)
	userApi.Update(e)
	followApi := api.NewFollowApi(client, auth)
	followApi.Update(e)
	d := must.Must(time.ParseDuration(conf.TTL))
	s := api.SignFunc([]byte(conf.Secret), d)
	handler := api.NewAccountHandler(client, s)
	handler.Update(e)
	oauth := api.NewOauth2(&conf.Oauth, client, s)
	oauth.Update(e)
	return e
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
	e.Use(otelecho.Middleware("shauser"))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	return e
}
