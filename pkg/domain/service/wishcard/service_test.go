package wishcard

import (
	"context"
	"os"
	"testing"
	"time"
	placeEntity "wantum/pkg/domain/entity/place"
	userEntity "wantum/pkg/domain/entity/user"
	wishCardEntity "wantum/pkg/domain/entity/wishcard"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/wishcard/mock_wish_card"
	"wantum/pkg/domain/repository/wishcardtag/mock_wish_card_tag"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	masterTx repository.MasterTx

	dummyDate        = time.Date(2020, 9, 1, 12, 0, 0, 0, time.Local)
	dummyActivity    = "sampleActivity"
	dummyDescription = "sampleDescription"
	dummyWishCardID  = 2
	dummyUserID      = 1
	dummyPlaceID     = 1
	dummyCategoryID  = 1
)

func TestMain(m *testing.M) {
	before()
	code := m.Run()
	os.Exit(code)
}

func before() {
	masterTx = repository.NewMockMasterTx()
}

func TestService_Create(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wcRepo := mock_wish_card.NewMockRepository(ctrl)
	wcRepo.EXPECT().Insert(ctx, masterTx, gomock.Any(), gomock.Any()).Return(1, nil)

	wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)
	wctRepo.EXPECT().BulkInsert(ctx, masterTx, gomock.Any(), gomock.Any()).Return(nil)

	service := New(wcRepo, wctRepo)
	result, err := service.Create(ctx, masterTx, dummyActivity, dummyDescription, &dummyDate, 1, 1, 1, []int{1, 2})

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
}

func TestService_Update(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyData := &wishCardEntity.Entity{
		ID: 1,
		Author: &userEntity.Entity{
			ID: 1,
		},
		Activity:    "act",
		Description: "desc",
		Date:        &dummyDate,
		DoneAt:      &dummyDate,
		CreatedAt:   &dummyDate,
		UpdatedAt:   &dummyDate,
		Place: &placeEntity.Entity{
			ID: 1,
		},
	}

	wcRepo := mock_wish_card.NewMockRepository(ctrl)
	wcRepo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(dummyData, nil)
	wcRepo.EXPECT().Update(ctx, masterTx, gomock.Any()).Return(nil)

	wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)

	service := New(wcRepo, wctRepo)
	result, err := service.Update(ctx, masterTx, 1, dummyActivity, dummyDescription, &dummyDate, &dummyDate, 1, 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, dummyActivity, result.Activity)
	assert.Equal(t, dummyDescription, result.Description)
}

func TestService_UpdateActivity(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dummyData := &wishCardEntity.Entity{
			ID: 1,
			Author: &userEntity.Entity{
				ID: 1,
			},
			Activity:    "act",
			Description: dummyDescription,
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		}

		wcRepo := mock_wish_card.NewMockRepository(ctrl)
		wcRepo.EXPECT().UpdateActivity(ctx, masterTx, gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		wcRepo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(dummyData, nil)

		wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)

		service := New(wcRepo, wctRepo)
		result, err := service.UpdateActivity(ctx, masterTx, dummyWishCardID, dummyActivity)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, dummyActivity, result.Activity)
		assert.NotEqual(t, dummyDate, result.UpdatedAt)
	})

	t.Run("failure_存在しないデータ", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		wcRepo := mock_wish_card.NewMockRepository(ctrl)
		wcRepo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(nil, nil)

		wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)

		service := New(wcRepo, wctRepo)
		_, err := service.UpdateActivity(ctx, masterTx, dummyWishCardID, dummyActivity)

		assert.Error(t, err)
	})
}

func TestService_UpdateDescription(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dummyData := &wishCardEntity.Entity{
			ID: 1,
			Author: &userEntity.Entity{
				ID: 1,
			},
			Activity:    dummyActivity,
			Description: "desc",
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		}

		wcRepo := mock_wish_card.NewMockRepository(ctrl)
		wcRepo.EXPECT().UpdateDescription(ctx, masterTx, gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		wcRepo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(dummyData, nil)

		wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)

		service := New(wcRepo, wctRepo)
		result, err := service.UpdateDescription(ctx, masterTx, dummyWishCardID, dummyDescription)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, dummyDescription, result.Description)
		assert.NotEqual(t, dummyDate, result.UpdatedAt)
	})

	t.Run("failure_存在しないデータ", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		wcRepo := mock_wish_card.NewMockRepository(ctrl)
		wcRepo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(nil, nil)

		wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)

		service := New(wcRepo, wctRepo)
		_, err := service.UpdateDescription(ctx, masterTx, dummyWishCardID, dummyDescription)

		assert.Error(t, err)
	})
}

