package wishboard

import (
	"context"
	"wantum/pkg/domain/entity/wishboard"
	"wantum/pkg/domain/repository"
)

type Repository interface {
	Insert(ctx context.Context, masterTx repository.MasterTx, b *wishboard.Entity) (*wishboard.Entity, error)
	SelectByPK(ctx context.Context, masterTx repository.MasterTx, wishBoardID int) (*wishboard.Entity, error)
	SelectByUserID(ctx context.Context, masterTx repository.MasterTx, userID int) (wishboard.EntitySlice, error)
}
