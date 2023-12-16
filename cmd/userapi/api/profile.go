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

	pb "codeberg.org/sixwaaaay/sharing-pb"
	"codeberg.org/sixwaaaay/sharing-pb/encoder"
	"github.com/labstack/echo/v4"
	"github.com/sixwaaaay/token/rpc"
)

func (u *UserApi) Profile(c echo.Context) error {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(403, "invalid user id")
	}
	getUserRequest := &pb.GetUserRequest{
		UserId: userID,
	}
	user, err := u.uc.GetUser(rpc.Ctx4H(c.Request()), getUserRequest)
	if err != nil {
		return err
	}
	return encoder.Marshal(c.Response(), user.User)
}

func (u *UserApi) UpdateProfile(ctx echo.Context) error {
	var req = new(pb.UpdateUserRequest)
	var err error
	if err := encoder.Unmarshal(ctx.Request().Body, req); err != nil {
		return echo.NewHTTPError(400, err)
	}
	r, err := u.uc.UpdateUser(rpc.Ctx4H(ctx.Request()), &pb.UpdateUserRequest{
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
