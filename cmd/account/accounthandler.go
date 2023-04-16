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

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sixwaaaay/sharing/pkg/pb"
	"github.com/sixwaaaay/sharing/pkg/sign"
	"regexp"
	"time"
)

type AccountHandler struct {
	client pb.UserServiceClient
	jwt    sign.JWT
}

func NewAccountHandler(client pb.UserServiceClient, jwt sign.JWT) *AccountHandler {
	return &AccountHandler{
		client: client,
		jwt:    jwt,
	}
}

func (h *AccountHandler) Update(e *echo.Echo) {
	e.POST("/sign/in", h.Login)
	e.POST("/sign/up", h.Register)
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

const email = `^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+((\.[a-zA-Z0-9_-]{2,3}){1,2})$`

var emailReg = regexp.MustCompile(email)

const password = `^[\x20-\x7E]+$`

var passwordReg = regexp.MustCompile(password)

func (r LoginRequest) Validate() error {
	// regex validation
	if !emailReg.MatchString(r.Email) {
		return echo.NewHTTPError(400, "invalid email")
	}
	// password validation should only contain printable characters

	if !passwordReg.MatchString(r.Password) {
		return echo.NewHTTPError(400, "invalid password")
	}
	return nil
}

type LoginReply struct {
	Account *pb.User `json:"account"`
	Token   string   `json:"token"`
}

func (h *AccountHandler) Login(ctx echo.Context) error {
	// fluent binding
	var req LoginRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(400, err)
	}
	if err := req.Validate(); err != nil {
		return ctx.JSON(400, err)
	}

	reply, err := h.client.Login(ctx.Request().Context(), &pb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return ctx.JSON(400, err)
	}
	u := reply.GetUser()
	option := sign.SignOption{
		Username: u.Name,
		UserID:   u.Id,
		Duration: time.Duration(h.jwt.TTL) * time.Second,
		Secret:   []byte(h.jwt.Secret),
	}
	token, err := sign.GenSignedToken(option)
	if err != nil {
		return ctx.JSON(400, err)
	}

	resp := &LoginReply{
		Account: u,
		Token:   token,
	}
	return ctx.JSON(200, resp)
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Username string `json:"username" validate:"required"`
}

type RegisterReply struct {
	Account *pb.User `json:"account"`
	Token   string   `json:"token"`
}

func (r *RegisterRequest) Validate() error {
	if !emailReg.MatchString(r.Email) {
		return echo.NewHTTPError(400, "invalid email")
	}
	if !passwordReg.MatchString(r.Password) {
		return echo.NewHTTPError(400, "invalid password")
	}
	if len(r.Username) < 3 || len(r.Username) > 20 {
		return echo.NewHTTPError(400, "invalid username")
	}
	return nil
}

func (h *AccountHandler) Register(ctx echo.Context) error {
	var req RegisterRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(400, err)
	}
	if err := req.Validate(); err != nil {
		return ctx.JSON(400, err)
	}
	reply, err := h.client.Register(ctx.Request().Context(), &pb.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Username,
	})
	if err != nil {
		return ctx.JSON(400, err)
	}
	u := reply.GetUser()
	option := sign.SignOption{
		Username: u.Name,
		UserID:   u.Id,
		Duration: time.Duration(h.jwt.TTL) * time.Second,
		Secret:   []byte(h.jwt.Secret),
	}
	token, err := sign.GenSignedToken(option)
	if err != nil {
		return ctx.JSON(400, err)
	}

	resp := &RegisterReply{
		Account: u,
		Token:   token,
	}
	return ctx.JSON(200, resp)
}
