package user

import (
	"testing"
	"wantum/pkg/domain/entity"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/user/mock_user"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	userID    = 1
	authID    = "authID"
	userName  = "userName"
	mail      = "test@test.com"
	name      = "name"
	thumbnail = "thumbnail"
	bio       = "bio"
	gender    = 1
	phone     = "000-0000-0000"
	place     = "place"
	birth     = "1998-05-03"
)

func TestService_CreateNewUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := repository.NewMockMasterTx()

	userRepository := mock_user.NewMockRepository(ctrl)
	userEntity := &entity.User{
		AuthID:   authID,
		UserName: userName,
		Mail:     mail,
		Profile: &entity.Profile{
			Name:      name,
			Thumbnail: thumbnail,
			Bio:       bio,
			Gender:    gender,
			Phone:     phone,
			Place:     place,
			Birth:     birth,
		},
	}
	userRepository.EXPECT().InsertUser(masterTx, userEntity).Return(userEntity, nil).Times(1)

	service := New(userRepository)
	createdUser, err := service.CreateNewUser(masterTx, authID, userName, mail, name, thumbnail, bio, phone, place, birth, gender)

	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
}

func TestService_GetByPK(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := repository.NewMockMasterTx()

	existedUser := &entity.User{
		ID:       userID,
		AuthID:   authID,
		UserName: userName,
		Mail:     mail,
		Profile: &entity.Profile{
			Name:      name,
			Thumbnail: thumbnail,
			Bio:       bio,
			Gender:    gender,
			Phone:     phone,
			Place:     place,
			Birth:     birth,
		},
	}

	userRepository := mock_user.NewMockRepository(ctrl)
	userRepository.EXPECT().SelectByPK(ctx, masterTx, userID).Return(existedUser, nil).Times(1)

	service := New(userRepository)
	users, err := service.GetByPK(ctx, masterTx, userID)

	assert.NoError(t, err)
	assert.NotNil(t, users)
}

func TestService_SelectAll(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := repository.NewMockMasterTx()

	existedUsers := entity.UserSlice{
		{
			ID:       userID,
			AuthID:   authID,
			UserName: userName,
			Mail:     mail,
			Profile: &entity.Profile{
				Name:      name,
				Thumbnail: thumbnail,
				Bio:       bio,
				Gender:    gender,
				Phone:     phone,
				Place:     place,
				Birth:     birth,
			},
		},
	}

	userRepository := mock_user.NewMockRepository(ctrl)
	userRepository.EXPECT().SelectAll(ctx, masterTx).Return(existedUsers, nil).Times(1)

	service := New(userRepository)
	users, err := service.GetAll(ctx, masterTx)

	assert.NoError(t, err)
	assert.NotNil(t, users)
}
