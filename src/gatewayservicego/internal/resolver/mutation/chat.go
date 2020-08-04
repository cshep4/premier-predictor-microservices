package mutation

import (
	"context"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/gatewayservice/internal/model"
)

type Chat struct {
	chatService gen.ChatServiceClient
}

func NewChat(opts ...ChatOption) Chat {
	c := Chat{}
	c.Apply(opts...)
	return c
}

func (c *Chat) Apply(opts ...ChatOption) {
	for _, opt := range opts {
		opt(c)
	}
}

func (c *Chat) SendMessage(ctx context.Context, request model.SendMessageInput) (bool, error) {
	panic("implement me")
}
