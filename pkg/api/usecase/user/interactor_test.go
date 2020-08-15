package user

import (
	"testing"
	"wantum/pkg/domain/entity"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/service/user/mock_user"

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

func TestIntereractor_CreateNewUser(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := repository.NewMockMasterTx()
	masterTxManager := repository.NewMockMasterTxManager(masterTx)

	createdUser := &entity.User{
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

	userService := mock_user.NewMockService(ctrl)
	userService.EXPECT().CreateNewUser(masterTx, authID, userName, mail, name, thumbnail, bio, phone, place, birth, gender).Return(createdUser, nil).Times(1)

	interactor := New(masterTxManager, userService)
	createdUser, err := interactor.CreateNewUser(ctx, authID, userName, mail, name, thumbnail, bio, phone, place, birth, gender)

	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
}

func TestIntereractor_GetUserProfile(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

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

	masterTx := repository.NewMockMasterTx()
	masterTxManager := repository.NewMockMasterTxManager(masterTx)

	userService := mock_user.NewMockService(ctrl)
	userService.EXPECT().GetByAuthID(ctx, masterTx, authID).Return(existedUser, nil).Times(1)

	interactor := New(masterTxManager, userService)
	userData, err := interactor.GetUserProfile(ctx, authID)

	assert.NoError(t, err)
	assert.NotNil(t, userData)
}

func TestIntereractor_GetAll(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

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

	masterTx := repository.NewMockMasterTx()
	masterTxManager := repository.NewMockMasterTxManager(masterTx)

	userService := mock_user.NewMockService(ctrl)
	userService.EXPECT().GetAll(ctx, masterTx).Return(existedUsers, nil).Times(1)

	interactor := New(masterTxManager, userService)
	users, err := interactor.GetAll(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, users)
}
