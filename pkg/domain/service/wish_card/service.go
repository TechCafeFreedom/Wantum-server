package wish_card

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"wantum/pkg/domain/entity"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/wish_card"
	"wantum/pkg/infrastructure/mysql/model"
	"wantum/pkg/werrors"
)

type Service interface {
	Create(ctx context.Context, masterTx repository.MasterTx, activity, description string, date *time.Time, userID, categoryID, placeID int) (*entity.WishCard, error)
	Update(ctx context.Context, masterTx repository.MasterTx, wishCardID int, activity, description string, date, doneAt *time.Time, userID, categoryID, placeID int) (*entity.WishCard, error)

	Delete(ctx context.Context, masterTx repository.MasterTx, wishCardID int) error
	UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (*entity.WishCard, error)
	DownDeleteFlag(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (*entity.WishCard, error)

	GetByID(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (*entity.WishCard, error)
	GetByIDs(ctx context.Context, masterTx repository.MasterTx, wishCardIDs []int) (entity.WishCardSlice, error)
	GetByCategoryID(ctx context.Context, masterTx repository.MasterTx, categoryID int) (entity.WishCardSlice, error)
}

type service struct {
	wishCardRepository wish_card.Repository
}

func New(repo wish_card.Repository) Service {
	return &service{
		wishCardRepository: repo,
	}
}

func (s *service) Create(ctx context.Context, masterTx repository.MasterTx, activity, description string, date *time.Time, userID, categoryID, placeID int) (*entity.WishCard, error) {
	createdAt := time.Now()
	wishCard := &model.WishCardModel{
		UserID:      userID,
		Activity:    activity,
		Description: description,
		Date:        date,
		CategoryID:  categoryID,
		PlaceID:     placeID,
		CreatedAt:   &createdAt,
		UpdatedAt:   &createdAt,
	}
	result, err := s.wishCardRepository.Insert(ctx, masterTx, wishCard)
	if err != nil {
		return nil, err
	}
	wishCard.ID = result
	return model.ConvertToWishCardEntiry(wishCard), nil
}

// NOTE: 空値があった時、元データが消滅する。
// NOTE: リクエストは、全フィールド埋める or 差分だけ
func (s *service) Update(ctx context.Context, masterTx repository.MasterTx, wishCardID int, activity, description string, date, doneAt *time.Time, userID, categoryID, placeID int) (*entity.WishCard, error) {
	wishCard, err := s.wishCardRepository.SelectByID(ctx, masterTx, wishCardID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	updatedAt := time.Now()
	wishCard.Activity = activity
	wishCard.Description = description
	wishCard.Date = date
	wishCard.DoneAt = doneAt
	wishCard.CategoryID = categoryID
	wishCard.PlaceID = placeID
	wishCard.UpdatedAt = &updatedAt

	err = s.wishCardRepository.Update(ctx, masterTx, wishCard)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return model.ConvertToWishCardEntiry(wishCard), nil
}

func (s *service) UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (*entity.WishCard, error) {
	wishCard, err := s.wishCardRepository.SelectByID(ctx, masterTx, wishCardID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	updatedAt := time.Now()
	wishCard.UpdatedAt = &updatedAt
	wishCard.DeletedAt = &updatedAt
	err = s.wishCardRepository.UpDeleteFlag(ctx, masterTx, wishCard)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return model.ConvertToWishCardEntiry(wishCard), nil
}

func (s *service) DownDeleteFlag(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (*entity.WishCard, error) {
	wishCard, err := s.wishCardRepository.SelectByID(ctx, masterTx, wishCardID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	updatedAt := time.Now()
	wishCard.UpdatedAt = &updatedAt
	wishCard.DeletedAt = nil
	err = s.wishCardRepository.DownDeleteFlag(ctx, masterTx, wishCard)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return model.ConvertToWishCardEntiry(wishCard), nil
}

func (s *service) Delete(ctx context.Context, masterTx repository.MasterTx, wishCardID int) error {
	wishCard, err := s.wishCardRepository.SelectByID(ctx, masterTx, wishCardID)
	if err != nil {
		return werrors.Stack(err)
	}
	if wishCard.DeletedAt == nil {
		return werrors.Newf(
			fmt.Errorf("can't delete this data. this data did not up a delete flag. wishCardID=%v", wishCardID),
			http.StatusBadRequest,
			"このデータは削除できません",
			"could not delete this place",
		)
	}
	err = s.wishCardRepository.Delete(ctx, masterTx, wishCardID)
	if err != nil {
		return werrors.Stack(err)
	}
	return nil
}

func (s *service) GetByID(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (*entity.WishCard, error) {
	wishCard, err := s.wishCardRepository.SelectByID(ctx, masterTx, wishCardID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return model.ConvertToWishCardEntiry(wishCard), nil
}

func (s *service) GetByIDs(ctx context.Context, masterTx repository.MasterTx, wishCardIDs []int) (entity.WishCardSlice, error) {
	idList := make([]string, 0, len(wishCardIDs))
	for _, id := range wishCardIDs {
		idList = append(idList, strconv.Itoa(id))
	}
	wishCards, err := s.wishCardRepository.SelectByIDs(ctx, masterTx, idList)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return model.ConvertToWishCardSliceEntity(wishCards), nil
}

func (s *service) GetByCategoryID(ctx context.Context, masterTx repository.MasterTx, categoryID int) (entity.WishCardSlice, error) {
	wishCards, err := s.wishCardRepository.SelectByCategoryID(ctx, masterTx, categoryID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return model.ConvertToWishCardSliceEntity(wishCards), nil
}
