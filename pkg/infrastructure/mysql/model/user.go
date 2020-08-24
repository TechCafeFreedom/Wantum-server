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
}

type UserModelSlice []*UserModel

func ConvertToUserEntity(userData *UserModel) *entity.User {
	if userData == nil {
		return nil
	}
	return &entity.User{
		ID:        userData.ID,
		AuthID:    userData.AuthID,
		UserName:  userData.UserName,
		Mail:      userData.Mail,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
		DeletedAt: userData.DeletedAt,
	}
}

func ConvertToUserSliceEntity(userSlice UserModelSlice) entity.UserSlice {
	res := make(entity.UserSlice, 0, len(userSlice))
	for _, userData := range userSlice {
		res = append(res, ConvertToUserEntity(userData))
	}
	return res
}
