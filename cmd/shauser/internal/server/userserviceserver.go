/*
 * Copyright (c) 2024 sixwaaaay.
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

package server

import (
	"context"

	"github.com/sixwaaaay/shauser/internal/logic"
	"github.com/sixwaaaay/shauser/user"
)

type UserServiceServer struct {
	GetUsersLogic     *logic.UsersLogic
	SignLogic         *logic.SignLogic
	FollowActionLogic *logic.FollowActionLogic
	GetFollowingLogic *logic.FollowQueryLogic
	UpdateUserLogic   *logic.UpdateUserLogic
	user.UnimplementedUserServiceServer
}

func NewUserServiceServer(getUsersLogic *logic.UsersLogic, signLogic *logic.SignLogic, followActionLogic *logic.FollowActionLogic, getFollowingLogic *logic.FollowQueryLogic, updateUserLogic *logic.UpdateUserLogic) *UserServiceServer {
	return &UserServiceServer{GetUsersLogic: getUsersLogic, SignLogic: signLogic, FollowActionLogic: followActionLogic, GetFollowingLogic: getFollowingLogic, UpdateUserLogic: updateUserLogic}
}

func (s *UserServiceServer) GetByMail(ctx context.Context, req *user.GetByMailReq) (*user.GetByMailReply, error) {
	return s.GetUsersLogic.GetByMail(ctx, req)
}

func (s *UserServiceServer) GetUser(ctx context.Context, in *user.GetUserRequest) (*user.GetUserReply, error) {
	return s.GetUsersLogic.GetUser(ctx, in)
}

func (s *UserServiceServer) GetUsers(ctx context.Context, in *user.GetUsersRequest) (*user.GetUsersReply, error) {
	return s.GetUsersLogic.GetUsers(ctx, in)
}

func (s *UserServiceServer) Login(ctx context.Context, in *user.LoginRequest) (*user.LoginReply, error) {
	return s.SignLogic.Login(ctx, in)
}

func (s *UserServiceServer) Register(ctx context.Context, in *user.RegisterRequest) (*user.RegisterReply, error) {
	return s.SignLogic.Register(ctx, in)
}

func (s *UserServiceServer) FollowAction(ctx context.Context, in *user.FollowActionRequest) (*user.FollowActionReply, error) {
	return s.FollowActionLogic.FollowAction(ctx, in)
}

func (s *UserServiceServer) GetFollowers(ctx context.Context, in *user.FollowQueryReq) (*user.UsersPage, error) {
	return s.GetFollowingLogic.GetFollowers(ctx, in)
}

func (s *UserServiceServer) GetFollowings(ctx context.Context, in *user.FollowQueryReq) (*user.UsersPage, error) {
	return s.GetFollowingLogic.GetFollowings(ctx, in)
}

func (s *UserServiceServer) UpdateUser(ctx context.Context, in *user.UpdateUserRequest) (*user.UpdateUserReply, error) {
	return s.UpdateUserLogic.UpdateUser(ctx, in)
}
