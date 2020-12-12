package token_test

import (
	"context"
	"errors"

	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/mocks/auth"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/token"

	pb "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type testError string

func (t testError) Error() string {
	return string(t)
}

func TestNew(t *testing.T) {
	t.Run("returns error if auth client is empty", func(t *testing.T) {
		g, err := token.New(nil)
		require.Error(t, err)
		assert.Empty(t, g)

		ipe, ok := err.(token.InvalidParameterError)
		require.True(t, ok)

		assert.Equal(t, "authClient", ipe.Parameter)
	})

	t.Run("returns generator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		authClient := auth_mock.NewMockAuthServiceClient(ctrl)

		g, err := token.New(authClient)
		require.NoError(t, err)
		assert.NotEmpty(t, g)
	})
}

func TestGenerator_Generate(t *testing.T) {
	t.Run("returns error if error from auth service", func(t *testing.T) {
		const (
			service           = "🏢"
			testErr testError = "error"
		)

		ctrl, ctx := gomock.WithContext(context.Background(), t)
		defer ctrl.Finish()

		authClient := auth_mock.NewMockAuthServiceClient(ctrl)

		generator, err := token.New(authClient)
		require.NoError(t, err)

		authClient.EXPECT().IssueServiceToken(ctx, &pb.IssueServiceTokenRequest{
			Audience: service,
		}).Return(nil, testErr)

		token, err := generator.Generate(ctx, service)
		require.Error(t, err)

		assert.Empty(t, token)
		assert.True(t, errors.Is(err, testErr))
	})

	t.Run("returns error if error from auth service", func(t *testing.T) {
		const (
			service   = "🏢"
			someToken = "🔑"
		)

		ctrl, ctx := gomock.WithContext(context.Background(), t)
		defer ctrl.Finish()

		authClient := auth_mock.NewMockAuthServiceClient(ctrl)

		generator, err := token.New(authClient)
		require.NoError(t, err)

		authClient.EXPECT().IssueServiceToken(ctx, &pb.IssueServiceTokenRequest{
			Audience: service,
		}).Return(&pb.IssueServiceTokenResponse{
			Token: someToken,
		}, nil)

		token, err := generator.Generate(ctx, service)
		require.NoError(t, err)

		assert.Equal(t, someToken, token)
	})
}
