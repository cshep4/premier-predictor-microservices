package auth

import (
	"context"
)

const tokenKey authKey = "token"

type authKey string

func SetTokenCtx(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenKey, token)
}

func GetTokenFromContext(ctx context.Context) (string, bool) {
	token, ok := ctx.Value(tokenKey).(string)
	return token, ok
}