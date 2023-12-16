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
	"strconv"
	"time"

	pb "codeberg.org/sixwaaaay/sharing-pb"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sixwaaaay/must"
	"github.com/sixwaaaay/token"
	"github.com/spf13/viper"
	_ "go.uber.org/automaxprocs"
	"golang.org/x/oauth2"
)

type Config struct {
	ListenOn           string
	UserServiceAddress string
	SECRET             string
	TTL                string
	Oauth              oauth2.Config
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

	d := must.Must(time.ParseDuration(config.TTL))

	s := signFunc([]byte(config.SECRET), d)

	handler := NewAccountHandler(client, s)
	handler.Update(e)
	oauth := NewOauth2(&config.Oauth, client, s)
	oauth.Update(e)

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
	must.RunE(e.Shutdown(ctx))
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

func sign(secret []byte, d time.Duration, id int64, name string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"exp":           time.Now().Add(d).Unix(), // seconds
			"iss":           "sharing",
			token.ClaimID:   strconv.FormatInt(id, 10),
			token.ClaimName: name,
		})
	s, err := t.SignedString(secret)
	return s, err
}

type signer = func(id int64, name string) (string, error)

func signFunc(secret []byte, d time.Duration) signer {
	return func(id int64, name string) (string, error) {
		return sign(secret, d, id, name)
	}
}
