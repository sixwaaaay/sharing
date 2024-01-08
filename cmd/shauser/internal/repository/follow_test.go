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
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sixwaaaay/shauser/internal/config"
)

func TestRelationFind(t *testing.T) {

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

	model := NewFollowQuery(gormDB)

	t.Run("FindFollowExits", func(t *testing.T) {
		like, err := model.FindFollowExits(context.TODO(), 1, []int64{})
		assertions.NoError(err)
		assertions.Empty(like)
	})

	t.Run("FindFollowExits success", func(t *testing.T) {
		_, err := model.FindFollowExits(context.TODO(), 1, []int64{1, 2})
		assertions.NoError(err)
	})

	t.Run("FindFollowing success", func(t *testing.T) {
		_, err := model.FindFollowing(context.TODO(), 1, 3, 1)
		assertions.NoError(err)
	})

	t.Run("FindFollowers success", func(t *testing.T) {
		_, err := model.FindFollowers(context.TODO(), 5, 1, 1)
		assertions.NoError(err)
	})

	command := NewFollowCommand(gormDB)

	t.Run("command", func(t *testing.T) {
		err := command.Insert(context.Background(), &Follow{
			UserID: 1,
			Target: 2,
		})
		assertions.NoError(err)

		err = command.Insert(context.Background(), &Follow{
			UserID: 1,
			Target: 2,
		})
		assertions.Error(err)

		err = command.Delete(context.Background(), 1, 2)
		assertions.NoError(err)

		err = command.Delete(context.Background(), 1, 2)
		assertions.Error(err)
	})
}
