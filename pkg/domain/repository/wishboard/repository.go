package wishboard

import (
	"context"
	"time"
	"wantum/pkg/domain/entity/wishboard"
	"wantum/pkg/domain/repository"
)

type Repository interface {
	Insert(ctx context.Context, masterTx repository.MasterTx, title, backgroundImageURL, inviteURL string, userID int, createdAt, updatedAt time.Time) (*wishboard.Entity, error)
	SelectByPK(ctx context.Context, masterTx repository.MasterTx, wishBoardID int) (*wishboard.Entity, error)
	SelectByPKs(ctx context.Context, masterTx repository.MasterTx, wishBoardIDs []int) (wishboard.EntitySlice, error)
	SelectByUserID(ctx context.Context, masterTx repository.MasterTx, userID int) (wishboard.EntitySlice, error)
	UpdateTitle(ctx context.Context, masterTx repository.MasterTx, wishBoardID int, title string, updatedAt time.Time) error
	UpdateBackgroundImageURL(ctx context.Context, masterTx repository.MasterTx, wishBoardID int, backgroundImageURL string, updatedAt time.Time) error
	Delete(ctx context.Context, masterTx repository.MasterTx, wishBoardID int, updatedAt, deletedAt time.Time) error
}
