package main

import (
	"context"
	"net"

	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

const (
	password  = "plaintextPassword123"
	createdId = "created user id"

	failUpdatePasswordID = "fail update password id"
	failUpdateSignature  = "fail update signature"

	jwtSecret = "some-jwt-secret"
)

func (s server) GetAllUsers(context.Context, *gen.GetAllUsersRequest) (*gen.GetAllUsersResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetAllUsers not implemented")
}

func (s server) GetUsersByIds(context.Context, *gen.GetUsersByIdsRequest) (*gen.GetUsersByIdsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetAllUsersByIds not implemented")
}

func (s server) GetOverallRank(context.Context, *gen.GetOverallRankRequest) (*gen.GetOverallRankResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetOverallRank not implemented")
}

func (s server) GetRankForGroup(context.Context, *gen.GetRankForGroupRequest) (*gen.GetRankForGroupResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetRankForGroup not implemented")
}

func (s server) GetUserCount(context.Context, *gen.GetUserCountRequest) (*gen.GetUserCountResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetUserCount not implemented")
}

func (s server) GetUserByEmail(_ context.Context, req *gen.GetUserByEmailRequest) (*gen.GetUserByEmailResponse, error) {
	switch req.Email {
	case "notfound@test.com", "failcreate@test.com":
		return nil, status.Error(codes.NotFound, "user not found")
	case "getusererror@test.com":
		return nil, status.Error(codes.Internal, "error getting user")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, status.Error(codes.Internal, "error hashing password")
	}

	signature, err := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.StandardClaims{}).
		SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, status.Error(codes.Internal, "error generating signature")
	}

	return &gen.GetUserByEmailResponse{
		User: &gen.User{
			Id:        getUserId(req.Email),
			FirstName: "first name",
			Surname:   "last name",
			Email:     req.Email,
			Password:  string(password),
			Signature: signature,
		},
	}, nil
}

func getUserId(email string) string {
	switch email {
	case "failupdatepassword@test.com":
		return failUpdatePasswordID
	case "failupdatesignature@test.com":
		return failUpdateSignature
	}

	return email
}

func (s server) UpdatePassword(_ context.Context, req *gen.UpdatePasswordRequest) (*gen.UpdatePasswordResponse, error) {
	if req.Id == failUpdatePasswordID {
		return nil, status.Error(codes.Internal, "error updating password")
	}

	return &gen.UpdatePasswordResponse{}, nil
}

func (s server) UpdateSignature(_ context.Context, req *gen.UpdateSignatureRequest) (*gen.UpdateSignatureResponse, error) {
	if req.Id == failUpdateSignature {
		return nil, status.Error(codes.Internal, "error updating signature")
	}

	return &gen.UpdateSignatureResponse{}, nil
}

func (s server) Create(_ context.Context, req *gen.CreateRequest) (*gen.CreateResponse, error) {
	if req.Email == "failcreate@test.com" {
		return nil, status.Error(codes.Internal, "error creating user")
	}

	return &gen.CreateResponse{
		Id: createdId,
	}, nil
}

func main() {
	l, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	s := grpc.NewServer()
	gen.RegisterUserServiceServer(s, server{})
	panic(s.Serve(l))
}
