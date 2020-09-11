package userprofile

import "time"

type Entity struct {
	UserID    int
	Name      string
	Thumbnail string
	Bio       string
	Gender    int
	Phone     string
	Place     string
	Birth     *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type EntitySlice []*Entity
