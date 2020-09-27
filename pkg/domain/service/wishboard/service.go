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

	newWishBoard := &wishboard.Entity{
		Title:              title,
		BackgroundImageURL: backgroundImageURL,
		InviteURL:          inviteURL,
		UserID:             userID,
		CreatedAt:          &now,
		UpdatedAt:          &now,
	}

	createdWishBoard, err := s.wishBoardRepository.Insert(ctx, masterTx, newWishBoard)
	if err != nil {
		return nil, werrors.Stack(err)
	}

	if err := s.userWishBoardRepository.Insert(ctx, masterTx, userID, createdWishBoard.ID); err != nil {
		return nil, werrors.Stack(err)
	}

	return createdWishBoard, nil
}

func (s *service) GetByPK(ctx context.Context, masterTx repository.MasterTx, wishBoardID int) (*wishboard.Entity, error) {
	wishBoardEntity, err := s.wishBoardRepository.SelectByPK(ctx, masterTx, wishBoardID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return wishBoardEntity, nil
}

func (s *service) GetMyBoards(ctx context.Context, masterTx repository.MasterTx, userID int) (wishboard.EntitySlice, error) {
	wishBoardIDs, err := s.userWishBoardRepository.SelectByUserID(ctx, masterTx, userID)
	if err != nil {
		return nil, werrors.Stack(err)
	}

	if len(wishBoardIDs) == 0 {
		return nil, nil
	}

	wishBoardSlice, err := s.wishBoardRepository.SelectByPKs(ctx, masterTx, wishBoardIDs)
	if err != nil {
		return nil, werrors.Stack(err)
	}

	return wishBoardSlice, nil
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

	if err := s.wishBoardRepository.UpdateTitle(ctx, masterTx, wishBoardID, title, &now); err != nil {
		return werrors.Stack(err)
	}
	return nil
}

func (s *service) UpdateBackgroundImageURL(ctx context.Context, masterTx repository.MasterTx, wishBoardID int, backgroundImageURL string) error {
	now := time.Now()

	if err := s.wishBoardRepository.UpdateBackgroundImageURL(ctx, masterTx, wishBoardID, backgroundImageURL, &now); err != nil {
		return werrors.Stack(err)
	}
	return nil
}

func (s *service) Delete(ctx context.Context, masterTx repository.MasterTx, wishBoardID int) error {
	now := time.Now()

	wishBoardEntity := &wishboard.Entity{
		ID:        wishBoardID,
		UpdatedAt: &now,
		DeletedAt: &now,
	}

	if err := s.wishBoardRepository.Delete(ctx, masterTx, wishBoardEntity); err != nil {
		return werrors.Stack(err)
	}
	return nil
}
