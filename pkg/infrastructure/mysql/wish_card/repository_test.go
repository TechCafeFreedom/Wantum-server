package wish_card

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os"
	"testing"
	"time"
	"wantum/pkg/domain/repository"
	wcrepo "wantum/pkg/domain/repository/wish_card"
	tx "wantum/pkg/infrastructure/mysql"
	"wantum/pkg/infrastructure/mysql/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

var db *sql.DB
var txManager repository.MasterTxManager
var repo wcrepo.Repository
var dummyDate time.Time

func TestMain(m *testing.M) {
	before()
	code := m.Run()
	after()
	os.Exit(code)
}

func before() {
	var err error
	// TODO: 環境変数とか使いたい気持ちもする
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/wantum?parseTime=true")
	if err != nil {
		log.Fatal("faild to connect db: ", err)
	}
	txManager = tx.NewDBMasterTxManager(db)
	repo = New(txManager)
	dummyDate = time.Date(2020, 9, 1, 12, 0, 0, 0, time.Local)
}

func after() {
	db.Close()
}

func TestInsert(t *testing.T) {
	t.Run("success to insert data", func(t *testing.T) {
		var err error
		ctx := context.Background()
		wishCard := &model.WishCardModel{
			UserID:      1,
			Activity:    "なんかしたい",
			Description: "何かがしたい",
			Date:        &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			CategoryID:  1,
			PlaceID:     1,
		}
		var result int
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			result, err = repo.Insert(ctx, masterTx, wishCard)
			return err
		})
		assert.NoError(t, err)
		assert.NotEqual(t, 0, result)
	})

	t.Run("failure to insert data. data is nil", func(t *testing.T) {
		var err error
		ctx := context.Background()

		var result int
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			result, err = repo.Insert(ctx, masterTx, nil)
			return err
		})
		assert.Error(t, err)
		assert.Equal(t, 0, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("success to update data", func(t *testing.T) {
		var err error
		ctx := context.Background()
		wishCard := &model.WishCardModel{
			ID:          1,
			UserID:      1,
			Activity:    "なんかしたい",
			Description: "何かがしたい",
			Date:        &dummyDate,
			DoneAt:      &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			CategoryID:  1,
			PlaceID:     1,
		}
		var result *model.WishCardModel
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			err = repo.Update(ctx, masterTx, wishCard)
			if err != nil {
				return err
			}

			result, _ = repo.SelectByID(ctx, masterTx, 1)
			return nil
		})
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "なんかしたい", result.Activity)
	})

	t.Run("success to update data. done at is null", func(t *testing.T) {
		var err error
		ctx := context.Background()
		wishCard := &model.WishCardModel{
			ID:          1,
			UserID:      1,
			Activity:    "なんかしたい",
			Description: "何かがしたい",
			Date:        &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			CategoryID:  1,
			PlaceID:     1,
		}
		var result *model.WishCardModel
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			err = repo.Update(ctx, masterTx, wishCard)
			if err != nil {
				return err
			}

			result, _ = repo.SelectByID(ctx, masterTx, 1)
			return nil
		})
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "なんかしたい", result.Activity)
	})

	t.Run("failure to update data. data is nil", func(t *testing.T) {
		var err error
		ctx := context.Background()

		var result *model.WishCardModel
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			err = repo.Update(ctx, masterTx, nil)
			return err
		})
		assert.Error(t, err)
		assert.Nil(t, result)
	})

}

