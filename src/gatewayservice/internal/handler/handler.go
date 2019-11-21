package handler

import (
	"context"
	graphqlhandler "github.com/99designs/gqlgen/handler"
	"github.com/cshep4/premier-predictor-microservices/src/common/health"
	"github.com/cshep4/premier-predictor-microservices/src/gatewayservice/internal/resolver"
	"github.com/cshep4/premier-predictor-microservices/src/gatewayservice/internal/schema"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

// Handler handles HTTP requests.
type Handler struct {
	router      *mux.Router
	resolver    schema.ResolverRoot
	graphqlopts []graphqlhandler.Option
}

// New returns a new http.Handler.
func New(opts ...Option) (*Handler, error) {
	h := &Handler{
		router: mux.NewRouter(),
		graphqlopts: []graphqlhandler.Option{
			graphqlhandler.WebsocketUpgrader(
				websocket.Upgrader{
					CheckOrigin: func(r *http.Request) bool {
						return true
					},
				},
			),
			graphqlhandler.WebsocketKeepAliveDuration(time.Second),
		},
	}

	if err := h.Apply(opts...); err != nil {
		return nil, err
	}

	h.router.HandleFunc("/graphql", graphqlhandler.GraphQL(
		schema.NewExecutableSchema(schema.Config{Resolvers: h.resolver}),
		h.graphqlopts...,
	)).Methods(http.MethodPost, http.MethodGet)

	h.router.Handle("/", graphqlhandler.Playground("GatewayService", "/gateway/graphql"))

	h.router.HandleFunc("/health", health.Health).Methods(http.MethodGet)

	return h, nil
}

func (h *Handler) Apply(opts ...Option) error {
	for _, opt := range opts {
		if err := opt(h); err != nil {
			return err
		}
	}
	return nil
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const (
		allowedHeaders = "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"
		allowedMethods = "POST, GET, OPTIONS, PUT, DELETE"
	)
	if origin, ok := r.Header["Origin"]; ok && len(origin) == 1 {
		w.Header().Set("Access-Control-Allow-Origin", origin[0])
	}
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Allow-Methods", allowedMethods)

	if r.Method != http.MethodOptions {
		token := r.Header.Get("Authorization")
		ctx := context.WithValue(r.Context(), resolver.TokenKey, token)

		h.router.ServeHTTP(w, r.WithContext(ctx))
	}
}
