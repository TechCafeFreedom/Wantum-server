package user

import (
	"context"
	"wantum/pkg/domain/repository"
	"wantum/pkg/infrastructure/mysql/model"
)

type Repository interface {
	InsertUser(masterTx repository.MasterTx, userModel *model.UserModel) (*model.UserModel, error)
	SelectByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*model.UserModel, error)
	SelectByAuthID(ctx context.Context, masterTx repository.MasterTx, authID string) (*model.UserModel, error)
	SelectAll(ctx context.Context, masterTx repository.MasterTx) (model.UserModelSlice, error)
}
