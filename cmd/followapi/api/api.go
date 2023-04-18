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
	"github.com/labstack/echo/v4"
	"github.com/sixwaaaay/sharing/pkg/encoder"
	"github.com/sixwaaaay/sharing/pkg/pb"
	"github.com/sixwaaaay/sharing/pkg/sign"
	"strconv"
)

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

func (u *FollowApi) Update(e *echo.Echo) {
	e.POST("/follow/following", u.Following, echo.WrapMiddleware(sign.Middleware(u.secret, false)))
	e.POST("/follow/followers", u.Followers, echo.WrapMiddleware(sign.Middleware(u.secret, false)))
	e.POST("/follow", u.Follow, echo.WrapMiddleware(sign.Middleware(u.secret, true)))
	e.POST("/follow/friends", u.Friends, echo.WrapMiddleware(sign.Middleware(u.secret, true)))
}

func (h *FollowApi) subjectId(ctx echo.Context) (int64, error) {
	subjectID, _ := ctx.Request().Context().Value("x-id").(string)
	id, err := strconv.ParseInt(subjectID, 10, 64)
	return id, err
}

func (u *FollowApi) Follow(c echo.Context) error {
	var req pb.FollowActionRequest
	if err := encoder.Unmarshal(c.Request().Body, &req); err != nil {
		return echo.NewHTTPError(403, "invalid request")
	}
	id, err := u.subjectId(c)
	if err != nil {
		return echo.NewHTTPError(403, "invalid token")
	}
	req.SubjectId = id
	reply, err := u.uc.FollowAction(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(403, err.Error())
	}
	return encoder.Marshal(c.Response().Writer, reply)
}

func (u *FollowApi) Friends(c echo.Context) error {
	var req pb.GetFriendsRequest
	if err := encoder.Unmarshal(c.Request().Body, &req); err != nil {
		return echo.NewHTTPError(403, "invalid request")
	}
	id, err := u.subjectId(c)
	if err != nil {
		return echo.NewHTTPError(403, "invalid token")
	}
	req.SubjectId = id
	reply, err := u.uc.GetFriends(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(403, err.Error())
	}
	return encoder.Marshal(c.Response().Writer, reply)
}
