package wishcard

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	placeEntity "wantum/pkg/domain/entity/place"
	tagEntity "wantum/pkg/domain/entity/tag"
	userEntity "wantum/pkg/domain/entity/user"
	profileEntity "wantum/pkg/domain/entity/userprofile"
	wishCardEntity "wantum/pkg/domain/entity/wishcard"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/service/place/mock_place"
	"wantum/pkg/domain/service/profile/mock_profile"
	"wantum/pkg/domain/service/tag/mock_tag"
	"wantum/pkg/domain/service/user/mock_user"
	"wantum/pkg/domain/service/wishcard/mock_wish_card"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	masterTx        repository.MasterTx
	masterTxManager repository.MasterTxManager

	dummyDate        = time.Date(2040, 9, 1, 12, 0, 0, 0, time.Local)
	dummyActivity    = "dummyActivity"
	dummyDescription = "dummyDescription"
	dummyTagName1    = "dummyTag1"
	dummyTagName2    = "dummyTag2"
	dummyPlaceName   = "dummyPlace"

	dummyProfile = &profileEntity.Entity{
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

	dummyUser = &userEntity.Entity{
		ID:        1,
		AuthID:    "dummyID",
		UserName:  "dummyUserName",
		Mail:      "hogehoge@example.com",
		CreatedAt: &dummyDate,
		UpdatedAt: &dummyDate,
		DeletedAt: &dummyDate,
	}

	dummyTag1 = &tagEntity.Entity{
		ID:        1,
		Name:      dummyTagName1,
		CreatedAt: &dummyDate,
		UpdatedAt: &dummyDate,
		DeletedAt: &dummyDate,
	}

	dummyTag2 = &tagEntity.Entity{
		ID:        2,
		Name:      dummyTagName2,
		CreatedAt: &dummyDate,
		UpdatedAt: &dummyDate,
		DeletedAt: nil,
	}

	dummyTagSlice = tagEntity.EntitySlice{
		dummyTag1,
		dummyTag2,
	}

	dummyPlace = &placeEntity.Entity{
		ID:        1,
		Name:      dummyPlaceName,
		CreatedAt: &dummyDate,
		UpdatedAt: &dummyDate,
		DeletedAt: nil,
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

		dummyWishCard := &wishCardEntity.Entity{
			ID:          1,
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			DeletedAt:   nil,
		}

		userService := mock_user.NewMockService(ctrl)
		userService.EXPECT().GetByPK(ctx, masterTx, gomock.Any()).Return(dummyUser, nil)

		profileService := mock_profile.NewMockService(ctrl)
		profileService.EXPECT().GetByUserID(ctx, masterTx, gomock.Any()).Return(dummyProfile, nil)

		placeService := mock_place.NewMockService(ctrl)
		placeService.EXPECT().Create(ctx, masterTx, gomock.Any()).Return(dummyPlace, nil)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().Create(ctx, masterTx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(dummyWishCard, nil)

		tagService := mock_tag.NewMockService(ctrl)
		tagService.EXPECT().GetByName(ctx, masterTx, dummyTagName1).Return(dummyTag1, nil)
		tagService.EXPECT().GetByName(ctx, masterTx, dummyTagName2).Return(nil, nil)
		tagService.EXPECT().Create(ctx, masterTx, dummyTagName2).Return(dummyTag2, nil)

		interactor := New(masterTxManager, wishCardService, userService, profileService, tagService, placeService)

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

		dummyWishCard := &wishCardEntity.Entity{
			ID:          1,
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			DeletedAt:   nil,
		}

		userService := mock_user.NewMockService(ctrl)
		userService.EXPECT().GetByPK(ctx, masterTx, gomock.Any()).Return(dummyUser, nil)

		profileService := mock_profile.NewMockService(ctrl)
		profileService.EXPECT().GetByUserID(ctx, masterTx, gomock.Any()).Return(dummyProfile, nil)

		placeService := mock_place.NewMockService(ctrl)
		placeService.EXPECT().Create(ctx, masterTx, gomock.Any()).Return(dummyPlace, nil)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().UpdateWithCategoryID(ctx, masterTx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(dummyWishCard, nil)

		tagService := mock_tag.NewMockService(ctrl)
		tagService.EXPECT().GetByName(ctx, masterTx, dummyTagName1).Return(dummyTag1, nil)
		tagService.EXPECT().GetByName(ctx, masterTx, dummyTagName2).Return(nil, nil)
		tagService.EXPECT().Create(ctx, masterTx, dummyTagName2).Return(dummyTag2, nil)

		interactor := New(masterTxManager, wishCardService, userService, profileService, tagService, placeService)

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

		userService := mock_user.NewMockService(ctrl)

		profileService := mock_profile.NewMockService(ctrl)

		placeService := mock_place.NewMockService(ctrl)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().Delete(ctx, masterTx, 1).Return(nil)

		tagService := mock_tag.NewMockService(ctrl)

		interactor := New(masterTxManager, wishCardService, userService, profileService, tagService, placeService)

		err := interactor.DeleteWishCardByID(ctx, 1)

		assert.NoError(t, err)
	})
}

func TestInteractor_GetByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		dummyWishCard := &wishCardEntity.Entity{
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
			DeletedAt:   nil,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		}

		userService := mock_user.NewMockService(ctrl)
		userService.EXPECT().GetByPK(ctx, masterTx, gomock.Any()).Return(dummyUser, nil)

		profileService := mock_profile.NewMockService(ctrl)
		profileService.EXPECT().GetByUserID(ctx, masterTx, gomock.Any()).Return(dummyProfile, nil)

		placeService := mock_place.NewMockService(ctrl)
		placeService.EXPECT().GetByID(ctx, masterTx, gomock.Any()).Return(dummyPlace, nil)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().GetByID(ctx, masterTx, gomock.Any()).Return(dummyWishCard, nil)

		tagService := mock_tag.NewMockService(ctrl)
		tagService.EXPECT().GetByWishCardID(ctx, masterTx, gomock.Any()).Return(dummyTagSlice, nil)

		interactor := New(masterTxManager, wishCardService, userService, profileService, tagService, placeService)

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
				DeletedAt:   nil,
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
				DeletedAt:   nil,
				Place: &placeEntity.Entity{
					ID: 1,
				},
			},
		}

		userService := mock_user.NewMockService(ctrl)
		userService.EXPECT().GetByPK(ctx, masterTx, gomock.Any()).Return(dummyUser, nil).Times(2)

		profileService := mock_profile.NewMockService(ctrl)
		profileService.EXPECT().GetByUserID(ctx, masterTx, gomock.Any()).Return(dummyProfile, nil).Times(2)

		placeService := mock_place.NewMockService(ctrl)
		placeService.EXPECT().GetByID(ctx, masterTx, gomock.Any()).Return(dummyPlace, nil).Times(2)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().GetByCategoryID(ctx, masterTx, gomock.Any()).Return(dummyWishCards, nil)

		tagService := mock_tag.NewMockService(ctrl)
		tagService.EXPECT().GetByWishCardID(ctx, masterTx, gomock.Any()).Return(dummyTagSlice, nil).Times(2)

		interactor := New(masterTxManager, wishCardService, userService, profileService, tagService, placeService)

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

		dummyWishCard := &wishCardEntity.Entity{
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
			DeletedAt:   nil,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		}

		userService := mock_user.NewMockService(ctrl)
		userService.EXPECT().GetByPK(ctx, masterTx, gomock.Any()).Return(dummyUser, nil)

		profileService := mock_profile.NewMockService(ctrl)
		profileService.EXPECT().GetByUserID(ctx, masterTx, gomock.Any()).Return(dummyProfile, nil)

		placeService := mock_place.NewMockService(ctrl)
		placeService.EXPECT().GetByID(ctx, masterTx, gomock.Any()).Return(dummyPlace, nil)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().UpdateActivity(ctx, masterTx, gomock.Any(), gomock.Any()).Return(dummyWishCard, nil)

		tagService := mock_tag.NewMockService(ctrl)
		tagService.EXPECT().GetByWishCardID(ctx, masterTx, gomock.Any()).Return(dummyTagSlice, nil)

		interactor := New(masterTxManager, wishCardService, userService, profileService, tagService, placeService)

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

		dummyWishCard := &wishCardEntity.Entity{
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
			DeletedAt:   nil,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		}

		userService := mock_user.NewMockService(ctrl)
		userService.EXPECT().GetByPK(ctx, masterTx, gomock.Any()).Return(dummyUser, nil)

		profileService := mock_profile.NewMockService(ctrl)
		profileService.EXPECT().GetByUserID(ctx, masterTx, gomock.Any()).Return(dummyProfile, nil)

		placeService := mock_place.NewMockService(ctrl)
		placeService.EXPECT().GetByID(ctx, masterTx, gomock.Any()).Return(dummyPlace, nil)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().UpdateDescription(ctx, masterTx, gomock.Any(), gomock.Any()).Return(dummyWishCard, nil)

		tagService := mock_tag.NewMockService(ctrl)
		tagService.EXPECT().GetByWishCardID(ctx, masterTx, gomock.Any()).Return(dummyTagSlice, nil)

		interactor := New(masterTxManager, wishCardService, userService, profileService, tagService, placeService)

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

		dummyWishCard := &wishCardEntity.Entity{
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
			DeletedAt:   nil,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		}

		userService := mock_user.NewMockService(ctrl)
		userService.EXPECT().GetByPK(ctx, masterTx, gomock.Any()).Return(dummyUser, nil)

		profileService := mock_profile.NewMockService(ctrl)
		profileService.EXPECT().GetByUserID(ctx, masterTx, gomock.Any()).Return(dummyProfile, nil)

		placeService := mock_place.NewMockService(ctrl)
		placeService.EXPECT().Create(ctx, masterTx, gomock.Any()).Return(dummyPlace, nil)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().UpdatePlace(ctx, masterTx, gomock.Any(), gomock.Any()).Return(dummyWishCard, nil)

		tagService := mock_tag.NewMockService(ctrl)
		tagService.EXPECT().GetByWishCardID(ctx, masterTx, gomock.Any()).Return(dummyTagSlice, nil)

		interactor := New(masterTxManager, wishCardService, userService, profileService, tagService, placeService)

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

		dummyWishCard := &wishCardEntity.Entity{
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
			DeletedAt:   nil,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		}

		userService := mock_user.NewMockService(ctrl)
		userService.EXPECT().GetByPK(ctx, masterTx, gomock.Any()).Return(dummyUser, nil)

		profileService := mock_profile.NewMockService(ctrl)
		profileService.EXPECT().GetByUserID(ctx, masterTx, gomock.Any()).Return(dummyProfile, nil)

		placeService := mock_place.NewMockService(ctrl)
		placeService.EXPECT().GetByID(ctx, masterTx, gomock.Any()).Return(dummyPlace, nil)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().UpdateDate(ctx, masterTx, gomock.Any(), gomock.Any()).Return(dummyWishCard, nil)

		tagService := mock_tag.NewMockService(ctrl)
		tagService.EXPECT().GetByWishCardID(ctx, masterTx, gomock.Any()).Return(dummyTagSlice, nil)

		interactor := New(masterTxManager, wishCardService, userService, profileService, tagService, placeService)

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

		dummyWishCard := &wishCardEntity.Entity{
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
			DeletedAt:   nil,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		}

		userService := mock_user.NewMockService(ctrl)
		userService.EXPECT().GetByPK(ctx, masterTx, gomock.Any()).Return(dummyUser, nil)

		profileService := mock_profile.NewMockService(ctrl)
		profileService.EXPECT().GetByUserID(ctx, masterTx, gomock.Any()).Return(dummyProfile, nil)

		placeService := mock_place.NewMockService(ctrl)
		placeService.EXPECT().GetByID(ctx, masterTx, gomock.Any()).Return(dummyPlace, nil)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().AddTags(ctx, masterTx, gomock.Any(), gomock.Any()).Return(dummyWishCard, nil)

		tagService := mock_tag.NewMockService(ctrl)
		tagService.EXPECT().GetByName(ctx, masterTx, dummyTagName1).Return(dummyTag1, nil)
		tagService.EXPECT().GetByName(ctx, masterTx, dummyTagName2).Return(nil, nil)
		tagService.EXPECT().Create(ctx, masterTx, dummyTagName2).Return(dummyTag2, nil)
		tagService.EXPECT().GetByWishCardID(ctx, masterTx, gomock.Any()).Return(dummyTagSlice, nil)

		interactor := New(masterTxManager, wishCardService, userService, profileService, tagService, placeService)

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

		dummyWishCard := &wishCardEntity.Entity{
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
			DeletedAt:   nil,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		}

		userService := mock_user.NewMockService(ctrl)
		userService.EXPECT().GetByPK(ctx, masterTx, gomock.Any()).Return(dummyUser, nil)

		profileService := mock_profile.NewMockService(ctrl)
		profileService.EXPECT().GetByUserID(ctx, masterTx, gomock.Any()).Return(dummyProfile, nil)

		placeService := mock_place.NewMockService(ctrl)
		placeService.EXPECT().GetByID(ctx, masterTx, gomock.Any()).Return(dummyPlace, nil)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().DeleteTags(ctx, masterTx, gomock.Any(), gomock.Any()).Return(dummyWishCard, nil)

		tagService := mock_tag.NewMockService(ctrl)
		tagService.EXPECT().GetByWishCardID(ctx, masterTx, gomock.Any()).Return(dummyTagSlice, nil)

		interactor := New(masterTxManager, wishCardService, userService, profileService, tagService, placeService)

		result, err := interactor.DeleteTags(ctx, 1, 1, []int{1, 2})

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

