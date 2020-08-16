package profile

import (
	"context"
	"wantum/pkg/domain/repository"
	"wantum/pkg/infrastructure/mysql/model"
)

type Repository interface {
	InsertProfile(ctx context.Context, masterTx repository.MasterTx, profileEntity *model.ProfileModel) (*model.ProfileModel, error)
}
