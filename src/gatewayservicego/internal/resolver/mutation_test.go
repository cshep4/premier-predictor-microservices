package resolver_test

import (
	"testing"

	"github.com/cshep4/premier-predictor-microservices/src/gatewayservice/internal/resolver"
	"github.com/cshep4/premier-predictor-microservices/src/gatewayservice/internal/schema"
)

func TestNewMutation(t *testing.T) {
	var i interface{} = new(resolver.Mutation)
	if _, ok := i.(schema.MutationResolver); !ok {
		t.Fail()
	}
}
