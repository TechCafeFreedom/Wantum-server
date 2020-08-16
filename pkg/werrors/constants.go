package werrors

import (
	"net/http"
)

var (
	AuthFail = &WantumError{
		ErrorCode:      http.StatusUnauthorized,
		ErrorMessageJP: "認証に失敗しました",
		ErrorMessageEN: "authorization failed",
	}
	ServerError = &WantumError{
		ErrorCode:      http.StatusInternalServerError,
		ErrorMessageJP: "サーバでエラーが発生しました",
		ErrorMessageEN: "error occurred in server",
	}
	UserNotFound = &WantumError{
		ErrorCode:      http.StatusNotFound,
		ErrorMessageJP: "リクエストされたユーザーはすでに削除されている可能性があります",
		ErrorMessageEN: "requested user is not found",
	}
)
