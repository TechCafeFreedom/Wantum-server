package userwishboard

import (
	"context"
	"database/sql"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/userwishboard"
	"wantum/pkg/infrastructure/mysql"
	"wantum/pkg/tlog"
	"wantum/pkg/werrors"
)

type repositoryImpliment struct {
	masterTxManager repository.MasterTxManager
}

func New(masterTxManager repository.MasterTxManager) userwishboard.Repository {
	return &repositoryImpliment{
		masterTxManager: masterTxManager,
	}
}

func (r *repositoryImpliment) Insert(ctx context.Context, masterTx repository.MasterTx, userID, wishBoardID int) error {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}

	_, err = tx.Exec(`
		INSERT INTO users_wish_boards(
			user_id, wish_board_id
		) VALUES (?, ?)
	`, userID, wishBoardID)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}

	return nil
}

func (r *repositoryImpliment) Exists(ctx context.Context, masterTx repository.MasterTx, userID, wishBoardID int) (bool, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return false, werrors.FromConstant(err, werrors.ServerError)
	}

	row := tx.QueryRow(`
		SELECT id FROM users_wish_boards WHERE user_id = ? AND wish_board_id = ?
	`, userID, wishBoardID)

	var i int
	err = row.Scan(&i)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, werrors.FromConstant(err, werrors.ServerError)
	}

	return true, nil
}

func (r *repositoryImpliment) SelectByUserID(ctx context.Context, masterTx repository.MasterTx, userID int) ([]int, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	rows, err := tx.Query(`
		SELECT wish_board_id FROM users_wish_boards WHERE user_id = ?
	`, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	wishBoardIDSlice := []int{}
	for rows.Next() {
		var wishBoardID int
		if err := rows.Scan(&wishBoardID); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, werrors.FromConstant(err, werrors.ServerError)
		}

		wishBoardIDSlice = append(wishBoardIDSlice, wishBoardID)
	}

	return wishBoardIDSlice, nil
}

func (r *repositoryImpliment) Delete(ctx context.Context, masterTx repository.MasterTx, userID, wishBoardID int) error {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}

	_, err = tx.Exec(`
		DELETE FROM users_wish_boards WHERE user_id = ? AND wish_board_id = ?
	`, userID, wishBoardID)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}

	return nil
}
