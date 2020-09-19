package wishcard

import (
	"context"
	"os"
	"testing"
	"time"

	placeEntity "wantum/pkg/domain/entity/place"
	tagEntity "wantum/pkg/domain/entity/tag"
	userEntity "wantum/pkg/domain/entity/user"
	profileEntity "wantum/pkg/domain/entity/userprofile"
	wishCardEntity "wantum/pkg/domain/entity/wishcard"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/service/place/mock_place"
	"wantum/pkg/domain/service/tag/mock_tag"
	"wantum/pkg/domain/service/wishcard/mock_wish_card"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	masterTx        repository.MasterTx
	masterTxManager repository.MasterTxManager

	dummyDate        = time.Date(2020, 9, 1, 12, 0, 0, 0, time.Local)
	dummyActivity    = "dummyActivity"
	dummyDescription = "dummyDescription"
	dummyTagName1    = "dummyTag1"
	dummyTagName2    = "dummyTag2"
	dummyPlaceName   = "dummyPlace"

	dummyProfile = profileEntity.Entity{
		UserID:    1,
		Name:      "dummyName",
		Thumbnail: "dummyThumbnail",
		Bio:       "dummyBio",
		Gender:    1,
		Phone:     "12345678901",
		Birth:     &dummyDate,
		CreatedAt: &dummyDate,
		UpdatedAt: &dummyDate,
		DeletedAt: &dummyDate,
	}

	dummyUser = userEntity.Entity{
		ID:        1,
		AuthID:    "dummyID",
		UserName:  "dummyUserName",
		Mail:      "hogehoge@example.com",
		CreatedAt: &dummyDate,
		UpdatedAt: &dummyDate,
		DeletedAt: &dummyDate,
		Profile:   &dummyProfile,
	}

	dummyTag1 = tagEntity.Entity{
		ID:        1,
		Name:      dummyTagName1,
		CreatedAt: &dummyDate,
		UpdatedAt: &dummyDate,
		DeletedAt: &dummyDate,
	}

	dummyTag2 = tagEntity.Entity{
		ID:        2,
		Name:      dummyTagName2,
		CreatedAt: &dummyDate,
		UpdatedAt: &dummyDate,
		DeletedAt: nil,
	}

	dummyTagSlice = tagEntity.EntitySlice{
		&dummyTag1,
		&dummyTag2,
	}

	dummyPlace = placeEntity.Entity{
		ID:        1,
		Name:      dummyPlaceName,
		CreatedAt: &dummyDate,
		UpdatedAt: &dummyDate,
		DeletedAt: nil,
	}

	dummyWishCard = wishCardEntity.Entity{
		ID:          1,
		Author:      &dummyUser,
		Activity:    dummyActivity,
		Description: dummyDescription,
		Date:        &dummyDate,
		DoneAt:      nil,
		CreatedAt:   &dummyDate,
		UpdatedAt:   &dummyDate,
		Place:       &dummyPlace,
		Tags:        dummyTagSlice,
	}
)

func TestMain(m *testing.M) {
	before()
	code := m.Run()
	after()
	os.Exit(code)
}

func before() {
	masterTx = repository.NewMockMasterTx()
	masterTxManager = repository.NewMockMasterTxManager(masterTx)
}

func after() {}

func TestInteractor_CreateNewWishCard(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		placeService := mock_place.NewMockService(ctrl)
		placeService.EXPECT().Create(ctx, masterTx, gomock.Any()).Return(&dummyPlace, nil)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().Create(ctx, masterTx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(&dummyWishCard, nil)

		tagService := mock_tag.NewMockService(ctrl)
		tagService.EXPECT().GetByName(ctx, masterTx, dummyTagName1).Return(&dummyTag1, nil)
		tagService.EXPECT().GetByName(ctx, masterTx, dummyTagName2).Return(nil, nil)
		tagService.EXPECT().Create(ctx, masterTx, dummyTagName2).Return(&dummyTag2, nil)

		interactor := New(masterTxManager, wishCardService, tagService, placeService)

		tags := []string{dummyTagName1, dummyTagName2}
		result, err := interactor.CreateNewWishCard(ctx, 1, 1, dummyActivity, dummyDescription, dummyPlaceName, &dummyDate, tags)

		assert.NoError(t, err)
		assert.NotNil(t, result)

		// validate time
		assert.Equal(t, (*time.Time)(nil), result.DeletedAt)
		assert.Equal(t, (*time.Time)(nil), result.DoneAt)
		// validate place
		assert.Equal(t, dummyPlaceName, result.Place.Name)
		// validate tag
		assert.Equal(t, 2, len(result.Tags))
		assert.Equal(t, dummyTagName1, result.Tags[0].Name)
	})
}

