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

type RegisterLogic struct {
	conf        *config.Config
	userCommand *data.UserCommand
}

type RegisterLogicOption struct {
	Config      *config.Config
	UserCommand *data.UserCommand
}

func NewRegisterLogic(opt RegisterLogicOption) *RegisterLogic {
	return &RegisterLogic{
		conf:        opt.Config,
		userCommand: opt.UserCommand,
	}
}

func (l *RegisterLogic) Register(ctx context.Context, in *user.RegisterRequest) (*user.RegisterReply, error) {
	if in.Name == "" || in.Password == "" || in.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	account := &data.Account{
		Username: in.Name,
		Email:    in.Email,
	}

	// 使用 bcrypt 加密密码
	pwd, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	account.Password = string(pwd)
	err = l.userCommand.Insert(ctx, account) // 尝试保存到数据库
	if err != nil {
		return nil, err
	}

	u := user.User{
		Id:   account.ID,
		Name: account.Username,
	}

	reply := &user.RegisterReply{
		User: &u,
	}

	return reply, nil
}
