package resolver

import (
	"context"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/gatewayservice/internal/model"
)

type Subscription struct {
	authService      gen.AuthServiceClient
	liveMatchService gen.LiveMatchServiceClient
}

func NewSubscription(opts ...SubscriptionOption) *Subscription {
	s := new(Subscription)
	s.Apply(opts...)
	return s
}

func (s *Subscription) Apply(opts ...SubscriptionOption) {
	for _, opt := range opts {
		opt(s)
	}
}

func (s *Subscription) UpcomingMatches(ctx context.Context) (<-chan *model.UpcomingMatchesResponse, error) {
	panic("implement me")
}
