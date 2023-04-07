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

type FollowActionLogic struct {
	conf   *config.Config
	follow *data.FollowCommand
}

type FollowActionLogicOption struct {
	Config *config.Config
	Follow *data.FollowCommand
}

func NewFollowActionLogic(opt FollowActionLogicOption) *FollowActionLogic {
	return &FollowActionLogic{
		conf:   opt.Config,
		follow: opt.Follow,
	}
}

func (l *FollowActionLogic) FollowAction(ctx context.Context, in *user.FollowActionRequest) (*user.FollowActionReply, error) {
	if in.UserId == 0 || in.SubjectId == 0 || in.UserId == in.SubjectId {
		return nil, status.Error(codes.InvalidArgument, "invalid user id")
	}
	if in.Action == 1 {
		rel := &data.Follow{
			UserID: in.SubjectId,
			Target: in.UserId,
		}
		err := l.follow.Insert(ctx, rel)
		if err != nil {
			return nil, err
		}
		return &user.FollowActionReply{}, nil
	} else if in.Action == 2 {
		err := l.follow.Delete(ctx, in.SubjectId, in.UserId)
		if err != nil {
			return nil, err
		}
		return &user.FollowActionReply{}, nil
	}
	return nil, status.Error(codes.InvalidArgument, "invalid action")
}