func TestService_UpdateDate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		_dummyDate := time.Date(2020, 10, 10, 10, 0, 0, 0, time.Local)

		dummyData := &wishCardEntity.Entity{
			ID: 1,
			Author: &userEntity.Entity{
				ID: 1,
			},
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &_dummyDate,
			DoneAt:      &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		}

		wcRepo := mock_wish_card.NewMockRepository(ctrl)
		wcRepo.EXPECT().UpdateDate(ctx, masterTx, gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		wcRepo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(dummyData, nil)

		wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)

		service := New(wcRepo, wctRepo)
		result, err := service.UpdateDate(ctx, masterTx, dummyWishCardID, &dummyDate)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, dummyDate, result.Date.Local())
		assert.NotEqual(t, dummyDate, result.UpdatedAt)
	})

	t.Run("failure_存在しないデータ", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		wcRepo := mock_wish_card.NewMockRepository(ctrl)
		wcRepo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(nil, nil)

		wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)

		service := New(wcRepo, wctRepo)
		_, err := service.UpdateDate(ctx, masterTx, dummyWishCardID, &dummyDate)

		assert.Error(t, err)
	})
}

func TestService_UpdateDoneAt(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		_dummyDate := time.Date(2020, 10, 10, 10, 0, 0, 0, time.Local)

		dummyData := &wishCardEntity.Entity{
			ID: 1,
			Author: &userEntity.Entity{
				ID: 1,
			},
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			DoneAt:      &_dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		}

		wcRepo := mock_wish_card.NewMockRepository(ctrl)
		wcRepo.EXPECT().UpdateDoneAt(ctx, masterTx, gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		wcRepo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(dummyData, nil)

		wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)

		service := New(wcRepo, wctRepo)
		result, err := service.UpdateDoneAt(ctx, masterTx, dummyWishCardID, &dummyDate)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, dummyDate, result.DoneAt.Local())
		assert.NotEqual(t, dummyDate, result.UpdatedAt)
	})

	t.Run("failure_存在しないデータ", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		wcRepo := mock_wish_card.NewMockRepository(ctrl)
		wcRepo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(nil, nil)

		wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)

		service := New(wcRepo, wctRepo)
		_, err := service.UpdateDoneAt(ctx, masterTx, dummyWishCardID, &dummyDate)

		assert.Error(t, err)
	})

}

func TestService_UpdateAuthor(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dummyData := &wishCardEntity.Entity{
			ID: 1,
			Author: &userEntity.Entity{
				ID: 5,
			},
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		}

		wcRepo := mock_wish_card.NewMockRepository(ctrl)
		wcRepo.EXPECT().UpdateUserID(ctx, masterTx, gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		wcRepo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(dummyData, nil)

		wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)

		service := New(wcRepo, wctRepo)
		result, err := service.UpdateAuthor(ctx, masterTx, dummyWishCardID, dummyUserID)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, dummyUserID, result.Author.ID)
		assert.NotEqual(t, dummyDate, result.UpdatedAt)
	})

	t.Run("failure_存在しないデータ", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		wcRepo := mock_wish_card.NewMockRepository(ctrl)
		wcRepo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(nil, nil)

		wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)

		service := New(wcRepo, wctRepo)
		_, err := service.UpdateAuthor(ctx, masterTx, dummyWishCardID, dummyUserID)

		assert.Error(t, err)
	})
}

func TestService_UpdatePlace(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dummyData := &wishCardEntity.Entity{
			ID: 1,
			Author: &userEntity.Entity{
				ID: 1,
			},
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			Place: &placeEntity.Entity{
				ID: 5,
			},
		}

		wcRepo := mock_wish_card.NewMockRepository(ctrl)
		wcRepo.EXPECT().UpdatePlaceID(ctx, masterTx, gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		wcRepo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(dummyData, nil)

		wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)

		service := New(wcRepo, wctRepo)
		result, err := service.UpdatePlace(ctx, masterTx, dummyWishCardID, dummyPlaceID)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, dummyPlaceID, result.Place.ID)
		assert.NotEqual(t, dummyDate, result.UpdatedAt)
	})

	t.Run("failure_存在しないデータ", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		wcRepo := mock_wish_card.NewMockRepository(ctrl)
		wcRepo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(nil, nil)

		wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)

		service := New(wcRepo, wctRepo)
		_, err := service.UpdatePlace(ctx, masterTx, dummyWishCardID, dummyPlaceID)

		assert.Error(t, err)
	})

}

