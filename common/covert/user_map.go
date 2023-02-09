package covert

import (
	"github.com/sixwaaaay/sharing/pkg/app/types"
)

func UserMap(users []types.User) map[int64]types.User {
	userMap := make(map[int64]types.User, len(users))
	for _, u := range users {
		userMap[u.ID] = u
	}
	return userMap
}
