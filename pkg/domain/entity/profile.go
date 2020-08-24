package entity

import "time"

type Profile struct {
	UserID    int
	Name      string
	Thumbnail string
	Bio       string
	Gender    int
	Phone     string
	Place     string
	Birth     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type ProfileSlice []*Profile
