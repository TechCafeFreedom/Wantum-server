package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"wantum/pkg/api/response"
	"wantum/pkg/constants"
	"wantum/pkg/tlog"
	"wantum/pkg/werrors"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

type FirebaseAuth interface {
	MiddlewareFunc(next http.Handler) http.Handler
}

type firebaseAuth struct {
	client *auth.Client
}

func (fa *firebaseAuth) MiddlewareFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Authorizationヘッダーからjwtトークンを取得
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			err := errors.New("ユーザのAuthorizationが空だっためエラーとしました。")
			tlog.PrintErrorLogWithCtx(r.Context(), err)
			response.Error(w, r, werrors.FromConstant(err, werrors.AuthFail))
			return
		}
		jwtToken := strings.Replace(authHeader, "Bearer ", "", 1)

		// JWT の検証
		authedUserToken, err := fa.client.VerifyIDToken(r.Context(), jwtToken)
		if err != nil {
			tlog.PrintErrorLogWithCtx(r.Context(), err)
			response.Error(w, r, werrors.FromConstant(err, werrors.AuthFail))
			return
		}
		// contextにuidを格納
		r = r.WithContext(context.WithValue(r.Context(), constants.AuthCtxKey, authedUserToken.UID))
		next.ServeHTTP(w, r)
	})
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
