package main

import (
	"fmt"
	"wantum/db/mysql"
	grpcMiddleware "wantum/pkg/api/middleware/grpc"
	restMiddleware "wantum/pkg/api/middleware/rest"
	tx "wantum/pkg/infrastructure/mysql"
	"wantum/pkg/tlog"

	"github.com/joho/godotenv"
)

func main() {
	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		tlog.GetAppLogger().Error(fmt.Sprintf("failed to load .env file: %v", err))
	}

	// DBインスタンスの作成
	dbInstance := mysql.CreateSQLInstance()
	defer dbInstance.Close()

	// トランザクションマネージャーの作成
	masterTxManager := tx.NewDBMasterTxManager(dbInstance)

	// firebase middlewareの作成
	restFirebaseClient := restMiddleware.CreateFirebaseInstance()
	grpcFirebaseClient := grpcMiddleware.CreateFirebaseInstance()

	// RESTサーバの立ち上げ
	go func() {
		StartRestAPIServer(restFirebaseClient, masterTxManager)
	}()

	// gRPCサーバーの立ち上げ
	StartGrpcServer(grpcFirebaseClient, masterTxManager)
}
