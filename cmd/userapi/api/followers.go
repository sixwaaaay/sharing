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
	pb "codeberg.org/sixwaaaay/sharing-pb"
	"github.com/labstack/echo/v4"
	"github.com/sixwaaaay/token/rpc"

	"github.com/sixwaaaay/sharing/pkg/encoder"
)

type GetFollowersReply struct {
	Users []*pb.User `json:"users"`
}

func (f *FollowApi) Followers(c echo.Context) error {
	var req pb.FollowQueryReq
	if err := encoder.Unmarshal(c.Request().Body, &req); err != nil {
		return echo.NewHTTPError(403, "invalid request")
	}
	reply, err := f.uc.GetFollowers(rpc.Ctx4H(c.Request()), &req)
	if err != nil {
		return echo.NewHTTPError(500, "internal error")
	}
	return encoder.Marshal(c.Response().Writer, reply)
}

func (f *FollowApi) Following(c echo.Context) error {
	var req pb.FollowQueryReq
	if err := encoder.Unmarshal(c.Request().Body, &req); err != nil {
		return echo.NewHTTPError(403, "invalid request")
	}
	users, err := f.uc.GetFollowings(rpc.Ctx4H(c.Request()), &req)
	if err != nil {
		return echo.NewHTTPError(403, "invalid token")
	}
	return encoder.Marshal(c.Response().Writer, users)
}
