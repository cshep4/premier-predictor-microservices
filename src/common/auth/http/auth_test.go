package http

import (
	"context"
	"errors"
	"testing"

	"github.com/cshep4/premier-predictor-microservices/src/common/auth/internal/context"
	"github.com/cshep4/premier-predictor-microservices/src/common/internal/mocks/auth"

	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthenticator_doAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authClient := auth_mocks.NewMockAuthServiceClient(ctrl)
	authorizer := auth_mocks.NewMockAuthorizer(ctrl)

	authenticator, err := New(authClient, authorizer)
	require.NoError(t, err)

	const token = "token"
	ctx := context.Background()

	t.Run("returns error if error from auth service", func(t *testing.T) {
		testErr := errors.New("error")

		authClient.EXPECT().Validate(ctx, &gen.ValidateRequest{Token: token}).Return(nil, testErr)

		ctx, err := authenticator.doAuth(ctx, token, "", gen.Role_ROLE_INVALID)
		require.Error(t, err)

		assert.Equal(t, testErr, err)
		assert.Nil(t, ctx)
	})

	t.Run("returns nil and adds token to ctx if no error from auth service", func(t *testing.T) {
		authClient.EXPECT().Validate(ctx, &gen.ValidateRequest{Token: token}).Return(nil, nil)

		ctx, err := authenticator.doAuth(ctx, token, "", gen.Role_ROLE_INVALID)
		require.NoError(t, err)

		authToken, ok := auth.GetTokenFromContext(ctx)

		assert.True(t, ok)
		assert.Equal(t, token, authToken)
	})
}
