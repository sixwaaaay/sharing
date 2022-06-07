// StatusCode generated by goctl. DO NOT EDIT!

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
	favoritesFieldNames          = builder.RawFieldNames(&Favorites{})
	favoritesRows                = strings.Join(favoritesFieldNames, ",")
	favoritesRowsExpectAutoSet   = strings.Join(stringx.Remove(favoritesFieldNames, "`create_time`", "`update_time`"), ",")
	favoritesRowsWithPlaceHolder = strings.Join(stringx.Remove(favoritesFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheFavoritesIdPrefix = "cache:favorites:id:"
)

type (
	favoritesModel interface {
		Insert(ctx context.Context, data *Favorites) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Favorites, error)
		Update(ctx context.Context, data *Favorites) error
		Delete(ctx context.Context, id int64) error
	}

	defaultFavoritesModel struct {
		sqlc.CachedConn
		table string
	}

	Favorites struct {
		Id       int64     `db:"id"`
		UserId   int64     `db:"user_id"`
		VideoId  int64     `db:"video_id"`
		Action   int64     `db:"action"`
		UpdateAt time.Time `db:"update_at"`
	}
)

func newFavoritesModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultFavoritesModel {
	return &defaultFavoritesModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`favorites`",
	}
}

func (m *defaultFavoritesModel) Insert(ctx context.Context, data *Favorites) (sql.Result, error) {
	favoritesIdKey := fmt.Sprintf("%s%v", cacheFavoritesIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, favoritesRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.UserId, data.VideoId, data.Action, data.UpdateAt)
	}, favoritesIdKey)
	return ret, err
}

func (m *defaultFavoritesModel) FindOne(ctx context.Context, id int64) (*Favorites, error) {
	favoritesIdKey := fmt.Sprintf("%s%v", cacheFavoritesIdPrefix, id)
	var resp Favorites
	err := m.QueryRowCtx(ctx, &resp, favoritesIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", favoritesRows, m.table)
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

func (m *defaultFavoritesModel) Update(ctx context.Context, data *Favorites) error {
	favoritesIdKey := fmt.Sprintf("%s%v", cacheFavoritesIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, favoritesRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.VideoId, data.Action, data.UpdateAt, data.Id)
	}, favoritesIdKey)
	return err
}

func (m *defaultFavoritesModel) Delete(ctx context.Context, id int64) error {
	favoritesIdKey := fmt.Sprintf("%s%v", cacheFavoritesIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, favoritesIdKey)
	return err
}

func (m *defaultFavoritesModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheFavoritesIdPrefix, primary)
}

func (m *defaultFavoritesModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", favoritesRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultFavoritesModel) tableName() string {
	return m.table
}
