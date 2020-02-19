package auth

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
)

type authKey string
const tokenKey authKey = "token"

func tokenCtx(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenKey, token)
}

func getTokenFromGrpcMetadata(ctx context.Context) (string, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("missing context metadata")
	}

	if len(meta["token"]) != 1 {
		return "", errors.New("invalid access token")
	}

	return meta["token"][0], nil
}

func MetadataFromContext(ctx context.Context) (context.Context, error) {
	token, ok := ctx.Value(tokenKey).(string)
	if !ok {
		return nil, errors.New("missing token")
	}

	return metadata.AppendToOutgoingContext(context.Background(), "token", token), nil
}