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
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sixwaaaay/cq"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/sixwaaaay/shauser/internal/config"
)

// UserQuery  is the struct for user query
type UserQuery struct {
	db     *gorm.DB
	cache  *cq.Cache[User]
	enable bool
}

// NewUserQuery creates a new UserQuery
func NewUserQuery(db *gorm.DB, conf *config.Config, logger *zap.Logger) *UserQuery {
	u := &UserQuery{
		db: db,
	}

	duration, err := time.ParseDuration(conf.Cache.TTL)
	if err != nil {
		logger.Warn("failed to parse cache ttl, use default ttl", zap.Error(err))
		duration = time.Minute
	}
	client := redis.NewUniversalClient(&conf.Redis)
	id := func(u *User) int64 {
		return u.ID
	}
	wrapRepo := cq.NewWrapRepo[User](u.findOne, u.findMany)
	cache := cq.NewCache[User](wrapRepo, id, "users", client, duration)
	u.cache = cache
	u.enable = conf.Cache.Enabled
	return u
}

// FindOne find one user by id
func (c *UserQuery) FindOne(ctx context.Context, id int64) (*User, error) {
	if c.enable {
		return c.cache.FindOne(ctx, id)
	}
	return c.findOne(ctx, id)
}

// FindMany find many users by ids
// Even if there is no any user matched, it will return an empty slice
func (c *UserQuery) FindMany(ctx context.Context, ids []int64) ([]User, error) {
	if c.enable {
		return c.cache.FindMany(ctx, ids)
	}
	return c.findMany(ctx, ids)
}

func (c *UserQuery) findOne(ctx context.Context, id int64) (*User, error) {
	var u User
	err := c.db.WithContext(ctx).Take(&u, id).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (c *UserQuery) findMany(ctx context.Context, ids []int64) ([]User, error) {
	var users []User
	err := c.db.WithContext(ctx).Where("id IN ?", ids).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// FindFollowing find following count by user id
func (c *UserQuery) FindFollowing(ctx context.Context, id int64) (int32, error) {
	var count int32
	err := c.db.WithContext(ctx).Raw("SELECT following FROM users WHERE id = ?", id).Scan(&count).Error
	return count, err
}

// FindFollowers find followers count by user id
func (c *UserQuery) FindFollowers(ctx context.Context, id int64) (int32, error) {
	var count int32
	err := c.db.WithContext(ctx).Raw("SELECT followers FROM users WHERE id = ?", id).Scan(&count).Error
	return count, err
}
