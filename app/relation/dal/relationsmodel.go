package dal

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RelationsModel = (*customRelationsModel)(nil)

type (
	// RelationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRelationsModel.
	RelationsModel interface {
		relationsModel
		// FindRelationEdge 查询 a 是否关注了 b, 使用 (a,b) 查询
		FindRelationEdge(ctx context.Context, selfId, userId int64) (bool, error)
		// FindFollower 查询 a 的粉丝，返回的是 a 的粉丝的 id
		FindFollower(ctx context.Context, a int64) ([]int64, error)
		// FindFollowed 查询 a 的关注，返回的是 a 的关注的 id
		FindFollowed(ctx context.Context, a int64) ([]int64, error)

		// DeleteUserRelation 删除用户关系
		DeleteUserRelation(ctx context.Context, selfId, userId int64) error

		// FindFollowRelation 查询指定列表中 selfId 的关注列表
		FindFollowRelation(ctx context.Context, selfId int64, userIds []int64) ([]int64, error)
	}

	customRelationsModel struct {
		*defaultRelationsModel
	}
)

// NewRelationsModel returns a model for the database table.
func NewRelationsModel(conn sqlx.SqlConn, c cache.CacheConf) RelationsModel {
	return &customRelationsModel{
		defaultRelationsModel: newRelationsModel(conn, c),
	}
}

func (m customRelationsModel) DeleteUserRelation(ctx context.Context, selfId, userId int64) error {
	query, args, err := squirrel.Delete(m.table).
		Where(squirrel.Eq{"Follower": selfId, "Followed": userId}).ToSql()
	if err != nil {
		return err
	}
	_, err = m.ExecNoCacheCtx(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (m customRelationsModel) FindFollower(ctx context.Context, userId int64) ([]int64, error) {
	query, args, err := squirrel.Select("followed").
		From(m.table).
		Where(squirrel.Eq{"Followed": userId}).ToSql()
	if err != nil {
		return nil, err
	}
	var ids []int64
	err = m.QueryRowsNoCacheCtx(ctx, &ids, query, args...)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func (m customRelationsModel) FindFollowed(ctx context.Context, userId int64) ([]int64, error) {
	query, args, err := squirrel.Select("follower").
		From(m.table).
		Where(squirrel.Eq{"Follower": userId}).ToSql()
	if err != nil {
		return nil, err
	}
	var ids []int64
	err = m.QueryRowsNoCacheCtx(ctx, &ids, query, args...)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func (m customRelationsModel) FindFollowRelation(ctx context.Context, selfId int64, userIds []int64) ([]int64, error) {

	logger := logx.WithContext(ctx)
	query, args, err := squirrel.Select("followed").From(m.table).
		Where(squirrel.Eq{"Follower": selfId, "Followed": userIds}).ToSql()
	if err != nil {
		logger.Errorf("find follow relation failed, %s", err)
		return nil, err
	}
	var ids []int64
	err = m.QueryRowsNoCacheCtx(ctx, &ids, query, args...)
	if err != nil {
		logger.Errorf("find follow relation failed, %s", err)
		return nil, err
	}
	return ids, nil
}

func (m customRelationsModel) FindRelationEdge(ctx context.Context, userId, toUserId int64) (bool, error) {
	query, args, err := squirrel.Select("id").
		From(m.table).
		Where(squirrel.Eq{"Follower": userId, "Followed": toUserId}).ToSql()
	if err != nil {
		return false, err
	}
	var id int64
	err = m.QueryRowNoCacheCtx(ctx, &id, query, args...)
	if err != nil {
		return false, err
	}
	return id > 0, nil
}