// validation test
func TestIntereractor_validateActivity(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err := validateActivity("activityyyy")
		assert.NoError(t, err)
	})

	t.Run("success_文字列制限ぴったり", func(t *testing.T) {
		err := validateActivity(strings.Repeat("a", 50))
		assert.NoError(t, err)
	})

	t.Run("failure_空値", func(t *testing.T) {
		err := validateActivity("")
		assert.Error(t, err)
	})

	t.Run("failure_文字列超過", func(t *testing.T) {
		err := validateActivity(strings.Repeat("a", 51))
		assert.Error(t, err)
	})
}

func TestIntereractor_validateDescription(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err := validateDescription("sampleDescription")
		assert.NoError(t, err)
	})

	t.Run("success_文字列制限ぴったり", func(t *testing.T) {
		err := validateDescription(strings.Repeat("a", 100))
		assert.NoError(t, err)
	})

	t.Run("success_空値", func(t *testing.T) {
		err := validateDescription("")
		assert.NoError(t, err)
	})

	t.Run("failure_文字列超過", func(t *testing.T) {
		err := validateDescription(strings.Repeat("a", 101))
		assert.Error(t, err)
	})
}

func TestIntereractor_validatePlace(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err := validatePlace("samplePlace")
		assert.NoError(t, err)
	})

	t.Run("success_文字列制限ぴったり", func(t *testing.T) {
		err := validatePlace(strings.Repeat("a", 200))
		assert.NoError(t, err)
	})

	t.Run("failure_空値", func(t *testing.T) {
		err := validatePlace("")
		assert.Error(t, err)
	})

	t.Run("failure_文字列超過", func(t *testing.T) {
		err := validatePlace(strings.Repeat("a", 201))
		assert.Error(t, err)
	})
}

