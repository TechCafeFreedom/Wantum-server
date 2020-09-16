package wishcard

import (
	"time"
	"wantum/pkg/domain/entity/place"
	"wantum/pkg/domain/entity/tag"
	"wantum/pkg/domain/entity/user"
)

type Entity struct {
	ID          int
	Author      *user.Entity
	Activity    string
	Description string
	Date        *time.Time
	DoneAt      *time.Time
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
	Place       *place.Entity
	Tags        tag.EntitySlice
}

type EntitySlice []*Entity
