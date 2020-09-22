package wishboard

import (
	"context"
	"wantum/pkg/domain/entity/wishboard"
	"wantum/pkg/domain/repository"
)

type Repository interface {
	Insert(ctx context.Context, masterTx repository.MasterTx, title, backgroundImageUrl, inviteUrl string, userID int) (*wishboard.Entity, error)
	SelectByPK(ctx context.Context, masterTx repository.MasterTx, wishBoardID int) (*wishboard.Entity, error)
	SelectByPKs(ctx context.Context, masterTx repository.MasterTx, wishBoardIDs []int) (wishboard.EntitySlice, error)
	SelectByUserID(ctx context.Context, masterTx repository.MasterTx, userID int) (wishboard.EntitySlice, error)
	UpdateTitle(ctx context.Context, masterTx repository.MasterTx, wishBoardID int, title string) error
	UpdateBackgroundImageUrl(ctx context.Context, masterTx repository.MasterTx, wishBoardID int, backgroundImageUrl string) error
	Delete(ctx context.Context, masterTx repository.MasterTx, wishBoardID int) error
}
