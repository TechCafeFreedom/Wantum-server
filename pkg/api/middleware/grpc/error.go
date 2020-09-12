package middleware

import (
	"context"
	"wantum/pkg/werrors"

	"golang.org/x/xerrors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func UnaryErrorHandling() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		resp, err = handler(ctx, req)

		if err != nil {
			// エラーレスポンスの送信
			var wantumError *werrors.WantumError
			if ok := xerrors.As(err, &wantumError); ok {
				st := status.New(wantumError.GrpcErrorCode, "some error occurred")
				errResJP := &errdetails.LocalizedMessage{
					Locale:  "ja-JP",
					Message: wantumError.ErrorMessageJP,
				}
				errResEN := &errdetails.LocalizedMessage{
					Locale:  "en-US",
					Message: wantumError.ErrorMessageEN,
				}
				detailsErr, _ := st.WithDetails(errResJP, errResEN)
				return nil, detailsErr.Err()
			}
			return nil, err

		}

		return resp, nil
	}
}

func StreamErrorHandling() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		err := handler(srv, ss)

		if err != nil {
			// エラーレスポンスの送信
			var wantumError *werrors.WantumError
			if ok := xerrors.As(err, &wantumError); ok {
				st := status.New(wantumError.GrpcErrorCode, "some error occurred")
				errResJP := &errdetails.LocalizedMessage{
					Locale:  "ja-JP",
					Message: wantumError.ErrorMessageJP,
				}
				errResEN := &errdetails.LocalizedMessage{
					Locale:  "en-US",
					Message: wantumError.ErrorMessageEN,
				}
				detailsErr, _ := st.WithDetails(errResJP, errResEN)
				return detailsErr.Err()
			}
			return err
		}

		return nil
	}
}
