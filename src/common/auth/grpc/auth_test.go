package grpc

import (
	"context"
	"errors"
	"testing"

	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	auth "github.com/cshep4/premier-predictor-microservices/src/common/auth/internal/context"
	grpccfg "github.com/cshep4/premier-predictor-microservices/src/common/grpc"
	"github.com/cshep4/premier-predictor-microservices/src/common/internal/mocks/auth"
)

const (
	token       = "token"
	response    = "test response"
	serviceName = "service name"
)

var ctx = metadata.NewIncomingContext(context.Background(), metadata.MD{"token": []string{token}})

func TestAuthenticator_GrpcUnary(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authClient := auth_mocks.NewMockAuthServiceClient(ctrl)

	authenticator, err := New(authClient, serviceName)
	require.NoError(t, err)

	t.Run("returns error if token not in metadata", func(t *testing.T) {
		_, err := authenticator.GrpcUnary(context.Background(), nil, nil, nil)
		require.Error(t, err)

		statusErr := status.Convert(err)

		assert.Equal(t, codes.Unauthenticated, statusErr.Code())
		assert.Equal(t, "missing context metadata", statusErr.Message())
	})

	t.Run("returns error if error from auth service", func(t *testing.T) {
		testErr := errors.New("error")

		req := &gen.ValidateRequest{
			Token:    token,
			Audience: serviceName,
			Role:     gen.Role_ROLE_SERVICE,
		}
		authClient.EXPECT().Validate(ctx, req).Return(nil, testErr)

		_, err := authenticator.GrpcUnary(ctx, nil, nil, nil)
		require.Error(t, err)

		statusErr := status.Convert(err)

		assert.Equal(t, codes.Unauthenticated, statusErr.Code())
		assert.Equal(t, testErr.Error(), statusErr.Message())
	})

	t.Run("authenticates and calls handler", func(t *testing.T) {
		req := &gen.ValidateRequest{
			Token:    token,
			Audience: serviceName,
			Role:     gen.Role_ROLE_SERVICE,
		}
		authClient.EXPECT().Validate(ctx, req).Return(nil, nil)

		called := false
		handler := func(ctx context.Context, req interface{}) (i interface{}, err error) {
			authToken, ok := auth.GetTokenFromContext(ctx)

			assert.True(t, ok)
			assert.Equal(t, token, authToken)

			called = true
			return response, nil
		}
		res, err := authenticator.GrpcUnary(ctx, nil, nil, handler)
		require.NoError(t, err)

		assert.Equal(t, response, res)
		assert.True(t, called)
	})
}

func TestAuthenticator_GrpcStreamInterceptor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authClient := auth_mocks.NewMockAuthServiceClient(ctrl)

	authenticator, err := New(authClient, serviceName)
	require.NoError(t, err)

	t.Run("returns error if token not in metadata", func(t *testing.T) {
		err := authenticator.GrpcStream(nil, &grpccfg.ContextServerStream{Ctx: context.Background()}, nil, nil)
		require.Error(t, err)

		statusErr := status.Convert(err)

		assert.Equal(t, codes.Unauthenticated, statusErr.Code())
		assert.Equal(t, "missing context metadata", statusErr.Message())
	})

	t.Run("returns error if error from auth service", func(t *testing.T) {
		testErr := errors.New("error")

		req := &gen.ValidateRequest{
			Token:    token,
			Audience: serviceName,
			Role:     gen.Role_ROLE_SERVICE,
		}
		authClient.EXPECT().Validate(ctx, req).Return(nil, testErr)

		err := authenticator.GrpcStream(nil, &grpccfg.ContextServerStream{Ctx: ctx}, nil, nil)
		require.Error(t, err)

		statusErr := status.Convert(err)

		assert.Equal(t, codes.Unauthenticated, statusErr.Code())
		assert.Equal(t, testErr.Error(), statusErr.Message())
	})

	t.Run("authenticates and calls handler", func(t *testing.T) {
		req := &gen.ValidateRequest{
			Token:    token,
			Audience: serviceName,
			Role:     gen.Role_ROLE_SERVICE,
		}
		authClient.EXPECT().Validate(ctx, req).Return(nil, nil)

		called := false
		handler := func(srv interface{}, stream grpc.ServerStream) error {
			authToken, ok := auth.GetTokenFromContext(stream.Context())

			assert.True(t, ok)
			assert.Equal(t, token, authToken)

			called = true

			return nil
		}
		err := authenticator.GrpcStream(nil, &grpccfg.ContextServerStream{Ctx: ctx}, nil, handler)
		require.NoError(t, err)

		assert.True(t, called)
	})
}

func TestAuthenticator_doAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authClient := auth_mocks.NewMockAuthServiceClient(ctrl)

	authenticator, err := New(authClient, serviceName)
	require.NoError(t, err)

	const token = "token"
	ctx := context.Background()

	t.Run("returns error if error from auth service", func(t *testing.T) {
		testErr := errors.New("error")

		authClient.EXPECT().Validate(ctx, &gen.ValidateRequest{Token: token}).Return(nil, testErr)

		ctx, err := authenticator.doAuth(ctx, token)
		require.Error(t, err)

		assert.Equal(t, testErr, err)
		assert.Nil(t, ctx)
	})

	t.Run("returns nil and adds token to ctx if no error from auth service", func(t *testing.T) {
		authClient.EXPECT().Validate(ctx, &gen.ValidateRequest{Token: token}).Return(nil, nil)

		ctx, err := authenticator.doAuth(ctx, token)
		require.NoError(t, err)

		authToken, ok := auth.GetTokenFromContext(ctx)

		assert.True(t, ok)
		assert.Equal(t, token, authToken)
	})
}

func TestGetTokenFromGrpcMetadata(t *testing.T) {
	t.Run("returns error if context not correct metadata", func(t *testing.T) {
		res, err := getTokenFromGrpcMetadata(context.Background())
		require.Error(t, err)

		assert.Empty(t, res)
		assert.Equal(t, "missing context metadata", err.Error())
	})

	t.Run("returns error if token not in metadata", func(t *testing.T) {
		res, err := getTokenFromGrpcMetadata(metadata.NewIncomingContext(context.Background(), metadata.MD{}))
		require.Error(t, err)

		assert.Empty(t, res)
		assert.Equal(t, "invalid access token", err.Error())
	})

	t.Run("returns token from incoming metadata", func(t *testing.T) {
		const token = "token"

		res, err := getTokenFromGrpcMetadata(metadata.NewIncomingContext(context.Background(), metadata.MD{"token": []string{token}}))
		require.NoError(t, err)

		assert.Equal(t, token, res)
	})
}
