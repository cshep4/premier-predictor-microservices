package auth_test

import (
	"context"
	auth "github.com/cshep4/premier-predictor-microservices/src/common/auth/internal/context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetTokenCtx(t *testing.T) {
	t.Run("returns context with token added", func(t *testing.T) {
		const token = "token"

		res := auth.SetTokenCtx(context.Background(), token)

		authToken, ok := auth.GetTokenFromContext(res)

		assert.True(t, ok)
		assert.Equal(t, token, authToken)
	})
}