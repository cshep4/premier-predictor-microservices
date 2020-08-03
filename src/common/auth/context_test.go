package auth

import (
	"context"
	"testing"

	auth "github.com/cshep4/premier-predictor-microservices/src/common/auth/internal/context"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

func TestMetadataFromContext(t *testing.T) {
	t.Run("returns error if token not in context", func(t *testing.T) {
		res, err := MetadataFromContext(context.Background())
		require.Error(t, err)

		assert.Empty(t, res)
		assert.Equal(t, "missing token", err.Error())
	})

	t.Run("returns outgoing metadata with token", func(t *testing.T) {
		const token = "token"
		ctx := auth.SetTokenCtx(context.Background(), token)

		res, err := MetadataFromContext(ctx)
		require.NoError(t, err)

		md, ok := metadata.FromOutgoingContext(res)
		assert.True(t, ok)
		assert.Equal(t, token, md["token"][0])
	})
}
