package resolver

import (
	"github.com/cshep4/premier-predictor-microservices/src/gatewayservice/internal/schema"
)

type contextKey string

const TokenKey = contextKey("token")

type Root struct {
	mutation     schema.MutationResolver
	query        schema.QueryResolver
	subscription schema.SubscriptionResolver
}

func NewRoot(opts ...RootOption) *Root {
	r := new(Root)
	r.Apply(opts...)
	return r
}

func (r *Root) Apply(opts ...RootOption) {
	for _, opt := range opts {
		opt(r)
	}
}

func (r *Root) Mutation() schema.MutationResolver         { return r.mutation }
func (r *Root) Query() schema.QueryResolver               { return r.query }
func (r *Root) Subscription() schema.SubscriptionResolver { return r.subscription }
