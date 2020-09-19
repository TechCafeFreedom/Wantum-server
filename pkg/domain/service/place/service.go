package place

import (
	"context"
	"fmt"
	"net/http"
	"time"
	placeEntity "wantum/pkg/domain/entity/place"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/place"
	"wantum/pkg/werrors"

	"google.golang.org/grpc/codes"
)

type Service interface {
	Create(ctx context.Context, masterTx repository.MasterTx, name string) (*placeEntity.Entity, error)
	Update(ctx context.Context, masterTx repository.MasterTx, placeID int, name string) (*placeEntity.Entity, error)
	Delete(ctx context.Context, masterTx repository.MasterTx, placeID int) error
	UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, placeID int) (*placeEntity.Entity, error)
	DownDeleteFlag(ctx context.Context, masterTx repository.MasterTx, placeID int) (*placeEntity.Entity, error)
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
	result, err := s.placeRepository.Insert(ctx, masterTx, place)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	place.ID = result
	return place, nil
}

// WARNING: 空値があった時、元データが消滅する。
// QUESTION: リクエストは、全フィールド埋める or 差分だけ
func (s *service) Update(ctx context.Context, masterTx repository.MasterTx, placeID int, name string) (*placeEntity.Entity, error) {
	place, err := s.placeRepository.SelectByID(ctx, masterTx, placeID)
	if err != nil {
		return nil, werrors.Stack(err)
	}

	now := time.Now()
	place.Name = name
	place.UpdatedAt = &now
	err = s.placeRepository.Update(ctx, masterTx, place)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return place, nil
}

func (s *service) UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, placeID int) (*placeEntity.Entity, error) {
	place, err := s.placeRepository.SelectByID(ctx, masterTx, placeID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	now := time.Now()
	place.UpdatedAt = &now
	place.DeletedAt = &now
	err = s.placeRepository.UpDeleteFlag(ctx, masterTx, place)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return place, nil
}

func (s *service) DownDeleteFlag(ctx context.Context, masterTx repository.MasterTx, placeID int) (*placeEntity.Entity, error) {
	place, err := s.placeRepository.SelectByID(ctx, masterTx, placeID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	now := time.Now()
	place.UpdatedAt = &now
	place.DeletedAt = nil
	err = s.placeRepository.DownDeleteFlag(ctx, masterTx, place)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return place, nil
}

func (s *service) Delete(ctx context.Context, masterTx repository.MasterTx, placeID int) error {
	place, err := s.placeRepository.SelectByID(ctx, masterTx, placeID)
	if err != nil {
		return werrors.Stack(err)
	}
	if place.DeletedAt == nil {
		return werrors.Newf(
			fmt.Errorf("can't delete this data. this data did not up a delete flag. placeID=%v", placeID),
			codes.FailedPrecondition,
			http.StatusBadRequest,
			"このデータは削除できません",
			"could not delete this place",
		)
	}
	err = s.placeRepository.Delete(ctx, masterTx, placeID)
	if err != nil {
		return werrors.Stack(err)
	}
	return nil
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
