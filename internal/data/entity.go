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

type UserDetail struct {
	ID            int64   `gorm:"column:id;"`
	Username      string  `gorm:"column:username;"`
	Email         string  `gorm:"column:email;"`
	Phone         string  `gorm:"column:phone;"`
	AvatarURL     string  `gorm:"column:avatar_url;"`
	BgURL         string  `gorm:"column:bg_url;"`
	Bio           string  `gorm:"column:bio;"`
	LikesGiven    int     `gorm:"column:likes_given;"`
	LikesReceived int     `gorm:"column:likes_received;"`
	VideosPosted  int     `gorm:"column:videos_posted;"`
	Gender        int     `gorm:"column:gender"`
	DateOfBirth   *string `gorm:"column:date_of_birth"`
	Nationality   string  `gorm:"column:nationality"`
	Following     int     `gorm:"column:following"`
	Followers     int     `gorm:"column:followers"`
}

type User struct {
	ID            int64  `gorm:"column:id;"`
	Username      string `gorm:"column:username;"`
	AvatarURL     string `gorm:"column:avatar_url;"`
	BgURL         string `gorm:"column:bg_url;"`
	Bio           string `gorm:"column:bio;"`
	LikesGiven    int32  `gorm:"column:likes_given;"`
	LikesReceived int32  `gorm:"column:likes_received;"`
	VideosPosted  int32  `gorm:"column:videos_posted;"`
	Nationality   string `gorm:"column:nationality"`
	Following     int32  `gorm:"column:following"`
	Followers     int32  `gorm:"column:followers"`
}

type Account struct {
	ID       int64  `gorm:"column:id;"`
	Username string `gorm:"column:username;"`
	Password string `gorm:"column:password;"`
	Email    string `gorm:"column:email;"`
	Phone    string `gorm:"column:phone;"`
}

type Follow struct {
	ID        int64     `gorm:"column:id;"`
	UserID    int64     `gorm:"column:user_id;"`
	Target    int64     `gorm:"column:target;"`
	CreatedAt time.Time `gorm:"column:created_at;"`
}
