package entity

import "time"

type User struct {
	ID        int
	AuthID    string
	UserName  string
	Mail      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	Profile   *Profile
}

type UserSlice []*User

type UserMap map[int]*User

func (umap *UserMap) Keys(userMap UserMap) []int {
	keys := make([]int, 0, len(userMap))
	for key := range userMap {
		keys = append(keys, key)
	}
	return keys
}