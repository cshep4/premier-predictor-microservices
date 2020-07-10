package http

import (
	"net/http"

	"github.com/cshep4/premier-predictor-microservices/src/common/log"
	"github.com/gorilla/mux"
)

type option func(*server)

func WithPort(p int) option {
	return func(s *server) {
		s.port = p
	}
}

func WithLogger(service, level string) option {
	return func(s *server) {
		logger := log.New(level)

		s.middlewares = append(s.middlewares, func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				next.ServeHTTP(w, r.WithContext(log.WithServiceName(r.Context(), logger, service)))
			})
		})
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
