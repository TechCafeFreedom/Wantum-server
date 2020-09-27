package place

import (
	"context"
	"os"
	"testing"
	"time"
	placeEntity "wantum/pkg/domain/entity/place"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/place/mock_place"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	masterTx  repository.MasterTx
	dummyDate time.Time

	dummyPlaceID   = 1
	dummyPlaceName = "tsushima"
)

func TestMain(m *testing.M) {
	before()
	code := m.Run()
	os.Exit(code)
}

func before() {
	dummyDate = time.Date(2020, 9, 1, 12, 0, 0, 0, time.Local)
	masterTx = repository.NewMockMasterTx()
}

func TestService_Create(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_place.NewMockRepository(ctrl)
	repo.EXPECT().Insert(ctx, masterTx, gomock.Any()).Return(1, nil)

	service := New(repo)
	result, err := service.Create(ctx, masterTx, "tokyo")

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestService_Update(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyData := &placeEntity.Entity{
		ID:        1,
		Name:      "tokyo",
		CreatedAt: &dummyDate,
		UpdatedAt: &dummyDate,
	}

	repo := mock_place.NewMockRepository(ctrl)
	repo.EXPECT().Update(ctx, masterTx, gomock.Any()).Return(nil)
	repo.EXPECT().SelectByID(ctx, masterTx, dummyData.ID).Return(dummyData, nil)

	service := New(repo)
	result, err := service.Update(ctx, masterTx, dummyData.ID, "shibuya")

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestService_UpdateName(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_place.NewMockRepository(ctrl)
	repo.EXPECT().UpdateName(ctx, masterTx, gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	service := New(repo)
	err := service.UpdateName(ctx, masterTx, dummyPlaceID, dummyPlaceName)

	assert.NoError(t, err)
}

func TestService_Delete(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyData := &placeEntity.Entity{
		ID:        1,
		Name:      "tokyo",
		CreatedAt: &dummyDate,
		UpdatedAt: &dummyDate,
	}

	repo := mock_place.NewMockRepository(ctrl)
	repo.EXPECT().UpDeleteFlag(ctx, masterTx, gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	repo.EXPECT().SelectByID(ctx, masterTx, dummyData.ID).Return(dummyData, nil)

	service := New(repo)
	result, err := service.Delete(ctx, masterTx, dummyData.ID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.DeletedAt)
}

func TestService_GetByID(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyData := &placeEntity.Entity{
		ID:        1,
		Name:      "tokyo",
		CreatedAt: &dummyDate,
		UpdatedAt: &dummyDate,
	}

	repo := mock_place.NewMockRepository(ctrl)
	repo.EXPECT().SelectByID(ctx, masterTx, dummyData.ID).Return(dummyData, nil)

	service := New(repo)
	result, err := service.GetByID(ctx, masterTx, dummyData.ID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestService_GetAll(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyData := placeEntity.EntitySlice{
		&placeEntity.Entity{
			ID:        1,
			Name:      "tokyo",
			CreatedAt: &dummyDate,
			UpdatedAt: &dummyDate,
		},
		&placeEntity.Entity{
			ID:        2,
			Name:      "shibuya",
			CreatedAt: &dummyDate,
			UpdatedAt: &dummyDate,
			DeletedAt: &dummyDate,
		},
	}

	repo := mock_place.NewMockRepository(ctrl)
	repo.EXPECT().SelectAll(ctx, masterTx).Return(dummyData, nil)

	service := New(repo)
	result, err := service.GetAll(ctx, masterTx)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}
