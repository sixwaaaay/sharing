package dal

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	relationsFieldNames          = builder.RawFieldNames(&Relations{})
	relationsRows                = strings.Join(relationsFieldNames, ",")
	relationsRowsExpectAutoSet   = strings.Join(stringx.Remove(relationsFieldNames, "`create_time`", "`update_time`"), ",")
	relationsRowsWithPlaceHolder = strings.Join(stringx.Remove(relationsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheRelationsIdPrefix = "cache:relations:id:"
)

type (
	relationsModel interface {
		Insert(ctx context.Context, data *Relations) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Relations, error)
		Update(ctx context.Context, data *Relations) error
		Delete(ctx context.Context, id int64) error
	}

	defaultRelationsModel struct {
		sqlc.CachedConn
		table string
	}

	Relations struct {
		Id       int64     `db:"id"`
		Follower int64     `db:"follower"`
		Followed int64     `db:"followed"`
		Status   int64     `db:"status"`
		UpdateAt time.Time `db:"update_at"`
	}
)

func newRelationsModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultRelationsModel {
	return &defaultRelationsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`relations`",
	}
}

func (m *defaultRelationsModel) Insert(ctx context.Context, data *Relations) (sql.Result, error) {
	relationsIdKey := fmt.Sprintf("%s%v", cacheRelationsIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, relationsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.Follower, data.Followed, data.Status, data.UpdateAt)
	}, relationsIdKey)
	return ret, err
}

func (m *defaultRelationsModel) FindOne(ctx context.Context, id int64) (*Relations, error) {
	relationsIdKey := fmt.Sprintf("%s%v", cacheRelationsIdPrefix, id)
	var resp Relations
	err := m.QueryRowCtx(ctx, &resp, relationsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", relationsRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRelationsModel) Update(ctx context.Context, data *Relations) error {
	relationsIdKey := fmt.Sprintf("%s%v", cacheRelationsIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, relationsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Follower, data.Followed, data.Status, data.UpdateAt, data.Id)
	}, relationsIdKey)
	return err
}

func (m *defaultRelationsModel) Delete(ctx context.Context, id int64) error {
	relationsIdKey := fmt.Sprintf("%s%v", cacheRelationsIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, relationsIdKey)
	return err
}

func (m *defaultRelationsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheRelationsIdPrefix, primary)
}

func (m *defaultRelationsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", relationsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultRelationsModel) tableName() string {
	return m.table
}
