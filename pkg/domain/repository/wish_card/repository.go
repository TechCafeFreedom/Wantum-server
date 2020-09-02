package wish_card

import (
	"context"
	"wantum/pkg/domain/repository"
	"wantum/pkg/infrastructure/mysql/model"
)

type Repository interface {
	Insert(ctx context.Context, masterTx repository.MasterTx, wishCard *model.WishCardModel) (int, error)

	Update(ctx context.Context, masterTx repository.MasterTx, wishCard *model.WishCardModel) error

	UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, wishCard *model.WishCardModel) error
	DownDeleteFlag(ctx context.Context, masterTx repository.MasterTx, wishCard *model.WishCardModel) error
	Delete(ctx context.Context, masterTx repository.MasterTx, wishCardID int) error

	SelectByID(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (*model.WishCardModel, error)
	SelectByIDs(ctx context.Context, masterTx repository.MasterTx, wishCardIDs []string) (model.WishCardModelSlice, error)
	SelectByCategoryID(ctx context.Context, masterTx repository.MasterTx, categryID int) (model.WishCardModelSlice, error)
	// TODO: SelectByCategoryIDsあっても良いかもと思ったが？
}