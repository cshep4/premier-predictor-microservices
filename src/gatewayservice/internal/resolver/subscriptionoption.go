package resolver

import (
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
)

type SubscriptionOption func(*Subscription)

func SubscriptionAuthService(a gen.AuthServiceClient) SubscriptionOption {
	return func(s *Subscription) {
		s.authService = a
	}
}

func SubscriptionLiveMatchService(l gen.LiveMatchServiceClient) SubscriptionOption {
	return func(s *Subscription) {
		s.liveMatchService = l
	}
}
