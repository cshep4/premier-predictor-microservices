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

func (a *authenticator) doAuth(token string) error {
	request := &gen.ValidateRequest{Token: token}

	_, err := a.client.Validate(context.Background(), request)

	if err != nil {
		return err
	}

	return nil
}