package user

import (
	"context"
	"wantum/pkg/domain/entity"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/user"
	"wantum/pkg/werrors"
)

type Service interface {
	CreateNewUser(ctx context.Context, masterTx repository.MasterTx, uid, name, thumbnail string) error
	GetByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*entity.User, error)
	GetByUID(ctx context.Context, masterTx repository.MasterTx, uid string) (*entity.User, error)
	GetAll(ctx context.Context, masterTx repository.MasterTx) (entity.UserSlice, error)
}

type service struct {
	userRepository user.Repository
}

func New(userRepository user.Repository) Service {
	return &service{
		userRepository: userRepository,
	}
}

func (s *service) CreateNewUser(ctx context.Context, masterTx repository.MasterTx, uid, name, thumbnail string) error {
	if err := s.userRepository.InsertUser(ctx, masterTx, uid, name, thumbnail); err != nil {
		return werrors.Stack(err)
	}
	return nil
}

func (s *service) GetByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*entity.User, error) {
	userData, err := s.userRepository.SelectByPK(ctx, masterTx, userID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return userData, nil
}

func (s *service) GetByUID(ctx context.Context, masterTx repository.MasterTx, uid string) (*entity.User, error) {
	userData, err := s.userRepository.SelectByUID(ctx, masterTx, uid)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return userData, nil
}

func (s *service) GetAll(ctx context.Context, masterTx repository.MasterTx) (entity.UserSlice, error) {
	userSlice, err := s.userRepository.SelectAll(ctx, masterTx)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return userSlice, nil
}
