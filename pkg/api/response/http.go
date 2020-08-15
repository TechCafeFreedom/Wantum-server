package response

import (
	"encoding/json"
	"net/http"
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
			tlog.PrintLogWithCtx(r.Context(), err)
			Error(w, r, werrors.Newf(nil, http.StatusInternalServerError, "サーバで予期せぬエラーが発生しました。", "Unexpected error occurred."))
		}
		w.Write(data)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	data, err := json.Marshal(err)
	if err != nil {
		tlog.PrintLogWithCtx(r.Context(), err)
		Error(w, r, werrors.Newf(nil, http.StatusInternalServerError, "サーバで予期せぬエラーが発生しました。", "Unexpected error occurred."))
	}
	w.Write(data)
}

func JSON(w http.ResponseWriter, r *http.Request, result interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(result)
	if err != nil {
		tlog.PrintLogWithCtx(r.Context(), err)
		Error(w, r, werrors.Newf(nil, http.StatusInternalServerError, "サーバで予期せぬエラーが発生しました。", "Unexpected error occurred."))
	}
	w.Write(data)
}
