package logic

import (
	"bytelite/app/basic"
	"bytelite/app/favorite/types"
	"bytelite/common/errorx"
	"bytelite/service"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic func(req *types.FavoriteListReq) (resp *types.FavoriteListResp, err error)

var NewFavoriteListLogic = newFavoriteListLogic

func newFavoriteListLogic(ctx context.Context, appCtx *service.AppContext) FavoriteListLogic {
	return func(req *types.FavoriteListReq) (resp *types.FavoriteListResp, err error) {
		logger := logx.WithContext(ctx)
		favoriteList, err := appCtx.FavoriteModel.FindByUserID(ctx, req.UserId)

		if err != nil {
			logger.Errorf("not found, error: %v", err)
			return nil, errorx.NewDefaultError("not found")
		}

		ids := make([]int64, 0, len(favoriteList))
		for _, v := range favoriteList {
			ids = append(ids, v.VideoId)
		}
		videos, err := basic.QueryMultiVideo(ctx, appCtx, req.UserId, ids)

		return &types.FavoriteListResp{
			StatusCode: 0,
			VideoList:  videos,
		}, nil
	}
}
