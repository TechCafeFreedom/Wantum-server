package user

import (
	"context"
	"database/sql"
	"net/http"
	"wantum/pkg/domain/entity"
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

func (u *userRepositoryImpliment) InsertUser(masterTx repository.MasterTx, userEntity *entity.User) (*entity.User, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintLogWithUID(userEntity.AuthID, err)
		return nil, werrors.Stack(err)
	}
	createdUser, err := tx.Exec(`
			INSERT INTO users(
			    auth_id, user_name, mail
			)
			VALUES (?, ?, ?)
	`, userEntity.AuthID, userEntity.UserName, userEntity.Mail)
	if err != nil {
		tlog.PrintLogWithUID(userEntity.AuthID, err)

		return nil, werrors.Newf(
			err,
			http.StatusInternalServerError,
			"DBインサート時にエラーが発生しました。",
			"Error occurred when DB insert.",
		)
	}

	createdUserID, err := createdUser.LastInsertId()
	if err != nil {
		tlog.PrintLogWithUID(userEntity.AuthID, err)

		return nil, werrors.Newf(
			err,
			http.StatusInternalServerError,
			"DBインサート時にエラーが発生しました。",
			"Error occurred when DB insert.",
		)
	}
	userEntity.ID = int(createdUserID)
	_, err = tx.Exec(`
		INSERT INTO profiles(
			    user_id, name, thumbnail, bio, gender, phone, place, birth
			)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`, userEntity.ID,
		userEntity.Profile.Name,
		userEntity.Profile.Thumbnail,
		userEntity.Profile.Bio,
		userEntity.Profile.Gender,
		userEntity.Profile.Phone,
		userEntity.Profile.Place,
		userEntity.Profile.Birth,
	)
	if err != nil {
		tlog.PrintLogWithUID(userEntity.AuthID, err)

		return nil, werrors.Newf(
			err,
			http.StatusInternalServerError,
			"DBインサート時にエラーが発生しました。",
			"Error occurred when DB insert.",
		)
	}

	return userEntity, nil
}

func (u *userRepositoryImpliment) SelectByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*entity.User, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintLogWithCtx(ctx, err)
		return nil, werrors.Stack(err)
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
		tlog.PrintLogWithCtx(ctx, err)

		if err == sql.ErrNoRows {
			return nil, werrors.Newf(err, http.StatusInternalServerError, "ユーザが見つかりませんでした。ユーザ登録されているか確認してください。", "User not found. Please make sure signup.")
		}
		return nil, werrors.Wrapf(err, http.StatusInternalServerError, "サーバでエラーが発生しました。", "Error occurred at server.")
	}

	return model.ConvertToUserEntity(&userData), nil
}

func (u *userRepositoryImpliment) SelectByAuthID(ctx context.Context, masterTx repository.MasterTx, authID string) (*entity.User, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintLogWithCtx(ctx, err)
		return nil, werrors.Stack(err)
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
		tlog.PrintLogWithCtx(ctx, err)

		if err == sql.ErrNoRows {
			return nil, werrors.Newf(err, http.StatusUnauthorized, "不正なユーザです。", "Invalid user.")
		}
		return nil, werrors.Wrapf(err, http.StatusInternalServerError, "サーバでエラーが発生しました。", "Error occured at server.")
	}

	return model.ConvertToUserEntity(&userData), nil
}

func (u *userRepositoryImpliment) SelectAll(ctx context.Context, masterTx repository.MasterTx) (entity.UserSlice, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintLogWithCtx(ctx, err)
		return nil, werrors.Stack(err)
	}

	rows, err := tx.Query(`
		SELECT *
		FROM users
		JOIN profiles
		ON users.id = profiles.user_id
	`)
	if err != nil {
		tlog.PrintLogWithCtx(ctx, err)

		if err == sql.ErrNoRows {
			return nil, werrors.Newf(err, http.StatusInternalServerError, "ユーザは1人も登録されていません。", "User doesn't exists.")
		}
		return nil, werrors.Wrapf(err, http.StatusInternalServerError, "サーバでエラーが発生しました。", "Error occured at server.")
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
			tlog.PrintLogWithCtx(ctx, err)
			return nil, werrors.Wrapf(err, http.StatusInternalServerError, "サーバでエラーが発生しました。", "Error occured at server.")
		}
		userSlice = append(userSlice, &userData)
	}

	return model.ConvertToUserSliceEntity(userSlice), nil
}
