package middleware

import (
	"context"
	"fmt"
	"wantum/pkg/constants"
	"wantum/pkg/tlog"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/api/option"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FirebaseAuth interface {
	MiddlewareFunc() grpc_auth.AuthFunc
}

type firebaseAuth struct {
	client *auth.Client
}

func (fa *firebaseAuth) MiddlewareFunc() grpc_auth.AuthFunc {
	return fa.middlewareImpl()
}

func (fa *firebaseAuth) middlewareImpl() grpc_auth.AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		// Authorizationヘッダーからjwtトークンを取得
		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			tlog.PrintErrorLogWithCtx(ctx, err)
			st := status.New(codes.Unauthenticated, "authentication error occurred")
			errResJP := &errdetails.LocalizedMessage{
				Locale:  "ja-JP",
				Message: "認証情報が見つかりませんでした。",
			}
			errResEN := &errdetails.LocalizedMessage{
				Locale:  "en-US",
				Message: "Authorization header is empty",
			}
			detailsErr, _ := st.WithDetails(errResJP, errResEN)
			return nil, detailsErr.Err()
		}

		// JWT の検証
		authedUserToken, err := fa.client.VerifyIDToken(ctx, token)
		if err != nil {
			tlog.PrintErrorLogWithCtx(ctx, err)
			st := status.New(codes.Unauthenticated, "authentication error occurred")
			errResJP := &errdetails.LocalizedMessage{
				Locale:  "ja-JP",
				Message: "トークンが有効ではありませんでした。",
			}
			errResEN := &errdetails.LocalizedMessage{
				Locale:  "en-US",
				Message: "Your token is invalid.",
			}
			detailsErr, _ := st.WithDetails(errResJP, errResEN)
			return nil, detailsErr.Err()
		}

		// ユーザデータの取得
		userData, err := fa.client.GetUser(ctx, authedUserToken.UID)
		if err != nil {
			tlog.PrintErrorLogWithCtx(ctx, err)
			st := status.New(codes.Internal, "authentication error occurred")
			errResJP := &errdetails.LocalizedMessage{
				Locale:  "ja-JP",
				Message: "サーバでエラーが発生しました。。",
			}
			errResEN := &errdetails.LocalizedMessage{
				Locale:  "en-US",
				Message: "Error occurred in server.",
			}
			detailsErr, _ := st.WithDetails(errResJP, errResEN)
			return nil, detailsErr.Err()
		}

		// contextにuidを格納
		ctx = context.WithValue(ctx, constants.AuthCtxKey, userData.UID)
		ctx = context.WithValue(ctx, constants.EmailCtxKey, userData.Email)
		return ctx, nil
	}
}

func CreateFirebaseInstance() FirebaseAuth {
	ctx := context.Background()

	// get credential of firebase
	var opt option.ClientOption
	gcpClient, err := secretmanager.NewClient(ctx)
	if err != nil {
		// local
		opt = option.WithCredentialsFile("wantum-firebase-adminsdk-cz9e4-4c4789f0f4.json")
	} else {
		// gcp
		opt = option.WithCredentialsJSON(getFirebaseCredentialJSON(ctx, gcpClient))
	}

	// firebase appの作成
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		tlog.GetAppLogger().Panic(fmt.Sprintf("error initializing firebase app: %v", err))
	}

	// firebase admin clientの作成
	client, err := app.Auth(ctx)
	if err != nil {
		tlog.GetAppLogger().Panic(fmt.Sprintf("error initialize firebase instance. %v", err))
	}

	return &firebaseAuth{
		client: client,
	}
}

// getFirebaseCredentialJSON firebaseの証明書をjsonで取得
func getFirebaseCredentialJSON(ctx context.Context, client *secretmanager.Client) []byte {
	projectID := "wantum-server"
	secretID := "fireauth-key"
	// requestの作成
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", projectID, secretID),
	}

	// get secret value
	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		tlog.GetAppLogger().Panic(fmt.Sprintf("failed to access secret version: %v", err))
	}

	return result.Payload.Data
}
