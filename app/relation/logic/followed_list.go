package logic

import (
	"bytelite/app/basic"
	"bytelite/app/relation/types"
	"bytelite/common/cotypes"
	"bytelite/common/errorx"
	"bytelite/service"
	"context"
)

type FollowedListLogic func(req *types.RelationReq) (*types.FollowListResp, error)

var NewFollowedListLogic = newFollowedListLogic

func newFollowedListLogic(ctx context.Context, appCtx *service.AppContext) FollowedListLogic {
	queryUserList := func(curUser int64, userId []int64) ([]cotypes.User, error) {
		info, err := basic.QueryMultiUserInfo(ctx, appCtx, curUser, userId)
		if err != nil {
			return nil, err
		}
		return info, nil
	}

	return func(req *types.RelationReq) (*types.FollowListResp, error) {

		followed, err := appCtx.RelationModel.FindFollowed(ctx, req.UserId)
		if err != nil {
			return nil, errorx.NewDefaultError("find followed error")
		}
		list, err := queryUserList(req.UserId, followed)
		if err != nil {
			return nil, errorx.NewDefaultError("find user info error")
		}
		return &types.FollowListResp{
			StatusCode: 0,
			UserList:   list,
		}, nil
	}
}
