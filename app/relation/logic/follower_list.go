package logic

import (
	"bytelite/app/basic"
	"bytelite/app/relation/types"
	"bytelite/common/cotypes"
	"bytelite/common/errorx"
	"bytelite/service"
	"context"
	"github.com/jinzhu/copier"
)

type FollowerListLogic func(req *types.RelationReq) (*types.FollowerListResp, error)

var NewFollowerListLogic = newFollowerListLogic

func newFollowerListLogic(ctx context.Context, appCtx *service.AppContext) FollowerListLogic {
	queryUserList := func(curUser int64, userId []int64) ([]cotypes.User, error) {
		info, err := basic.QueryMultiUserInfo(ctx, appCtx, curUser, userId)
		if err != nil {
			return nil, err
		}
		var userList []cotypes.User
		for _, v := range info {
			var user cotypes.User
			copier.Copy(&user, v)
			userList = append(userList, user)
		}
		return userList, nil
	}
	return func(req *types.RelationReq) (*types.FollowerListResp, error) {
		followed, err := appCtx.RelationModel.FindFollower(ctx, req.UserId)
		if err != nil {
			return nil, errorx.NewDefaultError("find followed error")
		}
		list, err := queryUserList(req.UserId, followed)
		if err != nil {
			return nil, errorx.NewDefaultError("find user info error")
		}
		return &types.FollowerListResp{
			StatusCode: 0,
			UserList:   list,
		}, nil
	}
}
