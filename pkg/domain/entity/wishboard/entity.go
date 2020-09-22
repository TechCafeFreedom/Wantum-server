package wishboard

import "time"

type Entity struct {
	ID                 int
	Title              string
	BackgroundImageUrl string
	InviteUrl          string
	UserID             int
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *time.Time
}

type EntitySlice []*Entity
