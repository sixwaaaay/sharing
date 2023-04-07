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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRelationFind(t *testing.T) {
	mock, gormDB := mockDB(t)
	model := NewFollowQuery(gormDB)
	assertions := assert.New(t)
	const FindFollowExits = "SELECT `target` FROM `follows` WHERE target IN (?,?) AND user_id = ?"
	t.Run("FindFollowExits success", func(t *testing.T) {
		// userid int64, videoID []int64
		mock.ExpectQuery(FindFollowExits).
			WithArgs(int64(1), int64(2), int64(1)).
			WillReturnRows(mock.NewRows([]string{"followed_id"}).AddRow(1))
		like, err := model.FindFollowExits(context.TODO(), 1, []int64{1, 2})
		assertions.NoError(err)
		assertions.ElementsMatch([]int64{1}, like)
	})
	const FindFollowing = "SELECT `target` FROM `follows` WHERE user_id = ? AND id > ? ORDER BY id desc LIMIT 1"
	t.Run("FindFollowing success", func(t *testing.T) {
		mock.ExpectQuery(FindFollowing).
			WithArgs(int64(1), int64(3)).
			WillReturnRows(mock.NewRows([]string{"follow_to"}).AddRow(1).AddRow(2).AddRow(3))
		like, err := model.FindFollowing(context.TODO(), 1, 3, 1)
		assertions.NoError(err)
		assertions.ElementsMatch([]int64{1, 2, 3}, like)
	})
	const FindFollowerFrom = "SELECT `user_id` FROM `follows` WHERE target = ? AND id > ? ORDER BY id desc LIMIT 1"
	t.Run("FindFollowers success", func(t *testing.T) {
		mock.ExpectQuery(FindFollowerFrom).
			WithArgs(int64(5), int64(1)).
			WillReturnRows(mock.NewRows([]string{"user_id"}).AddRow(1))
		like, err := model.FindFollowers(context.TODO(), 5, 1, 1)
		assertions.NoError(err)
		assertions.ElementsMatch([]int64{1}, like)
	})
	const FindFriends = "SELECT target FROM follows WHERE user_id = ? AND target IN (SELECT user_id FROM follows WHERE target = ?) AND id > ? ORDER BY id asc LIMIT ?"
	t.Run("FindFriends success", func(t *testing.T) {
		mock.ExpectQuery(FindFriends).
			WithArgs(int64(1), int64(1), int64(1), int64(1)).
			WillReturnRows(mock.NewRows([]string{"target"}).AddRow(1))
		like, err := model.FindFriends(context.TODO(), 1, 1, 1)
		assertions.NoError(err)
		assertions.ElementsMatch([]int64{1}, like)

	})
}
