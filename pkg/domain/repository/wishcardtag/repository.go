package wishcardtag

import (
	"context"
	"wantum/pkg/domain/repository"
)

type Repository interface {
	Insert(ctx context.Context, masterTx repository.MasterTx, wishCardID, tagID int) error
	BulkInsert(ctx context.Context, masterTx repository.MasterTx, wishCardID int, tagIDs []int) error
	Delete(ctx context.Context, masterTx repository.MasterTx, wishCardID, tagID int) error
	DeleteByWishCardID(ctx context.Context, masterTx repository.MasterTx, wishCardID int) error
	DeleteByIDs(ctx context.Context, masterTx repository.MasterTx, wishCardID int, tagIDs []int) error
}
