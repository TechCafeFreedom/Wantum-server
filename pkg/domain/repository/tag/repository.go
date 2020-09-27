package tag

import (
	"context"
	"time"
	"wantum/pkg/domain/entity/tag"
	"wantum/pkg/domain/repository"
)

type Repository interface {
	Insert(ctx context.Context, masterTx repository.MasterTx, tag *tag.Entity) (int, error)
	UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, tagID int, updatedAt, deletedAt *time.Time) error
	DownDeleteFlag(ctx context.Context, masterTx repository.MasterTx, tagID int, updatedAt *time.Time) error
	Delete(ctx context.Context, masterTx repository.MasterTx, tag *tag.Entity) error
	SelectByID(ctx context.Context, masterTx repository.MasterTx, tagID int) (*tag.Entity, error)
	SelectByIDs(ctx context.Context, masterTx repository.MasterTx, tagIDs []int) (tag.EntitySlice, error)
	SelectByName(ctx context.Context, masterTx repository.MasterTx, name string) (*tag.Entity, error)
	SelectByWishCardID(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (tag.EntitySlice, error)
	SelectByMemoryID(ctx context.Context, masterTx repository.MasterTx, memoryID int) (tag.EntitySlice, error)
}
