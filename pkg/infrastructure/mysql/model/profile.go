package model

import (
	"time"
	"wantum/pkg/domain/entity"
)

type ProfileModel struct {
	ID        int
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

type ProfileModelSlice []*ProfileModel

func ConvertToProfileEntity(profileData *ProfileModel) *entity.Profile {
	if profileData == nil {
		return nil
	}
	return &entity.Profile{
		UserID:    profileData.UserID,
		Name:      profileData.Name,
		Thumbnail: profileData.Thumbnail,
		Bio:       profileData.Bio,
		Gender:    profileData.Gender,
		Phone:     profileData.Phone,
		Place:     profileData.Place,
		Birth:     profileData.Birth,
		CreatedAt: profileData.CreatedAt,
		UpdatedAt: profileData.UpdatedAt,
		DeletedAt: profileData.DeletedAt,
	}
}

func ConvertToProfileSliceEntity(profileSlice ProfileModelSlice) entity.ProfileSlice {
	res := make(entity.ProfileSlice, 0, len(profileSlice))
	for _, profileData := range profileSlice {
		res = append(res, ConvertToProfileEntity(profileData))
	}
	return res
}
