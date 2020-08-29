package model

import (
	"time"
	"wantum/pkg/domain/entity"
)

type WishCardModel struct {
	ID          int
	UserID      int
	Activity    string
	Description string
	Date        *time.Time
	DoneAt      *time.Time
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
	CategoryID  int
	PlaceID     int
}

type WishCardModelSlice []*WishCardModel

func ConvertToWishCardEntiry(wishCard *WishCardModel) *entity.WishCard {
	if wishCard == nil {
		return nil
	}
	return &entity.WishCard{
		ID:          wishCard.ID,
		UserID:      wishCard.UserID,
		Activity:    wishCard.Activity,
		Description: wishCard.Description,
		Date:        wishCard.Date,
		DoneAt:      wishCard.DoneAt,
		CreatedAt:   wishCard.CreatedAt,
		UpdatedAt:   wishCard.UpdatedAt,
		DeletedAt:   wishCard.DeletedAt,
	}
}

func ConvertToWishCardSliceEntity(wishCards WishCardModelSlice) entity.WishCardSlice {
	if wishCards == nil {
		return nil
	}
	res := make(entity.WishCardSlice, 0, len(wishCards))
	for _, wishCard := range wishCards {
		res = append(res, ConvertToWishCardEntiry(wishCard))
	}
	return res
}
