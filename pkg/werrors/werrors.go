package werrors

import (
	"fmt"

	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
)

/**
gRPCのエラーコードとHTTPのエラーコードの対比表は、
https://www.notion.so/gRPC-5621cb76a3684491a7a6df37da71cd14#ca4c1e1a4fe5489ab5bb705254571d54 をチェック
*/

// WantumError サーバ-クライアント間エラーハンドリング用エラー
type WantumError struct {
	// gRPC用エラーコード
	GrpcErrorCode codes.Code
	// REST用エラーコード
	ErrorCode int
	// システムエラーメッセージ(日本語)
	ErrorMessageJP string
	// システムエラーメッセージ(英語)
	ErrorMessageEN string
	// xerrors拡張用フィールド
	err error
	// それぞれでfmt.Errorf("%w", err)を記述する必要があるためxerrors使う。
	frame xerrors.Frame
}

// New WantumErrorを生成する
func New(cause error, grpcErrCode codes.Code, errorCode int) error {
	return newError(cause, grpcErrCode, errorCode, "", "")
}

// Newf WantumErrorをエラーメッセージ付きで生成する
func Newf(cause error, grpcErrCode codes.Code, errorCode int, messageJP, messageEN string) error {
	return newError(cause, grpcErrCode, errorCode, messageJP, messageEN)
}

// Wrap エラーをWantumエラーでラップする
func Wrap(cause error, grpcErrCode codes.Code, errorCode int) error {
	return newError(cause, grpcErrCode, errorCode, "", "")
}

// Wrapf エラーをWantumエラーで、エラーメッセージ付きでラップする
func Wrapf(cause error, grpcErrCode codes.Code, errorCode int, messageJP, messageEN string) error {
	return newError(cause, grpcErrCode, errorCode, messageJP, messageEN)
}

// FromConstant AuthFailなどのように予め定数として用意してあるWantumエラーから生成する
func FromConstant(cause error, wantumError *WantumError) error {
	return newError(cause, wantumError.GrpcErrorCode, wantumError.ErrorCode, wantumError.ErrorMessageJP, wantumError.ErrorMessageEN)
}

func newError(cause error, grpcErrCode codes.Code, errorCode int, errorMessageJP, errorMessageEN string) error {
	return &WantumError{
		GrpcErrorCode:  grpcErrCode,
		ErrorCode:      errorCode,
		ErrorMessageJP: errorMessageJP,
		ErrorMessageEN: errorMessageEN,
		err:            cause,
		frame:          xerrors.Caller(2),
	}
}

// Stack エラーをStackする
// スタックフレームを明示的に積んでいく必要があるためエラー出力に記録したいエラーハンドリング箇所ではStackを行う
func Stack(err error) error {
	var grpcErrCode codes.Code
	var errorCode int
	var errorMessageJP, errorMessageEN string
	var wantumError *WantumError
	if ok := xerrors.As(err, &wantumError); ok {
		grpcErrCode = wantumError.GrpcErrorCode
		errorCode = wantumError.ErrorCode
		errorMessageJP = wantumError.ErrorMessageJP
		errorMessageEN = wantumError.ErrorMessageEN
	} else {
		return ServerError
	}
	return &WantumError{
		GrpcErrorCode:  grpcErrCode,
		ErrorCode:      errorCode,
		ErrorMessageJP: errorMessageJP,
		ErrorMessageEN: errorMessageEN,
		err:            err,
		frame:          xerrors.Caller(1),
	}
}

// Error エラーメッセージを取得する
func (e *WantumError) Error() string {
	return fmt.Sprintf("messageJP=%s, messageEN=%s", e.ErrorMessageJP, e.ErrorMessageEN)
}

func (e *WantumError) Unwrap() error {
	return e.err
}

func (e *WantumError) Format(s fmt.State, v rune) {
	xerrors.FormatError(e, s, v)
}

func (e *WantumError) FormatError(p xerrors.Printer) error {
	p.Print(e.ErrorMessageJP, e.ErrorMessageEN)
	e.frame.Format(p)
	return e.err
}
