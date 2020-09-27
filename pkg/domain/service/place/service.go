package place

import (
	"context"
	"time"
	placeEntity "wantum/pkg/domain/entity/place"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/place"
	"wantum/pkg/werrors"
)

type Service interface {
	Create(ctx context.Context, masterTx repository.MasterTx, name string) (*placeEntity.Entity, error)
	Update(ctx context.Context, masterTx repository.MasterTx, placeID int, name string) (*placeEntity.Entity, error)
	UpdateName(ctx context.Context, masterTx repository.MasterTx, placeID int, name string) (*placeEntity.Entity, error)
	Delete(ctx context.Context, masterTx repository.MasterTx, placeID int) (*placeEntity.Entity, error)
	GetByID(ctx context.Context, masterTx repository.MasterTx, placeID int) (*placeEntity.Entity, error)
	GetAll(ctx context.Context, masterTx repository.MasterTx) (placeEntity.EntitySlice, error)
}

type service struct {
	placeRepository place.Repository
}

func New(repo place.Repository) Service {
	return &service{
		placeRepository: repo,
	}
}

func (s *service) Create(ctx context.Context, masterTx repository.MasterTx, name string) (*placeEntity.Entity, error) {
	now := time.Now()
	place := &placeEntity.Entity{
		Name:      name,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	newID, err := s.placeRepository.Insert(ctx, masterTx, place)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	place.ID = newID
	return place, nil
}

// WARNING: 空値があった時、元データが消滅する。
func (s *service) Update(ctx context.Context, masterTx repository.MasterTx, placeID int, name string) (*placeEntity.Entity, error) {
	place, err := s.placeRepository.SelectByID(ctx, masterTx, placeID)
	if err != nil {
		return nil, werrors.Stack(err)
	}

	now := time.Now()
	place.Name = name
	place.UpdatedAt = &now
	if err = s.placeRepository.Update(ctx, masterTx, place); err != nil {
		return nil, werrors.Stack(err)
	}
	return place, nil
}

func (s *service) UpdateName(ctx context.Context, masterTx repository.MasterTx, placeID int, name string) (*placeEntity.Entity, error) {
	place, err := s.placeRepository.SelectByID(ctx, masterTx, placeID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	if place == nil {
		return nil, werrors.Stack(werrors.PlaceNotFound)
	}
	now := time.Now()
	place.UpdatedAt = &now
	place.Name = name
	if err = s.placeRepository.UpdateName(ctx, masterTx, placeID, place.Name, place.UpdatedAt); err != nil {
		return nil, werrors.Stack(err)
	}
	return place, nil
}

func (s *service) Delete(ctx context.Context, masterTx repository.MasterTx, placeID int) (*placeEntity.Entity, error) {
	place, err := s.placeRepository.SelectByID(ctx, masterTx, placeID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	now := time.Now()
	place.UpdatedAt = &now
	place.DeletedAt = &now
	if err = s.placeRepository.UpDeleteFlag(ctx, masterTx, place.ID, place.UpdatedAt, place.DeletedAt); err != nil {
		return nil, werrors.Stack(err)
	}
	return place, nil
}

func (s *service) GetByID(ctx context.Context, masterTx repository.MasterTx, placeID int) (*placeEntity.Entity, error) {
	place, err := s.placeRepository.SelectByID(ctx, masterTx, placeID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return place, nil
}

func (s *service) GetAll(ctx context.Context, masterTx repository.MasterTx) (placeEntity.EntitySlice, error) {
	places, err := s.placeRepository.SelectAll(ctx, masterTx)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return places, nil
}
