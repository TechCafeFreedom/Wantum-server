package tag

import (
	"context"
	"wantum/pkg/domain/repository"
	"wantum/pkg/infrastructure/mysql/model"
)

type Repository interface {
	Insert(ctx context.Context, masterTx repository.MasterTx, tag *model.TagModel) (int, error)

	UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, tag *model.TagModel) error
	DownDeleteFlag(ctx context.Context, masterTx repository.MasterTx, tag *model.TagModel) error
	Delete(ctx context.Context, masterTx repository.MasterTx, tagID int) error

	SelectByID(ctx context.Context, masterTx repository.MasterTx, tagID int) (*model.TagModel, error)
	SelectByName(ctx context.Context, masterTx repository.MasterTx, name string) (*model.TagModel, error)
	SelectByWishCardID(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (model.TagModelSlice, error)
	SelectByMemoryID(ctx context.Context, masterTx repository.MasterTx, memoryID int) (model.TagModelSlice, error)
}
