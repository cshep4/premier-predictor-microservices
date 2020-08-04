package handler_test

import (
	"context"
	"testing"

	"github.com/cshep4/premier-predictor-microservices/src/gatewayservice/internal/handler"
	"github.com/stretchr/testify/assert"
)

func TestListen(t *testing.T) {
	t.Run("should shutdown gracefully when cancelled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		// we can cancel now and the child errgroup should shutdown the server instantly.
		cancel()

		assert.NoError(t, handler.Listen(ctx, ":0", nil))
	})

	t.Run("should return errors returned by the server", func(t *testing.T) {
		assert.Error(t, handler.Listen(context.Background(), "invalidscheme://not:valid", nil))
	})
}
