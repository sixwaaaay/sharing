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

package data

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestUserFind(t *testing.T) {
	assertions := assert.New(t)
	mock, gormDB := mockDB(t)
	model := NewUserQuery(gormDB)
	const findOne = "SELECT `users`.`id`,`users`.`username`,`users`.`avatar_url`,`users`.`bg_url`,`users`.`bio`,`users`.`likes_given`,`users`.`likes_received`,`users`.`videos_posted`,`users`.`nationality`,`users`.`following`,`users`.`followers` " +
		"FROM `users` WHERE `users`.`id` = ? LIMIT 1"
	t.Run("FindOne success", func(t *testing.T) {
		mock.ExpectQuery(findOne).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "username"}).AddRow(1, "test"))
		user, err := model.FindOne(context.TODO(), 1)
		assertions.NoError(err)
		assertions.Equal("test", user.Username)
	})
	t.Run("FindOne error", func(t *testing.T) {
		mock.ExpectQuery(findOne).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "username"}))
		user, err := model.FindOne(context.TODO(), 1)
		// method will return gorm.ErrRecordNotFound if no record found
		assertions.Error(err)
		assertions.Equal(gorm.ErrRecordNotFound, err)
		assertions.Nil(user)
	})
	const findManyUser = "SELECT `users`.`id`,`users`.`username`,`users`.`avatar_url`,`users`.`bg_url`,`users`.`bio`,`users`.`likes_given`,`users`.`likes_received`,`users`.`videos_posted`,`users`.`nationality`,`users`.`following`,`users`.`followers` FROM `users` WHERE id IN (?,?)"
	t.Run("FindMany success", func(t *testing.T) {
		mock.ExpectQuery(findManyUser).
			WithArgs(1, 2).
			WillReturnRows(sqlmock.NewRows([]string{"id", "username"}).AddRow(1, "test1").AddRow(2, "test2"))
		users, err := model.FindMany(context.TODO(), []int64{1, 2})
		assertions.NoError(err)
		assertions.Equal(2, len(users))
		assertions.Equal("test1", users[0].Username)
		assertions.Equal("test2", users[1].Username)
	})

	t.Run("FindMany nothing", func(t *testing.T) {
		mock.ExpectQuery(findManyUser).
			WithArgs(1, 2).
			WillReturnRows(sqlmock.NewRows([]string{"id", "username"}))
		users, err := model.FindMany(context.TODO(), []int64{1, 2})
		assertions.NoError(err)
		assertions.Len(users, 0)
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

}

func mockDB(t *testing.T) (sqlmock.Sqlmock, *gorm.DB) {
	db, mock, err := sqlmock.New(
		sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual),
	)
	assert.NoError(t, err)
	opts := &gorm.Config{
		QueryFields:            true,
		SkipDefaultTransaction: true,
	}
	d := mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	})
	gormDB, err := gorm.Open(d, opts)
	assert.NoError(t, err)
	return mock, gormDB
}
