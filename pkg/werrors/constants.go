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
		ErrorMessageJP: "リクエストされたやりたいことボードは既に削除された可能性があります",
		ErrorMessageEN: "requested wish_board is not found",
	}
	BadRequest = &WantumError{
		GrpcErrorCode:  codes.InvalidArgument,
		ErrorCode:      http.StatusBadRequest,
		ErrorMessageJP: "リクエスト内容をもう一度見直してください",
		ErrorMessageEN: "Please check your request",
	}
)
