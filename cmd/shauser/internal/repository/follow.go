/*
 * Copyright (c) 2023-2024 sixwaaaay.
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

	"gorm.io/gorm"
)

type FollowQuery interface {
	FindFollowExits(ctx context.Context, userid int64, followTo []int64) ([]int64, error)
	FindFollowing(ctx context.Context, userid int64, token int64, limit int) ([]int64, error)
	FindFollowers(ctx context.Context, userid int64, token int64, limit int) ([]int64, error)
}

var _ FollowQuery = (*followQuery)(nil)

// followQuery is the struct for follow query
type followQuery struct {
	db *gorm.DB
}

// NewFollowQuery creates a new followQuery
func NewFollowQuery(db *gorm.DB) FollowQuery {
	return &followQuery{
		db: db,
	}
}

// FindFollowExits query whether the user follow the followTo ids
// return the followTo ids that the user follow
func (c *followQuery) FindFollowExits(ctx context.Context, userid int64, followTo []int64) ([]int64, error) {
	var result []int64
	if userid == 0 || len(followTo) == 0 {
		return result, nil
	}
	session := c.db.WithContext(ctx)
	session.Table("follows").
		Where("target IN ?", followTo).
		Where("user_id = ?", userid).
		Pluck("target", &result)
	err := session.Error
	return result, err
}

// FindFollowing query the user follow list
func (c *followQuery) FindFollowing(ctx context.Context, userid int64, token int64, limit int) ([]int64, error) {
	var result []int64
	session := c.db.WithContext(ctx)
	tx := session.Raw("SELECT target FROM follows WHERE user_id = ? AND id < ? ORDER BY id desc LIMIT ?", userid, token, limit).Scan(&result)
	err := tx.Error
	return result, err
}

// FindFollowers query the user's follower list
func (c *followQuery) FindFollowers(ctx context.Context, userid int64, token int64, limit int) ([]int64, error) {
	var result []int64
	session := c.db.WithContext(ctx)
	tx := session.Raw("SELECT user_id FROM follows WHERE target = ? AND id < ? ORDER BY id desc LIMIT ?", userid, token, limit).Scan(&result)
	err := tx.Error
	return result, err
}
