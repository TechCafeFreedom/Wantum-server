package place

import (
	"context"
	"time"
	"wantum/pkg/domain/entity/place"
	"wantum/pkg/domain/repository"
)

type Repository interface {
	Insert(ctx context.Context, masterTx repository.MasterTx, place *place.Entity) (int, error)
	Update(ctx context.Context, masterTx repository.MasterTx, place *place.Entity) error
	UpdateName(ctx context.Context, masterTx repository.MasterTx, placeID int, name string, updatedAt *time.Time) error
	Delete(ctx context.Context, masterTx repository.MasterTx, place *place.Entity) error
	UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, placeID int, updatedAt, deletedAt *time.Time) error
	DownDeleteFlag(ctx context.Context, masterTx repository.MasterTx, placeID int, updatedAt *time.Time) error
	SelectByID(ctx context.Context, masterTx repository.MasterTx, placeID int) (*place.Entity, error)
	SelectAll(ctx context.Context, masterTx repository.MasterTx) (place.EntitySlice, error)
}
