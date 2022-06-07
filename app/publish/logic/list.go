package logic

import (
	"bytelite/app/basic"
	"bytelite/app/publish/types"
	"bytelite/common/errorx"
	"bytelite/common/middleware"
	"bytelite/service"
	"context"
)

type PubListLogic func(req *types.PubListReq) (*types.PubListResp, error)

var NewPublishListLogic = newPubListLogic

func newPubListLogic(ctx context.Context, appCtx *service.AppContext) PubListLogic {
	return func(req *types.PubListReq) (*types.PubListResp, error) {
		selfId, _ := ctx.Value(middleware.UserClaimsKey).(int64)
		videos, err := basic.QueryUserVideo(ctx, appCtx, selfId, req.UserId)
		if err != nil {
			return nil, errorx.NewDefaultError("没有找到视频")
		}
		return &types.PubListResp{
			StatusCode: 0,
			VideoList:  videos,
		}, nil
	}
}
