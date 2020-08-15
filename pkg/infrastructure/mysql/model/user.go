package model

import (
	"time"
	"wantum/pkg/domain/entity"
)

type UserModel struct {
	ID        int
	AuthID    string
	UserName  string
	Mail      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	Profile   *ProfileModel
}

type UserModelSlice []*UserModel

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

func ConvertToUserEntity(userData *UserModel) *entity.User {
	return &entity.User{
		ID:        userData.ID,
		AuthID:    userData.AuthID,
		UserName:  userData.UserName,
		Mail:      userData.Mail,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
		DeletedAt: userData.DeletedAt,
		Profile: &entity.Profile{
			Name:      userData.Profile.Name,
			Thumbnail: userData.Profile.Thumbnail,
			Bio:       userData.Profile.Bio,
			Gender:    userData.Profile.Gender,
			Phone:     userData.Profile.Phone,
			Place:     userData.Profile.Place,
			Birth:     userData.Profile.Birth,
			CreatedAt: userData.Profile.CreatedAt,
			UpdatedAt: userData.Profile.UpdatedAt,
			DeletedAt: userData.Profile.DeletedAt,
		},
	}
}

func ConvertToUserSliceEntity(userSlice UserModelSlice) entity.UserSlice {
	res := make(entity.UserSlice, 0, len(userSlice))
	for _, userData := range userSlice {
		res = append(res, ConvertToUserEntity(userData))
	}
	return res
}
