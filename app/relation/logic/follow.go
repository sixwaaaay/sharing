package logic

import (
	"bytelite/app/basic"
	"bytelite/app/relation/types"
	"bytelite/common/errorx"
	"bytelite/common/middleware"
	"bytelite/service"
	"context"
)

type FollowLogic func(req *types.FollowActionReq) (*types.FollowActionResp, error)

var NewFollowActionLogic = newFollowActionLogic

func newFollowActionLogic(ctx context.Context, appCtx *service.AppContext) FollowLogic {
	return func(req *types.FollowActionReq) (*types.FollowActionResp, error) {
		selfId, _ := ctx.Value(middleware.UserClaimsKey).(int64)
		var resp types.FollowActionResp
		switch req.ActionType {
		case 1:
			err := basic.FollowUser(ctx, appCtx, selfId, req.ToUserId, req.ActionType)
			if err != nil {
				return nil, errorx.NewDefaultError("follow user failed")
			}
			resp.StatusCode = 0
			return &resp, nil
		case 2:
			err := basic.UnFollowUser(ctx, appCtx, selfId, req.ToUserId)
			if err != nil {
				return nil, errorx.NewDefaultError("unfollow user failed")
			}
		}
		// actionType not 1 or 2
		return nil, errorx.NewDefaultError("invalid action type")
	}
}
