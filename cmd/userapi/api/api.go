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

package api

import (
	"codeberg.org/sixwaaaay/sharing-pb"
	"codeberg.org/sixwaaaay/sharing-pb/encoder"
	"github.com/labstack/echo/v4"
	"github.com/sixwaaaay/token/rpc"
)

// UserApi defines '/users' api
type UserApi struct {
	uc pb.UserServiceClient

	auth echo.MiddlewareFunc
}

func NewUserApi(uc pb.UserServiceClient, auth echo.MiddlewareFunc) *UserApi {
	return &UserApi{uc: uc, auth: auth}
}

func (u *UserApi) Update(e *echo.Echo) {
	e.GET("/users/:id", u.Profile, u.auth)
	e.PATCH("/users", u.UpdateProfile, u.auth)
}

// FollowApi defines '/follow' api
type FollowApi struct {
	uc   pb.UserServiceClient
	auth echo.MiddlewareFunc
}

func NewFollowApi(uc pb.UserServiceClient, auth echo.MiddlewareFunc) *FollowApi {
	return &FollowApi{uc: uc, auth: auth}
}

func (f *FollowApi) Update(e *echo.Echo) {
	e.POST("/follow/following", f.Following, f.auth)
	e.POST("/follow/followers", f.Followers, f.auth)
	e.POST("/follow", f.Follow, f.auth)
}

func (f *FollowApi) Follow(c echo.Context) error {
	var req pb.FollowActionRequest
	if err := encoder.Unmarshal(c.Request().Body, &req); err != nil {
		return echo.NewHTTPError(403, "invalid request")
	}

	reply, err := f.uc.FollowAction(rpc.Ctx4H(c.Request()), &req)
	if err != nil {
		return echo.NewHTTPError(403, err.Error())
	}
	return encoder.Marshal(c.Response().Writer, reply)
}
