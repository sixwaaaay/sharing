package basic

import (
	"bytelite/app/relation/dal"
	"bytelite/service"
	"context"
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
	defer func() {
		_ = UpdateFollowRelationCount(ctx, appCtx, selfId, userId)
	}()

	if err != nil {
		return err
	}
	return nil
}

func UnFollowUser(ctx context.Context, appCtx *service.AppContext, selfId, userId int64) error {
	err := appCtx.RelationModel.DeleteUserRelation(ctx, selfId, userId)
	if err != nil {
		return err
	}
	defer func() {
		_ = UpdateUnFollowRelationCount(ctx, appCtx, selfId, userId)
	}()
	return nil
}

func QueryFollowedUser(ctx context.Context, appCtx *service.AppContext, selfId int64, userIds []int64) ([]int64, error) {
	return appCtx.RelationModel.FindFollowRelation(ctx, selfId, userIds)
}

func QueryIsFollowed(ctx context.Context, appCtx *service.AppContext, selfId, userId int64) (bool, error) {
	return appCtx.RelationModel.FindRelationEdge(ctx, selfId, userId)
}
