package grpc

import (
	"context"
	grpccfg "github.com/cshep4/premier-predictor-microservices/src/common/grpc"
	"github.com/cshep4/premier-predictor-microservices/src/common/log"
	"google.golang.org/grpc"
)

type option func(*server)

func WithRegisterer(r Registerer) option {
	return func(s *server) {
		s.registerers = append(s.registerers, r)
	}
}

func WithPort(p int) option {
	return func(s *server) {
		s.port = p
	}
}

func WithLogger(service, level string) option {
	return func(s *server) {
		logger := log.New(level)

		s.unaryInterceptors = append(s.unaryInterceptors, func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
			return handler(log.WithServiceName(ctx, logger, service), req)
		})
		s.streamInterceptors = append(s.streamInterceptors, func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			return handler(srv, &grpccfg.ContextServerStream{
				Ctx:          log.WithServiceName(ss.Context(), logger, service),
				ServerStream: ss,
			})
		})
	}
}

func WithUnaryInterceptor(i grpc.UnaryServerInterceptor) option {
	return func(s *server) {
		s.unaryInterceptors = append(s.unaryInterceptors, i)
	}
}

func WithStreamInterceptor(i grpc.StreamServerInterceptor) option {
	return func(s *server) {
		s.streamInterceptors = append(s.streamInterceptors, i)
	}
}
