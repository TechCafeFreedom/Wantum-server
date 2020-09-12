package tag

import (
	"context"
	"wantum/pkg/domain/entity/tag"
	"wantum/pkg/domain/repository"
)

type Repository interface {
	Insert(ctx context.Context, masterTx repository.MasterTx, tag *tag.Entity) (int, error)

	UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, tag *tag.Entity) error
	DownDeleteFlag(ctx context.Context, masterTx repository.MasterTx, tag *tag.Entity) error
	Delete(ctx context.Context, masterTx repository.MasterTx, tagID int) error

	SelectByID(ctx context.Context, masterTx repository.MasterTx, tagID int) (*tag.Entity, error)
	SelectByName(ctx context.Context, masterTx repository.MasterTx, name string) (*tag.Entity, error)
	SelectByWishCardID(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (tag.EntitySlice, error)
	SelectByMemoryID(ctx context.Context, masterTx repository.MasterTx, memoryID int) (tag.EntitySlice, error)
}
