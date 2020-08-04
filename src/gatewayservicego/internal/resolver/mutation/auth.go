package mutation

import (
	"context"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/gatewayservice/internal/model"
)

type Auth struct {
	authService gen.AuthServiceClient
}

func NewAuth(opts ...AuthOption) Auth {
	m := Auth{}
	m.Apply(opts...)
	return m
}

func (m *Auth) Apply(opts ...AuthOption) {
	for _, opt := range opts {
		opt(m)
	}
}

func (m *Auth) Login(ctx context.Context, request model.LoginInput) (*model.LoginResponse, error) {
	panic("implement me")
}

func (m *Auth) Register(ctx context.Context, request model.RegisterInput) (bool, error) {
	panic("implement me")
}

func (m *Auth) InitiatePasswordReset(ctx context.Context, email string) (bool, error) {
	panic("implement me")
}

func (m *Auth) ResetPassword(ctx context.Context, request model.ResetPasswordInput) (bool, error) {
	panic("implement me")
}
