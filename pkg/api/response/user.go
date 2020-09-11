package response

import (
	"sort"
	"wantum/pkg/domain/entity/user"
)

type UserResponse struct {
	ID       int    `json:"id"`
	AuthID   string `json:"auth_id"`
	UserName string `json:"user_name"`
	Mail     string `json:"mail"`

	// Profileデータ
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	Bio       string `json:"bio"`
	Gender    int    `json:"gender"`
	Phone     string `json:"phone"`
	Place     string `json:"place"`
	Birth     int64  `json:"birth"`
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

func ConvertToUsersResponse(userMap user.UserMap) UsersResponse {
	res := make(UsersResponse, 0, len(userMap))
	userIDs := userMap.Keys(userMap)
	for _, userID := range userIDs {
		res = append(res, ConvertToUserResponse(userMap[userID]))
	}
	// NOTE: user.idの昇順でレスポンスを並び替えています。
	sort.Slice(res, func(i, j int) bool {
		return res[i].ID < res[j].ID
	})
	return res
}

func ConvertToUserResponse(userData *user.User) *UserResponse {
	// nilチェック
	if userData == nil {
		return nil
	}
	if userData.Profile == nil {
		return &UserResponse{
			ID:       userData.ID,
			AuthID:   userData.AuthID,
			UserName: userData.UserName,
			Mail:     userData.Mail,
		}
	}

	return &UserResponse{
		ID:        userData.ID,
		AuthID:    userData.AuthID,
		UserName:  userData.UserName,
		Mail:      userData.Mail,
		Name:      userData.Profile.Name,
		Thumbnail: userData.Profile.Thumbnail,
		Bio:       userData.Profile.Bio,
		Gender:    userData.Profile.Gender,
		Phone:     userData.Profile.Phone,
		Place:     userData.Profile.Place,
		Birth:     userData.Profile.Birth.Unix(),
	}
}
