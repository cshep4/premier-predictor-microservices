// +build integration

package integration

import (
	"context"
	"testing"
	"time"

	model "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGRPC_Login(t *testing.T) {
	t.Run("will return error if email is empty", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			Login(ctx, &model.LoginRequest{})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Equal(t, "email is empty", st.Message())
	})

	t.Run("will return error if password is empty", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			Login(ctx, &model.LoginRequest{
				Email: validEmail,
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Equal(t, "password is empty", st.Message())
	})

	t.Run("will return error if error getting user", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			Login(ctx, &model.LoginRequest{
				Email:    getUserErrorEmail,
				Password: password,
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Unauthenticated, st.Code())
		assert.Equal(t, "could not login", st.Message())
	})

	t.Run("will return error if user not found", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			Login(ctx, &model.LoginRequest{
				Email:    notFoundErrorEmail,
				Password: password,
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Unauthenticated, st.Code())
		assert.Equal(t, "user not found", st.Message())
	})

	t.Run("will return error if password does not match", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			Login(ctx, &model.LoginRequest{
				Email:    validEmail,
				Password: "some different password",
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Unauthenticated, st.Code())
		assert.Equal(t, "password does not match", st.Message())
	})

	t.Run("will return user's ID and token if the credentials are correct", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		res, err := model.NewAuthServiceClient(conn).
			Login(ctx, &model.LoginRequest{
				Email:    validEmail,
				Password: password,
			})
		require.NoError(t, err)

		assert.Equal(t, validEmail, res.Id)
		assert.NotEmpty(t, res.Token)

		_, err = model.NewAuthServiceClient(conn).
			Validate(ctx, &model.ValidateRequest{
				Token:    res.Token,
				Audience: validEmail,
				Role:     model.Role_ROLE_USER,
			})
		require.NoError(t, err)
	})
}

