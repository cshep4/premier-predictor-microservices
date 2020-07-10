package grpc

import (
	"context"
	"google.golang.org/grpc"
)

type ContextServerStream struct {
	Ctx context.Context
	grpc.ServerStream
}

func (c *ContextServerStream) Context() context.Context {
	return c.Ctx
}
