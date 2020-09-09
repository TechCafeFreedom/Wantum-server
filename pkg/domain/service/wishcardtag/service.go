package wishcardtag

import (
	"context"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/wishcardtag"
	"wantum/pkg/werrors"
)

type Service interface {
	Create(ctx context.Context, masterTx repository.MasterTx, wishCardID, tagID int) error
	CreateMultipleTags(ctx context.Context, masterTx repository.MasterTx, wishCardID int, tagIDs []int) error
	Delete(ctx context.Context, masterTx repository.MasterTx, wishCardID, tagID int) error
	DeleteByWishCardID(ctx context.Context, masterTx repository.MasterTx, wishCardID int) error
}

type service struct {
	repository wishcardtag.Repository
}

func New(repo wishcardtag.Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) Create(ctx context.Context, masterTx repository.MasterTx, wishCardID, tagID int) error {
	err := s.repository.Insert(ctx, masterTx, wishCardID, tagID)
	if err != nil {
		return werrors.Stack(err)
	}
	return nil
}

func (s *service) CreateMultipleTags(ctx context.Context, masterTx repository.MasterTx, wishCardID int, tagIDs []int) error {
	err := s.repository.BulkInsert(ctx, masterTx, wishCardID, tagIDs)
	if err != nil {
		return werrors.Stack(err)
	}
	return nil
}

func (s *service) Delete(ctx context.Context, masterTx repository.MasterTx, wishCardID, tagID int) error {
	err := s.repository.Delete(ctx, masterTx, wishCardID, tagID)
	if err != nil {
		return werrors.Stack(err)
	}
	return nil
}

func (s *service) DeleteByWishCardID(ctx context.Context, masterTx repository.MasterTx, wishCardID int) error {
	err := s.repository.DeleteByWishCardID(ctx, masterTx, wishCardID)
	if err != nil {
		return werrors.Stack(err)
	}
	return nil
}
