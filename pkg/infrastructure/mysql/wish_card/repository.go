package wish_card

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/wish_card"
	"wantum/pkg/infrastructure/mysql"
	"wantum/pkg/infrastructure/mysql/model"
	"wantum/pkg/tlog"
	"wantum/pkg/werrors"
)

type wishCardRepositoryImplement struct {
	masterTxManager repository.MasterTxManager
}

func New(txManager repository.MasterTxManager) wish_card.Repository {
	return &wishCardRepositoryImplement{
		masterTxManager: txManager,
	}
}

func (repo *wishCardRepositoryImplement) Insert(ctx context.Context, masterTx repository.MasterTx, wishCard *model.WishCardModel) (int, error) {
	if err := checkIsNil(wishCard); err != nil {
		return 0, err
	}

	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return 0, werrors.FromConstant(err, werrors.ServerError)
	}
	result, err := tx.Exec(`
		INSERT INTO wish_cards(
			user_id, activity, description, date, created_at, updated_at, category_id, place_id
		) VALUES (?,?,?,?,?,?,?,?)
	`, wishCard.UserID,
		wishCard.Activity,
		wishCard.Description,
		wishCard.Date,
		wishCard.CreatedAt,
		wishCard.UpdatedAt,
		wishCard.CategoryID,
		wishCard.PlaceID,
	)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return 0, werrors.FromConstant(err, werrors.ServerError)
	}
	id, err := result.LastInsertId()
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return 0, werrors.FromConstant(err, werrors.ServerError)
	}
	return int(id), nil
}

func (repo *wishCardRepositoryImplement) Update(ctx context.Context, masterTx repository.MasterTx, wishCard *model.WishCardModel) error {
	if err := checkIsNil(wishCard); err != nil {
		return err
	}

	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	_, err = tx.Exec(`
		UPDATE wish_cards
		SET
			user_id=?,
			activity=?,
			description=?,
			date=?,
			done_at=?,
			updated_at=?,
			category_id=?,
			place_id=?
		WHERE id=?
	`, wishCard.UserID,
		wishCard.Activity,
		wishCard.Description,
		wishCard.Date,
		wishCard.DoneAt,
		wishCard.UpdatedAt,
		wishCard.CategoryID,
		wishCard.PlaceID,
		wishCard.ID,
	)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	return nil
}

func (repo *wishCardRepositoryImplement) UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, wishCard *model.WishCardModel) error {
	if err := checkIsNil(wishCard); err != nil {
		return err
	}
	if wishCard.DeletedAt == nil {
		return werrors.Newf(
			errors.New("deletedAt is nil"),
			werrors.ServerError.ErrorCode,
			werrors.ServerError.ErrorMessageJP,
			werrors.ServerError.ErrorMessageEN,
		)
	}
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	_, err = tx.Exec(`
		UPDATE wish_cards
		SET updated_at=?, deleted_at=?
		WHERE id=?
	`, wishCard.UpdatedAt,
		wishCard.DeletedAt,
		wishCard.ID,
	)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	return nil
}

func (repo *wishCardRepositoryImplement) DownDeleteFlag(ctx context.Context, masterTx repository.MasterTx, wishCard *model.WishCardModel) error {
	if err := checkIsNil(wishCard); err != nil {
		return err
	}
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	_, err = tx.Exec(`
		UPDATE wish_cards
		SET updated_at=?, deleted_at=?
		WHERE id=?
	`, wishCard.UpdatedAt,
		nil,
		wishCard.ID,
	)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	return nil
}

func (repo *wishCardRepositoryImplement) Delete(ctx context.Context, masterTx repository.MasterTx, wishCardID int) error {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	_, err = tx.Exec(`
		DELETE FROM wish_cards
		WHERE id=? and deleted_at is not null
	`, wishCardID)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	return nil
}

func (repo *wishCardRepositoryImplement) SelectByID(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (*model.WishCardModel, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	row := tx.QueryRow(`
		SELECT id, user_id, activity, description, date, done_at, created_at, updated_at, deleted_at, category_id, place_id
		FROM wish_cards
		WHERE id=?
	`, wishCardID)
	var result model.WishCardModel
	err = row.Scan(
		&result.ID,
		&result.UserID,
		&result.Activity,
		&result.Description,
		&result.Date,
		&result.DoneAt,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
		&result.CategoryID,
		&result.PlaceID)
	if err != nil {
		// TODO: これってno rowsでもえらーでおっけえなの？
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	return &result, nil
}

func (repo *wishCardRepositoryImplement) SelectByIDs(ctx context.Context, masterTx repository.MasterTx, wishCardIDs []string) (model.WishCardModelSlice, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	rows, err := tx.Query(`
		SELECT id, user_id, activity, description, date, done_at, created_at, updated_at, deleted_at, category_id, place_id
		FROM wish_cards
		WHERE id
		IN (` + strings.Join(wishCardIDs, ",") + `)
	`)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	var result model.WishCardModelSlice
	for rows.Next() {
		var record model.WishCardModel
		err = rows.Scan(
			&record.ID,
			&record.UserID,
			&record.Activity,
			&record.Description,
			&record.Date,
			&record.DoneAt,
			&record.CreatedAt,
			&record.UpdatedAt,
			&record.DeletedAt,
			&record.CategoryID,
			&record.PlaceID,
		)
		if err != nil {
			if err != sql.ErrNoRows {
				return nil, nil
			}
			tlog.PrintErrorLogWithCtx(ctx, err)
			return nil, werrors.FromConstant(err, werrors.ServerError)
		}
		result = append(result, &record)
	}
	return result, nil
}

func (repo *wishCardRepositoryImplement) SelectByCategoryID(ctx context.Context, masterTx repository.MasterTx, categryID int) (model.WishCardModelSlice, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	rows, err := tx.Query(`
		SELECT id, user_id, activity, description, date, done_at, created_at, updated_at, deleted_at, category_id, place_id
		FROM wish_cards
		WHERE category_id=?
	`, categryID)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	var result model.WishCardModelSlice
	for rows.Next() {
		var record model.WishCardModel
		err = rows.Scan(
			&record.ID,
			&record.UserID,
			&record.Activity,
			&record.Description,
			&record.Date,
			&record.DoneAt,
			&record.CreatedAt,
			&record.UpdatedAt,
			&record.DeletedAt,
			&record.CategoryID,
			&record.PlaceID,
		)
		if err != nil {
			if err != sql.ErrNoRows {
				return nil, nil
			}
			tlog.PrintErrorLogWithCtx(ctx, err)
			return nil, werrors.FromConstant(err, werrors.ServerError)
		}
		result = append(result, &record)
	}
	return result, nil
}

func checkIsNil(wishCard *model.WishCardModel) error {
	if wishCard == nil {
		return werrors.Newf(
			errors.New("required data(wishCard) is nil"),
			werrors.ServerError.ErrorCode,
			werrors.ServerError.ErrorMessageJP,
			werrors.ServerError.ErrorMessageEN,
		)
	}
	return nil
}
