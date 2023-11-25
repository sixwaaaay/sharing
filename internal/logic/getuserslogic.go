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

type GetUsersLogic struct {
	conf    *config.Config
	userQ   *data.UserQuery
	followQ *data.FollowQuery
}

type GetUsersLogicOption struct {
	Config  *config.Config
	UserQ   *data.UserQuery
	FollowQ *data.FollowQuery
}

func NewGetUsersLogic(opt GetUsersLogicOption) *GetUsersLogic {
	return &GetUsersLogic{
		conf:    opt.Config,
		userQ:   opt.UserQ,
		followQ: opt.FollowQ,
	}
}

func (l *GetUsersLogic) GetUsers(ctx context.Context, in *user.GetUsersRequest) (*user.GetUsersReply, error) {
	many, err := l.userQ.FindMany(ctx, in.UserIds)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "users not found")
	}

	var list []int64
	if in.SubjectId != 0 {
		list, err = l.followQ.FindFollowExits(ctx, in.SubjectId, in.UserIds)
		if err != nil {
			return nil, status.Errorf(codes.NotFound, "users not found")
		}
	}

	users := makeUsers(many, list)

	resp := &user.GetUsersReply{
		Users: users,
	}
	return resp, nil
}

func makeUsers(many []*data.User, list []int64) []*user.User {
	users := make([]*user.User, 0, len(many))
	m := idsToMap(list)
	for _, u := range many {
		t := covertUser(u)
		_, t.IsFollow = m[u.ID]
		users = append(users, t)
	}
	return users
}

func idsToMap(ids []int64) map[int64]struct{} {
	m := make(map[int64]struct{}, len(ids))
	for _, id := range ids {
		m[id] = struct{}{}
	}
	return m
}
