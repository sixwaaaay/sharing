/*
 * Copyright (c) 2024 sixwaaaay.
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
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sixwaaaay/shauser/internal/repository"

	"github.com/sixwaaaay/shauser/user"
)

// SignLogic is a struct that contains the configuration, user command and logger.
type SignLogic struct {
	userCommand repository.UserCommand // User command to be executed
	logger      *zap.Logger            // Logger to log information and errors
}

// NewSignLogic is a constructor for the SignLogic struct.
// It takes a configuration and a user command as parameters and returns a pointer to a SignLogic struct.
func NewSignLogic(userCommand repository.UserCommand) *SignLogic {
	return &SignLogic{userCommand: userCommand}
}

// Register is a method of the SignLogic struct.
// It takes a context and a RegisterRequest as parameters.
// It checks if the name, password and email in the RegisterRequest are not empty.
// If they are, it returns an error.
// If they are not, it creates a new account with the provided name and email,
// generates a hashed password from the provided password,
// inserts the new account into the database,
// and returns a RegisterReply containing the new user's ID and username.
func (l *SignLogic) Register(ctx context.Context, in *user.RegisterRequest) (*user.RegisterReply, error) {
	if in.Name == "" || in.Password == "" || in.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	account := &repository.Account{
		Username: in.Name,
		Email:    in.Email,
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		l.logger.Error("generate password failed", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}
	account.Password = string(pwd)
	err = l.userCommand.Insert(ctx, account)
	if err != nil {
		l.logger.Error("insert account failed", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}

	u := user.User{Id: account.ID, Name: account.Username}

	reply := &user.RegisterReply{User: &u}

	return reply, nil
}

// Login is a method of the SignLogic struct.
// It takes a context and a LoginRequest as parameters.
// It checks if the password and email in the LoginRequest are not empty.
// If they are, it returns an error.
// If they are not, it finds the account with the provided email,
// compares the provided password with the account's password,
// and returns a LoginReply containing the user's ID and username.
func (l *SignLogic) Login(ctx context.Context, in *user.LoginRequest) (*user.LoginReply, error) {
	if in.Password == "" || in.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}
	account := &repository.Account{
		Email: in.Email,
	}
	err := l.userCommand.FindAccount(ctx, account)
	if err != nil {
		l.logger.Error("find account failed", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(in.Password))
	if err != nil {
		l.logger.Error("compare password failed", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}

	u := user.User{
		Id:   account.ID,
		Name: account.Username,
	}

	reply := &user.LoginReply{
		User: &u,
	}

	return reply, nil
}
