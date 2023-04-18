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
)

type GetFollowingsRequest struct {
	UserId int64 `json:"user_id"`
	Limit  int64 `json:"limit"`
	Token  int64 `json:"token"`
}

func (r *GetFollowingsRequest) Validate() error {
	if r.UserId <= 0 {
		return echo.NewHTTPError(403, "invalid user id")
	}
	if r.Limit < 0 {
		return echo.NewHTTPError(403, "invalid limit")
	}
	if r.Token < 0 {
		return echo.NewHTTPError(403, "invalid pagination token")
	}
	return nil
}

type GetFollowingsReply struct {
	Users []*pb.User `json:"users"`
}

func (u *FollowApi) Following(c echo.Context) error {
	var req pb.GetFollowingsRequest
	if err := encoder.Unmarshal(c.Request().Body, &req); err != nil {
		return echo.NewHTTPError(403, "invalid request")
	}
	id, err := u.subjectId(c)
	if err != nil {
		return echo.NewHTTPError(403, "invalid token")
	}
	req.SubjectId = id
	users, err := u.uc.GetFollowings(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(403, "invalid token")
	}
	return encoder.Marshal(c.Response().Writer, users)
}
