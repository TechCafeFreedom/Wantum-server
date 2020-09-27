package wishcard

import (
	"context"
	"time"
	"wantum/pkg/domain/entity/wishcard"
	"wantum/pkg/domain/repository"
)

type Repository interface {
	Insert(ctx context.Context, masterTx repository.MasterTx, wishCard *wishcard.Entity, categoryID int) (int, error)
	Update(ctx context.Context, masterTx repository.MasterTx, wishCard *wishcard.Entity) error
	UpdateActivity(ctx context.Context, masterTx repository.MasterTx, wishCardID int, activity string, updatedAt *time.Time) error
	UpdateDescription(ctx context.Context, masterTx repository.MasterTx, wishCardID int, description string, updatedAt *time.Time) error
	UpdateDate(ctx context.Context, masterTx repository.MasterTx, wishCardID int, date, updatedAt *time.Time) error
	UpdateDoneAt(ctx context.Context, masterTx repository.MasterTx, wishCardID int, doneAt, updatedAt *time.Time) error
	UpdateUserID(ctx context.Context, masterTx repository.MasterTx, wishCardID int, userID int, updatedAt *time.Time) error
	UpdatePlaceID(ctx context.Context, masterTx repository.MasterTx, wishCardID int, placeID int, updatedAt *time.Time) error
	UpdateCategoryID(ctx context.Context, masterTx repository.MasterTx, wishCardID int, categoryID int, updatedAt *time.Time) error
	UpdateWithCategoryID(ctx context.Context, masterTx repository.MasterTx, wishCard *wishcard.Entity, categoryID int) error
	UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, wishCardID int, updatedAt, deletedAt *time.Time) error
	DownDeleteFlag(ctx context.Context, masterTx repository.MasterTx, wishCardID int, updatedAt *time.Time) error
	Delete(ctx context.Context, masterTx repository.MasterTx, wishCard *wishcard.Entity) error
	SelectByID(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (*wishcard.Entity, error)
	SelectByIDs(ctx context.Context, masterTx repository.MasterTx, wishCardIDs []int) (wishcard.EntitySlice, error)
	SelectByCategoryID(ctx context.Context, masterTx repository.MasterTx, categryID int) (wishcard.EntitySlice, error)
}
