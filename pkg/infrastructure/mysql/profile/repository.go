package profile

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
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
	result, err := tx.Exec(`
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
	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	profileModel.ID = int(lastInsertedID)
	return profileModel, nil
}

func (p *profileRepositoryImpliment) SelectByUserID(ctx context.Context, masterTx repository.MasterTx, userID int) (*model.ProfileModel, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	var profileData model.ProfileModel
	row := tx.QueryRow(`
		SELECT
			id, user_id, name, thumbnail, bio, gender, phone, place, birth, created_at, updated_at, deleted_at
		FROM profiles
		WHERE user_id = ?
	`, userID,
	)
	err = row.Scan(
		&profileData.ID,
		&profileData.UserID,
		&profileData.Name,
		&profileData.Thumbnail,
		&profileData.Bio,
		&profileData.Gender,
		&profileData.Phone,
		&profileData.Place,
		&profileData.Birth,
		&profileData.CreatedAt,
		&profileData.UpdatedAt,
		&profileData.DeletedAt,
	)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	return &profileData, nil
}

func (p *profileRepositoryImpliment) SelectByUserIDs(ctx context.Context, masterTx repository.MasterTx, userIDs []int) (model.ProfileModelSlice, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	var userIDsStr []string
	for _, userID := range userIDs {
		userIDsStr = append(userIDsStr, strconv.Itoa(userID))
	}

	rows, err := tx.Query(`
		SELECT
		       id, user_id, name, thumbnail, bio, gender, phone, place, birth, created_at, updated_at, deleted_at
		FROM profiles
		WHERE user_id
		IN (` + strings.Join(userIDsStr, ",") + `)
	`)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)

		if err == sql.ErrNoRows {
			return nil, nil // 一件もユーザが登録されていない場合は何も返さない
		}
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	var profileSlice model.ProfileModelSlice
	for rows.Next() {
		var profileData model.ProfileModel
		err = rows.Scan(
			&profileData.ID,
			&profileData.UserID,
			&profileData.Name,
			&profileData.Thumbnail,
			&profileData.Bio,
			&profileData.Gender,
			&profileData.Phone,
			&profileData.Place,
			&profileData.Birth,
			&profileData.CreatedAt,
			&profileData.UpdatedAt,
			&profileData.DeletedAt,
		)
		if err != nil {
			tlog.PrintErrorLogWithCtx(ctx, err)
			return nil, werrors.FromConstant(err, werrors.ServerError)
		}
		profileSlice = append(profileSlice, &profileData)
	}
	return profileSlice, nil
}
