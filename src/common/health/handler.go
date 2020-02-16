package health

import (
	"github.com/gorilla/mux"
	"net/http"
)

type healthServiceServer struct{}

func (s *healthServiceServer) createRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", Health).Methods("GET")

	return r
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
