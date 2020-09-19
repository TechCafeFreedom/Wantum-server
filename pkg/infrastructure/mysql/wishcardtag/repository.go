package wishcardtag

import (
	"context"
	"strconv"
	"strings"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/wishcardtag"
	"wantum/pkg/infrastructure/mysql"
	"wantum/pkg/tlog"
	"wantum/pkg/werrors"
)

type wishCardTagRepositoryImplement struct {
	masterTxManager repository.MasterTxManager
}

func New(txManager repository.MasterTxManager) wishcardtag.Repository {
	return &wishCardTagRepositoryImplement{
		masterTxManager: txManager,
	}
}

func (repo *wishCardTagRepositoryImplement) Insert(ctx context.Context, masterTx repository.MasterTx, wishCardID, tagID int) error {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	_, err = tx.Exec(`
		INSERT INTO wish_cards_tags(wish_card_id, tag_id)
		VALUES (?,?)
	`, wishCardID,
		tagID,
	)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	return nil
}

func (repo *wishCardTagRepositoryImplement) BulkInsert(ctx context.Context, masterTx repository.MasterTx, wishCardID int, tagIDs []int) error {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}

	// TODO: うまいやりかたとは？
	query := "INSERT INTO wish_cards_tags(wish_card_id, tag_id) VALUES "
	values := make([]interface{}, 0, len(tagIDs))
	for _, tagID := range tagIDs {
		query = query + "(?, ?),"
		values = append(values, wishCardID, tagID)
	}
	query = strings.TrimSuffix(query, ",")

	if _, err = tx.Exec(query, values...); err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	return nil
}

func (repo *wishCardTagRepositoryImplement) Delete(ctx context.Context, masterTx repository.MasterTx, wishCardID, tagID int) error {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	_, err = tx.Exec(`
		DELETE FROM wish_cards_tags
		WHERE wish_card_id = ? and tag_id = ?
	`, wishCardID,
		tagID,
	)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	return nil
}

func (repo *wishCardTagRepositoryImplement) DeleteByWishCardID(ctx context.Context, masterTx repository.MasterTx, wishCardID int) error {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	_, err = tx.Exec(`
		DELETE FROM wish_cards_tags
		WHERE wish_card_id = ?
	`, wishCardID,
	)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	return nil
}

func (repo *wishCardTagRepositoryImplement) DeleteByIDs(ctx context.Context, masterTx repository.MasterTx, wishCardID int, tagIDs []int) error {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}

	tagIDsStr := make([]string, 0, len(tagIDs))
	for _, id := range tagIDs {
		tagIDsStr = append(tagIDsStr, strconv.Itoa(id))
	}

	_, err = tx.Exec(`
		DELETE FROM wish_cards_tags
		WHERE wish_card_id = ?
			AND tag_id IN (`+strings.Join(tagIDsStr, ",")+`)
	`, wishCardID,
	)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	return nil
}
