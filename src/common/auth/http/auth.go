package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cshep4/premier-predictor-microservices/src/common/auth/internal/context"

	"github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/gorilla/mux"
)

var unauthenticatedEndpoints = map[string]struct{}{
	"/health": {},
}

type (
	Authorizer interface {
		IsUnauthenticatedEndpoint(path string) bool
		GetRequestAudience(r *http.Request) string
	}

	authenticator struct {
		authClient model.AuthServiceClient
		authorizer Authorizer
	}

	// InvalidParameterError is returned when a required parameter passed to New is invalid.
	InvalidParameterError struct {
		Parameter string
	}
)

func (i InvalidParameterError) Error() string {
	return fmt.Sprintf("invalid parameter %s", i.Parameter)
}

func New(authClient model.AuthServiceClient, authorizer Authorizer) (*authenticator, error) {
	switch {
	case authClient == nil:
		return nil, InvalidParameterError{Parameter: "authClient"}
	case authorizer == nil:
		return nil, InvalidParameterError{Parameter: "authorizer"}
	}

	return &authenticator{
		authClient: authClient,
		authorizer: authorizer,
	}, nil
}

func (a *authenticator) Http(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p, err := mux.CurrentRoute(r).GetPathTemplate()
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		_, ok := unauthenticatedEndpoints[p]
		if ok || a.authorizer.IsUnauthenticatedEndpoint(p) {
			next.ServeHTTP(w, r)
			return
		}

		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		audience := a.authorizer.GetRequestAudience(r)

		ctx, err := a.doAuth(r.Context(), token, audience, model.Role_ROLE_USER)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *authenticator) doAuth(ctx context.Context, token, audience string, role model.Role) (context.Context, error) {
	_, err := a.authClient.Validate(ctx, &model.ValidateRequest{
		Token:    token,
		Audience: audience,
		Role:     role,
	})
	if err != nil {
		return nil, err
	}

	return auth.SetTokenCtx(ctx, token), nil
}
