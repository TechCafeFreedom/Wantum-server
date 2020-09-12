package user

import (
	"context"
	userentity "wantum/pkg/domain/entity/user"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/user"
	"wantum/pkg/werrors"
)

type Service interface {
	CreateNewUser(masterTx repository.MasterTx, authID, userName, mail string) (*userentity.Entity, error)
	GetByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*userentity.Entity, error)
	GetByAuthID(ctx context.Context, masterTx repository.MasterTx, authID string) (*userentity.Entity, error)
	GetAll(ctx context.Context, masterTx repository.MasterTx) (userentity.EntitySlice, error)
}

type service struct {
	userRepository user.Repository
}

func New(userRepository user.Repository) Service {
	return &service{
		userRepository: userRepository,
	}
}

func (s *service) CreateNewUser(masterTx repository.MasterTx, authID, userName, mail string) (*userentity.Entity, error) {
	newUser := &userentity.Entity{
		AuthID:   authID,
		UserName: userName,
		Mail:     mail,
	}
	createdUser, err := s.userRepository.InsertUser(masterTx, newUser)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return createdUser, nil
}

func (s *service) GetByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*userentity.Entity, error) {
	userEntity, err := s.userRepository.SelectByPK(ctx, masterTx, userID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return userEntity, nil
}

func (s *service) GetByAuthID(ctx context.Context, masterTx repository.MasterTx, authID string) (*userentity.Entity, error) {
	userEntity, err := s.userRepository.SelectByAuthID(ctx, masterTx, authID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return userEntity, nil
}

func (s *service) GetAll(ctx context.Context, masterTx repository.MasterTx) (userentity.EntitySlice, error) {
	userSlice, err := s.userRepository.SelectAll(ctx, masterTx)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return userSlice, nil
}
