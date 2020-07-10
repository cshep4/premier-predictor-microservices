package auth

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (a *authenticator) Http(next http.Handler) http.Handler {
	unauthenticatedEndpoints := map[string]struct{}{
		"/health": {},
		"/legacy": {},
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p, _ := mux.CurrentRoute(r).GetPathTemplate()
		if _, ok := unauthenticatedEndpoints[p]; ok {
			next.ServeHTTP(w, r)
			return
		}

		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx, err := a.doAuth(r.Context(), token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
