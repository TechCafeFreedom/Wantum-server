package werrors

import (
	"fmt"
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
	BadRequest = &WantumError{
		GrpcErrorCode:  codes.InvalidArgument,
		ErrorCode:      http.StatusBadRequest,
		ErrorMessageJP: "リクエスト内容をもう一度見直してください",
		ErrorMessageEN: "Please check your request",
	}
	WishCardNotFound = &WantumError{
		err:            fmt.Errorf("Attempted to update non-existent data"),
		GrpcErrorCode:  codes.NotFound,
		ErrorCode:      http.StatusNotFound,
		ErrorMessageJP: "存在しない「やりたいこと」です。",
		ErrorMessageEN: "the wish card is not exists.",
	}
	PlaceNotFound = &WantumError{
		err:            fmt.Errorf("Attempted to update non-existent data"),
		GrpcErrorCode:  codes.NotFound,
		ErrorCode:      http.StatusNotFound,
		ErrorMessageJP: "存在しない「場所」です。",
		ErrorMessageEN: "the place is not exists.",
	}
)
