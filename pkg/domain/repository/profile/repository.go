package profile

import (
	"context"
	"wantum/pkg/domain/entity/userprofile"
	"wantum/pkg/domain/repository"
)

type Repository interface {
	InsertProfile(ctx context.Context, masterTx repository.MasterTx, userProfileEntity *userprofile.Entity) (*userprofile.Entity, error)
	SelectByUserID(ctx context.Context, masterTx repository.MasterTx, userID int) (*userprofile.Entity, error)
	SelectByUserIDs(ctx context.Context, masterTx repository.MasterTx, userIDs []int) (userprofile.EntitySlice, error)
}
