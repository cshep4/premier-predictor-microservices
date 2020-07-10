package auth

import (
	"context"
	"errors"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
)

type authenticator struct {
	client gen.AuthServiceClient
}

func New(client gen.AuthServiceClient) (*authenticator, error) {
	if client == nil {
		return nil, errors.New("client_is_nil")
	}

	return &authenticator{
		client: client,
	}, nil
}

func (a *authenticator) doAuth(ctx context.Context, token string) (context.Context, error) {
	_, err := a.client.Validate(ctx, &gen.ValidateRequest{Token: token})
	if err != nil {
		return nil, err
	}

	return tokenCtx(ctx, token), nil
}
