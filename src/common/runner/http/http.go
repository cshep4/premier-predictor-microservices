package http

import (
	"context"
	"fmt"
	"github.com/cshep4/premier-predictor-microservices/src/common/cors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const defaultPort = 8080

type (
	Router interface {
		Route(*mux.Router)
	}

	Registerer interface {
		Register(*mux.Router)
	}

	HandlerWrapper interface {
		Wrap(http.Handler) http.Handler
	}

	server struct {
		handlers    []HandlerWrapper
		routers     []Router
		registerers []Registerer
		middlewares []mux.MiddlewareFunc
		https       *http.Server
		port        int
	}
)

func New(opts ...option) *server {
	s := &server{
		port: defaultPort,
		https: &http.Server{
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s server) Start(ctx context.Context) error {
	path := fmt.Sprintf(":%d", s.port)

	router := mux.NewRouter()

	for _, m := range s.middlewares {
		router.Use(m)
	}
	for _, r := range s.routers {
		r.Route(router)
	}
	for _, r := range s.registerers {
		r.Register(router)
	}

	s.https.Addr = path
	s.https.Handler = cors.New().
		Handler(router)

	for _, h := range s.handlers {
		s.https.Handler = h.Wrap(s.https.Handler)
	}

	log.Printf("http_server_listening_on: %s", path)

	err := s.https.ListenAndServe()
	if err != nil {
		return fmt.Errorf("listen_and_serve: %v", err)
	}

	return nil
}

func (s server) Stop(ctx context.Context) error {
	err := s.https.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("shutdown: %v", err)
	}

	log.Println("http_server_stopped")

	return nil
}
