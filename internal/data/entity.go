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

import "time"

// UserDetail represents the detailed information of a user in the system.
type UserDetail struct {
	// ID is the unique identifier of the user.
	ID int64 `gorm:"column:id;"`
	// Username is the unique username of the user.
	Username string `gorm:"column:username;"`
	// Email is the email address of the user.
	Email string `gorm:"column:email;"`
	// Phone is the phone number of the user.
	Phone string `gorm:"column:phone;"`
	// AvatarURL is the URL of the user's avatar.
	AvatarURL string `gorm:"column:avatar_url;"`
	// BgURL is the URL of the user's background image.
	BgURL string `gorm:"column:bg_url;"`
	// Bio is a short biography written by the user.
	Bio string `gorm:"column:bio;"`
	// LikesGiven is the total number of likes given by the user to other users' posts.
	LikesGiven int `gorm:"column:likes_given;"`
	// LikesReceived is the total number of likes received by the user from other users' posts.
	LikesReceived int `gorm:"column:likes_received;"`
	// VideosPosted is the total number of videos posted by the user.
	VideosPosted int `gorm:"column:videos_posted;"`
	// Gender is the gender of the user. It's an integer where specific values represent specific genders.
	Gender int `gorm:"column:gender"`
	// DateOfBirth is the date of birth of the user.
	DateOfBirth *string `gorm:"column:date_of_birth"`
	// Nationality is the nationality of the user.
	Nationality string `gorm:"column:nationality"`
	// Following is the number of users that this user is following.
	Following int `gorm:"column:following"`
	// Followers is the number of users that are following this user.
	Followers int `gorm:"column:followers"`
}

type User struct {
	// ID is the unique identifier of the user.
	ID int64 `gorm:"column:id;"`
	// Username is the unique username of the user.
	Username string `gorm:"column:username;"`
	// Email is the email address of the user.
	Email string `gorm:"column:email;"`
	// Phone is the phone number of the user.
	Phone string `gorm:"column:phone;"`
	// AvatarURL is the URL of the user's avatar.
	AvatarURL string `gorm:"column:avatar_url;"`
	// BgURL is the URL of the user's background image.
	BgURL string `gorm:"column:bg_url;"`
	// Bio is a short biography written by the user.
	Bio string `gorm:"column:bio;"`
	// LikesGiven is the total number of likes given by the user to other users' posts.
	LikesGiven int32 `gorm:"column:likes_given;"`
	// LikesReceived is the total number of likes received by the user from other users' posts.
	LikesReceived int32 `gorm:"column:likes_received;"`
	// VideosPosted is the total number of videos posted by the user.
	VideosPosted int32 `gorm:"column:videos_posted;"`
	// Nationality is the nationality of the user.
	Nationality string `gorm:"column:nationality"`
	// Following is the number of users that this user is following.
	Following int32 `gorm:"column:following"`
	// Followers is the number of users that are following this user.
	Followers int32 `gorm:"column:followers"`
}

// Account represents the account information of a user in the system.
type Account struct {
	// ID is the unique identifier of the account.
	ID int64 `gorm:"column:id;"`
	// Username is the unique username of the account.
	Username string `gorm:"column:username;"`
	// Password is the password of the account.
	Password string `gorm:"column:password;"`
	// Email is the email address associated with the account.
	Email string `gorm:"column:email;"`
	// Phone is the phone number associated with the account.
	Phone string `gorm:"column:phone;"`
}

// Follow represents the follow relationship between users in the system.
type Follow struct {
	// ID is the unique identifier of the follow relationship.
	ID int64 `gorm:"column:id;"`
	// UserID is the identifier of the user who initiates the follow.
	UserID int64 `gorm:"column:user_id;"`
	// Target is the identifier of the user who is being followed.
	Target int64 `gorm:"column:target;"`
	// CreatedAt is the timestamp when the follow relationship was created.
	CreatedAt time.Time `gorm:"column:created_at;"`
}
