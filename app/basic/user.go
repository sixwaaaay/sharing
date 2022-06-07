package basic

import (
	"bytelite/app/user/dal"
	"bytelite/common/cotypes"
	"bytelite/common/covert"
	"bytelite/common/errorx"
	"bytelite/service"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

func QueryUserInfo(ctx context.Context, appCtx *service.AppContext, selfId int64, userId int64) (*cotypes.User, error) {
	user, err := appCtx.UsersModel.FindUserInfo(ctx, userId)
	if err != nil {
		logx.Infof("query user info failed, err: %v", err)
		return nil, errorx.NewDefaultError("find user failed")
	}
	userInfo := toUser(user)
	if userId == selfId {
		return userInfo, nil
	}
	// 查询关注关系
	exists, err := QueryIsFollowed(ctx, appCtx, selfId, userId)
	if err != nil {
		logx.Infof("query is followed failed, err: %v", err)
		return nil, errorx.NewDefaultError("find relation failed")
	}
	userInfo.IsFollow = exists
	return userInfo, nil
}

func QueryMultiUserInfo(ctx context.Context, appCtx *service.AppContext,
	selfId int64, UserIds []int64) ([]cotypes.User, error) {

	users, err := appCtx.UsersModel.FindMultiUserInfo(ctx, UserIds)
	if err != nil {
		return nil, errorx.NewDefaultError("find user failed")
	}
	userInfos := toUsers(users)
	// 查询关注关系

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

func UpdateFollowRelationCount(ctx context.Context, appCtx *service.AppContext, selfId int64, userId int64) error {
	return appCtx.UsersModel.UpdateWhenFollow(ctx, selfId, userId)
}

func UpdateUnFollowRelationCount(ctx context.Context, appCtx *service.AppContext, selfId int64, userId int64) error {
	return appCtx.UsersModel.UpdateWhenUnFollow(ctx, selfId, userId)
}

func toUsers(users []*dal.UserInfo) []cotypes.User {
	var userInfos []cotypes.User
	for _, user := range users {
		userInfos = append(userInfos, *toUser(user))
	}
	return userInfos
}
func userIds(users []*dal.UserInfo) []int64 {
	var ids []int64
	for _, user := range users {
		ids = append(ids, user.Id)
	}
	return ids
}

func toUser(user *dal.UserInfo) *cotypes.User {
	var userInfo cotypes.User
	userInfo.ID = user.Id
	userInfo.Name = user.Username
	userInfo.FollowCount = user.FollowedCount
	userInfo.FollowerCount = user.FollowerCount
	return &userInfo
}
