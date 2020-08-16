package profile

import (
	"context"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/profile"
	"wantum/pkg/infrastructure/mysql"
	"wantum/pkg/infrastructure/mysql/model"
	"wantum/pkg/tlog"
	"wantum/pkg/werrors"
)

type profileRepositoryImpliment struct {
	masterTxManager repository.MasterTxManager
}

func New(masterTxManager repository.MasterTxManager) profile.Repository {
	return &profileRepositoryImpliment{
		masterTxManager: masterTxManager,
	}
}

func (p *profileRepositoryImpliment) InsertProfile(ctx context.Context, masterTx repository.MasterTx, profileModel *model.ProfileModel) (*model.ProfileModel, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	_, err = tx.Exec(`
		INSERT INTO profiles(
			user_id, name, thumbnail, bio, gender, phone, place, birth
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`, profileModel.UserID,
		profileModel.Name,
		profileModel.Thumbnail,
		profileModel.Bio,
		profileModel.Gender,
		profileModel.Phone,
		profileModel.Place,
		profileModel.Birth,
	)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)

		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	return profileModel, nil
}
