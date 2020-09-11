package profile

import (
	"context"
	"time"
	"wantum/pkg/domain/entity/userprofile"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/profile"
	"wantum/pkg/werrors"
)

type Service interface {
	CreateNewProfile(ctx context.Context, masterTx repository.MasterTx, userID int, name, thumbnail, bio, phone, place string, birth *time.Time, gender int) (*userprofile.Entity, error)
	GetByUserID(ctx context.Context, masterTx repository.MasterTx, userID int) (*userprofile.Entity, error)
	GetByUserIDs(ctx context.Context, masterTx repository.MasterTx, userIDs []int) (userprofile.EntitySlice, error)
}

type service struct {
	profileRepository profile.Repository
}

func New(profileRepository profile.Repository) Service {
	return &service{
		profileRepository: profileRepository,
	}
}

func (s *service) CreateNewProfile(ctx context.Context, masterTx repository.MasterTx, userID int, name, thumbnail, bio, phone, place string, birth *time.Time, gender int) (*userprofile.Entity, error) {
	newProfile := &userprofile.Entity{
		UserID:    userID,
		Name:      name,
		Thumbnail: thumbnail,
		Bio:       bio,
		Gender:    gender,
		Phone:     phone,
		Place:     place,
		Birth:     birth,
	}

	createdProfile, err := s.profileRepository.InsertProfile(ctx, masterTx, newProfile)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return createdProfile, nil
}

func (s *service) GetByUserID(ctx context.Context, masterTx repository.MasterTx, userID int) (*userprofile.Entity, error) {
	profileData, err := s.profileRepository.SelectByUserID(ctx, masterTx, userID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return profileData, nil
}

func (s *service) GetByUserIDs(ctx context.Context, masterTx repository.MasterTx, userIDs []int) (userprofile.EntitySlice, error) {
	profileSlice, err := s.profileRepository.SelectByUserIDs(ctx, masterTx, userIDs)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return profileSlice, nil
}
