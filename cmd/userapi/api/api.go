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
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sixwaaaay/sharing/pkg/encoder"
	"github.com/sixwaaaay/sharing/pkg/pb"
	"github.com/sixwaaaay/sharing/pkg/sign"
)

// UserApi defines '/users' api
type UserApi struct {
	uc     pb.UserServiceClient
	secret []byte
}

func NewUserApi(uc pb.UserServiceClient, secret string) *UserApi {
	return &UserApi{
		uc:     uc,
		secret: []byte(secret),
	}
}

func (u *UserApi) Update(e *echo.Echo) {
	e.GET("/users/:id", u.Profile, echo.WrapMiddleware(sign.Middleware(u.secret, false)))
	e.PATCH("/users", u.UpdateProfile, echo.WrapMiddleware(sign.Middleware(u.secret, true)))
}

// FollowApi defines '/follow' api
type FollowApi struct {
	uc     pb.UserServiceClient
	secret []byte
}

func NewFollowApi(uc pb.UserServiceClient, secret string) *FollowApi {
	return &FollowApi{
		uc:     uc,
		secret: []byte(secret),
	}
}

func (f *FollowApi) Update(e *echo.Echo) {
	e.POST("/follow/following", f.Following, echo.WrapMiddleware(sign.Middleware(f.secret, false)))
	e.POST("/follow/followers", f.Followers, echo.WrapMiddleware(sign.Middleware(f.secret, false)))
	e.POST("/follow", f.Follow, echo.WrapMiddleware(sign.Middleware(f.secret, true)))
	e.POST("/follow/friends", f.Friends, echo.WrapMiddleware(sign.Middleware(f.secret, true)))
}

func (f *FollowApi) subjectId(ctx echo.Context) (int64, error) {
	subjectID, _ := ctx.Request().Context().Value("x-id").(string)
	id, err := strconv.ParseInt(subjectID, 10, 64)
	return id, err
}

func (f *FollowApi) Follow(c echo.Context) error {
	var req pb.FollowActionRequest
	if err := encoder.Unmarshal(c.Request().Body, &req); err != nil {
		return echo.NewHTTPError(403, "invalid request")
	}
	id, err := f.subjectId(c)
	if err != nil {
		return echo.NewHTTPError(403, "invalid token")
	}
	req.SubjectId = id
	reply, err := f.uc.FollowAction(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(403, err.Error())
	}
	return encoder.Marshal(c.Response().Writer, reply)
}

func (f *FollowApi) Friends(c echo.Context) error {
	var req pb.GetFriendsRequest
	if err := encoder.Unmarshal(c.Request().Body, &req); err != nil {
		return echo.NewHTTPError(403, "invalid request")
	}
	id, err := f.subjectId(c)
	if err != nil {
		return echo.NewHTTPError(403, "invalid token")
	}
	req.SubjectId = id
	reply, err := f.uc.GetFriends(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(403, err.Error())
	}
	return encoder.Marshal(c.Response().Writer, reply)
}
