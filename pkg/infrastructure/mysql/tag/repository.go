package tag

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"strings"
	tagEntity "wantum/pkg/domain/entity/tag"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/tag"
	"wantum/pkg/infrastructure/mysql"
	"wantum/pkg/tlog"
	"wantum/pkg/werrors"

	"google.golang.org/grpc/codes"
)

type tagRepositoryImplement struct {
	masterTxManager repository.MasterTxManager
}

func New(txManager repository.MasterTxManager) tag.Repository {
	return &tagRepositoryImplement{
		masterTxManager: txManager,
	}
}

func (repo *tagRepositoryImplement) Insert(ctx context.Context, masterTx repository.MasterTx, tag *tagEntity.Entity) (int, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return 0, werrors.FromConstant(err, werrors.ServerError)
	}
	result, err := tx.Exec(`
		INSERT INTO tags(name, created_at, updated_at)
		VALUES (?,?,?)
	`, tag.Name,
		tag.CreatedAt,
		tag.UpdatedAt,
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

func (repo *tagRepositoryImplement) UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, tag *tagEntity.Entity) error {
	if tag.DeletedAt == nil {
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
		UPDATE tags
		SET updated_at=?, deleted_at=?
		WHERE id=?
	`, tag.UpdatedAt,
		tag.DeletedAt,
		tag.ID,
	)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	return nil
}

func (repo *tagRepositoryImplement) DownDeleteFlag(ctx context.Context, masterTx repository.MasterTx, tag *tagEntity.Entity) error {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	_, err = tx.Exec(`
		UPDATE tags
		SET updated_at=?, deleted_at=?
		WHERE id=?
	`, tag.UpdatedAt,
		nil,
		tag.ID,
	)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	return nil
}

func (repo *tagRepositoryImplement) Delete(ctx context.Context, masterTx repository.MasterTx, tagID int) error {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	_, err = tx.Exec(`
		DELETE FROM tags
		WHERE id=? and deleted_at is not null
	`, tagID)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.ServerError)
	}
	return nil
}

func (repo *tagRepositoryImplement) SelectByID(ctx context.Context, masterTx repository.MasterTx, tagID int) (*tagEntity.Entity, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	row := tx.QueryRow(`
		SELECT id, name, created_at, updated_at, deleted_at
		FROM tags
		WHERE id=?
	`, tagID)
	result, err := convertToTagEntity(row)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	return result, nil
}

func (repo *tagRepositoryImplement) SelectByIDs(ctx context.Context, masterTx repository.MasterTx, tagIDs []int) (tagEntity.EntitySlice, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	tagIDsStr := make([]string, 0, len(tagIDs))
	for _, id := range tagIDs {
		tagIDsStr = append(tagIDsStr, strconv.Itoa(id))
	}

	rows, err := tx.Query(`
		SELECT id, name, created_at, updated_at, deleted_at
		FROM tags
		WHERE id
		IN (` + strings.Join(tagIDsStr, ",") + `)
	`)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	result, err := convertToTagEntitySlice(rows)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	return result, nil
}

func (repo *tagRepositoryImplement) SelectByName(ctx context.Context, masterTx repository.MasterTx, name string) (*tagEntity.Entity, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	row := tx.QueryRow(`
		SELECT id, name, created_at, updated_at, deleted_at
		FROM tags
		WHERE name=?
	`, name)
	result, err := convertToTagEntity(row)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	return result, nil
}

func (repo *tagRepositoryImplement) SelectByWishCardID(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (tagEntity.EntitySlice, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	rows, err := tx.Query(`
		SELECT tags.id, tags.name, tags.created_at, tags.updated_at, tags.deleted_at
		FROM wish_cards_tags as r
		INNER JOIN tags ON tags.id = r.tag_id
		WHERE r.wish_card_id=?
	`, wishCardID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	result, err := convertToTagEntitySlice(rows)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	return result, nil
}

func (repo *tagRepositoryImplement) SelectByMemoryID(ctx context.Context, masterTx repository.MasterTx, memoryID int) (tagEntity.EntitySlice, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	rows, err := tx.Query(`
		SELECT tags.id, tags.name, tags.created_at, tags.updated_at, tags.deleted_at
		FROM memories_tags as r
		INNER JOIN tags ON tags.id = r.tag_id
		WHERE r.memory_id=?
	`, memoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	result, err := convertToTagEntitySlice(rows)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	return result, nil
}

func convertToTagEntity(row *sql.Row) (*tagEntity.Entity, error) {
	var tag tagEntity.Entity
	if err := row.Scan(&tag.ID, &tag.Name, &tag.CreatedAt, &tag.UpdatedAt, &tag.DeletedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	return &tag, nil
}

func convertToTagEntitySlice(rows *sql.Rows) (tagEntity.EntitySlice, error) {
	var tags tagEntity.EntitySlice
	for rows.Next() {
		var tag tagEntity.Entity
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.CreatedAt, &tag.UpdatedAt, &tag.DeletedAt); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, werrors.FromConstant(err, werrors.ServerError)
		}
		tags = append(tags, &tag)
	}
	return tags, nil
}
