package covert

import "bytelite/common/cotypes"

func UserMap(users []cotypes.User) map[int64]cotypes.User {
	userMap := make(map[int64]cotypes.User, len(users))
	for _, u := range users {
		userMap[u.ID] = u
	}
	return userMap
}
