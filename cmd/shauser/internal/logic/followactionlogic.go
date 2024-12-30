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

type FollowActionLogic struct {
	follow repository.FollowCommand
	logger *zap.Logger
}

func NewFollowActionLogic(follow repository.FollowCommand, logger *zap.Logger) *FollowActionLogic {
	return &FollowActionLogic{follow: follow, logger: logger}
}

// FollowAction is a method of the FollowActionLogic struct.
// It takes a context and a FollowActionRequest as parameters and returns a FollowActionReply and an error.
//
// The FollowActionRequest contains the UserID, SubjectID, and Action.
// UserID and SubjectID should not be 0 and should not be the same.
// Action should be either 1 or 2.
//
// If the Action is 1, a new Follow relationship is created where the UserID is the follower and the SubjectID is the target.
// The new Follow relationship is then inserted into the database.
// If there is an error during the insertion, it is returned.
//
// If the Action is 2, the Follow relationship where the UserID is the follower and the SubjectID is the target is deleted from the database.
// If there is an error during the deletion, it is returned.
//
// If the Action is neither 1 nor 2, an error with the message "invalid action" is returned.
func (l *FollowActionLogic) FollowAction(ctx context.Context, in *user.FollowActionRequest) (*user.FollowActionReply, error) {
	userID, ok := token.ClaimStrI64(ctx, token.ClaimID)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	if in.UserId == 0 || in.UserId == userID {
		return nil, status.Errorf(codes.InvalidArgument, "invalid arguments")
	}

	if in.Action == 1 {
		err := l.follow.Insert(ctx, &repository.Follow{UserID: userID, Target: in.UserId})

		if err != nil {
			l.logger.Error("insert follow failed", zap.Error(err))
			return nil, status.Errorf(codes.Internal, "failed to follow user %v", in.UserId)
		}

		return &user.FollowActionReply{}, nil

	} else if in.Action == 2 {
		err := l.follow.Delete(ctx, userID, in.UserId)

		if err != nil {
			l.logger.Error("delete follow failed", zap.Error(err))
			return nil, status.Errorf(codes.Internal, "failed to unfollow user %v", in.UserId)
		}

		return &user.FollowActionReply{}, nil
	}

	return nil, status.Error(codes.InvalidArgument, "invalid action")
}
