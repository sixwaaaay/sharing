package logic

import (
	"bytelite/app/user/dal"
	"bytelite/app/user/types"
	"bytelite/common/errorx"
	"bytelite/common/secu"
	"bytelite/service"
	"context"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic func(req *types.UserReq) (resp *types.UserResp, err error)

var NewRegisterLogic = newRegisterLogic

func newRegisterLogic(ctx context.Context, appCtx *service.AppContext) RegisterLogic {
	logger := logx.WithContext(ctx)
	return func(req *types.UserReq) (resp *types.UserResp, err error) {
		if len(req.Username) > 32 || len(req.Password) > 32 || len(req.Username) < 1 || len(req.Password) < 6 {
			return nil, errorx.NewDefaultError("username or password is too long")
		}
		if strings.ContainsAny(req.Username, " ") || strings.ContainsAny(req.Password, " ") {
			return nil, errorx.NewDefaultError("username or password contains space")
		}
		var user dal.User
		user.Username = req.Username
		user.Salt = secu.GenSalt(6)
		user.Password, user.Salt = secu.GenHashedPassAndSalt(req.Password)
		// 写入数据库
		ret, err := appCtx.UsersModel.Insert(ctx, &user)
		if err != nil {
			logger.Error(err)
			return nil, errorx.NewDefaultError("register failed")
		}
		userId, err := ret.LastInsertId()
		if err != nil {
			logger.Error(err)
			return nil, errorx.NewDefaultError("register failed")
		}

		// 生成token
		token, err := appCtx.JWTSigner.GenerateToken(userId, 60*60*24*7)
		if err != nil {
			logger.Error(err)
			return nil, errorx.NewDefaultError("register failed")
		}

		return &types.UserResp{
			StatusCode: 0,
			StatusMsg:  "",
			Token:      token,
			UserID:     userId,
		}, nil
	}
}
