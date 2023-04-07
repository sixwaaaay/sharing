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
	"testing"
)

func TestFollowExec(t *testing.T) {
	mock, gormDB := mockDB(t)
	model := NewFollowCommand(gormDB)
	assertions := assert.New(t)
	const InsertFollow = "INSERT INTO `follows` (`user_id`,`target`,`create_at`,`id`) VALUES (?,?,?,?)"
	follow := &Follow{
		UserID: 1,
		Target: 1,
	}
	t.Run("Insert success", func(t *testing.T) {
		mock.ExpectExec(InsertFollow).
			WithArgs(follow.UserID, follow.Target, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
		err := model.Insert(context.TODO(), follow)
		assertions.NoError(err)
		assertions.NotZero(follow.CreatedAt)
		assertions.NotZero(follow.ID)
	})
	const DeleteFollow = "DELETE FROM `follows` WHERE user_id = ? AND target = ?"
	t.Run("Delete success", func(t *testing.T) {
		mock.ExpectExec(DeleteFollow).
			WithArgs(follow.UserID, follow.Target).
			WillReturnResult(sqlmock.NewResult(1, 1))
		err := model.Delete(context.TODO(), follow.UserID, follow.Target)
		assertions.NoError(err)
	})
}
