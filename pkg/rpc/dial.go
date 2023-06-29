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

package rpc

import (
	"context"
	"github.com/sixwaaaay/sharing/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func NewUserClient(conf GrpcConfig) (pb.UserServiceClient, error) {
	if conn, err := dial(conf); err != nil {
		return nil, err
	} else {
		return pb.NewUserServiceClient(conn), nil
	}
}

func NewCommentClient(conf GrpcConfig) (pb.CommentServiceClient, error) {
	if conn, err := dial(conf); err != nil {
		return nil, err
	} else {
		return pb.NewCommentServiceClient(conn), nil
	}
}

func NewVideoClient(conf GrpcConfig) (pb.VideoServiceClient, error) {
	if conn, err := dial(conf); err != nil {
		return nil, err
	} else {
		return pb.NewVideoServiceClient(conn), nil
	}
}

func NewMessageClient(conf GrpcConfig) (pb.MessageServiceClient, error) {
	if conn, err := dial(conf); err != nil {
		return nil, err
	} else {
		return pb.NewMessageServiceClient(conn), nil
	}
}

func dial(conf GrpcConfig) (*grpc.ClientConn, error) {
	var interceptors []grpc.UnaryClientInterceptor
	if conf.ServiceName != "" {
		interceptors = append(interceptors, DaprServiceInvoker(conf.ServiceName))
	}
	var options []grpc.DialOption
	options = append(options, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if len(interceptors) > 0 {
		options = append(options, grpc.WithChainUnaryInterceptor(interceptors...))
	}
	conn, err := grpc.Dial(conf.Address, options...)
	return conn, err
}

func DaprServiceInvoker(DaprServiceName string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if DaprServiceName != "" {
			ctx = metadata.AppendToOutgoingContext(ctx, "dapr-app-id", DaprServiceName)
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}