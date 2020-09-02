package place

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"wantum/pkg/domain/entity"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/place"
	"wantum/pkg/infrastructure/mysql/model"
	"wantum/pkg/werrors"
)

type Service interface {
	Create(ctx context.Context, masterTx repository.MasterTx, name string) (*entity.Place, error)
	Update(ctx context.Context, masterTx repository.MasterTx, placeID int, name string) (*entity.Place, error)

	Delete(ctx context.Context, masterTx repository.MasterTx, placeID int) error
	UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, placeID int) (*entity.Place, error)

	GetByID(ctx context.Context, masterTx repository.MasterTx, placeID int) (*entity.Place, error)
	GetAll(ctx context.Context, masterTx repository.MasterTx) (entity.PlaceSlice, error)
}

type service struct {
	placeRepository place.Repository
}

func New(repo place.Repository) Service {
	return &service{
		placeRepository: repo,
	}
}

func (s *service) Create(ctx context.Context, masterTx repository.MasterTx, name string) (*entity.Place, error) {
	createdAt := time.Now()
	place := &model.PlaceModel{
		Name:      name,
		CreatedAt: &createdAt,
		UpdatedAt: &createdAt,
	}
	result, err := s.placeRepository.Insert(ctx, masterTx, place)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	place.ID = result
	return model.ConvertToPlaceEntity(place), nil
}

// NOTE: 空値があった時、元データが消滅する。
// NOTE: リクエストは、全フィールド埋める or 差分だけ
func (s *service) Update(ctx context.Context, masterTx repository.MasterTx, placeID int, name string) (*entity.Place, error) {
	place, err := s.placeRepository.SelectByID(ctx, masterTx, placeID)
	if err != nil {
		return nil, werrors.Stack(err)
	}

	updatedAt := time.Now()
	place.Name = name
	place.UpdatedAt = &updatedAt
	err = s.placeRepository.Update(ctx, masterTx, place)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return model.ConvertToPlaceEntity(place), nil
}

func (s *service) UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, placeID int) (*entity.Place, error) {
	place, err := s.placeRepository.SelectByID(ctx, masterTx, placeID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	updatedAt := time.Now()
	place.UpdatedAt = &updatedAt
	place.DeletedAt = &updatedAt
	err = s.placeRepository.UpDeleteFlag(ctx, masterTx, place)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return model.ConvertToPlaceEntity(place), nil
}

func (s *service) Delete(ctx context.Context, masterTx repository.MasterTx, placeID int) error {
	place, err := s.placeRepository.SelectByID(ctx, masterTx, placeID)
	if err != nil {
		return werrors.Stack(err)
	}
	if place.DeletedAt == nil {
		return werrors.Newf(
			fmt.Errorf("can't delete this data. this data did not up a delete flag. placeID=%v", placeID),
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

func (s *service) GetByID(ctx context.Context, masterTx repository.MasterTx, placeID int) (*entity.Place, error) {
	place, err := s.placeRepository.SelectByID(ctx, masterTx, placeID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return model.ConvertToPlaceEntity(place), nil
}

func (s *service) GetAll(ctx context.Context, masterTx repository.MasterTx) (entity.PlaceSlice, error) {
	places, err := s.placeRepository.SelectAll(ctx, masterTx)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return model.ConvertToPlaceSliceEntity(places), nil
}
