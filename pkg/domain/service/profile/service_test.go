package profile

import (
	"context"
	"testing"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/profile/mock_profile"
	"wantum/pkg/infrastructure/mysql/model"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	userID    = 1
	name      = "name"
	thumbnail = "thumbnail"
	bio       = "bio"
	gender    = 1
	phone     = "000-0000-0000"
	place     = "place"
	birth     = "1998-05-03"
)

func TestService_CreateNewProfile(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := repository.NewMockMasterTx()

	profileRepository := mock_profile.NewMockRepository(ctrl)
	profileModel := &model.ProfileModel{
		UserID:    userID,
		Name:      name,
		Thumbnail: thumbnail,
		Bio:       bio,
		Gender:    gender,
		Phone:     phone,
		Place:     place,
		Birth:     birth,
	}
	profileRepository.EXPECT().InsertProfile(ctx, masterTx, profileModel).Return(profileModel, nil).Times(1)

	service := New(profileRepository)
	createdProfile, err := service.CreateNewProfile(ctx, masterTx, userID, name, thumbnail, bio, phone, place, birth, gender)

	assert.NoError(t, err)
	assert.NotNil(t, createdProfile)
}
