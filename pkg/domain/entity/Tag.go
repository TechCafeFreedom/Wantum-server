package entity

import "time"

type Tag struct {
	ID        int
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type TagSlice []*Tag
