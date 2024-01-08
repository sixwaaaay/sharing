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
	"math"

	"github.com/sixwaaaay/token"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sixwaaaay/shauser/internal/config"
	"github.com/sixwaaaay/shauser/internal/repository"
	"github.com/sixwaaaay/shauser/user"
)

type FollowQueryLogic struct {
	conf     *config.Config
	followQ  repository.FollowQuery
	userQ    repository.UserQuery
	getUsers *UsersLogic
	logger   *zap.Logger
}

func NewFollowQueryLogic(conf *config.Config, followQ repository.FollowQuery, userQ repository.UserQuery, getUsers *UsersLogic, logger *zap.Logger) *FollowQueryLogic {
	return &FollowQueryLogic{conf: conf, followQ: followQ, userQ: userQ, getUsers: getUsers, logger: logger}
}

// GetFollowings retrieves the followings of a specific user.
// It takes a context and a FollowQueryReq as parameters.
// The FollowQueryReq contains the user ID, the limit of followings to retrieve, and a token for pagination.
// It returns a UsersPage containing the retrieved users and an error if any occurred.
func (l *FollowQueryLogic) GetFollowings(ctx context.Context, in *user.FollowQueryReq) (*user.UsersPage, error) {

	userID, _ := token.ClaimStrI64(ctx, token.ClaimID)

	if in.Limit == 0 || in.Limit > l.conf.MaxLimit {
		in.Limit = l.conf.DefaultLimit
	}

	if in.Page == 0 {
		in.Page = math.MaxInt64
	}

	following, err := l.userQ.FindFollowing(ctx, in.UserId)
	if err != nil {
		l.logger.Error("get following failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to find following for user %v: %v", in.UserId, err)
	}

	list, err := l.followQ.FindFollowing(ctx, in.UserId, in.Page, int(in.Limit))
	if err != nil {
		l.logger.Error("get following failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to find following for user %v: %v", in.UserId, err)
	}

	users, err := l.users(ctx, list, userID)
	if err != nil {
		l.logger.Error("get users failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to find following for user %v: %v", in.UserId, err)
	}

	nextToken := l.nextPage(list, in.Limit)

	return &user.UsersPage{
		Users:    users,
		NextPage: nextToken,
		AllCount: following,
	}, nil

}

// GetFollowers retrieves the followers of a specific user.
// It takes a context and a FollowQueryReq as parameters.
// The FollowQueryReq contains the user ID, the limit of followers to retrieve, and a token for pagination.
// It returns a UsersPage containing the retrieved users and an error if any occurred.
func (l *FollowQueryLogic) GetFollowers(ctx context.Context, in *user.FollowQueryReq) (*user.UsersPage, error) {
	userID, _ := token.ClaimStrI64(ctx, token.ClaimID)

	if in.Limit == 0 || in.Limit > l.conf.MaxLimit {
		in.Limit = l.conf.DefaultLimit
	}

	if in.Page == 0 {
		in.Page = math.MaxInt64
	}

	followers, err := l.userQ.FindFollowers(ctx, in.UserId)
	if err != nil {
		l.logger.Error("get followers failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to find followers for user %v: %v", in.UserId, err)
	}

	list, err := l.followQ.FindFollowers(ctx, in.UserId, in.Page, int(in.Limit))
	if err != nil {
		l.logger.Error("get followers failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to find followers for user %v: %v", in.UserId, err)
	}

	users, err := l.users(ctx, list, userID)
	if err != nil {
		l.logger.Error("get users failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to find followers for user %v: %v", in.UserId, err)
	}
	nextToken := l.nextPage(list, in.Limit)

	l.logger.Info("GetFollowers completed")
	return &user.UsersPage{
		Users:    users,
		NextPage: nextToken,
		AllCount: followers,
	}, nil
}

// users is a method of the FollowQueryLogic struct.
// It takes a context, a list of user IDs, and a subject ID as parameters.
// It calls the makeUsers function with the provided parameters and the user and follow queries of the FollowQueryLogic struct.
// It returns a slice of User pointers and an error if any occurred.
func (l *FollowQueryLogic) users(ctx context.Context, list []int64, subjectId int64) ([]*user.User, error) {
	return makeUsers(ctx, list, subjectId, l.userQ, l.followQ)
}

// nextPage is a method of the FollowQueryLogic struct.
// It takes a list of user IDs and a length as parameters.
// If the length of the list is equal to the provided length, it returns the last ID in the list.
// If the length of the list is not equal to the provided length, it returns 0.
func (l *FollowQueryLogic) nextPage(list []int64, length int32) int64 {
	if len(list) == int(length) { //	be the last id of the user list
		return list[len(list)-1]
	}
	return 0
}
