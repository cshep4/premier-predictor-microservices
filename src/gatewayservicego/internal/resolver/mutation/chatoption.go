package mutation

import (
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
)

type ChatOption func(*Chat)

func ChatService(cs gen.ChatServiceClient) ChatOption {
	return func(c *Chat) {
		c.chatService = cs
	}
}