func TestGRPC_Register(t *testing.T) {
	t.Run("will return first name if email is empty", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			Register(ctx, &model.RegisterRequest{})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Equal(t, "first name is empty", st.Message())
	})

	t.Run("will return last name if email is empty", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			Register(ctx, &model.RegisterRequest{
				FirstName: "first name",
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Equal(t, "surname is empty", st.Message())
	})

	t.Run("will return error if email is empty", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			Register(ctx, &model.RegisterRequest{
				FirstName: "first name",
				Surname:   "last name",
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Equal(t, "email is empty", st.Message())
	})

	t.Run("will return error if password is empty", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			Register(ctx, &model.RegisterRequest{
				FirstName: "first name",
				Surname:   "last name",
				Email:     "email",
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Equal(t, "password is empty", st.Message())
	})

	t.Run("will return error if confirmation is empty", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			Register(ctx, &model.RegisterRequest{
				FirstName: "first name",
				Surname:   "last name",
				Email:     "email",
				Password:  "password",
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Equal(t, "confirmation is empty", st.Message())
	})

	t.Run("will return error if predicted winner is empty", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			Register(ctx, &model.RegisterRequest{
				FirstName:    "first name",
				Surname:      "last name",
				Email:        "email",
				Password:     "password",
				Confirmation: "confirmation",
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Equal(t, "predicted winner is empty", st.Message())
	})

	t.Run("will return error if email is invalid", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			Register(ctx, &model.RegisterRequest{
				FirstName:       "first name",
				Surname:         "last name",
				Email:           "email",
				Password:        "password",
				Confirmation:    "confirmation",
				PredictedWinner: "predicted winner",
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Equal(t, "email address is invalid", st.Message())
	})

	t.Run("will return error if password is invalid", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			Register(ctx, &model.RegisterRequest{
				FirstName:       "first name",
				Surname:         "last name",
				Email:           validEmail,
				Password:        "password",
				Confirmation:    "confirmation",
				PredictedWinner: "predicted winner",
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Equal(t, "password is invalid", st.Message())
	})

	t.Run("will return error if password and confirmation does not match", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			Register(ctx, &model.RegisterRequest{
				FirstName:       "first name",
				Surname:         "last name",
				Email:           validEmail,
				Password:        password,
				Confirmation:    "confirmation",
				PredictedWinner: "predicted winner",
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Equal(t, "password and confirmation do not match", st.Message())
	})

	t.Run("will return error if user already exists with email", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			Register(ctx, &model.RegisterRequest{
				FirstName:       "first name",
				Surname:         "last name",
				Email:           validEmail,
				Password:        password,
				Confirmation:    password,
				PredictedWinner: "predicted winner",
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.AlreadyExists, st.Code())
		assert.Equal(t, "email already exists", st.Message())
	})

	t.Run("will return error if error getting user by email", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			Register(ctx, &model.RegisterRequest{
				FirstName:       "first name",
				Surname:         "last name",
				Email:           getUserErrorEmail,
				Password:        password,
				Confirmation:    password,
				PredictedWinner: "predicted winner",
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Internal, st.Code())
		assert.Equal(t, "could not get user", st.Message())
	})

	t.Run("will return error if error creating user", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			Register(ctx, &model.RegisterRequest{
				FirstName:       "first name",
				Surname:         "last name",
				Email:           failCreateErrorEmail,
				Password:        password,
				Confirmation:    password,
				PredictedWinner: "predicted winner",
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Internal, st.Code())
		assert.Equal(t, "could not register", st.Message())
	})

	t.Run("will return user's ID and token if user created successfully", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		res, err := model.NewAuthServiceClient(conn).
			Register(ctx, &model.RegisterRequest{
				FirstName:       "first name",
				Surname:         "last name",
				Email:           notFoundErrorEmail,
				Password:        password,
				Confirmation:    password,
				PredictedWinner: "predicted winner",
			})
		require.NoError(t, err)

		assert.Equal(t, createdUserId, res.Id)
		assert.NotEmpty(t, res.Token)

		_, err = model.NewAuthServiceClient(conn).
			Validate(ctx, &model.ValidateRequest{
				Token:    res.Token,
				Audience: createdUserId,
				Role:     model.Role_ROLE_USER,
			})
		require.NoError(t, err)
	})
}

func TestGRPC_Validate(t *testing.T) {
	t.Run("will return error if token is empty", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			Validate(ctx, &model.ValidateRequest{})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Equal(t, "token is empty", st.Message())
	})

	t.Run("will return error if token is invalid", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			Validate(ctx, &model.ValidateRequest{
				Token: "token",
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Equal(t, "role is invalid", st.Message())
	})

	t.Run("will return unauthenticated if token is invalid", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			Validate(ctx, &model.ValidateRequest{
				Token: "invalid token",
				Role:  model.Role_ROLE_USER,
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Unauthenticated, st.Code())
		assert.Equal(t, "could not verify token", st.Message())
	})

	t.Run("will return authenticated if token signed with different secret", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.StandardClaims{}).
			SignedString([]byte("different secret"))
		require.NoError(t, err)

		_, err = model.NewAuthServiceClient(conn).
			Validate(ctx, &model.ValidateRequest{
				Token: token,
				Role:  model.Role_ROLE_USER,
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Unauthenticated, st.Code())
		assert.Equal(t, "could not verify token", st.Message())
	})

	t.Run("will return unauthenticated if expired", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		claims := &jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(-1, 0, 0).Unix(),
		}
		token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, claims).
			SignedString([]byte(jwtSecret))
		require.NoError(t, err)

		_, err = model.NewAuthServiceClient(conn).
			Validate(ctx, &model.ValidateRequest{
				Token: token,
				Role:  model.Role_ROLE_USER,
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Unauthenticated, st.Code())
		assert.Equal(t, "could not verify token", st.Message())
	})

	t.Run("will return unauthenticated if audience is incorrect", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		claims := &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second).Unix(),
			Audience:  "audience",
		}
		token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, claims).
			SignedString([]byte(jwtSecret))
		require.NoError(t, err)

		_, err = model.NewAuthServiceClient(conn).
			Validate(ctx, &model.ValidateRequest{
				Token:    token,
				Audience: "incorrect",
				Role:     model.Role_ROLE_USER,
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Unauthenticated, st.Code())
		assert.Equal(t, "could not verify token", st.Message())
	})

	t.Run("will return unauthenticated if role does not match", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		var claims = &struct {
			Role string `json:"role"`
			jwt.StandardClaims
		}{
			Role: "SERVICE",
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Second).Unix(),
				Audience:  "audience",
			},
		}
		token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, claims).
			SignedString([]byte(jwtSecret))
		require.NoError(t, err)

		_, err = model.NewAuthServiceClient(conn).
			Validate(ctx, &model.ValidateRequest{
				Token:    token,
				Audience: "audience",
				Role:     model.Role_ROLE_USER,
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Unauthenticated, st.Code())
		assert.Equal(t, "could not verify token", st.Message())
	})

	t.Run("will validate service token", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		var claims = &struct {
			Role string `json:"role"`
			jwt.StandardClaims
		}{
			Role: "SERVICE",
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Second).Unix(),
				Audience:  "audience",
			},
		}

		token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, claims).
			SignedString([]byte(jwtSecret))
		require.NoError(t, err)

		_, err = model.NewAuthServiceClient(conn).
			Validate(ctx, &model.ValidateRequest{
				Token:    token,
				Audience: "audience",
				Role:     model.Role_ROLE_SERVICE,
			})
		require.NoError(t, err)
	})

	t.Run("will validate user token", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		var claims = &struct {
			Role string `json:"role"`
			jwt.StandardClaims
		}{
			Role: "USER",
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Second).Unix(),
				Audience:  "audience",
			},
		}
		token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, claims).
			SignedString([]byte(jwtSecret))
		require.NoError(t, err)

		_, err = model.NewAuthServiceClient(conn).
			Validate(ctx, &model.ValidateRequest{
				Token:    token,
				Audience: "audience",
				Role:     model.Role_ROLE_USER,
			})
		require.NoError(t, err)
	})

	t.Run("will validate admin token", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		var claims = &struct {
			Role string `json:"role"`
			jwt.StandardClaims
		}{
			Role: "ADMIN",
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Second).Unix(),
				Audience:  "audience",
			},
		}
		token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, claims).
			SignedString([]byte(jwtSecret))
		require.NoError(t, err)

		_, err = model.NewAuthServiceClient(conn).
			Validate(ctx, &model.ValidateRequest{
				Token:    token,
				Audience: "audience",
				Role:     model.Role_ROLE_ADMIN,
			})
		require.NoError(t, err)
	})

	t.Run("will validate token if audience is empty", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		var claims = &struct {
			Role string `json:"role"`
			jwt.StandardClaims
		}{
			Role: "USER",
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Second).Unix(),
				Audience:  "audience",
			},
		}
		token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, claims).
			SignedString([]byte(jwtSecret))
		require.NoError(t, err)

		_, err = model.NewAuthServiceClient(conn).
			Validate(ctx, &model.ValidateRequest{
				Token: token,
				Role:  model.Role_ROLE_USER,
			})
		require.NoError(t, err)
	})
}

