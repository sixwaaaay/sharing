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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sixwaaaay/shauser/internal/data"

	"github.com/sixwaaaay/shauser/internal/config"
	"github.com/sixwaaaay/shauser/user"
)

type GetFollowersLogic struct {
	conf     *config.Config
	followQ  *data.FollowQuery
	getUsers *GetUsersLogic
}

func NewGetFollowersLogic(conf *config.Config, followQ *data.FollowQuery, getUsers *GetUsersLogic) *GetFollowersLogic {
	return &GetFollowersLogic{conf: conf, followQ: followQ, getUsers: getUsers}
}

func (l *GetFollowersLogic) GetFollowers(ctx context.Context, in *user.GetFollowersRequest) (*user.GetFollowersReply, error) {
	if in.Limit == 0 || in.Limit > l.conf.MaxLimit {
		in.Limit = l.conf.DefaultLimit
	}
	list, err := l.followQ.FindFollowers(ctx, in.UserId, in.Token, int(in.Limit))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get followers failed: %v", err)
	}
	usersReply, err := l.getUsers.GetUsers(ctx, &user.GetUsersRequest{
		UserIds:   list,
		SubjectId: in.SubjectId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get users failed: %v", err)
	}
	return &user.GetFollowersReply{
		Users: usersReply.Users,
	}, nil
}
