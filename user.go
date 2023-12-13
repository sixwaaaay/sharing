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
	"net"
	"os"
	"os/signal"
	"syscall"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/sixwaaaay/shauser/internal/config"
	"github.com/sixwaaaay/shauser/internal/data"
	"github.com/sixwaaaay/shauser/user"
)

var configFile = flag.String("f", "configs/config.yaml", "the config file")

func main() {
	flag.Parse()
	newConfig, err := config.NewConfig(*configFile)
	if err != nil {
		panic(err)
	}

	tp, err := TracerProvider(&newConfig)
	if err != nil {
		panic(err)
	}
	defer tp.Shutdown(context.Background())
	mp, err := MeterProvider(&newConfig)
	if err != nil {
		panic(err)
	}
	defer mp.Shutdown(context.Background())

	db, err := data.NewData(&newConfig)
	if err != nil {
		panic(err)
	}
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	server := NewServer(&newConfig, db, logger)
	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)
	user.RegisterUserServiceServer(grpcServer, server)
	ln, err := net.Listen("tcp", newConfig.ListenOn)
	if err != nil {
		panic(err)
	}

	// graceful shutdown
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-ch
		grpcServer.GracefulStop()
	}()
	if err := grpcServer.Serve(ln); err != nil {
		panic(err)
	}
}
