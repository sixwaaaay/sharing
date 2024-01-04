/*
 * Copyright (c) 2024 sixwaaaay.
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
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sixwaaaay/shauser/internal/config"
)

func TestName(t *testing.T) {

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
	command := NewUserCommand(gormDB)

	t.Run("Create success", func(t *testing.T) {
		err := command.Insert(context.Background(), &Account{
			Email: time.Now().String() + "@test.com",
		})
		assertions.NoError(err)
	})

	t.Run("FindAccount success", func(t *testing.T) {
		err := command.FindAccount(context.Background(), &Account{
			Email: "2@x.com",
		})
		assertions.NoError(err)
	})

	t.Run("UpdateUser success", func(t *testing.T) {
		err := command.UpdateUser(context.Background(), &User{
			ID:        1,
			AvatarURL: "http://example.com/1.jpg",
		})
		assertions.NoError(err)
	})
}
