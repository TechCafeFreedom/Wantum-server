package tlog

import (
	"context"
	"fmt"
	"runtime"
	"wantum/pkg/constants"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// アプリケーションログ用のロガー
var appLogger *zap.Logger

const AppLoggerKey = "AppLogger"

func init() {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	config.DisableStacktrace = false // スタックトレースONにしたい場合はfalseにする
	config.Encoding = "json"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	logger, _ := config.Build()
	appLogger = logger.Named(AppLoggerKey)
}

// getter
func GetAppLogger() *zap.Logger {
	return appLogger
}

func PrintErrorLogWithCtx(ctx context.Context, err error) {
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

func PrintErrorLogWithAuthID(authID string, err error) {
	// どこで起きたエラーかを特定するための情報を取得
	pt, file, line, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pt).Name()

	// エラーログ出力
	if authID == "" {
		GetAppLogger().Error(fmt.Sprintf("<[Unknown]Error:%+v, File: %s:%d, Function: %s>", err, file, line, funcName))
	} else {
		GetAppLogger().Error(fmt.Sprintf("<[%s]Error:%+v, File: %s:%d, Function: %s>", authID, err, file, line, funcName))
	}
}
