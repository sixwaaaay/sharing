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
	"context"
	"mime/multipart"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
	"github.com/sixwaaaay/sharing/pkg/encoder"
	"github.com/sixwaaaay/sharing/pkg/pb"
)

type UpdateProfileRequest struct {
	UserId    string `form:"user_id" validate:"required"`
	Name      string `form:"name" validate:"required"`
	Bio       string `form:"bio"`
	Avatar    *multipart.FileHeader
	Bg        *multipart.FileHeader
	AvatarUrl string `form:"avatar_url"`
	BgUrl     string `form:"bg_url"`
}

type UpdateProfileResponse struct {
	Profile *pb.User `json:"profile"`
}

func (u *UserApi) UpdateProfile(ctx echo.Context) error {
	var req UpdateProfileRequest
	var err error
	// get header value
	subjectId, ok := ctx.Request().Context().Value("x-id").(string)
	if !ok {
		return echo.NewHTTPError(403, "token is not valid")
	}
	req.UserId = subjectId
	if req.Avatar, err = ctx.FormFile("avatar"); err != nil {
	}
	if req.Bg, err = ctx.FormFile("background"); err != nil {
	}
	req.Name = ctx.FormValue("name")
	req.Bio = ctx.FormValue("bio")
	if req.Avatar != nil {
		req.AvatarUrl, err = u.uploadFile(ctx.Request().Context(), req.Avatar)
		if err != nil {
			return echo.NewHTTPError(500, err)
		}
	}
	//"github.com/golang/protobuf/jsonpb"
	if req.Bg != nil {
		req.BgUrl, err = u.uploadFile(ctx.Request().Context(), req.Bg)
		if err != nil {
			return echo.NewHTTPError(500, err)
		}
	}

	r, err := u.updateUser(ctx.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(500, err)
	}
	return encoder.Marshal(ctx.Response(), r)
}

func (u *UserApi) updateUser(ctx context.Context, req *UpdateProfileRequest) (*pb.UpdateUserReply, error) {
	id, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		return nil, err
	}
	reply, err := u.uc.UpdateUser(ctx, &pb.UpdateUserRequest{
		UserId:    id,
		Name:      req.Name,
		Bio:       req.Bio,
		AvatarUrl: req.AvatarUrl,
		BgUrl:     req.BgUrl,
	})
	if err != nil {
		return nil, err
	}
	return reply, err
}

func (u *UserApi) uploadFile(ctx context.Context, avatar *multipart.FileHeader) (string, error) {
	//	gen a uuid
	id := uuid.New().String()
	// open source file
	src, err := avatar.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	_, err = u.mc.PutObject(ctx, u.bucket, id, src, -1, minio.PutObjectOptions{
		ContentType: avatar.Header.Get("Content-Type"),
	})
	if err != nil {
		return "", err
	}
	return id, nil
}
