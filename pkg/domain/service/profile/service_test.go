package profile

import (
	"context"
	"testing"
	"time"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/profile/mock_profile"
	"wantum/pkg/infrastructure/mysql/model"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	profileID = 1
	userID    = 1
	name      = "name"
	thumbnail = "thumbnail"
	bio       = "bio"
	gender    = 1
	phone     = "000-0000-0000"
	place     = "place"
)

var (
	birth             = time.Now()
	dummyProfileModel = &model.ProfileModel{
		ID:        profileID,
		UserID:    userID,
		Name:      name,
		Thumbnail: thumbnail,
		Bio:       bio,
		Gender:    gender,
		Phone:     phone,
		Place:     place,
		Birth:     &birth,
	}

	dummyProfileModelWithoutID = &model.ProfileModel{
		UserID:    userID,
		Name:      name,
		Thumbnail: thumbnail,
		Bio:       bio,
		Gender:    gender,
		Phone:     phone,
		Place:     place,
		Birth:     &birth,
	}

	dummyProfileModelSlice = model.ProfileModelSlice{
		dummyProfileModel,
	}
)

func TestService_CreateNewProfile(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := repository.NewMockMasterTx()

	profileRepository := mock_profile.NewMockRepository(ctrl)

	profileRepository.EXPECT().InsertProfile(ctx, masterTx, dummyProfileModelWithoutID).Return(dummyProfileModelWithoutID, nil).Times(1)

	service := New(profileRepository)
	createdProfile, err := service.CreateNewProfile(ctx, masterTx, userID, name, thumbnail, bio, phone, place, &birth, gender)

	assert.NoError(t, err)
	assert.NotNil(t, createdProfile)
}

func TestService_GetByProfileID(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := repository.NewMockMasterTx()

	profileRepository := mock_profile.NewMockRepository(ctrl)
	profileRepository.EXPECT().SelectByUserID(ctx, masterTx, userID).Return(dummyProfileModel, nil).Times(1)

	service := New(profileRepository)
	profileData, err := service.GetByUserID(ctx, masterTx, userID)

	assert.NoError(t, err)
	assert.NotNil(t, profileData)
}

func TestService_GetByProfileIDs(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userIDs := []int{userID}

	masterTx := repository.NewMockMasterTx()

	profileRepository := mock_profile.NewMockRepository(ctrl)
	profileRepository.EXPECT().SelectByUserIDs(ctx, masterTx, userIDs).Return(dummyProfileModelSlice, nil).Times(1)

	service := New(profileRepository)
	profileSlice, err := service.GetByUserIDs(ctx, masterTx, userIDs)

	assert.NoError(t, err)
	assert.NotNil(t, profileSlice)
}
