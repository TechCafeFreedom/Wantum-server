package wishcard

import (
	"context"
	"wantum/pkg/domain/entity/wishcard"
	"wantum/pkg/domain/repository"
)

type Repository interface {
	Insert(ctx context.Context, masterTx repository.MasterTx, wishCard *wishcard.Entity, categoryID int) (int, error)

	Update(ctx context.Context, masterTx repository.MasterTx, wishCard *wishcard.Entity, categoryID int) error

	UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, wishCard *wishcard.Entity) error
	DownDeleteFlag(ctx context.Context, masterTx repository.MasterTx, wishCard *wishcard.Entity) error
	Delete(ctx context.Context, masterTx repository.MasterTx, wishCardID int) error

	SelectByID(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (*wishcard.Entity, error)
	SelectByIDs(ctx context.Context, masterTx repository.MasterTx, wishCardIDs []string) (wishcard.EntitySlice, error)
	SelectByCategoryID(ctx context.Context, masterTx repository.MasterTx, categryID int) (wishcard.EntitySlice, error)
	// TODO: SelectByCategoryIDsあっても良いかもと思ったが？
}
