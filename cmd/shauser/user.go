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

package main

import (
	"context"
	"errors"
	"flag"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sixwaaaay/must"
	"github.com/sixwaaaay/token/rpc"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/sixwaaaay/shauser/internal/config"
	"github.com/sixwaaaay/shauser/internal/repository"
	"github.com/sixwaaaay/shauser/internal/web"
	"github.com/sixwaaaay/shauser/user"
)

var configFile = flag.String("f", "configs/config.yaml", "the config file")

func main() {
	flag.Parse()
	ctx := context.Background()
	conf := must.Must(config.NewConfig(*configFile))
	tp := must.Must(TracerProvider(&conf))
	defer must.RunE(tp.Shutdown(ctx))

	mp := must.Must(MeterProvider(&conf))
	defer mp.Shutdown(ctx)

	db := must.Must(repository.NewDB(&conf))

	logger := must.Must(zap.NewDevelopment())
	defer logger.Sync()

	server := NewServer(&conf, db, logger)

	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.UnaryInterceptor(rpc.Handler([]byte(conf.Secret))))
	user.RegisterUserServiceServer(grpcServer, server)

	ln := must.Must(net.Listen("tcp", conf.ListenOn))

	m := web.M(conf, server)
	// graceful shutdown
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := m.Start(conf.HTTP); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("failed to start http server", zap.Error(err))
		}
	}()
	go func() {
		if err := grpcServer.Serve(ln); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			logger.Error("failed to start grpc server", zap.Error(err))
		}
	}()

	<-ch
	grpcServer.GracefulStop()
	must.RunE(m.Shutdown(ctx))
}