func TestIntereractor_validateDate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		future := time.Now().AddDate(1, 0, 0)
		err := validateDate(&future)
		assert.NoError(t, err)
	})

	t.Run("failure_空値", func(t *testing.T) {
		err := validateDate(nil)
		assert.Error(t, err)
	})

	t.Run("failure_過去の日付を指定", func(t *testing.T) {
		past := time.Date(1990, 9, 1, 12, 0, 0, 0, time.Local)
		err := validateDate(&past)
		assert.Error(t, err)
	})
}

func TestIntereractor_validateTags(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tags := []string{"sample1", "sample2", "sample3"}
		err := validateTags(tags)
		assert.NoError(t, err)
	})

	t.Run("failure_エラー混じり", func(t *testing.T) {
		tags := []string{"sample1", "sample2", strings.Repeat("a", 101)}
		err := validateTags(tags)
		assert.Error(t, err)
	})
}

func TestIntereractor_validateTag(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err := validateTag("sampleTag")
		assert.NoError(t, err)
	})

	t.Run("success_文字列制限ぴったり", func(t *testing.T) {
		err := validateTag(strings.Repeat("a", 100))
		assert.NoError(t, err)
	})

	t.Run("failure_空値", func(t *testing.T) {
		err := validateTag("")
		assert.Error(t, err)
	})

	t.Run("failure_文字列超過", func(t *testing.T) {
		err := validateTag(strings.Repeat("a", 101))
		assert.Error(t, err)
	})
}