func TestGRPC_IssueServiceToken(t *testing.T) {
	t.Run("will create service token for specified audience", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		res, err := model.NewAuthServiceClient(conn).
			IssueServiceToken(ctx, &model.IssueServiceTokenRequest{
				Audience: "audience",
			})
		require.NoError(t, err)

		_, err = model.NewAuthServiceClient(conn).
			Validate(ctx, &model.ValidateRequest{
				Token:    res.Token,
				Audience: "audience",
				Role:     model.Role_ROLE_USER,
			})
		require.Error(t, err)

		_, err = model.NewAuthServiceClient(conn).
			Validate(ctx, &model.ValidateRequest{
				Token:    res.Token,
				Audience: "audience",
				Role:     model.Role_ROLE_ADMIN,
			})
		require.Error(t, err)

		_, err = model.NewAuthServiceClient(conn).
			Validate(ctx, &model.ValidateRequest{
				Token:    res.Token,
				Audience: "wrong audience",
				Role:     model.Role_ROLE_SERVICE,
			})
		require.Error(t, err)

		_, err = model.NewAuthServiceClient(conn).
			Validate(ctx, &model.ValidateRequest{
				Token:    res.Token,
				Audience: "audience",
				Role:     model.Role_ROLE_SERVICE,
			})
		require.NoError(t, err)
	})
}

func TestGRPC_InitiatePasswordReset(t *testing.T) {
	t.Run("will return error if email is empty", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			InitiatePasswordReset(ctx, &model.InitiatePasswordResetRequest{})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Equal(t, "email is empty", st.Message())
	})

	t.Run("will return error if email is invalid", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			InitiatePasswordReset(ctx, &model.InitiatePasswordResetRequest{
				Email: "invalid email",
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Equal(t, "email address is invalid", st.Message())
	})

	t.Run("will return error if error getting user", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			InitiatePasswordReset(ctx, &model.InitiatePasswordResetRequest{
				Email: notFoundErrorEmail,
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Equal(t, "user not found", st.Message())
	})

	t.Run("will return error if error updating signature user", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			InitiatePasswordReset(ctx, &model.InitiatePasswordResetRequest{
				Email: failUpdateSignatureErrorEmail,
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Internal, st.Code())
		assert.Equal(t, "could not initiate password reset", st.Message())
	})

	t.Run("will return error if error sending email", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			InitiatePasswordReset(ctx, &model.InitiatePasswordResetRequest{
				Email: failEmailErrorEmail,
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Internal, st.Code())
		assert.Equal(t, "could not initiate password reset", st.Message())
	})

	t.Run("will update user's signature and send password reset email", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			InitiatePasswordReset(ctx, &model.InitiatePasswordResetRequest{
				Email: validEmail,
			})
		require.NoError(t, err)
	})
}

