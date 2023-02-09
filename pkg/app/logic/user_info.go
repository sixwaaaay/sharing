package logic

import (
	"context"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"
	"github.com/sixwaaaay/sharing/pkg/common/errorx"
	"github.com/sixwaaaay/sharing/pkg/common/middleware"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic func(req *types.UserInfoReq) (resp *types.UserInfoResp, err error)

var NewUserInfoLogic = newUserInfoLogic

func newUserInfoLogic(ctx context.Context, appCtx *service.AppContext) UserInfoLogic {
	logger := logx.WithContext(ctx)
	InfoQuery := func(selfId int64, userId int64) (user *types.User, err error) {
		return QueryUserInfo(ctx, appCtx, selfId, userId)
	}
	return func(req *types.UserInfoReq) (*types.UserInfoResp, error) {
		curUserId := ctx.Value(middleware.UserClaimsKey).(int64)

		info, err := InfoQuery(curUserId, req.UserID)
		if err != nil {
			logger.Errorf("query user info failed, err: %v", err)
			return nil, errorx.NewDefaultError("failed to query user info")
		}
		resp := &types.UserInfoResp{
			User:       info,
			StatusCode: 0,
		}
		return resp, nil
	}
}
