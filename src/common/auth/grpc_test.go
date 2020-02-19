package auth

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"testing"

	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	grpccfg "github.com/cshep4/premier-predictor-microservices/src/common/grpc"
	"github.com/cshep4/premier-predictor-microservices/src/common/internal/mocks/auth"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	token    = "token"
	response = "test response"
)

var ctx = metadata.NewIncomingContext(context.Background(), metadata.MD{"token": []string{token}})

func TestAuthenticator_GrpcUnary(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authClient := auth_mocks.NewMockAuthServiceClient(ctrl)

	auth, err := New(authClient)
	require.NoError(t, err)

	t.Run("returns error if token not in metadata", func(t *testing.T) {
		_, err := auth.GrpcUnary(context.Background(), nil, nil, nil)
		require.Error(t, err)

		statusErr := status.Convert(err)

		assert.Equal(t, codes.Unauthenticated, statusErr.Code())
		assert.Equal(t, "missing context metadata", statusErr.Message())
	})

	t.Run("returns error if error from auth service", func(t *testing.T) {
		testErr := errors.New("error")

		authClient.EXPECT().Validate(ctx, &gen.ValidateRequest{Token: token}).Return(nil, testErr)

		_, err := auth.GrpcUnary(ctx, nil, nil, nil)
		require.Error(t, err)

		statusErr := status.Convert(err)

		assert.Equal(t, codes.Unauthenticated, statusErr.Code())
		assert.Equal(t, testErr.Error(), statusErr.Message())
	})

	t.Run("authenticates and calls handler", func(t *testing.T) {
		authClient.EXPECT().Validate(ctx, &gen.ValidateRequest{Token: token}).Return(nil, nil)

		called := false
		handler := func(ctx context.Context, req interface{}) (i interface{}, err error) {
			authToken, ok := ctx.Value(tokenKey).(string)

			assert.True(t, ok)
			assert.Equal(t, token, authToken)

			called = true
			return response, nil
		}
		res, err := auth.GrpcUnary(ctx, nil, nil, handler)
		require.NoError(t, err)

		assert.Equal(t, response, res)
		assert.True(t, called)
	})
}

func TestAuthenticator_GrpcStreamInterceptor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authClient := auth_mocks.NewMockAuthServiceClient(ctrl)

	auth, err := New(authClient)
	require.NoError(t, err)

	t.Run("returns error if token not in metadata", func(t *testing.T) {
		err := auth.GrpcStream(nil, &grpccfg.ContextServerStream{Ctx: context.Background()}, nil, nil)
		require.Error(t, err)

		statusErr := status.Convert(err)

		assert.Equal(t, codes.Unauthenticated, statusErr.Code())
		assert.Equal(t, "missing context metadata", statusErr.Message())
	})

	t.Run("returns error if error from auth service", func(t *testing.T) {
		testErr := errors.New("error")

		authClient.EXPECT().Validate(ctx, &gen.ValidateRequest{Token: token}).Return(nil, testErr)

		err := auth.GrpcStream(nil, &grpccfg.ContextServerStream{Ctx: ctx}, nil, nil)
		require.Error(t, err)

		statusErr := status.Convert(err)

		assert.Equal(t, codes.Unauthenticated, statusErr.Code())
		assert.Equal(t, testErr.Error(), statusErr.Message())
	})

	t.Run("authenticates and calls handler", func(t *testing.T) {
		authClient.EXPECT().Validate(ctx, &gen.ValidateRequest{Token: token}).Return(nil, nil)

		called := false
		handler := func(srv interface{}, stream grpc.ServerStream) error {
			authToken, ok := stream.Context().Value(tokenKey).(string)

			assert.True(t, ok)
			assert.Equal(t, token, authToken)

			called = true

			return nil
		}
		err := auth.GrpcStream(nil, &grpccfg.ContextServerStream{Ctx: ctx}, nil, handler)
		require.NoError(t, err)

		assert.True(t, called)
	})
}
