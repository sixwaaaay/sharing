//go:build wireinject

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
	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/sixwaaaay/shauser/internal/config"
	"github.com/sixwaaaay/shauser/internal/data"
	"github.com/sixwaaaay/shauser/internal/logic"
	"github.com/sixwaaaay/shauser/internal/server"
)

func NewServer(config *config.Config, db *gorm.DB) *server.UserServiceServer {
	wire.Build(
		server.NewUserServiceServer,
		wire.Struct(new(server.ServerOption), "*"),
		logic.NewFollowActionLogic,
		logic.NewGetUserLogic,
		logic.NewGetUsersLogic,
		logic.NewGetFollowersLogic,
		logic.NewGetFollowingsLogic,
		logic.NewGetFriendsLogic,
		logic.NewRegisterLogic,
		logic.NewLoginLogic,
		logic.NewUpdateUserLogic,
		data.NewUserQuery,
		data.NewFollowQuery,
		data.NewFollowCommand,
		data.NewUserCommand,
	)
	return &server.UserServiceServer{}
}
