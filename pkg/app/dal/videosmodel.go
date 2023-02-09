package dal

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ VideosModel = (*customVideosModel)(nil)

type (
	// VideosModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideosModel.
	VideosModel interface {
		videosModel
		Delete(ctx context.Context, id int64) error

		FindMultiVideo(ctx context.Context, ids []int64) ([]*Videos, error)

		FindByTimestamp(ctx context.Context, timestamp int64) ([]*Videos, error)

		FindByUserID(ctx context.Context, userID int64) ([]*Videos, error)

		// UpdateFavoriteCount 更新视频的点赞数， +count or -count
		UpdateFavoriteCount(ctx context.Context, id int64, count int64) (sql.Result, error)

		// UpdateCommentCount 更新视频的评论数  +count or -count
		UpdateCommentCount(ctx context.Context, id int64, count int64) (sql.Result, error)
	}

	customVideosModel struct {
		*defaultVideosModel
	}
)

func (c customVideosModel) UpdateFavoriteCount(ctx context.Context, id int64, count int64) (sql.Result, error) {
	query, args, err := squirrel.Update("videos").
		Set("favorite_count", squirrel.Expr("favorite_count + ?", count)).
		Where("id = ?", id).ToSql()
	if err != nil {
		return nil, err
	}
	return c.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, query, args...)
	})

}

func (c customVideosModel) UpdateCommentCount(ctx context.Context, id int64, count int64) (sql.Result, error) {
	query, args, err := squirrel.Update("videos").
		Set("comment_count", squirrel.Expr("comment_count + ?", count)).
		Where("id = ?", id).ToSql()
	if err != nil {
		return nil, err
	}
	return c.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, query, args...)
	})
}

func (c customVideosModel) FindByTimestamp(ctx context.Context, timestamp int64) ([]*Videos, error) {
	query, args, err := squirrel.Select(videosFieldNames...).From(c.table).
		Where("created_at <= from_unixtime(?)", timestamp/1000).
		OrderBy("created_at DESC").
		Limit(30).
		ToSql()
	if err != nil {
		return nil, err
	}
	var videos []*Videos
	err = c.QueryRowsNoCacheCtx(ctx, &videos, query, args...)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (c customVideosModel) FindMultiVideo(ctx context.Context, ids []int64) ([]*Videos, error) {
	query, args, err := squirrel.Select(videosFieldNames...).From(c.table).Where(squirrel.Eq{"id": ids}).ToSql()
	if err != nil {
		return nil, err
	}
	var videos []*Videos
	err = c.QueryRowsNoCacheCtx(ctx, &videos, query, args...)
	if err != nil {
		return nil, err
	}
	return videos, nil

}

func (c customVideosModel) FindByUserID(ctx context.Context, userID int64) ([]*Videos, error) {
	query, args, err := squirrel.Select(videosFieldNames...).From(c.table).Where("user_id = ?", userID).ToSql()
	if err != nil {
		return nil, err
	}
	var videos []*Videos
	err = c.QueryRowsNoCacheCtx(ctx, &videos, query, args...)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

// NewVideosModel returns a model for the database table.
func NewVideosModel(conn sqlx.SqlConn, c cache.CacheConf) VideosModel {
	return &customVideosModel{
		defaultVideosModel: newVideosModel(conn, c),
	}
}
