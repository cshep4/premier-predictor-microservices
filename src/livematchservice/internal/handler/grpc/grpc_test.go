package grpc

import (
	"testing"
	"time"

	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/mocks/live"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestLiveMatchServiceServer_GetMatchSummary(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := live_mocks.NewMockServicer(ctrl)

	interval := 500 * time.Millisecond

	_, err := New(service, interval)
	require.NoError(t, err)

	t.Run("Gets match summary and updates match facts at given timer interval", func(t *testing.T) {

	})

	t.Run("Returns error if there is a problem", func(t *testing.T) {

	})

	t.Run("Return without error if there is a problem after sending initial response", func(t *testing.T) {

	})
}

func TestLiveMatchServiceServer_GetUpcomingMatches(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := live_mocks.NewMockServicer(ctrl)

	interval := 500 * time.Millisecond

	_, err := New(service, interval)
	require.NoError(t, err)

	t.Run("Gets upcoming matches and updates at given timer interval", func(t *testing.T) {

	})

	t.Run("Returns error if there is a problem", func(t *testing.T) {

	})

	t.Run("Return without error if there is a problem after sending initial response", func(t *testing.T) {

	})
}
