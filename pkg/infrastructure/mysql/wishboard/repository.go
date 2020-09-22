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

func (r *repositoryImpliment) Insert(ctx context.Context, masterTx repository.MasterTx, title, backgroundImageUrl, inviteUrl string, userID int) (*wishboard.Entity, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	now := time.Now()

	result, err := tx.Exec(`
		INSERT INTO wish_boards(
			title, background_image_url, invite_url, user_id, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?)
	`, title, backgroundImageUrl, inviteUrl, userID, now, now)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	insertID, err := result.LastInsertId()
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	return &wishboard.Entity{
		ID:                 int(insertID),
		Title:              title,
		BackgroundImageUrl: backgroundImageUrl,
		InviteUrl:          inviteUrl,
		UserID:             userID,
		CreatedAt:          now,
		UpdatedAt:          now,
		DeletedAt:          nil,
	}, nil
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

	b := wishboard.Entity{}
	err = row.Scan(
		&b.ID, &b.Title, &b.BackgroundImageUrl, &b.InviteUrl, &b.UserID, &b.CreatedAt, &b.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, werrors.FromConstant(err, werrors.WishBoardNotFound)
		}
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	return &b, nil
}

func (r *repositoryImpliment) SelectByPKs(ctx context.Context, masterTx repository.MasterTx, wishBoardIDs []int) (wishboard.EntitySlice, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

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
			return []*wishboard.Entity{}, nil
		}
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	bs := make(wishboard.EntitySlice, 0, 4)
	for rows.Next() {
		b := wishboard.Entity{}
		err := rows.Scan(
			&b.ID, &b.Title, &b.BackgroundImageUrl, &b.InviteUrl, &b.UserID, &b.CreatedAt, &b.UpdatedAt)

		if err != nil {
			if err == sql.ErrNoRows {
				return []*wishboard.Entity{}, nil
			}
			return nil, werrors.FromConstant(err, werrors.ServerError)
		}

		bs = append(bs, &b)
	}

	return bs, nil
}

func (r *repositoryImpliment) SelectByUserID(ctx context.Context, masterTx repository.MasterTx, userID int) (wishboard.EntitySlice, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	rows, err := tx.Query(`
		SELECT
			id, title, background_image_url, invite_url, user_id, created_at, updated_at
		FROM wish_boards
		WHERE user_id = ? AND deleted_at IS NULL
	`, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return []*wishboard.Entity{}, nil
		}
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	bs := make(wishboard.EntitySlice, 0, 4)
	for rows.Next() {
		b := wishboard.Entity{}
		err := rows.Scan(
			&b.ID, &b.Title, &b.BackgroundImageUrl, &b.InviteUrl, &b.UserID, &b.CreatedAt, &b.UpdatedAt)

		if err != nil {
			if err == sql.ErrNoRows {
				return []*wishboard.Entity{}, nil
			}
			return nil, werrors.FromConstant(err, werrors.ServerError)
		}

		bs = append(bs, &b)
	}

	return bs, nil
}

func (r *repositoryImpliment) UpdateTitle(ctx context.Context, masterTx repository.MasterTx, wishBoardID int, title string) error {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}

	now := time.Now()

	_, err = tx.Exec(`
		UPDATE wish_boards SET
			title=?,
			updated_at=?
		WHERE id = ?
	`, title, now, wishBoardID)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}

	return nil
}

func (r *repositoryImpliment) UpdateBackgroundImageUrl(ctx context.Context, masterTx repository.MasterTx, wishBoardID int, backgroundImageUrl string) error {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}

	now := time.Now()

	_, err = tx.Exec(`
		UPDATE wish_boards SET
			background_image_url=?,
			updated_at=?
		WHERE id = ?
	`, backgroundImageUrl, now, wishBoardID)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}

	return nil
}

func (r *repositoryImpliment) Delete(ctx context.Context, masterTx repository.MasterTx, wishBoardID int) error {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}

	now := time.Now()

	_, err = tx.Exec(`
		UPDATE wish_boards SET
			updated_at=?,
			deleted_at=?
		WHERE id = ?
	`, now, now, wishBoardID)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}

	return nil
}