func TestService_UpdateCategory(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dummyData := &wishCardEntity.Entity{
			ID: 1,
			Author: &userEntity.Entity{
				ID: 1,
			},
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		}

		wcRepo := mock_wish_card.NewMockRepository(ctrl)
		wcRepo.EXPECT().UpdateCategoryID(ctx, masterTx, gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		wcRepo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(dummyData, nil)

		wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)

		service := New(wcRepo, wctRepo)
		result, err := service.UpdateCategory(ctx, masterTx, dummyWishCardID, dummyCategoryID)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotEqual(t, dummyDate, result.UpdatedAt)
	})

	t.Run("failure_存在しないデータ", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		wcRepo := mock_wish_card.NewMockRepository(ctrl)
		wcRepo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(nil, nil)

		wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)

		service := New(wcRepo, wctRepo)
		_, err := service.UpdateCategory(ctx, masterTx, dummyWishCardID, dummyCategoryID)

		assert.Error(t, err)
	})
}

func TestService_UpdateWithCategoryID(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyData := &wishCardEntity.Entity{
		ID: 1,
		Author: &userEntity.Entity{
			ID: 1,
		},
		Activity:    "act",
		Description: "desc",
		Date:        &dummyDate,
		DoneAt:      &dummyDate,
		CreatedAt:   &dummyDate,
		UpdatedAt:   &dummyDate,
		Place: &placeEntity.Entity{
			ID: 1,
		},
	}

	wcRepo := mock_wish_card.NewMockRepository(ctrl)
	wcRepo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(dummyData, nil)
	wcRepo.EXPECT().UpdateWithCategoryID(ctx, masterTx, gomock.Any(), gomock.Any()).Return(nil)

	wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)
	wctRepo.EXPECT().BulkInsert(ctx, masterTx, gomock.Any(), gomock.Any()).Return(nil)
	wctRepo.EXPECT().DeleteByWishCardID(ctx, masterTx, gomock.Any()).Return(nil)

	service := New(wcRepo, wctRepo)
	result, err := service.UpdateWithCategoryID(ctx, masterTx, 1, dummyActivity, dummyDescription, &dummyDate, &dummyDate, 1, 1, 1, []int{1, 2})

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, dummyActivity, result.Activity)
	assert.Equal(t, dummyDescription, result.Description)
}

func TestService_UpDeleteFlag(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyData := &wishCardEntity.Entity{
		ID: 1,
		Author: &userEntity.Entity{
			ID: 1,
		},
		Activity:    dummyActivity,
		Description: dummyDescription,
		Date:        &dummyDate,
		DoneAt:      &dummyDate,
		CreatedAt:   &dummyDate,
		UpdatedAt:   &dummyDate,
		Place: &placeEntity.Entity{
			ID: 1,
		},
	}

	wcRepo := mock_wish_card.NewMockRepository(ctrl)
	wcRepo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(dummyData, nil)
	wcRepo.EXPECT().UpDeleteFlag(ctx, masterTx, gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)

	service := New(wcRepo, wctRepo)
	result, err := service.UpDeleteFlag(ctx, masterTx, 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.DeletedAt)
}

func TestService_DownDeleteFlag(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyData := &wishCardEntity.Entity{
		ID: 1,
		Author: &userEntity.Entity{
			ID: 1,
		},
		Activity:    dummyActivity,
		Description: dummyDescription,
		Date:        &dummyDate,
		DoneAt:      &dummyDate,
		CreatedAt:   &dummyDate,
		UpdatedAt:   &dummyDate,
		DeletedAt:   &dummyDate,
		Place: &placeEntity.Entity{
			ID: 1,
		},
	}

	wcRepo := mock_wish_card.NewMockRepository(ctrl)
	wcRepo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(dummyData, nil)
	wcRepo.EXPECT().DownDeleteFlag(ctx, masterTx, gomock.Any(), gomock.Any()).Return(nil)

	wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)

	service := New(wcRepo, wctRepo)
	result, err := service.DownDeleteFlag(ctx, masterTx, 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Nil(t, result.DeletedAt)
}

