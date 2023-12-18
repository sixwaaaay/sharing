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

package logic

import (
	"context"

	"github.com/sixwaaaay/shauser/internal/data"
	"github.com/sixwaaaay/shauser/user"
)

// makeUsers is a function that takes a context, a slice of user IDs, a subject ID, a UserQuery pointer, and a FollowQuery pointer as parameters.
// It finds many users with the provided user IDs using the FindMany method of the UserQuery.
// It finds if the subject follows the users with the provided user IDs using the FindFollowExits method of the FollowQuery.
// It composes the users with the results of the FindMany and FindFollowExits methods using the composeUsers function.
// It returns a slice of User pointers and an error if any occurred.
func makeUsers(ctx context.Context, UserIds []int64, subjectId int64, userQ *data.UserQuery, followQ *data.FollowQuery) ([]*user.User, error) {
	many, err := userQ.FindMany(ctx, UserIds)
	if err != nil {
		return nil, err
	}

	list, err := followQ.FindFollowExits(ctx, subjectId, UserIds)
	if err != nil {
		return nil, err
	}

	return composeUsers(many, list), nil
}

// composeUsers is a function that takes a slice of User pointers and a slice of user IDs as parameters.
// It converts the user IDs to a map using the idsToMap function.
// It converts each User in the slice to a User pointer using the covertUser function and checks if the user is followed by the subject.
// It returns a slice of User pointers.
func composeUsers(many []data.User, list []int64) []*user.User {
	users := make([]*user.User, 0, len(many))
	m := idsToMap(list)
	for _, u := range many {
		t := covertUser(&u)
		_, t.IsFollow = m[u.ID]
		users = append(users, t)
	}
	return users
}

// covertUser is a function that takes a User pointer as a parameter.
// It creates a new User pointer with the same fields as the provided User.
// It returns the new User pointer.
func covertUser(one *data.User) *user.User {
	u := &user.User{
		Id:            one.ID,
		Name:          one.Username,
		AvatarUrl:     one.AvatarURL,
		BgUrl:         one.BgURL,
		Bio:           one.Bio,
		LikesGiven:    one.LikesGiven,
		LikesReceived: one.LikesReceived,
		VideosPosted:  one.VideosPosted,
		Following:     one.Following,
		Followers:     one.Followers,
	}
	return u
}

// idsToMap is a function that takes a slice of user IDs as a parameter.
// It creates a map with the user IDs as keys and empty structs as values.
// It returns the map.
func idsToMap(ids []int64) map[int64]struct{} {
	m := make(map[int64]struct{}, len(ids))
	for _, id := range ids {
		m[id] = struct{}{}
	}
	return m
}
