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

	"github.com/sixwaaaay/must"
	"github.com/sixwaaaay/token/rpc"
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
	ctx := context.Background()
	newConfig := must.Must(config.NewConfig(*configFile))
	tp := must.Must(TracerProvider(&newConfig))
	defer must.RunE(tp.Shutdown(ctx))

	mp := must.Must(MeterProvider(&newConfig))
	defer must.RunE(mp.Shutdown(ctx))

	db := must.Must(data.NewData(&newConfig))

	logger := must.Must(zap.NewDevelopment())
	defer must.RunE(logger.Sync())

	server := NewServer(&newConfig, db, logger)

	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.UnaryInterceptor(rpc.Handler([]byte(newConfig.Secret))))
	user.RegisterUserServiceServer(grpcServer, server)

	ln := must.Must(net.Listen("tcp", newConfig.ListenOn))

	// graceful shutdown
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	go must.RunE(grpcServer.Serve(ln))

	<-ch
	grpcServer.GracefulStop()
}
