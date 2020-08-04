package mutation

import (
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
)

type AuthOption func(*Auth)

func AuthService(a gen.AuthServiceClient) AuthOption {
	return func(m *Auth) {
		m.authService = a
	}
}
