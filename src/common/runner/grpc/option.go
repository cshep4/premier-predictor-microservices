package grpc

import "google.golang.org/grpc"

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
