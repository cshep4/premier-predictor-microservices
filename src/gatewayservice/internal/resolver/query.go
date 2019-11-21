package resolver

import (
	"context"

	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/gatewayservice/internal/model"
)

type Query struct {
	userService gen.UserServiceClient
}

func NewQuery(opts ...QueryOption) *Query {
	q := new(Query)
	q.Apply(opts...)
	return q
}

func (q *Query) Apply(opts ...QueryOption) {
	for _, opt := range opts {
		opt(q)
	}
}

func (q *Query) UserRank(ctx context.Context, id string) (*model.UserRank, error) {
	panic("implement me")
}
