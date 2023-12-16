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
)

func (u *UserApi) UpdateProfile(ctx echo.Context) error {
	var req = new(pb.UpdateUserRequest)
	var err error // get header value
	if err := encoder.Unmarshal(ctx.Request().Body, req); err != nil {
		return echo.NewHTTPError(400, err)
	}
	subjectId, ok := ctx.Request().Context().Value("x-id").(string)
	if !ok {
		return echo.NewHTTPError(403, "token is not valid")
	}
	id, err := strconv.ParseInt(subjectId, 10, 64)
	if err != nil {
		return err
	}
	req.UserId = id
	r, err := u.uc.UpdateUser(ctx.Request().Context(), &pb.UpdateUserRequest{
		UserId:    id,
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
