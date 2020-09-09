package model

import (
	"testing"
	"time"
	"wantum/pkg/domain/entity"

	"github.com/stretchr/testify/assert"
)

func TestConvertToWishCardEntity(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		date := time.Date(2020, 9, 1, 10, 10, 10, 0, time.UTC)
		testData := &WishCardModel{
			ID:          1,
			UserID:      1,
			Activity:    "タピる",
			Description: "せつめいーーー",
			Date:        &date,
			DoneAt:      nil,
			CreatedAt:   &date,
			UpdatedAt:   &date,
			DeletedAt:   nil,
		}

		result := ConvertToWishCardEntiry(testData)

		assert.NotNil(t, result)
		assert.IsType(t, &entity.WishCard{}, result)
	})

	t.Run("failure_nil", func(t *testing.T) {
		result := ConvertToWishCardEntiry(nil)
		assert.Nil(t, result)
	})
}
func TestConvertToWishCardSliceEntity(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		date := time.Date(2020, 9, 1, 10, 10, 10, 0, time.UTC)
		data := &WishCardModel{
			ID:          1,
			UserID:      1,
			Activity:    "タピる",
			Description: "せつめいーーー",
			Date:        &date,
			DoneAt:      nil,
			CreatedAt:   &date,
			UpdatedAt:   &date,
			DeletedAt:   nil,
		}
		testData := WishCardModelSlice{data, data}

		result := ConvertToWishCardSliceEntity(testData)

		assert.NotNil(t, result)
		assert.IsType(t, entity.WishCardSlice{}, result)
	})

	t.Run("failure_nil", func(t *testing.T) {
		result := ConvertToWishCardSliceEntity(nil)
		assert.Nil(t, result)
	})
}
