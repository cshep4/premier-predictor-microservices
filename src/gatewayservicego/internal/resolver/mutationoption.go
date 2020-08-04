package resolver

import (
	"github.com/cshep4/premier-predictor-microservices/src/gatewayservice/internal/resolver/mutation"
)

type MutationOption func(*Mutation)

func MutationAuth(a mutation.Auth) MutationOption {
	return func(m *Mutation) {
		m.Auth = a
	}
}

func MutationChat(c mutation.Chat) MutationOption {
	return func(m *Mutation) {
		m.Chat = c
	}
}
