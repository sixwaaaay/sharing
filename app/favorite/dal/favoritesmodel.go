package dal

import (
	"bytelite/common/errorx"
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FavoritesModel = (*customFavoritesModel)(nil)

type (
	// FavoritesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFavoritesModel.
	FavoritesModel interface {
		favoritesModel
		FindByUserID(ctx context.Context, userID int64) ([]*Favorites, error)
		DeleteUserFavorite(ctx context.Context, selfId, videoId int64) error
		UpdateUserFavorite(ctx context.Context, selfId, videoId int64, actionType int64) error
		// IsFavorite 查询对某个视频是否点赞
		IsFavorite(ctx context.Context, selfId, videoId int64) (bool, error)
		QueryVideoFavorites(ctx context.Context, selfId int64, videoIds []int64) ([]int64, error)
	}

	customFavoritesModel struct {
		*defaultFavoritesModel
	}
)

func (c customFavoritesModel) IsFavorite(ctx context.Context, selfId, videoId int64) (bool, error) {
	query, args, err := squirrel.Select(favoritesFieldNames...).
		From(c.table).
		Where(squirrel.Eq{
			"user_id":  selfId,
			"video_id": videoId}).ToSql()
	if err != nil {
		return false, err
	}
	var favorites Favorites
	err = c.QueryRowNoCacheCtx(ctx, &favorites, query, args...)
	if err != nil {
		return false, err
	}
	return favorites.Id > 0, nil
}

func (c customFavoritesModel) UpdateUserFavorite(ctx context.Context, selfId, videoId, actionType int64) error {

	query, args, err := squirrel.Insert("favorites").
		Columns("user_id", "video_id", "action").
		Values(selfId, videoId, actionType).
		Suffix("ON DUPLICATE KEY UPDATE user_id = ?, video_id = ?", selfId, videoId).ToSql()
	if err != nil {
		return err
	}
	res, err := c.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, query, args...)
	})
	if err != nil {
		return err
	}
	if rowsAffected, err := res.RowsAffected(); err != nil {
		return err
	} else if rowsAffected == 0 {
		return errorx.NewDefaultError("no rows affected")
	}
	return nil
}

func (c customFavoritesModel) FindByUserID(ctx context.Context, userID int64) ([]*Favorites, error) {

	query, args, err := squirrel.Select(favoritesFieldNames...).From(c.table).Where("user_id = ?", userID).ToSql()
	if err != nil {
		return nil, err
	}
	var favorites []*Favorites
	err = c.QueryRowsNoCacheCtx(ctx, &favorites, query, args...)
	if err != nil {
		return nil, err
	}
	return favorites, nil
}

// NewFavoritesModel returns a model for the database table.
func NewFavoritesModel(conn sqlx.SqlConn, c cache.CacheConf) FavoritesModel {
	return &customFavoritesModel{
		defaultFavoritesModel: newFavoritesModel(conn, c),
	}
}

// DeleteUserFavorite 删除用户的点赞
func (c customFavoritesModel) DeleteUserFavorite(ctx context.Context, selfId, videoId int64) error {
	// 缓存使用的 key
	favoritesIdKey := fmt.Sprintf("%s%v", cacheFavoritesIdPrefix, videoId)
	// 缓存框架执行回调顺带删除缓存
	_, err := c.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		// 仅能删除自己的评论
		query, args, err := squirrel.Delete(c.table).Where(squirrel.Eq{
			"user_id":  selfId,
			"video_id": videoId,
		}).ToSql()
		if err != nil {
			return nil, err
		}
		return conn.ExecCtx(ctx, query, args...)
	}, favoritesIdKey)
	if err != nil {
		return err
	}
	return nil
}

func (c customFavoritesModel) QueryVideoFavorites(ctx context.Context, selfId int64, videoIds []int64) ([]int64, error) {
	if len(videoIds) == 0 {
		return nil, nil
	}
	query, args, err := squirrel.Select("video_id").
		From(c.table).
		Where(squirrel.Eq{
			"user_id":  selfId,
			"video_id": videoIds}).ToSql()
	if err != nil {
		return nil, err
	}
	var favorites []int64
	err = c.QueryRowsNoCacheCtx(ctx, &favorites, query, args...)
	if err != nil {
		return nil, err
	}
	return favorites, nil
}
