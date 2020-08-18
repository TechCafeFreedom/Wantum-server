package profile

import (
	"context"
	"wantum/pkg/domain/entity"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/profile"
	"wantum/pkg/infrastructure/mysql/model"
	"wantum/pkg/werrors"
)

type Service interface {
	CreateNewProfile(ctx context.Context, masterTx repository.MasterTx, userID int, name, thumbnail, bio, phone, place, birth string, gender int) (*entity.Profile, error)
	GetByUserID(ctx context.Context, masterTx repository.MasterTx, userID int) (*entity.Profile, error)
	GetByUserIDs(ctx context.Context, masterTx repository.MasterTx, userIDs []int) (entity.ProfileSlice, error)
}

type service struct {
	profileReposiroty profile.Repository
}

func New(profileRepository profile.Repository) Service {
	return &service{
		profileReposiroty: profileRepository,
	}
}

func (s *service) CreateNewProfile(ctx context.Context, masterTx repository.MasterTx, userID int, name, thumbnail, bio, phone, place, birth string, gender int) (*entity.Profile, error) {
	newProfile := &model.ProfileModel{
		UserID:    userID,
		Name:      name,
		Thumbnail: thumbnail,
		Bio:       bio,
		Gender:    gender,
		Phone:     phone,
		Place:     place,
		Birth:     birth,
	}

	createdProfile, err := s.profileReposiroty.InsertProfile(ctx, masterTx, newProfile)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return model.ConvertToProfileEntity(createdProfile), nil
}

func (s *service) GetByUserID(ctx context.Context, masterTx repository.MasterTx, userID int) (*entity.Profile, error) {
	profileData, err := s.profileReposiroty.SelectByUserID(ctx, masterTx, userID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return model.ConvertToProfileEntity(profileData), nil
}

func (s *service) GetByUserIDs(ctx context.Context, masterTx repository.MasterTx, userIDs []int) (entity.ProfileSlice, error) {
	profileSlice, err := s.profileReposiroty.SelectByUserIDs(ctx, masterTx, userIDs)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return model.ConvertToProfileSliceEntity(profileSlice), nil
}
