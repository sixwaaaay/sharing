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
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	"github.com/sixwaaaay/shauser/internal/config"
	"github.com/sixwaaaay/shauser/internal/data"
)

var configFile = flag.String("f", "configs/config.yaml", "the config file")

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	flag.Parse()
	newConfig, err := config.NewConfig(*configFile)
	if err != nil {
		panic(err)
	}

	db, err := data.NewData(&newConfig)
	if err != nil {
		panic(err)
	}

	api := NewServer(&newConfig, db)

	r.Method(http.MethodPost, "/sixwaaaay.user.UserService/GetUser", Handler(api.GetUserHandler))
	r.Method(http.MethodPost, "/sixwaaaay.user.UserService/GetUsers", Handler(api.GetUsersHandler))

	// graceful shutdown
	server := http.Server{Addr: newConfig.ListenOn, Handler: r}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-ch
		if err := server.Shutdown(context.Background()); err != nil {
			panic(err)
		}
	}()

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
