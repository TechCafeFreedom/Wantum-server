package user

import (
	"context"
	"database/sql"
	"time"
	userentity "wantum/pkg/domain/entity/user"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/user"
	"wantum/pkg/infrastructure/mysql"
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

func (u *userRepositoryImpliment) InsertUser(masterTx repository.MasterTx, userEntity *userentity.Entity) (*userentity.Entity, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithAuthID(userEntity.AuthID, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	result, err := tx.Exec(`
			INSERT INTO users(
			    auth_id, user_name, mail
			)
			VALUES (?, ?, ?)
	`, userEntity.AuthID, userEntity.UserName, userEntity.Mail)
	if err != nil {
		tlog.PrintErrorLogWithAuthID(userEntity.AuthID, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	createdUserID, err := result.LastInsertId()
	if err != nil {
		tlog.PrintErrorLogWithAuthID(userEntity.AuthID, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	userEntity.ID = int(createdUserID)

	return userEntity, nil
}

func (u *userRepositoryImpliment) SelectByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*userentity.Entity, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	row := tx.QueryRow(`
		SELECT
		       id,
		       auth_id,
		       user_name,
		       mail,
		       created_at,
		       updated_at,
		       deleted_at
		FROM users
		WHERE id = ?
	`, userID)

	userEntity, err := convertToUserEntity(row)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.Stack(err)
	}
	return userEntity, nil
}

func (u *userRepositoryImpliment) SelectByAuthID(ctx context.Context, masterTx repository.MasterTx, authID string) (*userentity.Entity, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	row := tx.QueryRow(`
		SELECT
		    id,
		    auth_id,
		    user_name,
		    mail,
		    created_at,
		    updated_at,
		    deleted_at
		FROM users
		WHERE auth_id = ?
	`, authID)

	userEntity, err := convertToUserEntity(row)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.Stack(err)
	}
	return userEntity, nil
}

func (u *userRepositoryImpliment) SelectAll(ctx context.Context, masterTx repository.MasterTx) (userentity.EntitySlice, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	rows, err := tx.Query(`
		SELECT
		   id,
		   auth_id,
		   user_name,
		   mail,
		   created_at,
		   updated_at,
		   deleted_at
		FROM users
	`)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // 一件もユーザが登録されていない場合は何も返さない
		}
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	userEntitySlice, err := convertToUserSliceEntity(rows)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	return userEntitySlice, nil
}

func convertToUserEntity(row *sql.Row) (*userentity.Entity, error) {
	var userID int
	var authID string
	var userName string
	var mail string
	var createdAt, updatedAt, deletedAt *time.Time

	if err := row.Scan(&userID, &authID, &userName, &mail, &createdAt, &updatedAt, &deletedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, werrors.FromConstant(err, werrors.UserNotFound)
		}
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	return &userentity.Entity{
		ID:        userID,
		AuthID:    authID,
		UserName:  userName,
		Mail:      mail,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}, nil
}

func convertToUserSliceEntity(rows *sql.Rows) (userentity.EntitySlice, error) {
	var userEntitySlice userentity.EntitySlice
	for rows.Next() {
		var userID int
		var authID string
		var userName string
		var mail string
		var createdAt, updatedAt, deletedAt *time.Time

		if err := rows.Scan(&userID, &authID, &userName, &mail, &createdAt, &updatedAt, &deletedAt); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil // 一件もユーザが登録されていない場合は何も返さない
			}
			return nil, werrors.FromConstant(err, werrors.ServerError)
		}

		userEntity := userentity.Entity{
			ID:        userID,
			AuthID:    authID,
			UserName:  userName,
			Mail:      mail,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			DeletedAt: deletedAt,
		}
		userEntitySlice = append(userEntitySlice, &userEntity)
	}

	return userEntitySlice, nil
}
