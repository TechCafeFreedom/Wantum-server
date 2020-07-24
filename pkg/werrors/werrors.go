package werrors

import (
	"fmt"
	"net/http"

	"golang.org/x/xerrors"
)

// WantumError サーバ-クライアント間エラーハンドリング用エラー
type WantumError struct {
	// エラーコード
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
func New(cause error, errorCode int) error {
	return newError(cause, errorCode, "", "")
}

// Newf WantumErrorをエラーメッセージ付きで生成する
func Newf(cause error, errorCode int, messageJP, messageEN string) error {
	return newError(cause, errorCode, messageJP, messageEN)
}

// Wrap エラーをWantumエラーでラップする
func Wrap(cause error, errorCode int) error {
	return newError(cause, errorCode, "", "")
}

// Wrapf エラーをWantumエラーで、エラーメッセージ付きでラップする
func Wrapf(cause error, errorCode int, messageJP, messageEN string) error {
	return newError(cause, errorCode, messageJP, messageEN)
}

func newError(cause error, errorCode int, errorMessageJP, errorMessageEN string) error {
	return &WantumError{
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
	var errorCode int
	var errorMessageJP, errorMessageEN string
	var wantumError *WantumError
	if ok := xerrors.As(err, &wantumError); ok {
		errorCode = wantumError.ErrorCode
		errorMessageJP = wantumError.ErrorMessageJP
		errorMessageEN = wantumError.ErrorMessageEN
	} else {
		return &WantumError{
			ErrorCode:      http.StatusInternalServerError,
			ErrorMessageJP: "エラーのコンバート時にエラーが発生しました",
			ErrorMessageEN: "Error occured at covert to original error",
			err:            err,
			frame:          xerrors.Caller(1),
		}
	}
	return &WantumError{
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