func TestInteractor_UpdateWishCardWithCategoryID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dummyWishCard := wishCardEntity.Entity{
			ID:          1,
			Author:      &dummyUser,
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			DeletedAt:   nil,
			Place:       &dummyPlace,
			Tags:        dummyTagSlice,
		}

		placeService := mock_place.NewMockService(ctrl)
		placeService.EXPECT().Create(ctx, masterTx, gomock.Any()).Return(&dummyPlace, nil)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().UpdateWithCategoryID(ctx, masterTx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(&dummyWishCard, nil)

		tagService := mock_tag.NewMockService(ctrl)
		tagService.EXPECT().GetByName(ctx, masterTx, dummyTagName1).Return(&dummyTag1, nil)
		tagService.EXPECT().GetByName(ctx, masterTx, dummyTagName2).Return(nil, nil)
		tagService.EXPECT().Create(ctx, masterTx, dummyTagName2).Return(&dummyTag2, nil)

		interactor := New(masterTxManager, wishCardService, tagService, placeService)

		tags := []string{dummyTagName1, dummyTagName2}
		result, err := interactor.UpdateWishCardWithCategoryID(ctx, 1, 1, dummyActivity, dummyDescription, dummyPlaceName, &dummyDate, &dummyDate, 1, tags)

		assert.NoError(t, err)
		assert.NotNil(t, result)

		// validate time
		assert.Equal(t, (*time.Time)(nil), result.DeletedAt)
		assert.NotEqual(t, (*time.Time)(nil), result.DoneAt)
		// validate place
		assert.Equal(t, dummyPlaceName, result.Place.Name)
		// validate tag
		assert.Equal(t, 2, len(result.Tags))
		assert.Equal(t, dummyTagName1, result.Tags[0].Name)
	})
}

func TestInteractor_DeleteWishCard(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		placeService := mock_place.NewMockService(ctrl)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().Delete(ctx, masterTx, 1).Return(nil)

		tagService := mock_tag.NewMockService(ctrl)

		interactor := New(masterTxManager, wishCardService, tagService, placeService)

		err := interactor.DeleteWishCardByID(ctx, 1)

		assert.NoError(t, err)
	})
}

func TestInteractor_GetByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		placeService := mock_place.NewMockService(ctrl)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().GetByID(ctx, masterTx, gomock.Any()).Return(&dummyWishCard, nil)

		tagService := mock_tag.NewMockService(ctrl)

		interactor := New(masterTxManager, wishCardService, tagService, placeService)

		result, err := interactor.GetByID(ctx, 1)

		assert.NoError(t, err)
		assert.NotNil(t, result)

		// validate place
		assert.Equal(t, dummyPlaceName, result.Place.Name)
		// validate tag
		assert.Equal(t, 2, len(result.Tags))
		assert.Equal(t, dummyTagName1, result.Tags[0].Name)
	})
}

func TestInteractor_GetByCategoryID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dummyWishCards := wishCardEntity.EntitySlice{
			&wishCardEntity.Entity{
				ID:          1,
				Author:      &dummyUser,
				Activity:    dummyActivity,
				Description: dummyDescription,
				Date:        &dummyDate,
				DoneAt:      &dummyDate,
				CreatedAt:   &dummyDate,
				UpdatedAt:   &dummyDate,
				DeletedAt:   nil,
				Place:       &dummyPlace,
				Tags:        dummyTagSlice,
			},
			&wishCardEntity.Entity{
				ID:          2,
				Author:      &dummyUser,
				Activity:    dummyActivity,
				Description: dummyDescription,
				Date:        &dummyDate,
				DoneAt:      nil,
				CreatedAt:   &dummyDate,
				UpdatedAt:   &dummyDate,
				DeletedAt:   nil,
				Place:       &dummyPlace,
				Tags:        dummyTagSlice,
			},
		}

		placeService := mock_place.NewMockService(ctrl)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().GetByCategoryID(ctx, masterTx, gomock.Any()).Return(dummyWishCards, nil)

		tagService := mock_tag.NewMockService(ctrl)

		interactor := New(masterTxManager, wishCardService, tagService, placeService)

		wishCards, err := interactor.GetByCategoryID(ctx, 1)

		assert.NoError(t, err)
		assert.NotNil(t, wishCards)
		assert.Equal(t, 2, len(wishCards))

		assert.Equal(t, dummyPlaceName, wishCards[0].Place.Name)
		assert.Equal(t, dummyPlaceName, wishCards[1].Place.Name)

		assert.Equal(t, dummyTagName1, wishCards[0].Tags[0].Name)
		assert.Equal(t, 2, len(wishCards[0].Tags))
		assert.Equal(t, dummyTagName1, wishCards[1].Tags[0].Name)
		assert.Equal(t, 2, len(wishCards[1].Tags))
	})
}

