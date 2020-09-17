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
	"wantum/pkg/domain/service/wishcardtag/mock_wish_card_tag"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// TODO: dummyデータを充実させる
// TODO: test名を日本語にする

var (
	masterTx        repository.MasterTx
	masterTxManager repository.MasterTxManager
)

var (
	dummyDate        = time.Date(2020, 9, 1, 12, 0, 0, 0, time.Local)
	dummyActivity    = "dummyActivity"
	dummyDescription = "dummyDescription"
)

var dummyProfile = profileEntity.Entity{
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

var dummyUser = userEntity.Entity{
	ID:        1,
	AuthID:    "dummyID",
	UserName:  "dummyUserName",
	Mail:      "hogehoge@example.com",
	CreatedAt: &dummyDate,
	UpdatedAt: &dummyDate,
	DeletedAt: &dummyDate,
	Profile:   nil,
}

var dummyTag1 = tagEntity.Entity{
	ID:        1,
	Name:      "tag1",
	CreatedAt: &dummyDate,
	UpdatedAt: &dummyDate,
	DeletedAt: &dummyDate,
}

var dummyTag2 = tagEntity.Entity{
	ID:        2,
	Name:      "tag2",
	CreatedAt: &dummyDate,
	UpdatedAt: &dummyDate,
	DeletedAt: nil,
}

var dummyTagSlice = tagEntity.EntitySlice{
	&dummyTag1,
	&dummyTag2,
}

var dummyPlace = placeEntity.Entity{
	ID:        1,
	Name:      "samplePlace",
	CreatedAt: &dummyDate,
	UpdatedAt: &dummyDate,
	DeletedAt: nil,
}

var dummyWishCard = wishCardEntity.Entity{
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
	t.Run("success to create data.", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		placeService := mock_place.NewMockService(ctrl)
		placeService.EXPECT().Create(ctx, masterTx, gomock.Any()).Return(&dummyPlace, nil)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().Create(ctx, masterTx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(&dummyWishCard, nil)

		tagService := mock_tag.NewMockService(ctrl)
		tagService.EXPECT().GetByName(ctx, masterTx, "tag1").Return(&dummyTag1, nil)
		tagService.EXPECT().GetByName(ctx, masterTx, "tag2").Return(nil, nil)
		tagService.EXPECT().Create(ctx, masterTx, "tag2").Return(&dummyTag2, nil)

		wishCardTagService := mock_wish_card_tag.NewMockService(ctrl)
		wishCardTagService.EXPECT().CreateMultipleRelation(ctx, masterTx, gomock.Any(), gomock.Any()).Return(nil)

		interactor := New(masterTxManager, wishCardService, tagService, placeService, wishCardTagService)

		tags := []string{"tag1", "tag2"}
		result, err := interactor.CreateNewWishCard(ctx, 1, "sampleActivity", "sampleDescription", "samplePlace", &dummyDate, 1, tags)

		assert.NoError(t, err)
		assert.NotNil(t, result)

		// validate time
		assert.Equal(t, (*time.Time)(nil), result.DeletedAt)
		assert.Equal(t, (*time.Time)(nil), result.DoneAt)
		// validate place
		assert.Equal(t, "samplePlace", result.Place.Name)
		// validate tag
		assert.Equal(t, 2, len(result.Tags))
		assert.Equal(t, "tag1", result.Tags[0].Name)
	})
}

func TestInteractor_UpdateWishCard(t *testing.T) {
	t.Run("success to update data.", func(t *testing.T) {
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
		placeService.EXPECT().Create(ctx, masterTx, "samplePlace").Return(&dummyPlace, nil)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().Update(ctx, masterTx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(&dummyWishCard, nil)

		tagService := mock_tag.NewMockService(ctrl)
		tagService.EXPECT().GetByName(ctx, masterTx, "tag1").Return(&dummyTag1, nil)
		tagService.EXPECT().GetByName(ctx, masterTx, "tag2").Return(nil, nil)
		tagService.EXPECT().Create(ctx, masterTx, "tag2").Return(&dummyTag2, nil)

		wishCardTagService := mock_wish_card_tag.NewMockService(ctrl)
		wishCardTagService.EXPECT().CreateMultipleRelation(ctx, masterTx, gomock.Any(), gomock.Any()).Return(nil)
		wishCardTagService.EXPECT().DeleteByWishCardID(ctx, masterTx, gomock.Any()).Return(nil)

		interactor := New(masterTxManager, wishCardService, tagService, placeService, wishCardTagService)

		tags := []string{"tag1", "tag2"}
		result, err := interactor.UpdateWishCard(ctx, 1, 1, "sampleActivity", "sampleDescription", "samplePlace", &dummyDate, &dummyDate, 1, tags)

		assert.NoError(t, err)
		assert.NotNil(t, result)

		// validate time
		assert.Equal(t, (*time.Time)(nil), result.DeletedAt)
		assert.NotEqual(t, (*time.Time)(nil), result.DoneAt)
		// validate place
		assert.Equal(t, "samplePlace", result.Place.Name)
		// validate tag
		assert.Equal(t, 2, len(result.Tags))
		assert.Equal(t, "tag1", result.Tags[0].Name)
	})
}

func TestInteractor_DeleteWishCard(t *testing.T) {
	t.Run("success to delete data.", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		placeService := mock_place.NewMockService(ctrl)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().Delete(ctx, masterTx, 1).Return(nil)

		tagService := mock_tag.NewMockService(ctrl)

		wishCardTagService := mock_wish_card_tag.NewMockService(ctrl)
		wishCardTagService.EXPECT().DeleteByWishCardID(ctx, masterTx, gomock.Any()).Return(nil)

		interactor := New(masterTxManager, wishCardService, tagService, placeService, wishCardTagService)

		err := interactor.DeleteWishCardByID(ctx, 1)

		assert.NoError(t, err)
	})
}

func TestInteractor_GetByID(t *testing.T) {
	t.Run("success to get data by id.", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		placeService := mock_place.NewMockService(ctrl)

		wishCardService := mock_wish_card.NewMockService(ctrl)
		wishCardService.EXPECT().GetByID(ctx, masterTx, gomock.Any()).Return(&dummyWishCard, nil)

		tagService := mock_tag.NewMockService(ctrl)

		wishCardTagService := mock_wish_card_tag.NewMockService(ctrl)

		interactor := New(masterTxManager, wishCardService, tagService, placeService, wishCardTagService)

		result, err := interactor.GetByID(ctx, 1)

		assert.NoError(t, err)
		assert.NotNil(t, result)

		// validate place
		assert.Equal(t, "samplePlace", result.Place.Name)
		// validate tag
		assert.Equal(t, 2, len(result.Tags))
		assert.Equal(t, "tag1", result.Tags[0].Name)
	})
}

func TestInteractor_GetByCategoryID(t *testing.T) {
	t.Run("success to get data by categoryID.", func(t *testing.T) {
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

		wishCardTagService := mock_wish_card_tag.NewMockService(ctrl)

		interactor := New(masterTxManager, wishCardService, tagService, placeService, wishCardTagService)

		wishCards, err := interactor.GetByCategoryID(ctx, 1)

		assert.NoError(t, err)
		assert.NotNil(t, wishCards)
		assert.Equal(t, 2, len(wishCards))

		assert.Equal(t, "samplePlace", wishCards[0].Place.Name)
		assert.Equal(t, "samplePlace", wishCards[1].Place.Name)

		assert.Equal(t, "tag1", wishCards[0].Tags[0].Name)
		assert.Equal(t, 2, len(wishCards[0].Tags))
		assert.Equal(t, "tag1", wishCards[1].Tags[0].Name)
		assert.Equal(t, 2, len(wishCards[1].Tags))
	})
}
