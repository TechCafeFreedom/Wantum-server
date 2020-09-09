package wish_card_tag

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"
	"time"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/wish_card_tag"
	tx "wantum/pkg/infrastructure/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

var db *sql.DB
var txManager repository.MasterTxManager
var repo wish_card_tag.Repository
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

		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			err = repo.Insert(ctx, masterTx, 1, 1)
			return err
		})
		assert.NoError(t, err)
	})
}

func TestBulkInsert(t *testing.T) {
	t.Run("success to bulk insert data", func(t *testing.T) {
		var err error
		ctx := context.Background()

		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			err = repo.BulkInsert(ctx, masterTx, 1, []int{2, 3, 4})
			return err
		})
		assert.NoError(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("success to delete data", func(t *testing.T) {
		var err error
		ctx := context.Background()

		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			err = repo.Delete(ctx, masterTx, 1, 1)
			return err
		})
		assert.NoError(t, err)
	})
}

func TestDeleteByWishCardID(t *testing.T) {
	t.Run("success to delete data", func(t *testing.T) {
		var err error
		ctx := context.Background()

		err = txManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
			err = repo.DeleteByWishCardID(ctx, masterTx, 1)
			return err
		})
		assert.NoError(t, err)
	})
}
