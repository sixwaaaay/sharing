/*
 * Copyright (c) 2023 sixwaaaay.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package repository

import (
	"context"
	"math"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/sixwaaaay/shauser/internal/config"
)

const dsn = "MYSQL_DSN"

type Ca struct {
	Enabled bool   `yaml:"enabled"`
	TTL     string `yaml:"ttl"`
}

func TestUserFind(t *testing.T) {
	assertions := assert.New(t)
	DSN := os.Getenv(dsn)

	c := config.Config{
		Cache: struct {
			Enabled bool   `yaml:"enabled"`
			TTL     string `yaml:"ttl"`
		}(Ca{Enabled: false}),
		MySQL: struct {
			DSN      string
			Replicas []string
		}{DSN: DSN},
	}

	gormDB, err := NewDB(&c)
	assertions.NoError(err)
	query := NewUserQuery(gormDB, &c, zap.L())
	t.Run("FindOne success", func(t *testing.T) {
		one, err := query.FindOne(context.Background(), 1)
		assertions.NoError(err)
		assertions.NotNil(one)
	})
	t.Run("FindOne error", func(t *testing.T) {
		_, err := query.FindOne(context.Background(), 0)
		assertions.Error(err)
	})
	t.Run("FindMany success", func(t *testing.T) {

		many, err := query.FindMany(context.Background(), []int64{1, 2})
		assertions.NoError(err)
		assertions.NotNil(many)
	})

	t.Run("FindMany nothing", func(t *testing.T) {
		many, err := query.FindMany(context.Background(), []int64{math.MaxInt64})
		assertions.NoError(err)
		assertions.NotNil(many)
	})

	t.Run("FindByEmail success", func(t *testing.T) {
		one, err := query.FindByMail(context.Background(), "1@x.com")
		assertions.NoError(err)
		assertions.NotNil(one)
	})

	t.Run("QueryFollowing success", func(t *testing.T) {
		one, err := query.FindFollowing(context.Background(), 1)
		assertions.NoError(err)
		assertions.NotNil(one)
	})

	t.Run("QueryFollowers success", func(t *testing.T) {
		one, err := query.FindFollowers(context.Background(), 1)
		assertions.NoError(err)
		assertions.NotNil(one)
	})

	c.Cache.Enabled = true
	query = NewUserQuery(gormDB, &c, zap.L())
	t.Run("FindOne success", func(t *testing.T) {
		one, err := query.FindOne(context.Background(), 1)
		assertions.NoError(err)
		assertions.NotNil(one)
	})

	t.Run("FindMany success", func(t *testing.T) {

		many, err := query.FindMany(context.Background(), []int64{1, 2})
		assertions.NoError(err)
		assertions.NotNil(many)
	})

	user := &User{
		ID:        1111,
		AvatarURL: "https://www.baidu.com",
		Username:  "test",
		BgURL:     "https://www.baidu.com",
	}
	tx := gormDB.WithContext(context.Background()).Table("users").Where("id = ?", 1111).Updates(user)
	//	generate sql
	s := tx.Statement.SQL.String()
	t.Log(s)
	assertions.NoError(tx.Error)

	c.MySQL.Replicas = []string{DSN}
	gormDB, err = NewDB(&c)
	assertions.NoError(err)
}
