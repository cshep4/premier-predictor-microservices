package resolver

import (
	"github.com/cshep4/premier-predictor-microservices/src/gatewayservice/internal/resolver/mutation"
)

type Mutation struct {
	mutation.Auth
	mutation.Chat
}

func NewMutation(opts ...MutationOption) *Mutation {
	m := new(Mutation)
	m.Apply(opts...)
	return m
}

func (m *Mutation) Apply(opts ...MutationOption) {
	for _, opt := range opts {
		opt(m)
	}
}
