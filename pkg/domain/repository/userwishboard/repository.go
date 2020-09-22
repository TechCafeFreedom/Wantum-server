package userwishboard

import (
	"context"
	"wantum/pkg/domain/repository"
)

type Repository interface {
	Insert(ctx context.Context, masterTx repository.MasterTx, userID, wishBoardID int) error
	SelectByUserIDAndWishBoardID(ctx context.Context, masterTx repository.MasterTx, userID, wishBoardID int) (bool, error)
	SelectByUserID(ctx context.Context, masterTx repository.MasterTx, userID int) ([]int, error)
	Delete(ctx context.Context, masterTx repository.MasterTx, userID, wishBoardID int) error
}
