package wishboard

import (
	"bytes"
	"context"
	"database/sql"
	"strconv"
	"time"
	"wantum/pkg/domain/entity/wishboard"
	"wantum/pkg/domain/repository"
	wishboardrepository "wantum/pkg/domain/repository/wishboard"
	"wantum/pkg/infrastructure/mysql"
	"wantum/pkg/tlog"
	"wantum/pkg/werrors"
)

type repositoryImpliment struct {
	masterTxManager repository.MasterTxManager
}

func New(masterTxManager repository.MasterTxManager) wishboardrepository.Repository {
	return &repositoryImpliment{
		masterTxManager: masterTxManager,
	}
}

func (r *repositoryImpliment) Insert(ctx context.Context, masterTx repository.MasterTx, wishBoardEntity *wishboard.Entity) (*wishboard.Entity, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	// CreatedAt, UpdatedAtの実体を渡す
	result, err := tx.Exec(`
		INSERT INTO wish_boards(
			title, background_image_url, invite_url, user_id, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?)
	`, wishBoardEntity.Title, wishBoardEntity.BackgroundImageURL, wishBoardEntity.InviteURL, wishBoardEntity.UserID, *wishBoardEntity.CreatedAt, *wishBoardEntity.UpdatedAt)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	insertID, err := result.LastInsertId()
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	wishBoardEntity.ID = int(insertID)

	return wishBoardEntity, nil
}

func (r *repositoryImpliment) SelectByPK(ctx context.Context, masterTx repository.MasterTx, wishBoardID int) (*wishboard.Entity, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	row := tx.QueryRow(`
		SELECT
			id, title, background_image_url, invite_url, user_id, created_at, updated_at
		FROM wish_boards
		WHERE id = ? AND deleted_at IS NULL
	`, wishBoardID)

	// ポインタ型のフィールドについては、あらかじめメモリ確保する
	wishBoardEntity := wishboard.Entity{CreatedAt: &time.Time{}, UpdatedAt: &time.Time{}}
	// CreatedAt, UpdatedAtはポインタなので「&」はつけずにそのまま渡す
	err = row.Scan(
		&wishBoardEntity.ID, &wishBoardEntity.Title, &wishBoardEntity.BackgroundImageURL, &wishBoardEntity.InviteURL, &wishBoardEntity.UserID, wishBoardEntity.CreatedAt, wishBoardEntity.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, werrors.FromConstant(err, werrors.WishBoardNotFound)
		}
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	return &wishBoardEntity, nil
}

func (r *repositoryImpliment) SelectByPKs(ctx context.Context, masterTx repository.MasterTx, wishBoardIDs []int) (wishboard.EntitySlice, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	// IDのリストからSQL文のIN句用の文字列を作成
	var buf bytes.Buffer
	for i, wishBoardID := range wishBoardIDs {
		if i == 0 {
			buf.WriteString(strconv.Itoa(wishBoardID))
		} else {
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(wishBoardID))
		}
	}

	rows, err := tx.Query(`
		SELECT
			id, title, background_image_url, invite_url, user_id, created_at, updated_at
		FROM wish_boards
		WHERE id IN (` + buf.String() +
		`) AND deleted_at IS NULL
	`)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, werrors.FromConstant(err, werrors.WishBoardNotFound)
		}
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	wishBoardEntitySlice := wishboard.EntitySlice{}
	for rows.Next() {
		// ポインタ型のフィールドについては、あらかじめメモリ確保する
		wishBoardEntity := wishboard.Entity{CreatedAt: &time.Time{}, UpdatedAt: &time.Time{}}
		// CreatedAt, UpdatedAtはポインタなので「&」はつけずにそのまま渡す
		err := rows.Scan(
			&wishBoardEntity.ID, &wishBoardEntity.Title, &wishBoardEntity.BackgroundImageURL, &wishBoardEntity.InviteURL, &wishBoardEntity.UserID, wishBoardEntity.CreatedAt, wishBoardEntity.UpdatedAt)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil, werrors.FromConstant(err, werrors.WishBoardNotFound)
			}
			return nil, werrors.FromConstant(err, werrors.ServerError)
		}

		wishBoardEntitySlice = append(wishBoardEntitySlice, &wishBoardEntity)
	}

	return wishBoardEntitySlice, nil
}

func (r *repositoryImpliment) UpdateTitle(ctx context.Context, masterTx repository.MasterTx, wishBoardID int, title string, updatedAt *time.Time) error {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}

	// UpdatedAtの実体を渡す
	_, err = tx.Exec(`
		UPDATE wish_boards SET
			title=?,
			updated_at=?
		WHERE id = ?
	`, title, *updatedAt, wishBoardID)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}

	return nil
}

func (r *repositoryImpliment) UpdateBackgroundImageURL(ctx context.Context, masterTx repository.MasterTx, wishBoardID int, backgroundImageURL string, updatedAt *time.Time) error {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}

	// UpdatedAtの実体を渡す
	_, err = tx.Exec(`
		UPDATE wish_boards SET
			background_image_url=?,
			updated_at=?
		WHERE id = ?
	`, backgroundImageURL, *updatedAt, wishBoardID)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}

	return nil
}

func (r *repositoryImpliment) Delete(ctx context.Context, masterTx repository.MasterTx, wishBoardEntity *wishboard.Entity) error {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}

	// UpdatedAt, DeletedAtの実体を渡す
	_, err = tx.Exec(`
		UPDATE wish_boards SET
			updated_at=?,
			deleted_at=?
		WHERE id = ?
	`, *wishBoardEntity.UpdatedAt, *wishBoardEntity.DeletedAt, wishBoardEntity.ID)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}

	return nil
}
