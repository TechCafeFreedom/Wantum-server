package user

import (
	"context"
	"testing"
	"time"
	"wantum/pkg/domain/entity"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/service/file/mock_file"
	"wantum/pkg/domain/service/profile/mock_profile"
	"wantum/pkg/domain/service/user/mock_user"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	userID      = 1
	authID      = "authID"
	userName    = "userName"
	mail        = "test@test.com"
	name        = "name"
	thumbnail   = "thumbnail"
	bio         = "bio"
	gender      = 1
	phone       = "000-0000-0000"
	place       = "place"
	birthLayout = "2006-01-02"
	birthDay    = "1998-05-03"
)

var (
	userIDs       = []int{userID}
	birth, _      = time.Parse(birthLayout, birthDay)
	birthJST      = birth.Local()
	birthUnix     = birth.Unix()
	thumbnailFile = []byte{1, 2, 3, 4}

	dummyUserEntity = &entity.User{
		ID:       userID,
		AuthID:   authID,
		UserName: userName,
		Mail:     mail,
	}

	dummyProfileEntity = &entity.Profile{
		UserID:    userID,
		Name:      name,
		Thumbnail: thumbnail,
		Bio:       bio,
		Gender:    gender,
		Phone:     phone,
		Place:     place,
		Birth:     &birth,
	}

	dummyProfileSlice = entity.ProfileSlice{
		dummyProfileEntity,
	}

	dummyUserSlice = entity.UserSlice{
		dummyUserEntity,
	}

	dummyUserMap = entity.UserMap{userID: dummyUserSlice[0]}
)

func TestIntereractor_CreateNewUser(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := repository.NewMockMasterTx()
	masterTxManager := repository.NewMockMasterTxManager(masterTx)

	userService := mock_user.NewMockService(ctrl)
	userService.EXPECT().CreateNewUser(masterTx, authID, userName, mail).Return(dummyUserEntity, nil).Times(1)

	profileService := mock_profile.NewMockService(ctrl)
	profileService.EXPECT().CreateNewProfile(ctx, masterTx, userID, name, thumbnail, bio, phone, place, &birthJST, gender).Return(dummyProfileEntity, nil).Times(1)

	fileService := mock_file.NewMockService(ctrl)
	fileService.EXPECT().UploadImageToLocalFolder(thumbnailFile).Return(thumbnail, nil).Times(1)

	interactor := New(masterTxManager, userService, profileService, fileService)
	userData, err := interactor.CreateNewUser(ctx, authID, userName, mail, name, bio, phone, place, thumbnailFile, int(birthUnix), gender)

	assert.NoError(t, err)
	assert.NotNil(t, userData)
	assert.Equal(t, dummyUserEntity, userData)
}

func TestIntereractor_GetUserProfile(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := repository.NewMockMasterTx()
	masterTxManager := repository.NewMockMasterTxManager(masterTx)

	userService := mock_user.NewMockService(ctrl)
	userService.EXPECT().GetByAuthID(ctx, masterTx, authID).Return(dummyUserEntity, nil).Times(1)

	profileService := mock_profile.NewMockService(ctrl)
	profileService.EXPECT().GetByUserID(ctx, masterTx, userID).Return(dummyProfileEntity, nil).Times(1)

	fileService := mock_file.NewMockService(ctrl)

	interactor := New(masterTxManager, userService, profileService, fileService)
	userData, err := interactor.GetAuthorizedUser(ctx, authID)

	assert.NoError(t, err)
	assert.NotNil(t, userData)
	assert.Equal(t, dummyUserEntity, userData)
}

func TestIntereractor_GetAll(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := repository.NewMockMasterTx()
	masterTxManager := repository.NewMockMasterTxManager(masterTx)

	userService := mock_user.NewMockService(ctrl)
	userService.EXPECT().GetAll(ctx, masterTx).Return(dummyUserSlice, nil).Times(1)

	profileService := mock_profile.NewMockService(ctrl)
	profileService.EXPECT().GetByUserIDs(ctx, masterTx, userIDs).Return(dummyProfileSlice, nil).Times(1)

	fileService := mock_file.NewMockService(ctrl)

	interactor := New(masterTxManager, userService, profileService, fileService)
	users, err := interactor.GetAll(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, dummyUserMap, users)
}
