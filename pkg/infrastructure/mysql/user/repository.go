package user

import (
	"context"
	"database/sql"
	"net/http"
	"time"
	"wantum/pkg/domain/entity"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/user"
	"wantum/pkg/infrastructure/mysql"
	"wantum/pkg/tlog"
	"wantum/pkg/werrors"
)

type userModel struct {
	id        int
	uid       string
	name      string
	thumbnail string
	createdAt time.Time
	updatedAt time.Time
}

type userModelSlice []*userModel

type userRepositoryImpliment struct {
	masterTxManager repository.MasterTxManager
}

func New(masterTxManager repository.MasterTxManager) user.Repository {
	return &userRepositoryImpliment{
		masterTxManager: masterTxManager,
	}
}

func (u *userRepositoryImpliment) InsertUser(masterTx repository.MasterTx, uid, name, thumbnail string) error {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintLogWithUID(uid, err)
		return werrors.Stack(err)
	}
	if _, err := tx.Exec("INSERT INTO users(uid, name, thumbnail) VALUES (?, ?, ?)", uid, name, thumbnail); err != nil {
		tlog.PrintLogWithUID(uid, err)

		return werrors.Newf(
			err,
			http.StatusInternalServerError,
			"DBインサート時にエラーが発生しました。",
			"Error occurred when DB insert.",
		)
	}

	return nil
}

func (u *userRepositoryImpliment) SelectByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*entity.User, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintLogWithCtx(ctx, err)
		return nil, werrors.Stack(err)
	}

	var userData userModel
	row := tx.QueryRow("SELECT * FROM users WHERE id = ?", userID)
	if err := row.Scan(&userData.id, &userData.uid, &userData.name, &userData.thumbnail, &userData.createdAt, &userData.updatedAt); err != nil {
		tlog.PrintLogWithCtx(ctx, err)

		if err == sql.ErrNoRows {
			return nil, werrors.Newf(err, http.StatusInternalServerError, "ユーザが見つかりませんでした。ユーザ登録されているか確認してください。", "User not found. Please make sure signup.")
		}
		return nil, werrors.Wrapf(err, http.StatusInternalServerError, "サーバでエラーが発生しました。", "Error occurred at server.")
	}

	return ConvertToUserEntity(&userData), nil
}

func (u *userRepositoryImpliment) SelectByUID(ctx context.Context, masterTx repository.MasterTx, uid string) (*entity.User, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintLogWithCtx(ctx, err)
		return nil, werrors.Stack(err)
	}

	var userData userModel
	row := tx.QueryRow("SELECT * FROM users WHERE uid = ?", uid)
	if err := row.Scan(&userData.id, &userData.uid, &userData.name, &userData.thumbnail, &userData.createdAt, &userData.updatedAt); err != nil {
		tlog.PrintLogWithCtx(ctx, err)

		if err == sql.ErrNoRows {
			return nil, werrors.Newf(err, http.StatusUnauthorized, "不正なユーザです。", "Invalid user.")
		}
		return nil, werrors.Wrapf(err, http.StatusInternalServerError, "サーバでエラーが発生しました。", "Error occured at server.")
	}

	return ConvertToUserEntity(&userData), nil
}

func (u *userRepositoryImpliment) SelectAll(ctx context.Context, masterTx repository.MasterTx) (entity.UserSlice, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintLogWithCtx(ctx, err)
		return nil, werrors.Stack(err)
	}

	rows, err := tx.Query("SELECT * FROM users")
	if err != nil {
		tlog.PrintLogWithCtx(ctx, err)

		if err == sql.ErrNoRows {
			return nil, werrors.Newf(err, http.StatusInternalServerError, "ユーザは1人も登録されていません。", "User doesn't exists.")
		}
		return nil, werrors.Wrapf(err, http.StatusInternalServerError, "サーバでエラーが発生しました。", "Error occured at server.")
	}

	var userSlice userModelSlice
	for rows.Next() {
		var userData userModel
		if err := rows.Scan(&userData.id, &userData.uid, &userData.name, &userData.thumbnail, &userData.createdAt, &userData.updatedAt); err != nil {
			tlog.PrintLogWithCtx(ctx, err)
			return nil, werrors.Wrapf(err, http.StatusInternalServerError, "サーバでエラーが発生しました。", "Error occured at server.")
		}
		userSlice = append(userSlice, &userData)
	}

	return ConvertToUserSliceEntity(userSlice), nil
}

func ConvertToUserEntity(userData *userModel) *entity.User {
	return &entity.User{
		ID:        userData.id,
		Name:      userData.name,
		Thumbnail: userData.thumbnail,
	}
}

func ConvertToUserSliceEntity(userSlice userModelSlice) entity.UserSlice {
	res := make(entity.UserSlice, 0, len(userSlice))
	for _, userData := range userSlice {
		res = append(res, ConvertToUserEntity(userData))
	}
	return res
}
