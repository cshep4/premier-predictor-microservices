package resolver_test

import (
	"testing"

	"github.com/cshep4/premier-predictor-microservices/src/gatewayservice/internal/resolver"
	"github.com/cshep4/premier-predictor-microservices/src/gatewayservice/internal/schema"
	"github.com/stretchr/testify/assert"
)

func TestNewRoot(t *testing.T) {
	t.Run("should implement schema.ResolverRoot", func(t *testing.T) {
		var i interface{} = &resolver.Root{}
		if _, ok := i.(schema.ResolverRoot); !ok {
			t.Fail()
		}
	})

	t.Run("should make new root and return resolvers", func(t *testing.T) {
		r := resolver.NewRoot(
			resolver.RootQuery(),
			resolver.RootMutation(),
			resolver.RootSubscription(),
		)
		assert.NotNil(t, r.Mutation())
		assert.NotNil(t, r.Query())
		assert.NotNil(t, r.Subscription())
	})
}
