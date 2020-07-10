package http

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Handle struct {
	path    string
	f       func(http.ResponseWriter, *http.Request)
	methods []string
}

func NewHandle(path string, f func(http.ResponseWriter, *http.Request), methods ...string) Handle {
	return Handle{
		path:    path,
		f:       f,
		methods: methods,
	}
}

func (h Handle) Register(router *mux.Router) {
	router.HandleFunc(h.path, h.f).
		Methods(h.methods...)
}

func Health() Handle {
	return NewHandle("/health",
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		},
		http.MethodGet,
	)
}