func TestService_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dummyData := &wishCardEntity.Entity{
			ID: 1,
			Author: &userEntity.Entity{
				ID: 1,
			},
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			DeletedAt:   &dummyDate,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		}

		wcRepo := mock_wish_card.NewMockRepository(ctrl)
		wcRepo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(dummyData, nil)
		wcRepo.EXPECT().Delete(ctx, masterTx, gomock.Any()).Return(nil)

		wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)
		wctRepo.EXPECT().DeleteByWishCardID(ctx, masterTx, gomock.Any()).Return(nil)

		service := New(wcRepo, wctRepo)
		err := service.Delete(ctx, masterTx, 1)

		assert.NoError(t, err)

	})

	t.Run("failure_deleteフラグがたってない", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dummyData := &wishCardEntity.Entity{
			ID: 1,
			Author: &userEntity.Entity{
				ID: 1,
			},
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		}
		wcRepo := mock_wish_card.NewMockRepository(ctrl)
		wcRepo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(dummyData, nil)

		wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)

		service := New(wcRepo, wctRepo)
		err := service.Delete(ctx, masterTx, 1)

		assert.Error(t, err)
	})
}

func TestService_GetByID(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyData := &wishCardEntity.Entity{
		ID: 1,
		Author: &userEntity.Entity{
			ID: 1,
		},
		Activity:    dummyActivity,
		Description: dummyDescription,
		Date:        &dummyDate,
		DoneAt:      &dummyDate,
		CreatedAt:   &dummyDate,
		UpdatedAt:   &dummyDate,
		Place: &placeEntity.Entity{
			ID: 1,
		},
	}

	wcRepo := mock_wish_card.NewMockRepository(ctrl)
	wcRepo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(dummyData, nil)

	wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)

	service := New(wcRepo, wctRepo)
	result, err := service.GetByID(ctx, masterTx, 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestService_GetByIDs(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyData := wishCardEntity.EntitySlice{
		&wishCardEntity.Entity{
			ID: 1,
			Author: &userEntity.Entity{
				ID: 1,
			},
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		},
		&wishCardEntity.Entity{
			ID: 2,
			Author: &userEntity.Entity{
				ID: 1,
			},
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		},
	}

	wcRepo := mock_wish_card.NewMockRepository(ctrl)
	wcRepo.EXPECT().SelectByIDs(ctx, masterTx, gomock.Any()).Return(dummyData, nil)

	wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)

	service := New(wcRepo, wctRepo)
	result, err := service.GetByIDs(ctx, masterTx, []int{1, 2})

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(result))
}

func TestService_GetByCategoryID(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyData := wishCardEntity.EntitySlice{
		&wishCardEntity.Entity{
			ID: 1,
			Author: &userEntity.Entity{
				ID: 1,
			},
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		},
		&wishCardEntity.Entity{
			ID: 2,
			Author: &userEntity.Entity{
				ID: 1,
			},
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		},
	}

	wcRepo := mock_wish_card.NewMockRepository(ctrl)
	wcRepo.EXPECT().SelectByCategoryID(ctx, masterTx, gomock.Any()).Return(dummyData, nil)

	wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)

	service := New(wcRepo, wctRepo)
	result, err := service.GetByCategoryID(ctx, masterTx, 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(result))
}

func TestService_AddTags(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyData := &wishCardEntity.Entity{
		ID: 1,
		Author: &userEntity.Entity{
			ID: 1,
		},
		Activity:    dummyActivity,
		Description: dummyDescription,
		Date:        &dummyDate,
		DoneAt:      &dummyDate,
		CreatedAt:   &dummyDate,
		UpdatedAt:   &dummyDate,
		Place: &placeEntity.Entity{
			ID: 1,
		},
	}

	wcRepo := mock_wish_card.NewMockRepository(ctrl)
	wcRepo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(dummyData, nil)

	wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)
	wctRepo.EXPECT().BulkInsert(ctx, masterTx, gomock.Any(), gomock.Any()).Return(nil)

	service := New(wcRepo, wctRepo)
	result, err := service.AddTags(ctx, masterTx, 1, []int{1, 2})

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestService_DeleteTags(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyData := &wishCardEntity.Entity{
		ID: 1,
		Author: &userEntity.Entity{
			ID: 1,
		},
		Activity:    dummyActivity,
		Description: dummyDescription,
		Date:        &dummyDate,
		DoneAt:      &dummyDate,
		CreatedAt:   &dummyDate,
		UpdatedAt:   &dummyDate,
		Place: &placeEntity.Entity{
			ID: 1,
		},
	}

	wcRepo := mock_wish_card.NewMockRepository(ctrl)
	wcRepo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(dummyData, nil)

	wctRepo := mock_wish_card_tag.NewMockRepository(ctrl)
	wctRepo.EXPECT().DeleteByIDs(ctx, masterTx, gomock.Any(), gomock.Any()).Return(nil)

	service := New(wcRepo, wctRepo)
	result, err := service.DeleteTags(ctx, masterTx, 1, []int{1, 2})

	assert.NoError(t, err)
	assert.NotNil(t, result)
}
