package wishboard

import "time"

type Entity struct {
	ID                 int
	Title              string
	BackgroundImageURL string
	InviteURL          string
	UserID             int
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
	DeletedAt          *time.Time
	// WishCategories
}

type EntitySlice []*Entity
