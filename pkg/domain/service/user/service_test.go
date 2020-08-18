package user

import (
	"context"
	"testing"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/user/mock_user"
	"wantum/pkg/infrastructure/mysql/model"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	userID   = 1
	authID   = "authID"
	userName = "userName"
	mail     = "test@test.com"
)

func TestService_CreateNewUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := repository.NewMockMasterTx()

	userRepository := mock_user.NewMockRepository(ctrl)
	userModel := &model.UserModel{
		AuthID:   authID,
		UserName: userName,
		Mail:     mail,
	}
	userRepository.EXPECT().InsertUser(masterTx, userModel).Return(userModel, nil).Times(1)

	service := New(userRepository)
	createdUser, err := service.CreateNewUser(masterTx, authID, userName, mail)

	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
}

func TestService_GetByPK(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := repository.NewMockMasterTx()

	existedUser := &model.UserModel{
		ID:       userID,
		AuthID:   authID,
		UserName: userName,
		Mail:     mail,
	}

	userRepository := mock_user.NewMockRepository(ctrl)
	userRepository.EXPECT().SelectByPK(ctx, masterTx, userID).Return(existedUser, nil).Times(1)

	service := New(userRepository)
	user, err := service.GetByPK(ctx, masterTx, userID)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestService_SelectAll(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := repository.NewMockMasterTx()

	existedUsers := model.UserModelSlice{
		{
			ID:       userID,
			AuthID:   authID,
			UserName: userName,
			Mail:     mail,
		},
	}

	userRepository := mock_user.NewMockRepository(ctrl)
	userRepository.EXPECT().SelectAll(ctx, masterTx).Return(existedUsers, nil).Times(1)

	service := New(userRepository)
	users, err := service.GetAll(ctx, masterTx)

	assert.NoError(t, err)
	assert.NotNil(t, users)
}
