package dal

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/builder"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel

		// FindUserInfo 仅查询用户信息
		FindUserInfo(ctx context.Context, id int64) (*UserInfo, error)

		// FindMultiUserInfo  查询多个用户
		FindMultiUserInfo(ctx context.Context, ids []int64) ([]*UserInfo, error)

		UpdateWhenUnFollow(ctx context.Context, selfId int64, userId int64) error
		UpdateWhenFollow(ctx context.Context, selfId int64, userId int64) error
	}

	customUsersModel struct {
		*defaultUsersModel
	}
	UserInfo struct {
		Id            int64  `db:"id"`
		Username      string `db:"username"`
		FollowedCount int64  `db:"followed_count"`
		FollowerCount int64  `db:"follower_count"`
	}
)

var userInfoFields = builder.RawFieldNames(&UserInfo{})

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn, c),
	}
}

func (m *customUsersModel) FindMultiUserInfo(ctx context.Context, ids []int64) ([]*UserInfo, error) {

	query, args, err := squirrel.Select(userInfoFields...).From(m.table).Where(squirrel.Eq{"id": ids}).ToSql()
	if err != nil {
		return nil, err
	}
	var user []*UserInfo
	err = m.QueryRowsNoCacheCtx(ctx, &user, query, args...)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *customUsersModel) FindUserInfo(ctx context.Context, id int64) (*UserInfo, error) {
	logger := logx.WithContext(ctx)
	query, args, err := squirrel.Select(userInfoFields...).From(m.table).Where(squirrel.Eq{"id": id}).ToSql()

	if err != nil {
		logger.Errorf("生成sql语句失败: %+v", err)
		return nil, err
	}
	var user UserInfo
	err = m.QueryRowNoCacheCtx(ctx, &user, query, args...)
	if err != nil {
		logger.Errorf("查询失败: %+v", err)
		return nil, err
	}
	return &user, nil
}

func (m *customUsersModel) UpdateWhenFollow(ctx context.Context, selfId int64, userId int64) error {
	err := m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		// update target user's follower count
		query, args, err := squirrel.Update(m.table).Set("follower_count", squirrel.Expr("follower_count + 1")).Where(squirrel.Eq{"id": userId}).ToSql()
		if err != nil {
			return err
		}
		_, err = session.ExecCtx(ctx, query, args...)
		if err != nil {
			return err
		}
		// update self followed count
		query, args, err = squirrel.Update(m.table).Set("followed_count", squirrel.Expr("followed_count + 1")).Where(squirrel.Eq{"id": selfId}).ToSql()
		if err != nil {
			return err
		}
		_, err = session.ExecCtx(ctx, query, args...)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (m *customUsersModel) UpdateWhenUnFollow(ctx context.Context, selfId int64, userId int64) error {
	err := m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		// update target user's follower count
		query, args, err := squirrel.Update(m.table).Set("follower_count", squirrel.Expr("follower_count - 1")).Where(squirrel.Eq{"id": userId}).ToSql()
		if err != nil {
			return err
		}
		_, err = session.Exec(query, args...)
		if err != nil {
			return err
		}
		// update self followed count
		query, args, err = squirrel.Update(m.table).
			Set("followed_count", squirrel.Expr("followed_count - 1")).
			Where(squirrel.Eq{"id": selfId}).ToSql()
		if err != nil {
			return err
		}
		_, err = session.Exec(query, args...)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
