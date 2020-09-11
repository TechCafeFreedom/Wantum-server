package user

import (
	"context"
	"wantum/pkg/domain/entity/user"
	"wantum/pkg/domain/repository"
)

type Repository interface {
	InsertUser(masterTx repository.MasterTx, userEntity *user.User) (*user.User, error)
	SelectByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*user.User, error)
	SelectByAuthID(ctx context.Context, masterTx repository.MasterTx, authID string) (*user.User, error)
	SelectAll(ctx context.Context, masterTx repository.MasterTx) (user.UserSlice, error)
}
