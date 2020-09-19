package wishcard

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"
	placeEntity "wantum/pkg/domain/entity/place"
	userEntity "wantum/pkg/domain/entity/user"
	wishCardEntity "wantum/pkg/domain/entity/wishcard"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/wishcard"
	"wantum/pkg/infrastructure/mysql"
	"wantum/pkg/tlog"
	"wantum/pkg/werrors"

	"google.golang.org/grpc/codes"
)

type wishCardRepositoryImplement struct {
	masterTxManager repository.MasterTxManager
}

func New(txManager repository.MasterTxManager) wishcard.Repository {
	return &wishCardRepositoryImplement{
		masterTxManager: txManager,
	}
}

func (repo *wishCardRepositoryImplement) Insert(ctx context.Context, masterTx repository.MasterTx, wishCard *wishCardEntity.Entity, categoryID int) (int, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return 0, werrors.FromConstant(err, werrors.ServerError)
	}
	result, err := tx.Exec(`
		INSERT INTO wish_cards(
			user_id, activity, description, date, created_at, updated_at, category_id, place_id
		) VALUES (?,?,?,?,?,?,?,?)
	`, wishCard.Author.ID,
		wishCard.Activity,
		wishCard.Description,
		wishCard.Date,
		wishCard.CreatedAt,
		wishCard.UpdatedAt,
		categoryID,
		wishCard.Place.ID,
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

func (repo *wishCardRepositoryImplement) Update(ctx context.Context, masterTx repository.MasterTx, wishCard *wishCardEntity.Entity, categoryID int) error {
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
	`, wishCard.Author.ID,
		wishCard.Activity,
		wishCard.Description,
		wishCard.Date,
		wishCard.DoneAt,
		wishCard.UpdatedAt,
		categoryID,
		wishCard.Place.ID,
		wishCard.ID,
	)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	return nil
}

func (repo *wishCardRepositoryImplement) UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, wishCard *wishCardEntity.Entity) error {
	if wishCard.DeletedAt == nil {
		return werrors.Newf(
			errors.New("can't up delete flag. deletedAt is nil"),
			codes.Internal,
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

func (repo *wishCardRepositoryImplement) DownDeleteFlag(ctx context.Context, masterTx repository.MasterTx, wishCard *wishCardEntity.Entity) error {
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

func (repo *wishCardRepositoryImplement) SelectByID(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (*wishCardEntity.Entity, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	row := tx.QueryRow(`
		SELECT id, user_id, activity, description, date, done_at, created_at, updated_at, deleted_at, place_id
		FROM wish_cards
		WHERE id=?
	`, wishCardID)
	var result wishCardEntity.Entity
	var place placeEntity.Entity
	var user userEntity.Entity
	err = row.Scan(
		&result.ID,
		&user.ID,
		&result.Activity,
		&result.Description,
		&result.Date,
		&result.DoneAt,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
		&place.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	result.Author = &user
	result.Place = &place
	log.Println(result.Place)
	return &result, nil
}

func (repo *wishCardRepositoryImplement) SelectByIDs(ctx context.Context, masterTx repository.MasterTx, wishCardIDs []string) (wishCardEntity.EntitySlice, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	rows, err := tx.Query(`
		SELECT id, user_id, activity, description, date, done_at, created_at, updated_at, deleted_at, place_id
		FROM wish_cards
		WHERE id
		IN (` + strings.Join(wishCardIDs, ",") + `)
	`)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	var result wishCardEntity.EntitySlice
	for rows.Next() {
		var wishCard wishCardEntity.Entity
		var place placeEntity.Entity
		var user userEntity.Entity
		err = rows.Scan(
			&wishCard.ID,
			&user.ID,
			&wishCard.Activity,
			&wishCard.Description,
			&wishCard.Date,
			&wishCard.DoneAt,
			&wishCard.CreatedAt,
			&wishCard.UpdatedAt,
			&wishCard.DeletedAt,
			&place.ID,
		)
		if err != nil {
			if err != sql.ErrNoRows {
				return nil, nil
			}
			tlog.PrintErrorLogWithCtx(ctx, err)
			return nil, werrors.FromConstant(err, werrors.ServerError)
		}
		wishCard.Author = &user
		wishCard.Place = &place
		result = append(result, &wishCard)
	}
	return result, nil
}

func (repo *wishCardRepositoryImplement) SelectByCategoryID(ctx context.Context, masterTx repository.MasterTx, categryID int) (wishCardEntity.EntitySlice, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	rows, err := tx.Query(`
		SELECT id, user_id, activity, description, date, done_at, created_at, updated_at, deleted_at, place_id
		FROM wish_cards
		WHERE category_id=?
	`, categryID)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	var result wishCardEntity.EntitySlice
	for rows.Next() {
		var wishCard wishCardEntity.Entity
		var place placeEntity.Entity
		var user userEntity.Entity
		err = rows.Scan(
			&wishCard.ID,
			&user.ID,
			&wishCard.Activity,
			&wishCard.Description,
			&wishCard.Date,
			&wishCard.DoneAt,
			&wishCard.CreatedAt,
			&wishCard.UpdatedAt,
			&wishCard.DeletedAt,
			&place.ID,
		)
		if err != nil {
			if err != sql.ErrNoRows {
				return nil, nil
			}
			tlog.PrintErrorLogWithCtx(ctx, err)
			return nil, werrors.FromConstant(err, werrors.ServerError)
		}
		wishCard.Author = &user
		wishCard.Place = &place
		result = append(result, &wishCard)
	}
	return result, nil
}
