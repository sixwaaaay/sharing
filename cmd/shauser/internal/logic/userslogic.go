/*
 * Copyright (c) 2023-2024 sixwaaaay.
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

	"github.com/sixwaaaay/token"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sixwaaaay/shauser/internal/repository"
	"github.com/sixwaaaay/shauser/user"
)

type UsersLogic struct {
	userQ   repository.UserQuery
	followQ repository.FollowQuery
	logger  *zap.Logger
}

func NewUsersLogic(userQ repository.UserQuery, followQ repository.FollowQuery, logger *zap.Logger) *UsersLogic {
	return &UsersLogic{userQ: userQ, followQ: followQ, logger: logger}
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
// Then it converts the user repository to a different format using the covertUser function.
// It also sets the IsFollow field of the user repository to true if the user is following the subject.
// Finally, it returns a GetUserReply containing the user repository.
func (l *UsersLogic) GetUser(ctx context.Context, in *user.GetUserRequest) (*user.GetUserReply, error) {
	userID, _ := token.ClaimStrI64(ctx, token.ClaimID)

	if in.UserId == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid user id")
	}

	one, err := l.userQ.FindOne(ctx, in.UserId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	list, err := l.followQ.FindFollowExits(ctx, userID, []int64{in.UserId})
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

	userID, _ := token.ClaimStrI64(ctx, token.ClaimID)

	users, err := makeUsers(ctx, in.UserIds, userID, l.userQ, l.followQ)
	if err != nil {
		return nil, err
	}

	return &user.GetUsersReply{Users: users}, nil
}

// GetByMail is a method of the UsersLogic struct.
// It takes a context and a GetByMailReq as parameters.
// It returns a GetByMailReply and an error.
//
// The method first calls the FindByMail method of the userQ object, passing the context and the email from the request.
// If there is an error in this process, it returns the error.
// Then it converts the user repository to a different format using the covertUser function.
// Finally, it returns a GetByMailReply containing the user repository.
func (l *UsersLogic) GetByMail(ctx context.Context, req *user.GetByMailReq) (*user.GetByMailReply, error) {
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid email")
	}
	one, err := l.userQ.FindByMail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	return &user.GetByMailReply{User: covertUser(one)}, nil
}
