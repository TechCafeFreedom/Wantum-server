package wishboard

import (
	"context"
	"wantum/pkg/domain/entity/wishboard"
	"wantum/pkg/domain/repository"
	userwishboardrepository "wantum/pkg/domain/repository/userwishboard"
	wishboardrepository "wantum/pkg/domain/repository/wishboard"
	"wantum/pkg/werrors"
)

type Service interface {
	Create(ctx context.Context, masterTx repository.MasterTx, title, backgroundImageUrl, inviteUrl string, userID int) (*wishboard.Entity, error)
	GetByPK(ctx context.Context, masterTx repository.MasterTx, wishBoardID int) (*wishboard.Entity, error)
	GetByUserID(ctx context.Context, masterTx repository.MasterTx, userID int) (wishboard.EntitySlice, error)
	UserBelongs(ctx context.Context, masterTx repository.MasterTx, userID, wishBoardID int) (bool, error)
}

type service struct {
	wishBoardRepository     wishboardrepository.Repository
	userWishBoardRepository userwishboardrepository.Repository
}

func New(wishBoardRepository wishboardrepository.Repository, userWishBoardRepository userwishboardrepository.Repository) Service {
	return &service{
		wishBoardRepository:     wishBoardRepository,
		userWishBoardRepository: userWishBoardRepository,
	}
}

func (s *service) Create(ctx context.Context, masterTx repository.MasterTx, title, backgroundImageUrl, inviteUrl string, userID int) (*wishboard.Entity, error) {
	b, err := s.wishBoardRepository.Insert(ctx, masterTx, title, backgroundImageUrl, inviteUrl, userID)
	if err != nil {
		return nil, werrors.Stack(err)
	}

	err = s.userWishBoardRepository.Insert(ctx, masterTx, userID, b.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}

	return b, nil
}

func (s *service) GetByPK(ctx context.Context, masterTx repository.MasterTx, wishBoardID int) (*wishboard.Entity, error) {
	b, err := s.wishBoardRepository.SelectByPK(ctx, masterTx, wishBoardID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return b, nil
}

func (s *service) GetByUserID(ctx context.Context, masterTx repository.MasterTx, userID int) (wishboard.EntitySlice, error) {
	bs, err := s.wishBoardRepository.SelectByUserID(ctx, masterTx, userID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return bs, err
}

func (s *service) UserBelongs(ctx context.Context, masterTx repository.MasterTx, userID, wishBoardID int) (bool, error) {
	exists, err := s.userWishBoardRepository.Select(ctx, masterTx, userID, wishBoardID)
	if err != nil {
		return false, werrors.Stack(err)
	}
	return exists, nil
}
