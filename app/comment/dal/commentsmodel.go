package dal

import (
	"bytelite/common/errorx"
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CommentsModel = (*customCommentsModel)(nil)

type (
	// CommentsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentsModel.
	CommentsModel interface {
		commentsModel
		// FindCommentsByVideoID 获取对应视频的评论
		FindCommentsByVideoID(ctx context.Context, videoID int64) ([]*Comments, error)
		// DeleteUserComment 删除用户的评论
		DeleteUserComment(ctx context.Context, userID, commentID int64) error
	}

	customCommentsModel struct {
		*defaultCommentsModel
	}
)

func (m customCommentsModel) DeleteUserComment(ctx context.Context, userID, commentID int64) error {
	query, args, err := squirrel.Delete(m.table).Where(
		squirrel.Eq{
			"user_id": userID,
			"id":      commentID,
		}).ToSql()
	if err != nil {
		return err
	}
	result, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, query, args...)
	})
	if err != nil {
		return err
	}
	if rows, err := result.RowsAffected(); err != nil {
		return err
	} else if rows == 0 {
		return errorx.NewDefaultError("comment not found")
	}
	return nil
}

func (m customCommentsModel) FindCommentsByVideoID(ctx context.Context, videoID int64) ([]*Comments, error) {
	query, args, err := squirrel.Select("*").From(m.table).Where(
		squirrel.Eq{
			"video_id": videoID,
		}).OrderBy("created_at desc").ToSql()
	if err != nil {
		return nil, err
	}
	var comments []*Comments
	err = m.QueryRowsNoCacheCtx(ctx, &comments, query, args...)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// NewCommentsModel returns a model for the database table.
func NewCommentsModel(conn sqlx.SqlConn, c cache.CacheConf) CommentsModel {
	return &customCommentsModel{
		defaultCommentsModel: newCommentsModel(conn, c),
	}
}
