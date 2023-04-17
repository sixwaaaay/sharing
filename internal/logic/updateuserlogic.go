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
	"github.com/sixwaaaay/shauser/internal/config"
	"github.com/sixwaaaay/shauser/internal/data"
	"github.com/sixwaaaay/shauser/user"
)

type UpdateUserLogic struct {
	conf        *config.Config
	userCommand *data.UserCommand
}

type UpdateUserLogicOption struct {
	Config      *config.Config
	UserCommand *data.UserCommand
}

func NewUpdateUserLogic(opt UpdateUserLogicOption) *UpdateUserLogic {
	return &UpdateUserLogic{
		conf:        opt.Config,
		userCommand: opt.UserCommand,
	}
}

func (l *UpdateUserLogic) UpdateUser(ctx context.Context, in *user.UpdateUserRequest) (*user.UpdateUserReply, error) {
	var u data.User
	u.Username = in.Name
	u.Bio = in.Bio
	u.AvatarURL = in.AvatarUrl
	u.BgURL = in.BgUrl
	if err := l.userCommand.UpdateUser(ctx, &u); err != nil {
		return nil, err
	}
	return &user.UpdateUserReply{}, nil
}
