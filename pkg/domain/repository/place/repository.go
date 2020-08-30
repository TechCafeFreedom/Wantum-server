package place

import (
	"context"
	"wantum/pkg/domain/repository"
	"wantum/pkg/infrastructure/mysql/model"
)

type Repository interface {
	// TODO: intでよくね？
	Insert(ctx context.Context, masterTx repository.MasterTx, place *model.PlaceModel) (*model.PlaceModel, error)

	Update(ctx context.Context, masterTx repository.MasterTx, place *model.PlaceModel) error

	Delete(ctx context.Context, masterTx repository.MasterTx, placeID int) error
	UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, place *model.PlaceModel) error

	SelectByID(ctx context.Context, masterTx repository.MasterTx, placeID int) (*model.PlaceModel, error)
	SelectAll(ctx context.Context, masterTx repository.MasterTx) (model.PlaceModelSlice, error)
}
