package wishboard

import (
	"context"
	"time"
	"wantum/pkg/domain/entity/wishboard"
	"wantum/pkg/domain/repository"
	userwishboardrepository "wantum/pkg/domain/repository/userwishboard"
	wishboardrepository "wantum/pkg/domain/repository/wishboard"
	"wantum/pkg/werrors"
)

type Service interface {
	Create(ctx context.Context, masterTx repository.MasterTx, title, backgroundImageURL, inviteURL string, userID int) (*wishboard.Entity, error)
	GetByPK(ctx context.Context, masterTx repository.MasterTx, wishBoardID int) (*wishboard.Entity, error)
	GetMyBoards(ctx context.Context, masterTx repository.MasterTx, userID int) (wishboard.EntitySlice, error)
	IsMember(ctx context.Context, masterTx repository.MasterTx, userID, wishBoardID int) (bool, error)
	UpdateTitle(ctx context.Context, masterTx repository.MasterTx, wishBoardID int, title string) error
	UpdateBackgroundImageURL(ctx context.Context, masterTx repository.MasterTx, wishBoardID int, backgroundImageURL string) error
	Delete(ctx context.Context, masterTx repository.MasterTx, wishBoardID int) error
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

func (s *service) Create(ctx context.Context, masterTx repository.MasterTx, title, backgroundImageURL, inviteURL string, userID int) (*wishboard.Entity, error) {
	now := time.Now()

	b := &wishboard.Entity{
		Title:              title,
		BackgroundImageURL: backgroundImageURL,
		InviteURL:          inviteURL,
		UserID:             userID,
		CreatedAt:          &now,
		UpdatedAt:          &now,
	}

	b, err := s.wishBoardRepository.Insert(ctx, masterTx, b)
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

func (s *service) GetMyBoards(ctx context.Context, masterTx repository.MasterTx, userID int) (wishboard.EntitySlice, error) {
	wishBoardIDs, err := s.userWishBoardRepository.SelectByUserID(ctx, masterTx, userID)
	if err != nil {
		return nil, werrors.Stack(err)
	}

	if len(wishBoardIDs) == 0 {
		return nil, nil
	}

	bs, err := s.wishBoardRepository.SelectByPKs(ctx, masterTx, wishBoardIDs)
	if err != nil {
		return nil, werrors.Stack(err)
	}

	return bs, nil
}

func (s *service) IsMember(ctx context.Context, masterTx repository.MasterTx, userID, wishBoardID int) (bool, error) {
	exists, err := s.userWishBoardRepository.Exists(ctx, masterTx, userID, wishBoardID)
	if err != nil {
		return false, werrors.Stack(err)
	}
	return exists, nil
}

func (s *service) UpdateTitle(ctx context.Context, masterTx repository.MasterTx, wishBoardID int, title string) error {
	now := time.Now()

	err := s.wishBoardRepository.UpdateTitle(ctx, masterTx, wishBoardID, title, &now)
	if err != nil {
		return werrors.Stack(err)
	}
	return nil
}

func (s *service) UpdateBackgroundImageURL(ctx context.Context, masterTx repository.MasterTx, wishBoardID int, backgroundImageURL string) error {
	now := time.Now()

	err := s.wishBoardRepository.UpdateBackgroundImageURL(ctx, masterTx, wishBoardID, backgroundImageURL, &now)
	if err != nil {
		return werrors.Stack(err)
	}
	return nil
}

func (s *service) Delete(ctx context.Context, masterTx repository.MasterTx, wishBoardID int) error {
	now := time.Now()

	b := &wishboard.Entity{ID: wishBoardID, UpdatedAt: &now, DeletedAt: &now}

	err := s.wishBoardRepository.Delete(ctx, masterTx, b)
	if err != nil {
		return werrors.Stack(err)
	}
	return nil
}
