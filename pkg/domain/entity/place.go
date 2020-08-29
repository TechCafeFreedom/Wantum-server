package entity

import "time"

type Place struct {
	ID        int
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

// TODO: 使わなさそう
type PlaceSlice []*Place
