//go:generate mockgen -destination=../auth/mocks/mock_auth_client.go -package=authmocks github.com/cshep4/premier-predictor-microservices/src/common/interfaces Authenticator

package interfaces

import (
	"context"
	"google.golang.org/grpc"
	"net/http"
)

// Deprecated: interface moved to new runner packages
type Authenticator interface {
	GrpcUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
	GrpcStreamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error
	HttpMiddleware(next http.Handler) http.Handler
}
