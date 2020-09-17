package wishcard

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os"
	"testing"
	"time"
	placeEntity "wantum/pkg/domain/entity/place"
	userEntity "wantum/pkg/domain/entity/user"
	wishCardEntity "wantum/pkg/domain/entity/wishcard"
	"wantum/pkg/domain/repository"
	wcrepo "wantum/pkg/domain/repository/wishcard"
	tx "wantum/pkg/infrastructure/mysql"
	"wantum/pkg/testutil"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

var db *sql.DB
var txManager repository.MasterTxManager
var repo wcrepo.Repository
var dummyDate time.Time

var dummyActivity = "sampleActivity"
var dummyDescription = "sampleDescription"

func TestMain(m *testing.M) {
	before()
	code := m.Run()
	after()
	os.Exit(code)
}

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

func after() {
	db.Close()
}

func TestInsert(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var err error
		ctx := context.Background()
		wishCard := &wishCardEntity.Entity{
			Author: &userEntity.Entity{
				ID: 1,
			},
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		}
		var result int
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			result, err = repo.Insert(ctx, masterTx, wishCard, 1)
			return err
		})
		assert.NoError(t, err)
		assert.NotEqual(t, 0, result)
	})

	t.Run("failure_データがnil", func(t *testing.T) {
		var err error
		ctx := context.Background()

		var result int
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			result, err = repo.Insert(ctx, masterTx, nil, 1)
			return err
		})
		assert.Error(t, err)
		assert.Equal(t, 0, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var err error
		ctx := context.Background()
		wishCard := &wishCardEntity.Entity{
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
		var result *wishCardEntity.Entity
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			err = repo.Update(ctx, masterTx, wishCard, 1)
			if err != nil {
				return err
			}

			result, _ = repo.SelectByID(ctx, masterTx, 1)
			return nil
		})
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, dummyActivity, result.Activity)
	})

	t.Run("success_doneAtがnil", func(t *testing.T) {
		var err error
		ctx := context.Background()
		wishCard := &wishCardEntity.Entity{
			ID: 1,
			Author: &userEntity.Entity{
				ID: 1,
			},
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		}
		var result *wishCardEntity.Entity
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			err = repo.Update(ctx, masterTx, wishCard, 1)
			if err != nil {
				return err
			}

			result, _ = repo.SelectByID(ctx, masterTx, 1)
			return nil
		})
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, dummyActivity, result.Activity)
	})

	t.Run("failure_データがnil", func(t *testing.T) {
		var err error
		ctx := context.Background()

		var result *wishCardEntity.Entity
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			err = repo.Update(ctx, masterTx, nil, 1)
			return err
		})
		assert.Error(t, err)
		assert.Nil(t, result)
	})

}

func TestUpDeleteFlag(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var err error
		ctx := context.Background()
		wishCard := &wishCardEntity.Entity{
			Author: &userEntity.Entity{
				ID: 1,
			},
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		}
		var result *wishCardEntity.Entity
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			newID, _ := repo.Insert(ctx, masterTx, wishCard, 1)

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

	t.Run("failure_deletedAtがnil", func(t *testing.T) {
		var err error
		ctx := context.Background()
		wishCard := &wishCardEntity.Entity{
			Author: &userEntity.Entity{
				ID: 1,
			},
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		}
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			newID, _ := repo.Insert(ctx, masterTx, wishCard, 1)

			wishCard.ID = newID
			err = repo.UpDeleteFlag(ctx, masterTx, wishCard)
			return err
		})
		assert.Error(t, err)
	})

	t.Run("failure_データがnil", func(t *testing.T) {
		var err error
		ctx := context.Background()

		var result *wishCardEntity.Entity
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			err = repo.UpDeleteFlag(ctx, masterTx, nil)
			return err
		})
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestDownDeleteFlag(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var err error
		ctx := context.Background()
		wishCard := &wishCardEntity.Entity{
			Author: &userEntity.Entity{
				ID: 1,
			},
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			DeletedAt:   &dummyDate,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		}
		var result *wishCardEntity.Entity
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			newID, _ := repo.Insert(ctx, masterTx, wishCard, 1)

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

	t.Run("failure_データがnil", func(t *testing.T) {
		var err error
		ctx := context.Background()

		var result *wishCardEntity.Entity
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			err = repo.UpDeleteFlag(ctx, masterTx, nil)
			return err
		})
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var err error
		ctx := context.Background()
		wishCard := &wishCardEntity.Entity{
			Author: &userEntity.Entity{
				ID: 1,
			},
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		}
		var result *wishCardEntity.Entity
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			newID, _ := repo.Insert(ctx, masterTx, wishCard, 1)
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
	t.Run("success", func(t *testing.T) {
		var err error
		ctx := context.Background()
		wishCard := &wishCardEntity.Entity{
			Author: &userEntity.Entity{
				ID: 1,
			},
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		}
		var result *wishCardEntity.Entity
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			newID, _ := repo.Insert(ctx, masterTx, wishCard, 1)

			result, err = repo.SelectByID(ctx, masterTx, newID)
			return err
		})
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("failure_存在しないデータ", func(t *testing.T) {
		var err error
		ctx := context.Background()

		var result *wishCardEntity.Entity
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			result, err = repo.SelectByID(ctx, masterTx, -1)
			return err
		})
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestSelectByIDs(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var err error
		ctx := context.Background()

		ids := []string{"1", "2", "3"}

		var result wishCardEntity.EntitySlice
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
	t.Run("success", func(t *testing.T) {
		var err error
		ctx := context.Background()
		wishCard := &wishCardEntity.Entity{
			Author: &userEntity.Entity{
				ID: 1,
			},
			Activity:    dummyActivity,
			Description: dummyDescription,
			Date:        &dummyDate,
			CreatedAt:   &dummyDate,
			UpdatedAt:   &dummyDate,
			Place: &placeEntity.Entity{
				ID: 1,
			},
		}
		var result wishCardEntity.EntitySlice
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			repo.Insert(ctx, masterTx, wishCard, 1)

			result, err = repo.SelectByCategoryID(ctx, masterTx, 1)
			return err
		})
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("success_存在しないカテゴリ", func(t *testing.T) {
		var err error
		ctx := context.Background()

		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {

			_, err = repo.SelectByCategoryID(ctx, masterTx, -1)
			return err
		})
		assert.NoError(t, err)
	})
}
