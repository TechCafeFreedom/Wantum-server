package user

import (
	"context"
	"wantum/pkg/domain/entity/user"
	"wantum/pkg/domain/repository"
)

type Repository interface {
	InsertUser(masterTx repository.MasterTx, userEntity *user.Entity) (*user.Entity, error)
	SelectByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*user.Entity, error)
	SelectByAuthID(ctx context.Context, masterTx repository.MasterTx, authID string) (*user.Entity, error)
	SelectAll(ctx context.Context, masterTx repository.MasterTx) (user.EntitySlice, error)
}
