package logic

import (
	"context"
	"github.com/sixwaaaay/sharing/pkg/app/dal"
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/sixwaaaay/sharing/pkg/common/errorx"
	"time"
)

func FollowUser(ctx context.Context, appCtx *service.AppContext, selfId, userId, actionType int64) error {
	relations := &dal.Relations{
		Follower: selfId,
		Followed: userId,
		Status:   actionType,
		UpdateAt: time.Now(),
	}
	_, err := appCtx.RelationModel.Insert(ctx, relations)
	if err != nil {
		return errorx.NewDefaultError("unsupported operation")
	}
	return nil
}

func UnFollowUser(ctx context.Context, appCtx *service.AppContext, selfId, userId int64) error {
	err := appCtx.RelationModel.DeleteUserRelation(ctx, selfId, userId)
	if err != nil {
		return errorx.NewDefaultError("an unsupported operation")
	}
	return nil
}

func QueryFollowedUser(ctx context.Context, appCtx *service.AppContext, selfId int64, userIds []int64) ([]int64, error) {
	relation, err := appCtx.RelationModel.FindFollowRelation(ctx, selfId, userIds)
	if err != nil {
		return nil, errorx.NewDefaultError("query failed")
	}
	return relation, nil
}

func QueryIsFollowed(ctx context.Context, appCtx *service.AppContext, selfId, userId int64) (bool, error) {
	followed, err := appCtx.RelationModel.FindRelationEdge(ctx, selfId, userId)
	if err != nil {
		return false, errorx.NewDefaultError("query follow relation failed")
	}
	return followed, nil
}
