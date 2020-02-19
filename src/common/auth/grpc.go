package auth

import (
	"context"

	grpccfg "github.com/cshep4/premier-predictor-microservices/src/common/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *authenticator) GrpcUnary(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	token, err := getTokenFromGrpcMetadata(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	ctx, err = a.doAuth(ctx, token)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return handler(ctx, req)
}

func (a *authenticator) GrpcStream(srv interface{}, stream grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	token, err := getTokenFromGrpcMetadata(stream.Context())
	if err != nil {
		return status.Error(codes.Unauthenticated, err.Error())
	}

	ctx, err := a.doAuth(stream.Context(), token)
	if err != nil {
		return status.Error(codes.Unauthenticated, err.Error())
	}

	return handler(srv, &grpccfg.ContextServerStream{
		Ctx:          ctx,
		ServerStream: stream,
	})
}
