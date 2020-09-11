package user

import (
	"time"
	"wantum/pkg/domain/entity/userprofile"
)

type Entity struct {
	ID        int
	AuthID    string
	UserName  string
	Mail      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	Profile   *userprofile.Entity
}

type EntitySlice []*Entity

type EntityMap map[int]*Entity

func (umap *EntityMap) Keys(userMap EntityMap) []int {
	keys := make([]int, 0, len(userMap))
	for key := range userMap {
		keys = append(keys, key)
	}
	return keys
}
