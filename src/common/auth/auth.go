package auth

import (
	"context"
	"net/http"

	grpcAuth "github.com/cshep4/premier-predictor-microservices/src/common/auth/grpc"
	httpAuth "github.com/cshep4/premier-predictor-microservices/src/common/auth/http"

	"github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"google.golang.org/grpc"
)

type (
	httpAuthenticator interface {
		Http(next http.Handler) http.Handler
	}

	grpcAuthenticator interface {
		GrpcUnary(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
		GrpcStream(srv interface{}, stream grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error
	}

	authenticator struct {
		httpAuthenticator
		grpcAuthenticator
	}
)

func New(authClient model.AuthServiceClient, serviceName string, authorizer httpAuth.Authorizer) (*authenticator, error) {
	httpAuthenticator, err := httpAuth.New(authClient, authorizer)
	if err != nil {
		return nil, err
	}

	grpcAuthenticator, err := grpcAuth.New(authClient, serviceName)
	if err != nil {
		return nil, err
	}

	return &authenticator{
		httpAuthenticator: httpAuthenticator,
		grpcAuthenticator: grpcAuthenticator,
	}, nil
}
