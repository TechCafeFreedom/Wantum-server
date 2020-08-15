package response

import (
	"wantum/pkg/domain/entity"
)

type UserResponse struct {
	ID       int          `json:"id"`
	AuthID   string       `json:"auth_id"`
	UserName string       `json:"user_name"`
	Mail     string       `json:"mail"`
	Profile  *UserProfile `json:"profile"`
}

type UsersResponse []*UserResponse

type UserProfile struct {
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	Bio       string `json:"bio"`
	Gender    int    `json:"gender"`
	Phone     string `json:"phone"`
	Place     string `json:"place"`
	Birth     string `json:"birth"`
}

func ConvertToUsersResponse(userSlice entity.UserSlice) UsersResponse {
	res := make(UsersResponse, 0, len(userSlice))
	for _, userData := range userSlice {
		res = append(res, ConvertToUserResponse(userData))
	}
	return res
}

func ConvertToUserResponse(userData *entity.User) *UserResponse {
	// nilチェック
	if userData == nil {
		return nil
	}

	return &UserResponse{
		ID:       userData.ID,
		AuthID:   userData.AuthID,
		UserName: userData.UserName,
		Mail:     userData.Mail,
		Profile: &UserProfile{
			Name:      userData.Profile.Name,
			Thumbnail: userData.Profile.Thumbnail,
			Bio:       userData.Profile.Bio,
			Gender:    userData.Profile.Gender,
			Phone:     userData.Profile.Phone,
			Place:     userData.Profile.Place,
			Birth:     userData.Profile.Birth,
		},
	}
}