func TestInteractor_UpdateActivity(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dummyWishCard := wishCardEntity.Entity{
			ID:          1,
			Author:      &dummyUser,
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			DeletedAt:   nil,
			Place:       &dummyPlace,
			Tags:        dummyTagSlice,
		}

		placeService := mock_place.NewMockService(ctrl)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().GetByID(ctx, masterTx, gomock.Any()).Return(&dummyWishCard, nil)
		wishCardService.EXPECT().Update(ctx, masterTx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(&dummyWishCard, nil)

		tagService := mock_tag.NewMockService(ctrl)

		interactor := New(masterTxManager, wishCardService, tagService, placeService)

		result, err := interactor.UpdateActivity(ctx, 1, 1, dummyActivity)

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestInteractor_UpdateDescription(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dummyWishCard := wishCardEntity.Entity{
			ID:          1,
			Author:      &dummyUser,
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			DeletedAt:   nil,
			Place:       &dummyPlace,
			Tags:        dummyTagSlice,
		}

		placeService := mock_place.NewMockService(ctrl)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().GetByID(ctx, masterTx, gomock.Any()).Return(&dummyWishCard, nil)
		wishCardService.EXPECT().Update(ctx, masterTx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(&dummyWishCard, nil)

		tagService := mock_tag.NewMockService(ctrl)

		interactor := New(masterTxManager, wishCardService, tagService, placeService)

		result, err := interactor.UpdateDescription(ctx, 1, 1, dummyDescription)

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestInteractor_UpdatePlace(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dummyWishCard := wishCardEntity.Entity{
			ID:          1,
			Author:      &dummyUser,
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			DeletedAt:   nil,
			Place:       &dummyPlace,
			Tags:        dummyTagSlice,
		}

		placeService := mock_place.NewMockService(ctrl)
		placeService.EXPECT().Create(ctx, masterTx, gomock.Any()).Return(&dummyPlace, nil)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().GetByID(ctx, masterTx, gomock.Any()).Return(&dummyWishCard, nil)
		wishCardService.EXPECT().Update(ctx, masterTx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(&dummyWishCard, nil)

		tagService := mock_tag.NewMockService(ctrl)

		interactor := New(masterTxManager, wishCardService, tagService, placeService)

		result, err := interactor.UpdatePlace(ctx, 1, 1, dummyPlaceName)

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestInteractor_UpdateDate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dummyWishCard := wishCardEntity.Entity{
			ID:          1,
			Author:      &dummyUser,
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			DeletedAt:   nil,
			Place:       &dummyPlace,
			Tags:        dummyTagSlice,
		}

		placeService := mock_place.NewMockService(ctrl)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().GetByID(ctx, masterTx, gomock.Any()).Return(&dummyWishCard, nil)
		wishCardService.EXPECT().Update(ctx, masterTx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(&dummyWishCard, nil)

		tagService := mock_tag.NewMockService(ctrl)

		interactor := New(masterTxManager, wishCardService, tagService, placeService)

		result, err := interactor.UpdateDate(ctx, 1, 1, &dummyDate)

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestInteractor_AddTags(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dummyWishCard := wishCardEntity.Entity{
			ID:          1,
			Author:      &dummyUser,
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			DeletedAt:   nil,
			Place:       &dummyPlace,
			Tags:        dummyTagSlice,
		}

		placeService := mock_place.NewMockService(ctrl)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().AddTags(ctx, masterTx, gomock.Any(), gomock.Any()).Return(&dummyWishCard, nil)

		tagService := mock_tag.NewMockService(ctrl)
		tagService.EXPECT().GetByName(ctx, masterTx, dummyTagName1).Return(&dummyTag1, nil)
		tagService.EXPECT().GetByName(ctx, masterTx, dummyTagName2).Return(nil, nil)
		tagService.EXPECT().Create(ctx, masterTx, dummyTagName2).Return(&dummyTag2, nil)

		interactor := New(masterTxManager, wishCardService, tagService, placeService)

		result, err := interactor.AddTags(ctx, 1, 1, []string{dummyTagName1, dummyTagName2})

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestInteractor_DeleteTags(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dummyWishCard := wishCardEntity.Entity{
			ID:          1,
			Author:      &dummyUser,
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			DeletedAt:   nil,
			Place:       &dummyPlace,
			Tags:        dummyTagSlice,
		}

		placeService := mock_place.NewMockService(ctrl)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().DeleteTags(ctx, masterTx, gomock.Any(), gomock.Any()).Return(&dummyWishCard, nil)

		tagService := mock_tag.NewMockService(ctrl)

		interactor := New(masterTxManager, wishCardService, tagService, placeService)

		result, err := interactor.DeleteTags(ctx, 1, 1, []int{1, 2})

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}
