package logic

import (
	"context"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sixwaaaay/shauser/internal/data"

	"github.com/sixwaaaay/shauser/internal/config"
	"github.com/sixwaaaay/shauser/user"
)

type LoginLogic struct {
	conf        *config.Config
	userCommand *data.UserCommand
}

func NewLoginLogic(conf *config.Config, userCommand *data.UserCommand) *LoginLogic {
	return &LoginLogic{conf: conf, userCommand: userCommand}
}

func (l *LoginLogic) Login(ctx context.Context, in *user.LoginRequest) (*user.LoginReply, error) {
	if in.Password == "" || in.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}
	account := &data.Account{
		Email: in.Email,
	}
	err := l.userCommand.FindAccount(ctx, account)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(in.Password)) // bcrypt 验证密码
	if err != nil {
		return nil, err
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
