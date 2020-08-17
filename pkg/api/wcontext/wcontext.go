package wcontext

import (
	"context"
	"errors"
	"wantum/pkg/constants"
)

func GetAuthIDFromContext(ctx context.Context) (string, error) {
	authID, ok := ctx.Value(constants.AuthCtxKey).(string)
	if !ok {
		return "", errors.New("コンテキストのAuthIDキャストでエラーが発生しました。")
	}
	return authID, nil
}
