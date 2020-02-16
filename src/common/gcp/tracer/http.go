package tracer

import (
	"net/http"

	"go.opencensus.io/plugin/ochttp"
)

func (tracer) Wrap(handler http.Handler) http.Handler {
	return &ochttp.Handler{
		Handler:          handler,
		IsPublicEndpoint: true,
		IsHealthEndpoint: func(r *http.Request) bool {
			return r.URL.Path == "/health"
		},
	}
}
