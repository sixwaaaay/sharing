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
)

func NewServer(conf *config.Config, db *gorm.DB) *WebApi {
	wire.Build(
		NewWebApi,
		wire.Struct(new(logic.GetUserLogicOption), "*"),
		logic.NewGetUserLogic,
		wire.Struct(new(logic.GetUsersLogicOption), "*"),
		logic.NewGetUsersLogic,
		data.NewUserQuery,
		data.NewFollowQuery,
	)
	return &WebApi{}
}
