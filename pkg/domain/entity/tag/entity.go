package tag

import "time"

type Entity struct {
	ID        int
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type EntitySlice []*Entity
