package model

import (
	"testing"
	"time"
	"wantum/pkg/domain/entity"

	"github.com/stretchr/testify/assert"
)

func TestConvertToTagEntity(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		date := time.Date(2020, 9, 1, 10, 10, 10, 0, time.Local)
		testData := &TagModel{
			ID:        1,
			Name:      "disney land",
			CreatedAt: &date,
			UpdatedAt: &date,
			DeletedAt: &date,
		}

		result := ConvertToTagEntity(testData)

		assert.NotNil(t, result)
		assert.IsType(t, &entity.Tag{}, result)
	})

	t.Run("failure_nil", func(t *testing.T) {
		result := ConvertToTagEntity(nil)
		assert.Nil(t, result)
	})
}
func TestConvertToTagSliceEntity(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		date := time.Date(2020, 9, 1, 10, 10, 10, 0, time.Local)
		data := &TagModel{
			ID:        1,
			Name:      "disney land",
			CreatedAt: &date,
			UpdatedAt: &date,
			DeletedAt: &date,
		}
		testData := TagModelSlice{data, data}

		result := ConvertToTagSliceEntity(testData)

		assert.NotNil(t, result)
		assert.IsType(t, entity.TagSlice{}, result)
	})

	t.Run("failure_nil", func(t *testing.T) {
		result := ConvertToTagSliceEntity(nil)
		assert.Nil(t, result)
	})
}
