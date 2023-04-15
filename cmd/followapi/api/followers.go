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
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/sixwaaaay/sharing/pkg/pb"
)

type GetFollowersRequest struct {
	UserId int64 `json:"user_id"`
	Limit  int64 `json:"limit"`
	Token  int64 `json:"token"`
}

func (r *GetFollowersRequest) Validate() error {
	if r.UserId <= 0 {
		return echo.NewHTTPError(403, "invalid user id")
	}
	if r.Limit < 0 {
		return echo.NewHTTPError(403, "invalid limit")
	}
	if r.Token < 0 { // for pagination
		return echo.NewHTTPError(403, "invalid pagination token")
	}
	return nil
}

type GetFollowersReply struct {
	Users []*pb.User `json:"users"`
}

func (u *FollowApi) Followers(c echo.Context) error {
	var req GetFollowersRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return echo.NewHTTPError(403, "invalid request")
	}
	if err := req.Validate(); err != nil {
		return err
	}
	id, err := u.subjectId(c)
	if err != nil {
		return echo.NewHTTPError(403, "invalid token")
	}
	reply, err := u.uc.GetFollowers(c.Request().Context(), &pb.GetFollowersRequest{
		UserId:    req.UserId,
		SubjectId: id,
		Limit:     req.Limit,
		Token:     req.Token,
	})
	if err != nil {
		return echo.NewHTTPError(500, "internal error")
	}
	return c.JSON(200, GetFollowersReply{
		Users: reply.Users,
	})
}
