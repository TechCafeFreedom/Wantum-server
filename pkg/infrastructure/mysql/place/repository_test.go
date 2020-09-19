package place

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os"
	"testing"
	"time"
	placeEntity "wantum/pkg/domain/entity/place"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/place"
	tx "wantum/pkg/infrastructure/mysql"
	"wantum/pkg/testutil"

	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db         *sql.DB
	txManager  repository.MasterTxManager
	repo       place.Repository
	dummyDate  time.Time
	dummyPlace = "samplePlace"
)

// TODO: それぞれの関数で使っているdummyDataの切り出し

func TestMain(m *testing.M) {
	before()
	code := m.Run()
	after()
	os.Exit(code)
}

// repositoryを作ってもらう
func before() {
	var err error
	db, err = testutil.ConnectLocalDB()
	if err != nil {
		log.Fatal("faild to connect db: ", err)
	}
	txManager = tx.NewDBMasterTxManager(db)
	repo = New(txManager)
	dummyDate = time.Date(2020, 9, 1, 12, 0, 0, 0, time.Local)
}

// dbのコネクションを閉じる
func after() {
	db.Close()
}

func TestInsert(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var err error
		ctx := context.Background()
		place := &placeEntity.Entity{
			Name:      dummyPlace,
			CreatedAt: &dummyDate,
			UpdatedAt: &dummyDate,
		}

		var result int
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			result, err = repo.Insert(ctx, masterTx, place)
			return err
		})

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var err error
		ctx := context.Background()
		place := &placeEntity.Entity{
			ID:        1,
			Name:      dummyPlace,
			CreatedAt: &dummyDate,
			UpdatedAt: &dummyDate,
		}

		var result *placeEntity.Entity
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			if err = repo.Update(ctx, masterTx, place); err != nil {
				return err
			}

			result, err = repo.SelectByID(ctx, masterTx, 1)
			return err
		})

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, dummyPlace, result.Name)
	})
}

func TestUpDeleteFlag(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var err error
		ctx := context.Background()
		place := &placeEntity.Entity{
			Name:      dummyPlace,
			CreatedAt: &dummyDate,
			UpdatedAt: &dummyDate,
		}

		var result *placeEntity.Entity
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			newPlaceID, _ := repo.Insert(ctx, masterTx, place)

			place.ID = newPlaceID
			place.DeletedAt = &dummyDate
			if err = repo.UpDeleteFlag(ctx, masterTx, place); err != nil {
				return err
			}

			result, _ = repo.SelectByID(ctx, masterTx, place.ID)
			return nil
		})

		assert.NoError(t, err)
		assert.NotNil(t, result.DeletedAt)
	})

	t.Run("failure_deletedAtがnil", func(t *testing.T) {
		var err error
		ctx := context.Background()
		place := &placeEntity.Entity{
			Name:      dummyPlace,
			CreatedAt: &dummyDate,
			UpdatedAt: &dummyDate,
		}

		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			newPlaceID, _ := repo.Insert(ctx, masterTx, place)

			place.ID = newPlaceID
			return repo.UpDeleteFlag(ctx, masterTx, place)
		})

		assert.Error(t, err)
	})
}

func TestDownDeleteFlag(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var err error
		ctx := context.Background()
		place := &placeEntity.Entity{
			Name:      dummyPlace,
			CreatedAt: &dummyDate,
			UpdatedAt: &dummyDate,
			DeletedAt: &dummyDate,
		}

		var result *placeEntity.Entity
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			newPlaceID, _ := repo.Insert(ctx, masterTx, place)

			place.ID = newPlaceID
			place.DeletedAt = nil
			if err = repo.DownDeleteFlag(ctx, masterTx, place); err != nil {
				return err
			}

			result, _ = repo.SelectByID(ctx, masterTx, place.ID)
			return nil
		})

		assert.NoError(t, err)
		assert.Nil(t, result.DeletedAt)
	})
}

func TestDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var err error
		ctx := context.Background()
		place := &placeEntity.Entity{
			Name:      dummyPlace,
			CreatedAt: &dummyDate,
			UpdatedAt: &dummyDate,
		}

		var result *placeEntity.Entity
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			newPlaceID, _ := repo.Insert(ctx, masterTx, place)
			place.ID = newPlaceID
			place.DeletedAt = &dummyDate
			repo.UpDeleteFlag(ctx, masterTx, place)

			if err = repo.Delete(ctx, masterTx, place.ID); err != nil {
				return err
			}

			result, err = repo.SelectByID(ctx, masterTx, place.ID)
			if err == nil {
				return errors.New("削除されたデータが引っかかった")
			}
			return nil
		})

		assert.NoError(t, err)
		assert.Nil(t, result)
	})
}

func TestSelectByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var err error
		ctx := context.Background()
		place := &placeEntity.Entity{
			Name:      dummyPlace,
			CreatedAt: &dummyDate,
			UpdatedAt: &dummyDate,
		}

		var result *placeEntity.Entity
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			newPlaceID, _ := repo.Insert(ctx, masterTx, place)

			result, err = repo.SelectByID(ctx, masterTx, newPlaceID)
			return err
		})

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

}

func TestSelectAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var err error
		ctx := context.Background()

		var result placeEntity.EntitySlice
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			result, err = repo.SelectAll(ctx, masterTx)
			return err
		})

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}
