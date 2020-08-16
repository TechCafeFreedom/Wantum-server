package user

import (
	"context"
	"wantum/pkg/domain/entity"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/user"
	"wantum/pkg/infrastructure/mysql/model"
	"wantum/pkg/werrors"
)

type Service interface {
	CreateNewUser(masterTx repository.MasterTx, authID, userName, mail string) (*entity.User, error)
	GetByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*entity.User, error)
	GetByAuthID(ctx context.Context, masterTx repository.MasterTx, authID string) (*entity.User, error)
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

func (s *service) CreateNewUser(masterTx repository.MasterTx, authID, userName, mail string) (*entity.User, error) {
	newUser := &model.UserModel{
		AuthID:   authID,
		UserName: userName,
		Mail:     mail,
	}
	createdUser, err := s.userRepository.InsertUser(masterTx, newUser)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return model.ConvertToUserEntity(createdUser), nil
}

func (s *service) GetByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*entity.User, error) {
	userData, err := s.userRepository.SelectByPK(ctx, masterTx, userID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return model.ConvertToUserEntity(userData), nil
}

func (s *service) GetByAuthID(ctx context.Context, masterTx repository.MasterTx, authID string) (*entity.User, error) {
	userData, err := s.userRepository.SelectByAuthID(ctx, masterTx, authID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return model.ConvertToUserEntity(userData), nil
}

func (s *service) GetAll(ctx context.Context, masterTx repository.MasterTx) (entity.UserSlice, error) {
	userSlice, err := s.userRepository.SelectAll(ctx, masterTx)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return model.ConvertToUserSliceEntity(userSlice), nil
}
