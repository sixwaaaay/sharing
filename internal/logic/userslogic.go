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

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sixwaaaay/shauser/internal/data"

	"github.com/sixwaaaay/shauser/internal/config"
	"github.com/sixwaaaay/shauser/user"
)

type UsersLogic struct {
	conf    *config.Config
	userQ   *data.UserQuery
	followQ *data.FollowQuery
	logger  *zap.Logger
}

func NewUsersLogic(conf *config.Config, userQ *data.UserQuery, followQ *data.FollowQuery, logger *zap.Logger) *UsersLogic {
	return &UsersLogic{conf: conf, userQ: userQ, followQ: followQ, logger: logger}
}

// GetUser is a method of the UsersLogic struct.
// It takes a context and a GetUserRequest as parameters.
// It returns a GetUserReply and an error.
//
// The method first checks if the UserId in the GetUserRequest is 0, and if so, returns an error.
// Then it tries to find a user with the UserId in the database.
// If the user is not found, it returns an error.
// Then it checks if the user is following the subject (another user) by calling the FindFollowExits method.
// If there is an error in this process, it returns the error.
// Then it converts the user data to a different format using the covertUser function.
// It also sets the IsFollow field of the user data to true if the user is following the subject.
// Finally, it returns a GetUserReply containing the user data.
func (l *UsersLogic) GetUser(ctx context.Context, in *user.GetUserRequest) (*user.GetUserReply, error) {
	if in.UserId == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid user id")
	}

	one, err := l.userQ.FindOne(ctx, in.UserId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	list, err := l.followQ.FindFollowExits(ctx, in.SubjectId, []int64{in.UserId})
	if err != nil {
		return nil, err
	}

	u := covertUser(one)
	u.IsFollow = len(list) > 0

	return &user.GetUserReply{User: u}, nil
}

// GetUsers is a method of the UsersLogic struct.
// It takes a context and a GetUsersRequest as parameters.
// It returns a GetUsersReply and an error.
//
// The method first calls the makeUsers function to get a list of users.
// The makeUsers function takes the context, a list of user IDs, the subject ID, and two query objects as parameters.
// If there is an error in this process, it returns the error.
// Finally, it returns a GetUsersReply containing the list of users.
func (l *UsersLogic) GetUsers(ctx context.Context, in *user.GetUsersRequest) (*user.GetUsersReply, error) {

	users, err := makeUsers(ctx, in.UserIds, in.SubjectId, l.userQ, l.followQ)
	if err != nil {
		return nil, err
	}

	return &user.GetUsersReply{Users: users}, nil
}
