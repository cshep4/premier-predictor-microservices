package token

import (
	"context"
	"fmt"

	pb "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
)

type (
	generator struct {
		authClient pb.AuthServiceClient
	}

	// InvalidParameterError is returned when a required parameter passed to New is invalid.
	InvalidParameterError struct {
		Parameter string
	}
)

func (i InvalidParameterError) Error() string {
	return fmt.Sprintf("invalid parameter %s", i.Parameter)
}

func New(authClient pb.AuthServiceClient) (*generator, error) {
	if authClient == nil {
		return nil, InvalidParameterError{Parameter: "authClient"}
	}

	return &generator{
		authClient: authClient,
	}, nil
}

func (g *generator) Generate(ctx context.Context, service string) (string, error) {
	res, err := g.authClient.IssueServiceToken(ctx, &pb.IssueServiceTokenRequest{
		Audience: service,
	})
	if err != nil {
		return "", fmt.Errorf("issue_service_token: %w", err)
	}

	return res.GetToken(), nil
}