func TestUpDeleteFlag(t *testing.T) {
	t.Run("success to up delete flag", func(t *testing.T) {
		var err error
		ctx := context.Background()
		wishCard := &model.WishCardModel{
			UserID:      1,
			Activity:    "なんかしたい",
			Description: "何かがしたい",
			Date:        &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			CategoryID:  1,
			PlaceID:     1,
		}
		var result *model.WishCardModel
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			newID, _ := repo.Insert(ctx, masterTx, wishCard)

			wishCard.ID = newID
			wishCard.DeletedAt = &dummyDate
			err = repo.UpDeleteFlag(ctx, masterTx, wishCard)
			if err != nil {
				return err
			}
			result, _ = repo.SelectByID(ctx, masterTx, wishCard.ID)
			return nil
		})
		assert.NoError(t, err)
		assert.NotNil(t, result.DeletedAt)
	})

	t.Run("failure to up delete flag. flag is nil", func(t *testing.T) {
		var err error
		ctx := context.Background()
		wishCard := &model.WishCardModel{
			UserID:      1,
			Activity:    "なんかしたい",
			Description: "何かがしたい",
			Date:        &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			CategoryID:  1,
			PlaceID:     1,
		}
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			newID, _ := repo.Insert(ctx, masterTx, wishCard)

			wishCard.ID = newID
			err = repo.UpDeleteFlag(ctx, masterTx, wishCard)
			return err
		})
		assert.Error(t, err)
	})

	t.Run("failure to up delete flag. data is nil", func(t *testing.T) {
		var err error
		ctx := context.Background()

		var result *model.WishCardModel
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			err = repo.UpDeleteFlag(ctx, masterTx, nil)
			return err
		})
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestDownDeleteFlag(t *testing.T) {
	t.Run("success to down delete flag", func(t *testing.T) {
		var err error
		ctx := context.Background()
		wishCard := &model.WishCardModel{
			UserID:      1,
			Activity:    "なんかしたい",
			Description: "何かがしたい",
			Date:        &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			DeletedAt:   &dummyDate,
			CategoryID:  1,
			PlaceID:     1,
		}
		var result *model.WishCardModel
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			newID, _ := repo.Insert(ctx, masterTx, wishCard)

			wishCard.ID = newID
			err = repo.DownDeleteFlag(ctx, masterTx, wishCard)
			if err != nil {
				return err
			}
			result, _ = repo.SelectByID(ctx, masterTx, wishCard.ID)
			return nil
		})
		assert.NoError(t, err)
		assert.Nil(t, result.DeletedAt)
	})

	t.Run("failure to up delete flag. data is nil", func(t *testing.T) {
		var err error
		ctx := context.Background()

		var result *model.WishCardModel
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			err = repo.UpDeleteFlag(ctx, masterTx, nil)
			return err
		})
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestDelete(t *testing.T) {
	t.Run("success to up delete flag", func(t *testing.T) {
		var err error
		ctx := context.Background()
		wishCard := &model.WishCardModel{
			UserID:      1,
			Activity:    "なんかしたい",
			Description: "何かがしたい",
			Date:        &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			CategoryID:  1,
			PlaceID:     1,
		}
		var result *model.WishCardModel
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			newID, _ := repo.Insert(ctx, masterTx, wishCard)
			wishCard.ID = newID
			wishCard.DeletedAt = &dummyDate
			repo.UpDeleteFlag(ctx, masterTx, wishCard)

			err = repo.Delete(ctx, masterTx, wishCard.ID)
			if err != nil {
				return err
			}

			result, err = repo.SelectByID(ctx, masterTx, wishCard.ID)
			if err == nil {
				return errors.New("削除されたデータが見つかった")
			}
			return nil
		})
		assert.NoError(t, err)
		assert.Nil(t, result)
	})
}

func TestSelectByID(t *testing.T) {
	t.Run("success to select data", func(t *testing.T) {
		var err error
		ctx := context.Background()
		wishCard := &model.WishCardModel{
			UserID:      1,
			Activity:    "なんかしたい",
			Description: "何かがしたい",
			Date:        &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			CategoryID:  1,
			PlaceID:     1,
		}
		var result *model.WishCardModel
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			newID, _ := repo.Insert(ctx, masterTx, wishCard)

			result, err = repo.SelectByID(ctx, masterTx, newID)
			return err
		})
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("failure to select data. data is not exist", func(t *testing.T) {
		var err error
		ctx := context.Background()

		var result *model.WishCardModel
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			result, err = repo.SelectByID(ctx, masterTx, -1)
			return err
		})
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestSelectByIDs(t *testing.T) {
	t.Run("success to select data", func(t *testing.T) {
		var err error
		ctx := context.Background()

		ids := []string{"1", "2", "3"}

		var result model.WishCardModelSlice
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {

			result, err = repo.SelectByIDs(ctx, masterTx, ids)
			return err
		})
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 3, len(result))
	})
}

func TestCategoryID(t *testing.T) {
	t.Run("success to select data", func(t *testing.T) {
		var err error
		ctx := context.Background()
		wishCard := &model.WishCardModel{
			UserID:      1,
			Activity:    "なんかしたい",
			Description: "何かがしたい",
			Date:        &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			CategoryID:  1,
			PlaceID:     1,
		}
		var result model.WishCardModelSlice
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			repo.Insert(ctx, masterTx, wishCard)

			result, err = repo.SelectByCategoryID(ctx, masterTx, 1)
			return err
		})
		assert.NoError(t, err)
		assert.NotNil(t, result)
		for _, row := range result {
			assert.Equal(t, 1, row.CategoryID)
		}
	})

	t.Run("success to select data. category is not exist", func(t *testing.T) {
		var err error
		ctx := context.Background()

		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {

			_, err = repo.SelectByCategoryID(ctx, masterTx, -1)
			return err
		})
		assert.NoError(t, err)
	})
}
