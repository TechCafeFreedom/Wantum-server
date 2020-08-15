package entity

import "time"

type User struct {
	ID        int
	AuthID    string
	UserName  string
	Mail      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	Profile   *Profile
}

type UserSlice []*User

type Profile struct {
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
