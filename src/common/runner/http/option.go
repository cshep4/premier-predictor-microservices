package http

import "github.com/gorilla/mux"

type option func(*server)

func WithPort(p int) option {
	return func(s *server) {
		s.port = p
	}
}

func WithMiddleware(m mux.MiddlewareFunc) option {
	return func(s *server) {
		s.middlewares = append(s.middlewares, m)
	}
}

func WithRouter(r Router) option {
	return func(s *server) {
		s.routers = append(s.routers, r)
	}
}

func WithRegisterer(r Registerer) option {
	return func(s *server) {
		s.registerers = append(s.registerers, r)
	}
}

func WithHandler(h HandlerWrapper) option {
	return func(s *server) {
		s.handlers = append(s.handlers, h)
	}
}
