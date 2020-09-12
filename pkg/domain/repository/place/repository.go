package place

import (
	"context"
	"wantum/pkg/domain/entity/place"
	"wantum/pkg/domain/repository"
)

// implement: wantum/pkg/infrastructure/mysql/place:placeRepositoryImpletent
type Repository interface {
	Insert(ctx context.Context, masterTx repository.MasterTx, place *place.Entity) (int, error)

	Update(ctx context.Context, masterTx repository.MasterTx, place *place.Entity) error

	Delete(ctx context.Context, masterTx repository.MasterTx, placeID int) error
	UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, place *place.Entity) error
	DownDeleteFlag(ctx context.Context, masterTx repository.MasterTx, place *place.Entity) error

	SelectByID(ctx context.Context, masterTx repository.MasterTx, placeID int) (*place.Entity, error)
	SelectAll(ctx context.Context, masterTx repository.MasterTx) (place.EntitySlice, error)
}
