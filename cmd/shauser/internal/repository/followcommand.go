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

	"github.com/sony/sonyflake"
	"gorm.io/gorm"
)

type FollowCommand interface {
	Insert(ctx context.Context, f *Follow) error
	Delete(ctx context.Context, userid, followTo int64) error
}

// followCommand is a struct that represents a command to follow a user.
// It contains a database connection and a unique ID generator.
type followCommand struct {
	// db is a pointer to a gorm.DB object, which represents a database connection.
	db *gorm.DB
	// uniqueID is a pointer to a sonyflake.Sonyflake object, which is used to generate unique IDs.
	uniqueID *sonyflake.Sonyflake
}

// NewFollowCommand is a constructor function for followCommand.
// It takes a database connection as an argument and returns a pointer to a followCommand object.
func NewFollowCommand(db *gorm.DB) FollowCommand {
	return &followCommand{
		// Initialize the db field with the provided database connection.
		db: db,
		// Initialize the uniqueID field with a new Sonyflake object.
		uniqueID: sonyflake.NewSonyflake(sonyflake.Settings{}),
	}
}

// Insert is a method of followCommand that inserts a new follow relationship into the database.
// It takes a context and a pointer to a Follow object as arguments.
// It returns an error if the insertion fails.
func (c *followCommand) Insert(ctx context.Context, f *Follow) (err error) {
	// Generate a new unique ID.
	id, err := c.uniqueID.NextID()
	if err != nil {
		return err
	}
	// Set the ID of the Follow object to the generated ID.
	f.ID = int64(id)

	// Start a new database transaction.
	return c.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Insert the Follow object into the database.
		if err := tx.Create(f).Error; err != nil {
			return err
		}
		// Increment the follow count of the user who is following.
		res := tx.Raw("UPDATE users SET following = following + 1 WHERE id = ?", f.UserID)
		if res.Error != nil {
			return res.Error
		}
		// Increment the follower count of the user who is being followed.
		res = tx.Raw("UPDATE users SET followers = followers + 1 WHERE id = ?", f.Target)
		return res.Error
	})
}

// Delete is a method of followCommand that deletes a follow relationship from the database.
// It takes a context and two user IDs as arguments: the ID of the user who is following and the ID of the user who is being followed.
// It returns an error if the deletion fails.
func (c *followCommand) Delete(ctx context.Context, userid, followTo int64) error {
	// Start a new database transaction.
	err := c.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Delete the follow relationship from the database.
		res := tx.Exec("DELETE FROM follows WHERE user_id = ? AND target = ?", userid, followTo)
		if res.Error != nil {
			return res.Error
		} else if res.RowsAffected == 0 { // If no rows were affected by the deletion, return an error.
			return gorm.ErrRecordNotFound
		}
		// Decrement the follow count of the user who was following.
		res = tx.Exec("UPDATE users SET following = following - 1 WHERE id = ?", userid)
		if res.Error != nil {
			return res.Error
		}
		// Decrement the follower count of the user who was being followed.
		res = tx.Exec("UPDATE users SET followers = followers - 1 WHERE id = ?", followTo)
		return res.Error
	})
	return err
}
