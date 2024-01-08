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

package api

import (
	"codeberg.org/sixwaaaay/sharing-pb/encoder"
	"github.com/labstack/echo/v4"

	"github.com/sixwaaaay/shauser/internal/server"
	"github.com/sixwaaaay/shauser/user"
)

// UserApi defines '/users' api
type UserApi struct {
	uc *server.UserServiceServer

	auth echo.MiddlewareFunc
}

func NewUserApi(uc *server.UserServiceServer, auth echo.MiddlewareFunc) *UserApi {
	return &UserApi{uc: uc, auth: auth}
}

func (u *UserApi) Update(e *echo.Echo) {
	e.GET("/users/:user_id", u.Profile, u.auth)
	e.GET("/users", u.MultipleProfile, u.auth)
	e.PATCH("/users", u.UpdateProfile, u.auth)
}

func (u *UserApi) Profile(c echo.Context) error {
	var req = new(user.GetUserRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(400, err)
	}
	userReply, err := u.uc.GetUser(c.Request().Context(), req)
	if err != nil {
		return err
	}
	return encoder.Marshal(c.Response(), userReply)
}

func (u *UserApi) MultipleProfile(c echo.Context) error {
	var req = new(user.GetUsersRequest) // ids is the url query param
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(400, err)
	}
	r, err := u.uc.GetUsers(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(500, err)
	}
	return encoder.Marshal(c.Response(), r)
}

func (u *UserApi) UpdateProfile(ctx echo.Context) error {
	var req = new(user.UpdateUserRequest)
	if err := encoder.Unmarshal(ctx.Request().Body, req); err != nil {
		return echo.NewHTTPError(400, err)
	}
	r, err := u.uc.UpdateUser(ctx.Request().Context(), &user.UpdateUserRequest{
		Name:      req.Name,
		Bio:       req.Bio,
		AvatarUrl: req.AvatarUrl,
		BgUrl:     req.BgUrl,
	})
	if err != nil {
		return echo.NewHTTPError(500, err)
	}
	return encoder.Marshal(ctx.Response(), r)
}

// FollowApi defines '/follow' api
type FollowApi struct {
	uc   *server.UserServiceServer
	auth echo.MiddlewareFunc
}

func NewFollowApi(uc *server.UserServiceServer, auth echo.MiddlewareFunc) *FollowApi {
	return &FollowApi{uc: uc, auth: auth}
}

func (f *FollowApi) Update(e *echo.Echo) {
	e.GET("/users/:user_id/following", f.Following, f.auth)
	e.GET("/users/:user_id/followers", f.Followers, f.auth)
	e.POST("/follow", f.Follow, f.auth)
}

func (f *FollowApi) Follow(c echo.Context) error {
	var req user.FollowActionRequest
	if err := encoder.Unmarshal(c.Request().Body, &req); err != nil {
		return echo.NewHTTPError(403, "invalid request")
	}

	reply, err := f.uc.FollowAction(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(403, err.Error())
	}
	return encoder.Marshal(c.Response().Writer, reply)
}

func (f *FollowApi) Followers(c echo.Context) error {
	var req user.FollowQueryReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(403, "invalid request")
	}
	reply, err := f.uc.GetFollowers(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(500, "internal error")
	}
	return encoder.Marshal(c.Response().Writer, reply)
}

func (f *FollowApi) Following(c echo.Context) error {
	var req user.FollowQueryReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(403, "invalid request")
	}
	users, err := f.uc.GetFollowings(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(403, "invalid token")
	}
	return encoder.Marshal(c.Response().Writer, users)
}
