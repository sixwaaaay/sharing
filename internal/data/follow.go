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

	"gorm.io/gorm"
)

// FollowQuery is the struct for follow query
type FollowQuery struct {
	db *gorm.DB
}

// NewFollowQuery creates a new FollowQuery
func NewFollowQuery(db *gorm.DB) *FollowQuery {
	return &FollowQuery{
		db: db,
	}
}

// FindFollowExits query whether the user follow the followTo ids
// return the followTo ids that the user follow
func (c *FollowQuery) FindFollowExits(ctx context.Context, userid int64, followTo []int64) ([]int64, error) {
	var result []int64
	session := c.db.WithContext(ctx)
	session.Table("follows").
		Where("target IN ?", followTo).
		Where("user_id = ?", userid).
		Pluck("target", &result)
	err := session.Error
	return result, err
}

// FindFollowing query the user follow list
func (c *FollowQuery) FindFollowing(ctx context.Context, userid int64, token int64, limit int) ([]int64, error) {
	var result []int64
	session := c.db.WithContext(ctx)
	session.Table("follows").
		Where("user_id = ?", userid).
		Where("id > ?", token).
		Order("id desc").
		Limit(limit).
		Pluck("target", &result)
	err := session.Error
	return result, err
}

// FindFollowers query the user's follower list
func (c *FollowQuery) FindFollowers(ctx context.Context, userid int64, token int64, limit int) ([]int64, error) {
	var result []int64
	session := c.db.WithContext(ctx)
	session.Table("follows").
		Where("target = ?", userid).
		Where("id > ?", token).
		Order("id desc").
		Limit(limit).
		Pluck("user_id", &result)
	err := session.Error
	return result, err
}

// FindFriends query the user's friend list
func (c *FollowQuery) FindFriends(ctx context.Context, userid int64, token int64, limit int) ([]int64, error) {
	var result []int64
	session := c.db.WithContext(ctx)
	tx := session.Raw("SELECT target FROM follows WHERE user_id = ? AND target IN (SELECT user_id FROM follows WHERE target = ?) AND id > ? ORDER BY id asc LIMIT ?", userid, userid, token, limit).Scan(&result)
	err := tx.Error
	return result, err
}