func TestGRPC_ResetPassword(t *testing.T) {
	t.Run("will return error if email is empty", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			ResetPassword(ctx, &model.ResetPasswordRequest{})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Equal(t, "email is empty", st.Message())
	})

	t.Run("will return error if signature is empty", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			ResetPassword(ctx, &model.ResetPasswordRequest{
				Email: validEmail,
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Equal(t, "signature is empty", st.Message())
	})

	t.Run("will return error if password is empty", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			ResetPassword(ctx, &model.ResetPasswordRequest{
				Email:     validEmail,
				Signature: "signature",
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Equal(t, "password is empty", st.Message())
	})

	t.Run("will return error if confirmation is empty", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			ResetPassword(ctx, &model.ResetPasswordRequest{
				Email:     validEmail,
				Signature: "signature",
				Password:  "password",
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Equal(t, "confirmation is empty", st.Message())
	})

	t.Run("will return error if password is invalid", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			ResetPassword(ctx, &model.ResetPasswordRequest{
				Email:        validEmail,
				Signature:    "signature",
				Password:     "password",
				Confirmation: "confirmation",
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Equal(t, "password is invalid", st.Message())
	})

	t.Run("will return error if password and confirmation does not match", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			ResetPassword(ctx, &model.ResetPasswordRequest{
				Email:        validEmail,
				Signature:    "signature",
				Password:     password,
				Confirmation: "different password",
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Equal(t, "password and confirmation do not match", st.Message())
	})

	t.Run("will return error if signature not valid", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		signature, err := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.StandardClaims{}).
			SignedString([]byte("different secret"))
		require.NoError(t, err)

		_, err = model.NewAuthServiceClient(conn).
			ResetPassword(ctx, &model.ResetPasswordRequest{
				Email:        validEmail,
				Signature:    signature,
				Password:     password,
				Confirmation: password,
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Internal, st.Code())
		assert.Equal(t, "could not reset password", st.Message())
	})

	t.Run("will return error if error getting user", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		signature, err := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.StandardClaims{}).
			SignedString([]byte(jwtSecret))
		require.NoError(t, err)

		_, err = model.NewAuthServiceClient(conn).
			ResetPassword(ctx, &model.ResetPasswordRequest{
				Email:        getUserErrorEmail,
				Signature:    signature,
				Password:     password,
				Confirmation: password,
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Internal, st.Code())
		assert.Equal(t, "could not reset password", st.Message())
	})

	t.Run("will return error if signature does not match", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		signature, err := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.StandardClaims{
			Audience: "will create different signature",
		}).
			SignedString([]byte(jwtSecret))
		require.NoError(t, err)

		_, err = model.NewAuthServiceClient(conn).
			ResetPassword(ctx, &model.ResetPasswordRequest{
				Email:        validEmail,
				Signature:    signature,
				Password:     password,
				Confirmation: password,
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Equal(t, "signature does not match", st.Message())
	})

	t.Run("will return error if error updating password", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		signature, err := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.StandardClaims{}).
			SignedString([]byte(jwtSecret))
		require.NoError(t, err)

		_, err = model.NewAuthServiceClient(conn).
			ResetPassword(ctx, &model.ResetPasswordRequest{
				Email:        failUpdatePasswordErrorEmail,
				Signature:    signature,
				Password:     password,
				Confirmation: password,
			})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Internal, st.Code())
		assert.Equal(t, "could not reset password", st.Message())
	})

	t.Run("will successfully update user's password", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		signature, err := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.StandardClaims{}).
			SignedString([]byte(jwtSecret))
		require.NoError(t, err)

		_, err = model.NewAuthServiceClient(conn).
			ResetPassword(ctx, &model.ResetPasswordRequest{
				Email:        validEmail,
				Signature:    signature,
				Password:     password,
				Confirmation: password,
			})
		require.NoError(t, err)
	})
}
