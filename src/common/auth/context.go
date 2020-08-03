package auth

import (
	"context"
	"errors"

	auth "github.com/cshep4/premier-predictor-microservices/src/common/auth/internal/context"

	"google.golang.org/grpc/metadata"
)

func MetadataFromContext(ctx context.Context) (context.Context, error) {
	token, ok := auth.GetTokenFromContext(ctx)
	if !ok {
		return nil, errors.New("missing token")
	}

	return metadata.AppendToOutgoingContext(context.Background(), "token", token), nil
}
