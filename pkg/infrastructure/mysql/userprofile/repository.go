package userprofile

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
	"time"
	"wantum/pkg/domain/entity/userprofile"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/profile"
	"wantum/pkg/infrastructure/mysql"
	"wantum/pkg/tlog"
	"wantum/pkg/werrors"
)

type profileRepositoryImpliment struct {
	masterTxManager repository.MasterTxManager
}

func New(masterTxManager repository.MasterTxManager) profile.Repository {
	return &profileRepositoryImpliment{
		masterTxManager: masterTxManager,
	}
}

func (p *profileRepositoryImpliment) InsertProfile(ctx context.Context, masterTx repository.MasterTx, userProfileEntity *userprofile.Entity) (*userprofile.Entity, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	_, err = tx.Exec(`
		INSERT INTO profiles(
			user_id, name, thumbnail, bio, gender, phone, place, birth
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`, userProfileEntity.UserID,
		userProfileEntity.Name,
		userProfileEntity.Thumbnail,
		userProfileEntity.Bio,
		userProfileEntity.Gender,
		userProfileEntity.Phone,
		userProfileEntity.Place,
		userProfileEntity.Birth,
	)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	return userProfileEntity, nil
}

func (p *profileRepositoryImpliment) SelectByUserID(ctx context.Context, masterTx repository.MasterTx, userID int) (*userprofile.Entity, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}
	row := tx.QueryRow(`
		SELECT
			id, user_id, name, thumbnail, bio, gender, phone, place, birth, created_at, updated_at, deleted_at
		FROM profiles
		WHERE user_id = ?
	`, userID,
	)
	userProfileEntity, err := convertToUserProfileEntity(row)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.Stack(err)
	}
	return userProfileEntity, nil
}

func (p *profileRepositoryImpliment) SelectByUserIDs(ctx context.Context, masterTx repository.MasterTx, userIDs []int) (userprofile.EntitySlice, error) {
	tx, err := mysql.ExtractTx(masterTx)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	userIDsStr := make([]string, 0, len(userIDs))
	for _, userID := range userIDs {
		userIDsStr = append(userIDsStr, strconv.Itoa(userID))
	}

	rows, err := tx.Query(`
		SELECT
			id, user_id, name, thumbnail, bio, gender, phone, place, birth, created_at, updated_at, deleted_at
		FROM profiles
		WHERE user_id IN (` + strings.Join(userIDsStr, ",") + `)
	`)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)

		if err == sql.ErrNoRows {
			return nil, nil // 一件もユーザが登録されていない場合は何も返さない
		}
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	userProfileEntitySlice, err := convertToUserProfileSliceEntity(rows)
	if err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.Stack(err)
	}
	return userProfileEntitySlice, nil
}

func convertToUserProfileEntity(row *sql.Row) (*userprofile.Entity, error) {
	var userID int
	var name string
	var thumbnail string
	var bio string
	var gender int
	var phone string
	var place string
	var birth *time.Time
	var createdAt, updatedAt, deletedAt *time.Time

	if err := row.Scan(&userID, &name, &thumbnail, &bio, &gender, &phone, &place, &birth, &createdAt, &updatedAt, &deletedAt); err != nil {
		return nil, werrors.FromConstant(err, werrors.ServerError)
	}

	return &userprofile.Entity{
		UserID:    userID,
		Name:      name,
		Thumbnail: thumbnail,
		Bio:       bio,
		Gender:    gender,
		Phone:     phone,
		Place:     place,
		Birth:     birth,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}, nil
}

func convertToUserProfileSliceEntity(rows *sql.Rows) (userprofile.EntitySlice, error) {
	var userProfileEntitySlice userprofile.EntitySlice
	for rows.Next() {
		var userID int
		var name string
		var thumbnail string
		var bio string
		var gender int
		var phone string
		var place string
		var birth *time.Time
		var createdAt, updatedAt, deletedAt *time.Time

		if err := rows.Scan(&userID, &name, &thumbnail, &bio, &gender, &phone, &place, &birth, &createdAt, &updatedAt, &deletedAt); err != nil {
			return nil, werrors.FromConstant(err, werrors.ServerError)
		}

		userProfileData := userprofile.Entity{
			UserID:    userID,
			Name:      name,
			Thumbnail: thumbnail,
			Bio:       bio,
			Gender:    gender,
			Phone:     phone,
			Place:     place,
			Birth:     birth,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			DeletedAt: deletedAt,
		}
		userProfileEntitySlice = append(userProfileEntitySlice, &userProfileData)
	}

	return userProfileEntitySlice, nil
}
