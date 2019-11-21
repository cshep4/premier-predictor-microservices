package resolver

import (
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
)

type QueryOption func(*Query)

func QueryUserService(u gen.UserServiceClient) QueryOption {
	return func(q *Query) {
		q.userService = u
	}
}
