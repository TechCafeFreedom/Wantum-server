package tag

import (
	"context"
	"os"
	"testing"
	"time"
	tagEntity "wantum/pkg/domain/entity/tag"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/tag/mock_tag"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	masterTx  repository.MasterTx
	dummyDate time.Time

	dummyTag = &tagEntity.Entity{
		ID:        1,
		Name:      "sampleTag",
		CreatedAt: &dummyDate,
		UpdatedAt: &dummyDate,
	}

	dummyTagSlice = tagEntity.EntitySlice{
		&tagEntity.Entity{
			ID:        1,
			Name:      "sampleTag1",
			CreatedAt: &dummyDate,
			UpdatedAt: &dummyDate,
		},
		&tagEntity.Entity{
			ID:        2,
			Name:      "sampleTag2",
			CreatedAt: &dummyDate,
			UpdatedAt: &dummyDate,
		},
	}
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

	repo := mock_tag.NewMockRepository(ctrl)
	repo.EXPECT().Insert(ctx, masterTx, gomock.Any()).Return(1, nil)

	service := New(repo)
	result, err := service.Create(ctx, masterTx, "sampleTag")

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestService_Delete(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_tag.NewMockRepository(ctrl)
	repo.EXPECT().UpDeleteFlag(ctx, masterTx, gomock.Any()).Return(nil)
	repo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(dummyTag, nil)

	service := New(repo)
	result, err := service.Delete(ctx, masterTx, dummyTag.ID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.DeletedAt)
}

func TestService_GetByID(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_tag.NewMockRepository(ctrl)
	repo.EXPECT().SelectByID(ctx, masterTx, gomock.Any()).Return(dummyTag, nil)

	service := New(repo)
	result, err := service.GetByID(ctx, masterTx, dummyTag.ID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestService_GetByWishCardID(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_tag.NewMockRepository(ctrl)
	repo.EXPECT().SelectByWishCardID(ctx, masterTx, 1).Return(dummyTagSlice, nil)

	service := New(repo)
	result, err := service.GetByWishCardID(ctx, masterTx, 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestService_GetByMemoryID(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_tag.NewMockRepository(ctrl)
	repo.EXPECT().SelectByMemoryID(ctx, masterTx, 1).Return(dummyTagSlice, nil)

	service := New(repo)
	result, err := service.GetByMemoryID(ctx, masterTx, 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}
