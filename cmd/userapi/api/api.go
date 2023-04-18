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
	"github.com/minio/minio-go/v7"
	"github.com/sixwaaaay/sharing/pkg/pb"
	"github.com/sixwaaaay/sharing/pkg/sign"
)

type UserApi struct {
	mc     *minio.Client
	uc     pb.UserServiceClient
	bucket string
	secret []byte
}

func NewUserApi(mc *minio.Client, uc pb.UserServiceClient, bucket string, secret string) *UserApi {
	return &UserApi{
		mc:     mc,
		uc:     uc,
		bucket: bucket,
		secret: []byte(secret),
	}
}

func (u *UserApi) Update(e *echo.Echo) {
	e.GET("/users/:id", u.Profile, echo.WrapMiddleware(sign.Middleware(u.secret, false)))
	e.PATCH("/users", u.UpdateProfile, echo.WrapMiddleware(sign.Middleware(u.secret, true)))
}
