package werrors

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

var (
	AuthFail = &WantumError{
		GrpcErrorCode:  codes.Unauthenticated,
		ErrorCode:      http.StatusUnauthorized,
		ErrorMessageJP: "認証に失敗しました",
		ErrorMessageEN: "authorization failed",
	}
	ServerError = &WantumError{
		GrpcErrorCode:  codes.Internal,
		ErrorCode:      http.StatusInternalServerError,
		ErrorMessageJP: "サーバでエラーが発生しました",
		ErrorMessageEN: "error occurred in server",
	}
	UserNotFound = &WantumError{
		GrpcErrorCode:  codes.NotFound,
		ErrorCode:      http.StatusNotFound,
		ErrorMessageJP: "リクエストされたユーザーはすでに削除されている可能性があります",
		ErrorMessageEN: "requested user is not found",
	}
	WishBoardNotFound = &WantumError{
		GrpcErrorCode:  codes.NotFound,
		ErrorCode:      http.StatusNotFound,
		ErrorMessageJP: "ボードは既に削除された可能性があります",
		ErrorMessageEN: "requested wish_board is not found",
	}
	WishBoardPermissionDenied = &WantumError{
		GrpcErrorCode:  codes.PermissionDenied,
		ErrorCode:      http.StatusForbidden,
		ErrorMessageJP: "ボードへのアクセスが許可されていません",
		ErrorMessageEN: "you don't have access to wish_board",
	}
	BadRequest = &WantumError{
		GrpcErrorCode:  codes.InvalidArgument,
		ErrorCode:      http.StatusBadRequest,
		ErrorMessageJP: "リクエスト内容をもう一度見直してください",
		ErrorMessageEN: "Please check your request",
	}
)
