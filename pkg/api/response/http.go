package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"wantum/pkg/tlog"
	"wantum/pkg/werrors"

	"golang.org/x/xerrors"
)

type key string

const (
	AuthCtxKey key = "AUTHED_UID"
)

func Error(w http.ResponseWriter, r *http.Request, err error) {
	// エラーレスポンスの送信
	var wantumError *werrors.WantumError
	if ok := xerrors.As(err, &wantumError); ok {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(wantumError.ErrorCode)
		data, err := json.Marshal(err)
		if err != nil {
			// どこで起きたエラーかを特定するための情報を取得
			pt, file, line, _ := runtime.Caller(0)
			funcName := runtime.FuncForPC(pt).Name()

			// エラーログ出力
			uid, ok := r.Context().Value(AuthCtxKey).(string)
			if !ok {
				tlog.GetAppLogger().Error(fmt.Sprintf("<[Unknown]Error:%+v, File: %s:%d, Function: %s>", err, file, line, funcName))
			} else {
				tlog.GetAppLogger().Error(fmt.Sprintf("<[%s]Error:%+v, File: %s:%d, Function: %s>", uid, err, file, line, funcName))
			}
			Error(w, r, werrors.Newf(nil, http.StatusInternalServerError, "サーバで予期せぬエラーが発生しました。", "Unexpected error occurred."))
		}
		_, _ = w.Write(data)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	data, err := json.Marshal(err)
	if err != nil {
		// どこで起きたエラーかを特定するための情報を取得
		pt, file, line, _ := runtime.Caller(0)
		funcName := runtime.FuncForPC(pt).Name()

		// エラーログ出力
		uid, ok := r.Context().Value(AuthCtxKey).(string)
		if !ok {
			tlog.GetAppLogger().Error(fmt.Sprintf("<[Unknown]Error:%+v, File: %s:%d, Function: %s>", err, file, line, funcName))
		} else {
			tlog.GetAppLogger().Error(fmt.Sprintf("<[%s]Error:%+v, File: %s:%d, Function: %s>", uid, err, file, line, funcName))
		}
		Error(w, r, werrors.Newf(nil, http.StatusInternalServerError, "サーバで予期せぬエラーが発生しました。", "Unexpected error occurred."))
	}
	_, _ = w.Write(data)
}

func JSON(w http.ResponseWriter, r *http.Request, result interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(result)
	if err != nil {
		// どこで起きたエラーかを特定するための情報を取得
		pt, file, line, _ := runtime.Caller(0)
		funcName := runtime.FuncForPC(pt).Name()

		// エラーログ出力
		uid, ok := r.Context().Value(AuthCtxKey).(string)
		if !ok {
			tlog.GetAppLogger().Error(fmt.Sprintf("<[Unknown]Error:%+v, File: %s:%d, Function: %s>", err, file, line, funcName))
		} else {
			tlog.GetAppLogger().Error(fmt.Sprintf("<[%s]Error:%+v, File: %s:%d, Function: %s>", uid, err, file, line, funcName))
		}
		Error(w, r, werrors.Newf(nil, http.StatusInternalServerError, "サーバで予期せぬエラーが発生しました。", "Unexpected error occurred."))
	}
	_, _ = w.Write(data)
}
