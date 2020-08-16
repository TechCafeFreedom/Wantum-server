package user

import (
	"context"
	"database/sql"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/user"
	"wantum/pkg/infrastructure/mysql"
	"wantum/pkg/infrastructure/mysql/model"
	"wantum/pkg/tlog"
	"wantum/pkg/werrors"
)

type userRepositoryImpliment struct {
	masterTxManager repository.MasterTxManager
}

func New(masterTxManager repository.MasterTxManager) user.Repository {
	return &userRepositoryImpliment{
		masterTxManager: masterTxManager,
	}
}

func (u *userRepositoryImpliment) InsertUser(masterTx repository.MasterTx, userModel *model.UserModel) (*model.UserModel, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithAuthID(userModel.AuthID, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	result, err := tx.Exec(`
			INSERT INTO users(
			    auth_id, user_name, mail
			)
			VALUES (?, ?, ?)
	`, userModel.AuthID, userModel.UserName, userModel.Mail)
	if err != nil {
		tlog.PrintErrorLogWithAuthID(userModel.AuthID, err)

		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	createdUserID, err := result.LastInsertId()
	if err != nil {
		tlog.PrintErrorLogWithAuthID(userModel.AuthID, err)

		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	userModel.ID = int(createdUserID)

	return userModel, nil
}

func (u *userRepositoryImpliment) SelectByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*model.UserModel, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	var userData model.UserModel
	row := tx.QueryRow(`
		SELECT *
		FROM users
		JOIN profiles
		ON users.id = profiles.user_id
		WHERE users.id = ?
	`, userID)
	err = row.Scan(
		&userData.ID,
		&userData.AuthID,
		&userData.UserName,
		&userData.Mail,
		&userData.CreatedAt,
		&userData.UpdatedAt,
		&userData.DeletedAt,
		&userData.Profile.ID,
		&userData.Profile.UserID,
		&userData.Profile.Name,
		&userData.Profile.Thumbnail,
		&userData.Profile.Bio,
		&userData.Profile.Gender,
		&userData.Profile.Phone,
		&userData.Profile.Place,
		&userData.Profile.Birth,
		&userData.Profile.CreatedAt,
		&userData.Profile.UpdatedAt,
		&userData.Profile.DeletedAt,
	)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)

		if err == sql.ErrNoRows {
			return nil, werrors.FromConstant(err, werrors.UserNotFound)
		}
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	return &userData, nil
}

func (u *userRepositoryImpliment) SelectByAuthID(ctx context.Context, masterTx repository.MasterTx, authID string) (*model.UserModel, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	var userData model.UserModel
	row := tx.QueryRow(`
		SELECT *
		FROM users
		JOIN profiles
		ON users.id = profiles.user_id
		WHERE auth_id = ?
	`, authID)
	err = row.Scan(
		&userData.ID,
		&userData.AuthID,
		&userData.UserName,
		&userData.Mail,
		&userData.CreatedAt,
		&userData.UpdatedAt,
		&userData.DeletedAt,
		&userData.Profile.ID,
		&userData.Profile.UserID,
		&userData.Profile.Name,
		&userData.Profile.Thumbnail,
		&userData.Profile.Bio,
		&userData.Profile.Gender,
		&userData.Profile.Phone,
		&userData.Profile.Place,
		&userData.Profile.Birth,
		&userData.Profile.CreatedAt,
		&userData.Profile.UpdatedAt,
		&userData.Profile.DeletedAt,
	)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)

		if err == sql.ErrNoRows {
			return nil, werrors.FromConstant(err, werrors.UserNotFound)
		}
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	return &userData, nil
}

func (u *userRepositoryImpliment) SelectAll(ctx context.Context, masterTx repository.MasterTx) (model.UserModelSlice, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	rows, err := tx.Query(`
		SELECT *
		FROM users
		JOIN profiles
		ON users.id = profiles.user_id
	`)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)

		if err == sql.ErrNoRows {
			return nil, nil // 一件もユーザが登録されていない場合は何も返さない
		}
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	var userSlice model.UserModelSlice
	for rows.Next() {
		var userData model.UserModel
		err = rows.Scan(
			&userData.ID,
			&userData.AuthID,
			&userData.UserName,
			&userData.Mail,
			&userData.CreatedAt,
			&userData.UpdatedAt,
			&userData.DeletedAt,
			&userData.Profile.ID,
			&userData.Profile.UserID,
			&userData.Profile.Name,
			&userData.Profile.Thumbnail,
			&userData.Profile.Bio,
			&userData.Profile.Gender,
			&userData.Profile.Phone,
			&userData.Profile.Place,
			&userData.Profile.Birth,
			&userData.Profile.CreatedAt,
			&userData.Profile.UpdatedAt,
			&userData.Profile.DeletedAt,
		)
		if err != nil {
			tlog.PrintErrorLogWithCtx(ctx, err)
			return nil, werrors.FromConstant(err, werrors.ServerError)
		}
		userSlice = append(userSlice, &userData)
	}

	return userSlice, nil
}
