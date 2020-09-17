package place

import (
	"context"
	"database/sql"
	"errors"
	placeEntity "wantum/pkg/domain/entity/place"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/place"
	"wantum/pkg/infrastructure/mysql"
	"wantum/pkg/tlog"
	"wantum/pkg/werrors"

	"google.golang.org/grpc/codes"
)

type placeRepositoryImplement struct {
	masterTxManager repository.MasterTxManager
}

func New(txManager repository.MasterTxManager) place.Repository {
	return &placeRepositoryImplement{
		masterTxManager: txManager,
	}
}

func (repo *placeRepositoryImplement) Insert(ctx context.Context, masterTx repository.MasterTx, place *placeEntity.Entity) (int, error) {
	// NOTE: nilで降りてきた用対策。いらないかも
	if err := checkIsNil(place); err != nil {
		return 0, err
	}

	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return 0, werrors.FromConstant(err, werrors.ServerError)
	}
	result, err := tx.Exec(`
		INSERT INTO places(
			name, created_at, updated_at
		) VALUES (?, ?, ?)
	`, place.Name,
		place.CreatedAt,
		place.UpdatedAt,
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

func (repo *placeRepositoryImplement) Update(ctx context.Context, masterTx repository.MasterTx, place *placeEntity.Entity) error {
	// NOTE: nilで降りてきた用対策。いらないかも
	if err := checkIsNil(place); err != nil {
		return err
	}

	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	_, err = tx.Exec(`
		UPDATE places
		SET name=?, updated_at=?
		WHERE id=?
	`, place.Name,
		place.UpdatedAt,
		place.ID,
	)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	return nil
}

func (repo *placeRepositoryImplement) UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, place *placeEntity.Entity) error {
	// NOTE: nilで降りてきた用対策。いらないかも
	if err := checkIsNil(place); err != nil {
		return err
	}
	if place.DeletedAt == nil {
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
		UPDATE places
		SET updated_at=?, deleted_at=?
		WHERE id=?
	`, place.UpdatedAt,
		place.DeletedAt,
		place.ID,
	)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	return nil
}

func (repo *placeRepositoryImplement) DownDeleteFlag(ctx context.Context, masterTx repository.MasterTx, place *placeEntity.Entity) error {
	if err := checkIsNil(place); err != nil {
		return err
	}

	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	_, err = tx.Exec(`
		UPDATE places
		SET updated_at=?, deleted_at=?
		WHERE id=?
	`, place.UpdatedAt,
		nil,
		place.ID,
	)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	return nil
}

func (repo *placeRepositoryImplement) Delete(ctx context.Context, masterTx repository.MasterTx, placeID int) error {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	_, err = tx.Exec(`
		DELETE FROM places
		WHERE id=? and deleted_at is not null
	`, placeID)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	return nil
}

func (repo *placeRepositoryImplement) SelectByID(ctx context.Context, masterTx repository.MasterTx, placeID int) (*placeEntity.Entity, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	row := tx.QueryRow(`
		SELECT id, name, created_at, updated_at, deleted_at
		FROM places
		WHERE id=?
	`, placeID)
	var result placeEntity.Entity
	err = row.Scan(
		&result.ID,
		&result.Name,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt)
	if err != nil {
		// TODO: no rows無くて大丈夫？
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	return &result, nil
}

func (repo *placeRepositoryImplement) SelectAll(ctx context.Context, masterTx repository.MasterTx) (placeEntity.EntitySlice, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	rows, err := tx.Query(`
		SELECT id, name, created_at, updated_at, deleted_at
		FROM places
	`)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	var result placeEntity.EntitySlice
	for rows.Next() {
		var place placeEntity.Entity
		err = rows.Scan(
			&place.ID,
			&place.Name,
			&place.CreatedAt,
			&place.UpdatedAt,
			&place.DeletedAt)
		if err != nil {
			tlog.PrintErrorLogWithCtx(ctx, err)
			return nil, werrors.FromConstant(err, werrors.ServerError)
		}
		result = append(result, &place)
	}
	return result, nil
}

func checkIsNil(place *placeEntity.Entity) error {
	if place == nil {
		return werrors.Newf(
			errors.New("required data is nil"),
			codes.Unknown,
			werrors.ServerError.ErrorCode,
			werrors.ServerError.ErrorMessageJP,
			werrors.ServerError.ErrorMessageEN,
		)
	}
	return nil
}
