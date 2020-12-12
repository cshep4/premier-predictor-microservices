package grpc

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	auth "github.com/cshep4/premier-predictor-microservices/src/common/auth/internal/context"
	grpccfg "github.com/cshep4/premier-predictor-microservices/src/common/grpc"
)

const authUserAgent = "kauthservice/1.0.0"

type (
	authenticator struct {
		authClient  model.AuthServiceClient
		serviceName string
	}

	// InvalidParameterError is returned when a required parameter passed to New is invalid.
	InvalidParameterError struct {
		Parameter string
	}
)

func (i InvalidParameterError) Error() string {
	return fmt.Sprintf("invalid parameter %s", i.Parameter)
}

func New(authClient model.AuthServiceClient, serviceName string) (*authenticator, error) {
	switch {
	case authClient == nil:
		return nil, InvalidParameterError{Parameter: "authClient"}
	case serviceName == "":
		return nil, InvalidParameterError{Parameter: "serviceName"}
	}

	return &authenticator{
		authClient:  authClient,
		serviceName: serviceName,
	}, nil
}

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

func (a *authenticator) doAuth(ctx context.Context, token string) (context.Context, error) {
	if getUserAgentFromGrpcMetadata(ctx) == authUserAgent {
		return auth.SetTokenCtx(ctx, token), nil
	}

	role := model.Role_ROLE_USER
	audience, ok := getAudienceFromGrpcMetadata(ctx)
	if !ok {
		role = model.Role_ROLE_SERVICE
		audience = a.serviceName
	}

	_, err := a.authClient.Validate(ctx, &model.ValidateRequest{
		Token:    token,
		Audience: audience,
		Role:     role,
	})
	if err != nil {
		return nil, err
	}

	return auth.SetTokenCtx(ctx, token), nil
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

func getUserAgentFromGrpcMetadata(ctx context.Context) string {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	if len(meta["user-agent"]) != 1 {
		return ""
	}

	return meta["user-agent"][0]
}

func getAudienceFromGrpcMetadata(ctx context.Context) (string, bool) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}

	if len(meta["audience"]) != 1 {
		return "", false
	}

	return meta["audience"][0], true
}
