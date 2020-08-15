package user

import (
	"context"
	"wantum/pkg/domain/entity"
	"wantum/pkg/domain/repository"
)

type Repository interface {
	InsertUser(masterTx repository.MasterTx, userEntity *entity.User) (*entity.User, error)
	SelectByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*entity.User, error)
	SelectByAuthID(ctx context.Context, masterTx repository.MasterTx, authID string) (*entity.User, error)
	SelectAll(ctx context.Context, masterTx repository.MasterTx) (entity.UserSlice, error)
}
