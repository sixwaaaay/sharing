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
	"strconv"
)

type ProfileRequest struct {
	UserId    int64
	SubjectID int64
}

func (r *ProfileRequest) Validate() error {
	if r.UserId <= 0 {
		return echo.NewHTTPError(403, "invalid user id")
	}
	if r.SubjectID < 0 {
		return echo.NewHTTPError(403, "invalid token")
	}
	return nil
}

type ProfileResponse struct {
	Profile *pb.User `json:"profile"`
}

func (u *UserApi) Profile(c echo.Context) error {
	var req ProfileRequest
	var err error
	subjectId, _ := c.Request().Context().Value("x-id").(string)
	if req.UserId, err = strconv.ParseInt(c.Param("id"), 10, 64); err != nil {
		return echo.NewHTTPError(403, "invalid user id")
	}
	if req.SubjectID, err = strconv.ParseInt(subjectId, 10, 64); err != nil {
		return echo.NewHTTPError(403, "invalid token")
	}
	if err := req.Validate(); err != nil {
		return err
	}

	getUserRequest := &pb.GetUserRequest{
		UserId:    req.UserId,
		SubjectId: req.SubjectID,
	}
	user, err := u.uc.GetUser(c.Request().Context(), getUserRequest)
	if err != nil {
		return err
	}
	return encoder.Marshal(c.Response(), user.User)
}
