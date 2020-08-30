package place

import (
	"context"
	"wantum/pkg/domain/repository"
	"wantum/pkg/infrastructure/mysql/model"
)

// implement: wantum/pkg/infrastructure/mysql/place:placeRepositoryImpletent
type Repository interface {
	Insert(ctx context.Context, masterTx repository.MasterTx, place *model.PlaceModel) (int, error)

	Update(ctx context.Context, masterTx repository.MasterTx, place *model.PlaceModel) error

	Delete(ctx context.Context, masterTx repository.MasterTx, placeID int) error
	UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, place *model.PlaceModel) error

	SelectByID(ctx context.Context, masterTx repository.MasterTx, placeID int) (*model.PlaceModel, error)
	SelectAll(ctx context.Context, masterTx repository.MasterTx) (model.PlaceModelSlice, error)
}
