package basic

import (
	"bytelite/service"
	"context"
)

func AddFavorite(ctx context.Context, appCtx *service.AppContext, selfId, videoId, actionType int64) error {

	err := appCtx.FavoriteModel.UpdateUserFavorite(ctx, selfId, videoId, actionType)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = UpdateVideoFavoriteCount(ctx, appCtx, videoId, 1)
		}
	}()
	return nil
}

func RemoveFavorite(ctx context.Context, appCtx *service.AppContext, selfId, videoId int64) error {
	err := appCtx.FavoriteModel.DeleteUserFavorite(ctx, selfId, videoId)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			return
		}
		_ = UpdateVideoFavoriteCount(ctx, appCtx, videoId, -1)
	}()
	return nil
}

// QueryIsFavorite 查询指定视频是否被收藏
func QueryIsFavorite(ctx context.Context, appCtx *service.AppContext, selfId, videoId int64) (bool, error) {
	return appCtx.FavoriteModel.IsFavorite(ctx, selfId, videoId)
}

// QueryVideoFavorites 查询指定视频列表中哪些收藏了
func QueryVideoFavorites(ctx context.Context, appCtx *service.AppContext, selfId int64, videoIds []int64) ([]int64, error) {
	return appCtx.FavoriteModel.QueryVideoFavorites(ctx, selfId, videoIds)
}
