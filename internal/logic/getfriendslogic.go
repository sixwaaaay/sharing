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

package logic

import (
	"context"
	"github.com/sixwaaaay/shauser/internal/data"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sixwaaaay/shauser/internal/config"
	"github.com/sixwaaaay/shauser/user"
)

type GetFriendsLogic struct {
	conf    *config.Config
	userQ   *data.UserQuery
	followQ *data.FollowQuery
}

type GetFriendsLogicOption struct {
	Config  *config.Config
	UserQ   *data.UserQuery
	FollowQ *data.FollowQuery
}

func NewGetFriendsLogic(opt GetFriendsLogicOption) *GetFriendsLogic {
	return &GetFriendsLogic{
		conf:    opt.Config,
		userQ:   opt.UserQ,
		followQ: opt.FollowQ,
	}
}

func (l *GetFriendsLogic) GetFriends(ctx context.Context, in *user.GetFriendsRequest) (*user.GetFriendsReply, error) {
	if in.Limit == 0 || in.Limit > l.conf.MaxLimit {
		in.Limit = l.conf.DefaultLimit
	}
	list, err := l.followQ.FindFriends(ctx, in.UserId, in.Token, int(in.Limit))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "users not found")
	}
	many, err := l.userQ.FindMany(ctx, list)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "users not found")
	}
	users := makeUsers(many, list)
	return &user.GetFriendsReply{
		Users: users,
	}, nil
}
