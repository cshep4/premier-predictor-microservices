package tracer

import (
	"context"

	grpccfg "github.com/cshep4/premier-predictor-microservices/src/common/grpc"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"
)

func (tracer) GrpcUnary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx, span := trace.StartSpan(ctx, info.FullMethod)
	defer span.End()

	return handler(ctx, req)
}

func (tracer) GrpcStream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ctx, span := trace.StartSpan(stream.Context(), info.FullMethod)
	defer span.End()

	return handler(srv, &grpccfg.ContextServerStream{
		Ctx:          ctx,
		ServerStream: stream,
	})
}
