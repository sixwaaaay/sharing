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
	"github.com/sony/sonyflake"
	"gorm.io/gorm"
)

// UserCommand is the implementation of dal.userCommand
type UserCommand struct {
	db       *gorm.DB
	uniqueID *sonyflake.Sonyflake
}

// NewUserCommand creates a new user command model
func NewUserCommand(db *gorm.DB) *UserCommand {
	return &UserCommand{
		db:       db,
		uniqueID: sonyflake.NewSonyflake(sonyflake.Settings{}),
	}
}

// Insert insert a user
func (c *UserCommand) Insert(ctx context.Context, u *Account) error {
	uid, err := c.uniqueID.NextID()
	if err != nil {
		return err
	}
	u.ID = int64(uid)
	session := c.db.WithContext(ctx)
	err = session.Table("users").Create(u).Error
	return err
}

// FindAccount find an account
// currently only support find by email
func (c *UserCommand) FindAccount(ctx context.Context, u *Account) error {
	session := c.db.WithContext(ctx)
	return session.Table("users").Where("email = ?", u.Email).First(u).Error
}
