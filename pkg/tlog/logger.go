package tlog

import (
	"context"
	"fmt"
	"runtime"
	"wantum/pkg/constants"

	"go.uber.org/zap"
)

// アプリケーションログ用のロガー
var appLogger *zap.Logger

const AppLoggerKey = "AppLogger"

func init() {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	config.DisableStacktrace = false // スタックトレースONにしたい場合はfalseにする
	logger, _ := config.Build()
	appLogger = logger.Named(AppLoggerKey)
}

// getter
func GetAppLogger() *zap.Logger {
	return appLogger
}

func PrintLogWithCtx(ctx context.Context, err error) {
	// どこで起きたエラーかを特定するための情報を取得
	pt, file, line, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pt).Name()

	// エラーログ出力
	uid, ok := ctx.Value(constants.AuthCtxKey).(string)
	if !ok {
		GetAppLogger().Error(fmt.Sprintf("<[Unknown]Error:%+v, File: %s:%d, Function: %s>", err, file, line, funcName))
	} else {
		GetAppLogger().Error(fmt.Sprintf("<[%s]Error:%+v, File: %s:%d, Function: %s>", uid, err, file, line, funcName))
	}
}

func PrintLogWithUID(uid string, err error) {
	// どこで起きたエラーかを特定するための情報を取得
	pt, file, line, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pt).Name()

	// エラーログ出力
	if uid == "" {
		GetAppLogger().Error(fmt.Sprintf("<[Unknown]Error:%+v, File: %s:%d, Function: %s>", err, file, line, funcName))
	} else {
		GetAppLogger().Error(fmt.Sprintf("<[%s]Error:%+v, File: %s:%d, Function: %s>", uid, err, file, line, funcName))
	}
}
