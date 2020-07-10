package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/cshep4/premier-predictor-microservices/src/common/grpc/options"
	"github.com/cshep4/premier-predictor-microservices/src/common/log"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

const defaultPort = 3000

type (
	Registerer interface {
		Register(*grpc.Server)
	}

	server struct {
		registerers        []Registerer
		unaryInterceptors  []grpc.UnaryServerInterceptor
		streamInterceptors []grpc.StreamServerInterceptor
		grpcs              *grpc.Server
		port               int
	}
)

func New(opts ...option) *server {
	s := &server{
		port: defaultPort,
	}

	for _, opt := range opts {
		opt(s)
	}

	s.grpcs = grpc.NewServer(
		options.ServerKeepAlive,
		grpc_middleware.WithUnaryServerChain(s.unaryInterceptors...),
		grpc_middleware.WithStreamServerChain(s.streamInterceptors...),
	)

	return s
}

func (s *server) Start(ctx context.Context) error {
	path := fmt.Sprintf(":%d", s.port)

	lis, err := net.Listen("tcp", path)
	if err != nil {
		return fmt.Errorf("listen: %v", err)
	}

	log.Info(ctx, "grpc_server_listening", log.SafeParam("path", path))

	for i := range s.registerers {
		s.registerers[i].Register(s.grpcs)
	}

	err = s.grpcs.Serve(lis)
	if err != nil {
		return fmt.Errorf("start_server: %v", err)
	}

	return nil
}

func (s *server) Stop(ctx context.Context) error {
	s.grpcs.GracefulStop()
	log.Info(ctx, "grpc_server_stopped")

	return nil
}
