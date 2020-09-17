package tag

import (
	"context"
	"crypto/rand"
	"database/sql"
	"errors"
	"log"
	"os"
	"testing"
	"time"
	tagEntity "wantum/pkg/domain/entity/tag"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/tag"
	tx "wantum/pkg/infrastructure/mysql"
	"wantum/pkg/testutil"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

var (
	db        *sql.DB
	txManager repository.MasterTxManager
	repo      tag.Repository
	dummyDate time.Time
)

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
		name, _ := makeRandomStr(10)
		tag := &tagEntity.Entity{
			Name:      name,
			CreatedAt: &dummyDate,
			UpdatedAt: &dummyDate,
		}

		var result int
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			result, err = repo.Insert(ctx, masterTx, tag)
			return err
		})

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("failure_データがnil", func(t *testing.T) {
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

func TestUpDeleteFlag(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var err error
		ctx := context.Background()
		name, _ := makeRandomStr(10)
		tag := &tagEntity.Entity{
			Name:      name,
			CreatedAt: &dummyDate,
			UpdatedAt: &dummyDate,
		}

		var result *tagEntity.Entity
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			newTagID, _ := repo.Insert(ctx, masterTx, tag)

			tag.ID = newTagID
			tag.DeletedAt = &dummyDate
			err = repo.UpDeleteFlag(ctx, masterTx, tag)
			if err != nil {
				return err
			}

			result, _ = repo.SelectByID(ctx, masterTx, tag.ID)
			return nil
		})

		assert.NoError(t, err)
		assert.NotNil(t, result.DeletedAt)
	})

	t.Run("failure_deletedAtがnil", func(t *testing.T) {
		var err error
		ctx := context.Background()
		name, _ := makeRandomStr(10)
		tag := &tagEntity.Entity{
			ID:        1,
			Name:      name,
			CreatedAt: &dummyDate,
			UpdatedAt: &dummyDate,
		}

		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			err = repo.UpDeleteFlag(ctx, masterTx, tag)
			return err
		})

		assert.Error(t, err)
	})
}

func TestDownDeleteFlag(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var err error
		ctx := context.Background()
		name, _ := makeRandomStr(10)
		tag := &tagEntity.Entity{
			ID:        1,
			Name:      name,
			CreatedAt: &dummyDate,
			UpdatedAt: &dummyDate,
			DeletedAt: &dummyDate,
		}

		var result *tagEntity.Entity
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {

			tag.DeletedAt = nil
			err = repo.DownDeleteFlag(ctx, masterTx, tag)
			if err != nil {
				return err
			}

			result, _ = repo.SelectByID(ctx, masterTx, tag.ID)
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
		name, _ := makeRandomStr(10)
		tag := &tagEntity.Entity{
			Name:      name,
			CreatedAt: &dummyDate,
			UpdatedAt: &dummyDate,
		}

		var result *tagEntity.Entity
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			newTagID, _ := repo.Insert(ctx, masterTx, tag)
			tag.ID = newTagID
			tag.DeletedAt = &dummyDate
			repo.UpDeleteFlag(ctx, masterTx, tag)

			err = repo.Delete(ctx, masterTx, tag.ID)
			if err != nil {
				return err
			}

			result, err = repo.SelectByID(ctx, masterTx, tag.ID)
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
		name, _ := makeRandomStr(10)
		tag := &tagEntity.Entity{
			Name:      name,
			CreatedAt: &dummyDate,
			UpdatedAt: &dummyDate,
		}

		var result *tagEntity.Entity
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			newTagID, _ := repo.Insert(ctx, masterTx, tag)

			result, err = repo.SelectByID(ctx, masterTx, newTagID)
			return err
		})

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestSelectByIDs(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var err error
		ctx := context.Background()

		var result tagEntity.EntitySlice
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {

			result, err = repo.SelectByIDs(ctx, masterTx, []int{1, 2, 3})
			return err
		})

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 3, len(result))
	})
}

func TestSelectByName(t *testing.T) {
	t.Run("successe", func(t *testing.T) {
		var err error
		ctx := context.Background()
		name, _ := makeRandomStr(10)
		tag := &tagEntity.Entity{
			Name:      name,
			CreatedAt: &dummyDate,
			UpdatedAt: &dummyDate,
		}

		var result *tagEntity.Entity
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			repo.Insert(ctx, masterTx, tag)

			result, err = repo.SelectByName(ctx, masterTx, name)
			return err
		})

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("failure_存在しないデータ", func(t *testing.T) {
		var err error
		ctx := context.Background()
		name, _ := makeRandomStr(10)

		var result *tagEntity.Entity
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {

			result, err = repo.SelectByName(ctx, masterTx, name)
			return err
		})

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestSelectByWishCardID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var err error
		ctx := context.Background()

		var result tagEntity.EntitySlice
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {

			result, err = repo.SelectByWishCardID(ctx, masterTx, 4)
			return err
		})

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestSelectByMemoryID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var err error
		ctx := context.Background()

		var result tagEntity.EntitySlice
		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {

			result, err = repo.SelectByMemoryID(ctx, masterTx, 4)
			return err
		})

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func makeRandomStr(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error...")
	}

	var result string
	for _, v := range b {
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}
