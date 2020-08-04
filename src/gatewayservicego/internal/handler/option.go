package handler

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	graphqlhandler "github.com/99designs/gqlgen/handler"
	"github.com/cshep4/premier-predictor-microservices/src/gatewayservice/internal/schema"
	"log"
)

type Option func(*Handler) error

func Resolver(r schema.ResolverRoot) Option {
	return func(h *Handler) error {
		h.resolver = r
		return nil
	}
}

func GraphQLOptions(opts ...graphqlhandler.Option) Option {
	return func(h *Handler) error {
		h.graphqlopts = append(h.graphqlopts, opts...)
		return nil
	}
}

func Middleware() Option {
	m := func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
		res, err = next(ctx)
		if err != nil {
			log.Println(fmt.Sprintf("an error occurred with request: %s", err))
		}
		return res, err
	}

	return func(h *Handler) error {
		return GraphQLOptions(graphqlhandler.ResolverMiddleware(m))(h)
	}
}
