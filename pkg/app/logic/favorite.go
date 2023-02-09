package logic

import (
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"
	"github.com/sixwaaaay/sharing/pkg/common/errorx"
	"github.com/sixwaaaay/sharing/pkg/common/middleware"

	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic func(req *types.FavoriteReq) (resp *types.FavoriteResp, err error)

var NewFavoriteLogic = newFavoriteLogic

func newFavoriteLogic(ctx context.Context, appCtx *service.AppContext) FavoriteActionLogic {
	return func(req *types.FavoriteReq) (resp *types.FavoriteResp, err error) {
		logger := logx.WithContext(ctx)
		selfId, _ := ctx.Value(middleware.UserClaimsKey).(int64)
		switch req.ActionType {
		case 1:
			err = AddFavorite(ctx, appCtx, selfId, req.VideoId, req.ActionType)
			if err != nil {
				logger.Errorf("add favorite failed, err: %v", err)
				return nil, errorx.NewDefaultError("add favorite failed")
			}
		case 2:
			err = RemoveFavorite(ctx, appCtx, selfId, req.VideoId)
			if err != nil {
				logger.Errorf("remove favorite failed, err: %v", err)
				return nil, errorx.NewDefaultError("remove favorite failed")
			}
		default:
			return nil, errorx.NewDefaultError("invalid action type")
		}
		return
	}
}
