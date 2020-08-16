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
	var profileData model.ProfileModel
	row := tx.QueryRow(`
		SELECT
		       u.id,
		       u.auth_id,
		       u.user_name,
		       u.mail,
		       u.created_at,
		       u.updated_at,
		       u.deleted_at,
		       p.id,
		       p.user_id,
		       p.name,
		       p.thumbnail,
		       p.bio,
		       p.gender,
		       p.phone,
		       p.place,
		       p.birth,
		       p.created_at,
		       p.updated_at,
		       p.deleted_at
		FROM users AS u
		JOIN profiles AS p
		ON u.id = p.user_id
		WHERE u.id = ?
	`, userID)
	err = row.Scan(
		&userData.ID,
		&userData.AuthID,
		&userData.UserName,
		&userData.Mail,
		&userData.CreatedAt,
		&userData.UpdatedAt,
		&userData.DeletedAt,
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

		if err == sql.ErrNoRows {
			return nil, werrors.FromConstant(err, werrors.UserNotFound)
		}
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	if profileData.ID != 0 {
		userData.Profile = &profileData
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
	var profileData model.ProfileModel
	row := tx.QueryRow(`
		SELECT
		       u.id,
		       u.auth_id,
		       u.user_name,
		       u.mail,
		       u.created_at,
		       u.updated_at,
		       u.deleted_at,
		       p.id,
		       p.user_id,
		       p.name,
		       p.thumbnail,
		       p.bio,
		       p.gender,
		       p.phone,
		       p.place,
		       p.birth,
		       p.created_at,
		       p.updated_at,
		       p.deleted_at
		FROM users AS u
		JOIN profiles AS p
		ON u.id = p.user_id
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

		if err == sql.ErrNoRows {
			return nil, werrors.FromConstant(err, werrors.UserNotFound)
		}
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	if profileData.ID != 0 {
		userData.Profile = &profileData
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
		SELECT
		       u.id,
		       u.auth_id,
		       u.user_name,
		       u.mail,
		       u.created_at,
		       u.updated_at,
		       u.deleted_at,
		       p.id,
		       p.user_id,
		       p.name,
		       p.thumbnail,
		       p.bio,
		       p.gender,
		       p.phone,
		       p.place,
		       p.birth,
		       p.created_at,
		       p.updated_at,
		       p.deleted_at
		FROM users AS u
		JOIN profiles AS p
		ON u.id = p.user_id
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
		var profileData model.ProfileModel
		err = rows.Scan(
			&userData.ID,
			&userData.AuthID,
			&userData.UserName,
			&userData.Mail,
			&userData.CreatedAt,
			&userData.UpdatedAt,
			&userData.DeletedAt,
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
		if profileData.ID != 0 {
			userData.Profile = &profileData
		}
		userSlice = append(userSlice, &userData)
	}

	return userSlice, nil
}
