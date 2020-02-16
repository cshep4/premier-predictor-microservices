package tracer

import (
	"context"

	"go.opencensus.io/trace"
	"google.golang.org/grpc"
)

type serverStream struct {
		ctx context.Context
		grpc.ServerStream
	}

func (tracer) GrpcUnary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx, span := trace.StartSpan(ctx, info.FullMethod)
	defer span.End()

	return handler(ctx, req)
}

func (tracer) GrpcStream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ctx, span := trace.StartSpan(stream.Context(), info.FullMethod)
	defer span.End()

	s := &serverStream{
		ctx:          ctx,
		ServerStream: stream,
	}

	return handler(srv, s)
}
