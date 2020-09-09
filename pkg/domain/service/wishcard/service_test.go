package wishcard

import (
	"context"
	"os"
	"testing"
	"time"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/wishcard/mock_wish_card"
	"wantum/pkg/infrastructure/mysql/model"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	masterTx repository.MasterTx
)

var (
	dummyDate        = time.Date(2020, 9, 1, 12, 0, 0, 0, time.Local)
	dummyActivity    = "sampleActivity"
	dummyDescription = "sampleDescription"
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

	repo := mock_wish_card.NewMockRepository(ctrl)
	repo.EXPECT().Insert(ctx, masterTx, gomock.Any()).Return(1, nil)

	service := New(repo)
	result, err := service.Create(ctx, masterTx, dummyActivity, dummyDescription, &dummyDate, 1, 1, 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
}

func TestService_Update(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyData := &model.WishCardModel{
		ID:          1,
		UserID:      1,
		Activity:    "act",
		Description: "desc",
		Date:        &dummyDate,
		DoneAt:      &dummyDate,
		CategoryID:  1,
		PlaceID:     1,
		CreatedAt:   &dummyDate,
		UpdatedAt:   &dummyDate,
	}

	repo := mock_wish_card.NewMockRepository(ctrl)
	repo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(dummyData, nil)
	repo.EXPECT().Update(ctx, masterTx, gomock.Any()).Return(nil)

	service := New(repo)
	result, err := service.Update(ctx, masterTx, 1, dummyActivity, dummyDescription, &dummyDate, &dummyDate, 1, 1, 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, dummyActivity, result.Activity)
	assert.Equal(t, dummyDescription, result.Description)

}

func TestService_UpDeleteFlag(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyData := &model.WishCardModel{
		ID:          1,
		UserID:      1,
		Activity:    dummyActivity,
		Description: dummyDescription,
		Date:        &dummyDate,
		DoneAt:      &dummyDate,
		CategoryID:  1,
		PlaceID:     1,
		CreatedAt:   &dummyDate,
		UpdatedAt:   &dummyDate,
	}

	repo := mock_wish_card.NewMockRepository(ctrl)
	repo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(dummyData, nil)
	repo.EXPECT().UpDeleteFlag(ctx, masterTx, gomock.Any()).Return(nil)

	service := New(repo)
	result, err := service.UpDeleteFlag(ctx, masterTx, 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.DeletedAt)
}

func TestService_DownDeleteFlag(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyData := &model.WishCardModel{
		ID:          1,
		UserID:      1,
		Activity:    dummyActivity,
		Description: dummyDescription,
		Date:        &dummyDate,
		DoneAt:      &dummyDate,
		CategoryID:  1,
		PlaceID:     1,
		CreatedAt:   &dummyDate,
		UpdatedAt:   &dummyDate,
		DeletedAt:   &dummyDate,
	}

	repo := mock_wish_card.NewMockRepository(ctrl)
	repo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(dummyData, nil)
	repo.EXPECT().DownDeleteFlag(ctx, masterTx, gomock.Any()).Return(nil)

	service := New(repo)
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

		dummyData := &model.WishCardModel{
			ID:          1,
			UserID:      1,
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CategoryID:  1,
			PlaceID:     1,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			DeletedAt:   &dummyDate,
		}

		repo := mock_wish_card.NewMockRepository(ctrl)
		repo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(dummyData, nil)
		repo.EXPECT().Delete(ctx, masterTx, gomock.Any()).Return(nil)

		service := New(repo)
		err := service.Delete(ctx, masterTx, 1)

		assert.NoError(t, err)
	})

	t.Run("failure_deleteフラグがたってない", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dummyData := &model.WishCardModel{
			ID:          1,
			UserID:      1,
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CategoryID:  1,
			PlaceID:     1,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
		}

		repo := mock_wish_card.NewMockRepository(ctrl)
		repo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(dummyData, nil)

		service := New(repo)
		err := service.Delete(ctx, masterTx, 1)

		assert.Error(t, err)
	})
}

func TestService_GetByID(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyData := &model.WishCardModel{
		ID:          1,
		UserID:      1,
		Activity:    dummyActivity,
		Description: dummyDescription,
		Date:        &dummyDate,
		DoneAt:      &dummyDate,
		CategoryID:  1,
		PlaceID:     1,
		CreatedAt:   &dummyDate,
		UpdatedAt:   &dummyDate,
	}

	repo := mock_wish_card.NewMockRepository(ctrl)
	repo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(dummyData, nil)

	service := New(repo)
	result, err := service.GetByID(ctx, masterTx, 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestService_GetByIDs(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyData := model.WishCardModelSlice{
		&model.WishCardModel{
			ID:          1,
			UserID:      1,
			Activity:    "activity1",
			Description: "desc2",
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CategoryID:  1,
			PlaceID:     1,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
		},
		&model.WishCardModel{
			ID:          2,
			UserID:      1,
			Activity:    "activity2",
			Description: "desc2",
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CategoryID:  2,
			PlaceID:     2,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
		},
	}

	repo := mock_wish_card.NewMockRepository(ctrl)
	repo.EXPECT().SelectByIDs(ctx, masterTx, gomock.Any()).Return(dummyData, nil)

	service := New(repo)
	result, err := service.GetByIDs(ctx, masterTx, []int{1, 2})

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(result))
}

func TestService_GetByCategoryID(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyData := model.WishCardModelSlice{
		&model.WishCardModel{
			ID:          1,
			UserID:      1,
			Activity:    "activity1",
			Description: "desc1",
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CategoryID:  1,
			PlaceID:     1,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
		},
		&model.WishCardModel{
			ID:          2,
			UserID:      1,
			Activity:    "activity2",
			Description: "desc2",
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CategoryID:  1,
			PlaceID:     2,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
		},
	}

	repo := mock_wish_card.NewMockRepository(ctrl)
	repo.EXPECT().SelectByCategoryID(ctx, masterTx, gomock.Any()).Return(dummyData, nil)

	service := New(repo)
	result, err := service.GetByCategoryID(ctx, masterTx, 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(result))
}
