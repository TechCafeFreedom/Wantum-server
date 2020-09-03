package entity

import "time"

type WishCard struct {
	ID          int
	UserID      int
	Activity    string
	Description string
	Date        *time.Time
	DoneAt      *time.Time
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
	Place       *Place
	Tags        TagSlice
}

type WishCardSlice []*WishCard
