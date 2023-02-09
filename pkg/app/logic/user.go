package logic

import (
	"context"
	"github.com/sixwaaaay/sharing/common/covert"
	"github.com/sixwaaaay/sharing/common/errorx"
	"github.com/sixwaaaay/sharing/common/itertool"
	"github.com/sixwaaaay/sharing/pkg/app/dal"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/app/types"
)

func QueryUserInfo(ctx context.Context, appCtx *service.AppContext, selfId int64, userId int64) (*types.User, error) {
	user, err := appCtx.UsersModel.FindUserInfo(ctx, userId)
	if err != nil {
		return nil, errorx.NewDefaultError("find user failed")
	}
	userInfo := toUser(user)
	if userId == selfId {
		return userInfo, nil
	}
	// 查询关注关系
	userInfo.IsFollow, err = QueryIsFollowed(ctx, appCtx, selfId, userId)
	if err != nil {
		return nil, errorx.NewDefaultError("find relation failed")
	}
	return userInfo, nil
}

func QueryMultiUserInfo(ctx context.Context, appCtx *service.AppContext,
	selfId int64, UserIds []int64) ([]types.User, error) {
	users, err := appCtx.UsersModel.FindMultiUserInfo(ctx, UserIds)
	if err != nil {
		return nil, errorx.NewDefaultError("find user failed")
	}
	userInfos := toUsers(users) // 查询关注关系
	followedList, err := QueryFollowedUser(ctx, appCtx, selfId, userIds(users))
	if err != nil {
		return nil, errorx.NewDefaultError("find relation failed")
	}
	followedMap := covert.Int64SliceToMap(followedList)
	for _, user := range userInfos {
		if _, ok := followedMap[user.ID]; ok {
			user.IsFollow = true
		}
	}
	return userInfos, nil
}

func toUsers(users []*dal.UserInfo) []types.User {
	return itertool.Reduce(users, func(agg []types.User, it *dal.UserInfo, _ int) []types.User {
		return append(agg, *toUser(it))
	}, []types.User{})
}

func userIds(users []*dal.UserInfo) []int64 {
	return itertool.Reduce(users, func(ids []int64, it *dal.UserInfo, _ int) []int64 {
		return append(ids, it.Id)
	}, []int64{})
}

func toUser(user *dal.UserInfo) *types.User {
	var userInfo types.User
	userInfo.ID = user.Id
	userInfo.Name = user.Username
	userInfo.FollowCount = user.FollowedCount
	userInfo.FollowerCount = user.FollowerCount
	return &userInfo
}
