package place

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os"
	"testing"
	"time"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/place"
	tx "wantum/pkg/infrastructure/mysql"
	"wantum/pkg/infrastructure/mysql/model"

	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var txManager repository.MasterTxManager
var repo place.Repository

func TestMain(m *testing.M) {
	before()
	code := m.Run()
	after()
	os.Exit(code)
}

// repositoryを作ってもらう
func before() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/wantum?parseTime=true")
	if err != nil {
		log.Fatal("faild to connect db: ", err)
	}
	txManager = tx.NewDBMasterTxManager(db)
	repo = New(txManager)
}

// dbのコネクションを閉じる
func after() {
	db.Close()
}

func TestInsert(t *testing.T) {
	t.Run("success to insert data", func(t *testing.T) {
		var err error
		ctx := context.Background()
		date := time.Date(2020, 9, 1, 12, 0, 0, 0, time.Local)
		place := &model.PlaceModel{
			Name:      "sample place",
			CreatedAt: &date,
			UpdatedAt: &date,
		}

		var result *model.PlaceModel
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			result, err = repo.Insert(ctx, masterTx, place)
			return err
		})

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotNil(t, result.ID)
	})

	t.Run("failed to insert data. data is nil", func(t *testing.T) {
		var err error
		ctx := context.Background()

		var result *model.PlaceModel
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			result, err = repo.Insert(ctx, masterTx, nil)
			return err
		})

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("success to update data", func(t *testing.T) {
		var err error
		ctx := context.Background()
		date := time.Date(2020, 9, 1, 12, 0, 0, 0, time.Local)
		place := &model.PlaceModel{
			ID:        1,
			Name:      "sample place",
			CreatedAt: &date,
			UpdatedAt: &date,
		}

		var result *model.PlaceModel
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			err = repo.Update(ctx, masterTx, place)
			if err != nil {
				return err
			}

			result, err = repo.SelectByID(ctx, masterTx, 1)
			return err
		})

		assert.NoError(t, err)
		assert.Equal(t, "sample place", result.Name)
	})

	t.Run("failure to update data. data is nil", func(t *testing.T) {
		var err error
		ctx := context.Background()

		var result *model.PlaceModel
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			err = repo.Update(ctx, masterTx, nil)
			if err != nil {
				return err
			}

			result, err = repo.SelectByID(ctx, masterTx, 1)
			return err
		})

		assert.Error(t, err)
		assert.Nil(t, result)

	})
}

func TestUpDeleteFlag(t *testing.T) {
	t.Run("success to up deleteFlag", func(t *testing.T) {
		var err error
		ctx := context.Background()
		date := time.Date(2020, 9, 1, 12, 0, 0, 0, time.Local)
		place := &model.PlaceModel{
			Name:      "sample place",
			CreatedAt: &date,
			UpdatedAt: &date,
		}

		var result *model.PlaceModel
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			newPlace, _ := repo.Insert(ctx, masterTx, place)
			assert.Nil(t, newPlace.DeletedAt)

			newPlace.DeletedAt = &date
			newPlace.UpdatedAt = &date
			err = repo.UpDeleteFlag(ctx, masterTx, newPlace)
			if err != nil {
				return err
			}

			result, _ = repo.SelectByID(ctx, masterTx, newPlace.ID)
			return nil
		})

		assert.NoError(t, err)
		assert.Nil(t, date, result.DeletedAt)
	})
}

func TestDelete(t *testing.T) {
	t.Run("success to delete", func(t *testing.T) {
		var err error
		ctx := context.Background()
		date := time.Date(2020, 9, 1, 12, 0, 0, 0, time.Local)
		place := &model.PlaceModel{
			Name:      "sample place",
			CreatedAt: &date,
			UpdatedAt: &date,
		}

		var result *model.PlaceModel
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			newData, _ := repo.Insert(ctx, masterTx, place)
			newData.DeletedAt = &date
			repo.UpDeleteFlag(ctx, masterTx, newData)

			err = repo.Delete(ctx, masterTx, result.ID)
			if err != nil {
				return err
			}

			result, err = repo.SelectByID(ctx, masterTx, result.ID)
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
	t.Run("success to select by id", func(t *testing.T) {
		var err error
		ctx := context.Background()
		date := time.Date(2020, 9, 1, 12, 0, 0, 0, time.Local)
		place := &model.PlaceModel{
			Name:      "sample place",
			CreatedAt: &date,
			UpdatedAt: &date,
		}

		var result *model.PlaceModel
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			newData, _ := repo.Insert(ctx, masterTx, place)

			result, err = repo.SelectByID(ctx, masterTx, newData.ID)
			return err
		})

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("failure to select by id. id is not exist", func(t *testing.T) {
		var err error
		ctx := context.Background()
		date := time.Date(2020, 9, 1, 12, 0, 0, 0, time.Local)
		place := &model.PlaceModel{
			Name:      "sample place",
			CreatedAt: &date,
			UpdatedAt: &date,
		}

		var result *model.PlaceModel
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			newData, _ := repo.Insert(ctx, masterTx, place)
			repo.Delete(ctx, masterTx, newData.ID)

			result, err = repo.SelectByID(ctx, masterTx, newData.ID)
			return err
		})

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestSelectAll(t *testing.T) {
	t.Run("success to select all", func(t *testing.T) {
		var err error
		ctx := context.Background()

		var result model.PlaceModelSlice
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			result, err = repo.SelectAll(ctx, masterTx)
			return err
		})

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}
