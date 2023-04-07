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
	"github.com/sixwaaaay/shauser/internal/config"
	"github.com/sixwaaaay/shauser/internal/data"
	"github.com/sixwaaaay/shauser/internal/logic"
	"github.com/sixwaaaay/shauser/internal/server"
	"gorm.io/gorm"
)

func NewServer(config *config.Config, db *gorm.DB) *server.UserServiceServer {
	wire.Build(
		server.NewUserServiceServer,
		wire.Struct(new(server.ServerOption), "*"),
		wire.Struct(new(logic.FollowActionLogicOption), "*"),
		logic.NewFollowActionLogic,
		wire.Struct(new(logic.GetUserLogicOption), "*"),
		logic.NewGetUserLogic,
		wire.Struct(new(logic.GetUsersLogicOption), "*"),
		logic.NewGetUsersLogic,
		wire.Struct(new(logic.GetFollowersLogicOption), "*"),
		logic.NewGetFollowersLogic,
		wire.Struct(new(logic.GetFollowingsLogicOption), "*"),
		logic.NewGetFollowingsLogic,
		wire.Struct(new(logic.GetFriendsLogicOption), "*"),
		logic.NewGetFriendsLogic,
		wire.Struct(new(logic.RegisterLogicOption), "*"),
		logic.NewRegisterLogic,
		wire.Struct(new(logic.LoginLogicOption), "*"),
		logic.NewLoginLogic,
		data.NewUserQuery,
		data.NewFollowQuery,
		data.NewFollowCommand,
		data.NewUserCommand,
	)
	return &server.UserServiceServer{}
}
