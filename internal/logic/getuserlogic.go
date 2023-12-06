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

type GetUserLogic struct {
	conf    *config.Config
	userQ   *data.UserQuery
	followQ *data.FollowQuery
}

func NewGetUserLogic(conf *config.Config, userQ *data.UserQuery, followQ *data.FollowQuery) *GetUserLogic {
	return &GetUserLogic{conf: conf, userQ: userQ, followQ: followQ}
}

func (l *GetUserLogic) GetUser(ctx context.Context, in *user.GetUserRequest) (*user.GetUserReply, error) {
	one, err := l.userQ.FindOne(ctx, in.UserId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "u not found")
	}
	list, err := l.followQ.FindFollowExits(ctx, in.SubjectId, []int64{in.UserId})
	if err != nil {
		return nil, err
	}
	u := covertUser(one)
	u.IsFollow = len(list) > 0
	reply := &user.GetUserReply{
		User: u,
	}
	return reply, nil
}

func covertUser(one *data.User) *user.User {
	u := &user.User{
		Id:            one.ID,
		Name:          one.Username,
		AvatarUrl:     one.AvatarURL,
		BgUrl:         one.BgURL,
		Bio:           one.Bio,
		LikesGiven:    one.LikesGiven,
		LikesReceived: one.LikesReceived,
		VideosPosted:  one.VideosPosted,
		Following:     one.Following,
		Followers:     one.Followers,
	}
	return u
}
