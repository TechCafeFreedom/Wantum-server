package model

import (
	"testing"
	"time"
	"wantum/pkg/domain/entity"

	"github.com/stretchr/testify/assert"
)

func TestConvertToPlaceEntity(t *testing.T) {
	t.Run("ver success", func(t *testing.T) {
		date := time.Date(2020, 9, 1, 10, 10, 10, 0, time.UTC)
		testData := &PlaceModel{
			ID:        1,
			Name:      "desney land",
			CreatedAt: &date,
			UpdatedAt: &date,
			DeletedAt: &date,
		}

		result := ConvertToPlaceEntity(testData)

		assert.NotNil(t, result)
		assert.IsType(t, &entity.Place{}, result)
	})

	t.Run("ver nil", func(t *testing.T) {
		result := ConvertToPlaceEntity(nil)
		assert.Nil(t, result)
	})
}
func TestConvertToPlaceSliceEntity(t *testing.T) {
	t.Run("ver success", func(t *testing.T) {
		date := time.Date(2020, 9, 1, 10, 10, 10, 0, time.UTC)
		data := &PlaceModel{
			ID:        1,
			Name:      "desney land",
			CreatedAt: &date,
			UpdatedAt: &date,
			DeletedAt: &date,
		}
		testData := PlaceModelSlice{data, data}

		result := ConvertToPlaceSliceEntity(testData)

		assert.NotNil(t, result)
		assert.IsType(t, entity.PlaceSlice{}, result)
	})

	t.Run("ver nil", func(t *testing.T) {
		result := ConvertToPlaceSliceEntity(nil)
		assert.Nil(t, result)
	})
}
