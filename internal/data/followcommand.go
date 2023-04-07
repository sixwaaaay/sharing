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

// FollowCommand is the implementation of dal.FollowCommand
type FollowCommand struct {
	db       *gorm.DB
	uniqueID *sonyflake.Sonyflake
}

// NewFollowCommand creates a new comment relation model
func NewFollowCommand(db *gorm.DB) *FollowCommand {
	return &FollowCommand{
		db:       db,
		uniqueID: sonyflake.NewSonyflake(sonyflake.Settings{}),
	}
}

// Insert create a relation record
func (c *FollowCommand) Insert(ctx context.Context, f *Follow) error {
	id, err := c.uniqueID.NextID()
	if err != nil {
		return err
	}
	f.ID = int64(id)
	session := c.db.WithContext(ctx)
	err = session.Create(f).Error
	return err
}

// Delete a relation record by userid and followTo
func (c *FollowCommand) Delete(ctx context.Context, userid, followTo int64) error {
	session := c.db.WithContext(ctx)
	res := session.
		Where("user_id = ?", userid).
		Where("target = ?", followTo).
		Delete(&Follow{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
