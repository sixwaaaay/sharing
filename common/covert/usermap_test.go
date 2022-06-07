package covert

import (
	"bytelite/common/cotypes"
	"strconv"
	"testing"
)

func TestUserMap(t *testing.T) {
	var userSlice []cotypes.User
	for i := 0; i < 10; i++ {
		userSlice = append(userSlice, cotypes.User{
			ID:   int64(i),
			Name: "name" + strconv.Itoa(i),
		})
	}
	userMap := UserMap(userSlice)
	for _, v := range userSlice {
		if _, ok := userMap[v.ID]; !ok {
			t.Error("userMap[v.ID] not found")
		}
	}
}
