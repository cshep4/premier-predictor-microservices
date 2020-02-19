package auth

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
	"testing"
)

func TestTokenCtx(t *testing.T) {
	t.Run("returns context with token added", func(t *testing.T) {
		const token = "token"

		res := tokenCtx(context.Background(), token)

		authToken, ok := res.Value(tokenKey).(string)

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

func TestMetadataFromContext(t *testing.T) {
	t.Run("returns error if token not in context", func(t *testing.T) {
		res, err := MetadataFromContext(context.Background())
		require.Error(t, err)

		assert.Empty(t, res)
		assert.Equal(t, "missing token", err.Error())
	})

	t.Run("returns outgoing metadata with token", func(t *testing.T) {
		const token = "token"
		ctx := context.WithValue(ctx, tokenKey, token)

		res, err := MetadataFromContext(context.WithValue(ctx, tokenKey, token))
		require.NoError(t, err)

		md, ok := metadata.FromOutgoingContext(res)
		assert.True(t, ok)
		assert.Equal(t, token, md["token"][0])
	})
}
