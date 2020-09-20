package userwishboard

import (
	"context"
	"wantum/pkg/domain/repository"
)

type Repository interface {
	Insert(ctx context.Context, masterTx repository.MasterTx, userID, wishBoardID int) error
	Select(ctx context.Context, masterTx repository.MasterTx, userID, wishBoardID int) (bool, error)
	Delete(ctx context.Context, masterTx repository.MasterTx, userID, wishBoardID int) error
}
