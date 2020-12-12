package grpc

import (
	"context"
	"errors"
	"fmt"

	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/handler"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/model"

	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	server struct {
		service handler.Service
	}

	// InvalidParameterError is returned when a required parameter passed to New is invalid.
	InvalidParameterError struct {
		Parameter string
	}
)

func (i InvalidParameterError) Error() string {
	return fmt.Sprintf("invalid parameter %s", i.Parameter)
}

func New(service handler.Service) (*server, error) {
	if service == nil {
		return nil, InvalidParameterError{Parameter: "service"}
	}

	return &server{
		service: service,
	}, nil
}

func (s *server) Register(g *grpc.Server) {
	gen.RegisterUserServiceServer(g, s)
}

func (s *server) GetAllUsers(ctx context.Context, _ *gen.GetAllUsersRequest) (*gen.GetAllUsersResponse, error) {
	users, err := s.service.GetAllUsers(ctx)
	if err != nil {
		log.Error(ctx, "error_getting_all_users", log.ErrorParam(err))
		return nil, status.Error(codes.Internal, "could not get users")
	}

	var usrs []*gen.User
	for _, u := range users {
		usrs = append(usrs, model.UserToGrpc(u))
	}

	return &gen.GetAllUsersResponse{
		Users: usrs,
	}, nil
}

func (s *server) GetUsersByIds(ctx context.Context, req *gen.GetUsersByIdsRequest) (*gen.GetUsersByIdsResponse, error) {
	users, err := s.service.GetAllUsersByIds(ctx, req.Ids)
	if err != nil {
		log.Error(ctx, "error_getting_all_users_by_ids", log.ErrorParam(err))
		return nil, status.Error(codes.Internal, "could not get users")
	}

	var usrs []*gen.User
	for _, u := range users {
		usrs = append(usrs, model.UserToGrpc(u))
	}

	return &gen.GetUsersByIdsResponse{
		Users: usrs,
	}, nil
}

func (s *server) GetOverallRank(ctx context.Context, req *gen.GetOverallRankRequest) (*gen.GetOverallRankResponse, error) {
	rank, err := s.service.GetOverallRank(ctx, req.Id)
	if err != nil {
		log.Error(ctx, "error_getting_overall_rank", log.ErrorParam(err))
		return nil, status.Error(codes.Internal, "could not get overall rank")
	}

	return &gen.GetOverallRankResponse{
		Rank: rank,
	}, nil
}

func (s *server) GetRankForGroup(ctx context.Context, req *gen.GetRankForGroupRequest) (*gen.GetRankForGroupResponse, error) {
	rank, err := s.service.GetRankForGroup(ctx, req.Id, req.Ids)
	if err != nil {
		log.Error(ctx, "error_getting_rank_for_group", log.ErrorParam(err))
		return nil, status.Error(codes.Internal, "could not get rank for group")
	}

	return &gen.GetRankForGroupResponse{
		Rank: rank,
	}, nil
}

func (s *server) GetUserCount(ctx context.Context, _ *gen.GetUserCountRequest) (*gen.GetUserCountResponse, error) {
	count, err := s.service.GetUserCount(ctx)
	if err != nil {
		log.Error(ctx, "error_getting_user_count", log.ErrorParam(err))
		return nil, status.Error(codes.Internal, "could not get user count")
	}

	return &gen.GetUserCountResponse{
		Count: count,
	}, nil
}

func (s *server) GetUserByEmail(ctx context.Context, req *gen.GetUserByEmailRequest) (*gen.GetUserByEmailResponse, error) {
	user, err := s.service.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, model.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		log.Error(ctx, "error_getting_user_by_email", log.ErrorParam(err))
		return nil, status.Error(codes.Internal, "could not get user by email")
	}

	return &gen.GetUserByEmailResponse{
		User: model.UserToGrpc(user),
	}, nil
}

func (s *server) UpdatePassword(ctx context.Context, req *gen.UpdatePasswordRequest) (*gen.UpdatePasswordResponse, error) {
	err := s.service.UpdatePassword(ctx, req.Id, req.Password)
	if err != nil {
		log.Error(ctx, "error_updating_password", log.ErrorParam(err))
		return nil, status.Error(codes.Internal, "could not update password")
	}

	return &gen.UpdatePasswordResponse{}, nil
}

func (s *server) UpdateSignature(ctx context.Context, req *gen.UpdateSignatureRequest) (*gen.UpdateSignatureResponse, error) {
	err := s.service.UpdateSignature(ctx, req.Id, req.Signature)
	if err != nil {
		log.Error(ctx, "error_updating_signature", log.ErrorParam(err))
		return nil, status.Error(codes.Internal, "could not update signature")
	}

	return &gen.UpdateSignatureResponse{}, nil
}

func (s *server) Create(ctx context.Context, req *gen.CreateRequest) (*gen.CreateResponse, error) {
	id, err := s.service.StoreUser(ctx, model.UserFromCreateReq(*req))
	if err != nil {
		log.Error(ctx, "error_creating_user", log.ErrorParam(err))
		return nil, status.Error(codes.Internal, "could not create user")
	}

	return &gen.CreateResponse{
		Id: id,
	}, nil
}

func (s *server) GetUser(ctx context.Context, req *gen.GetUserRequest) (*gen.GetUserResponse, error) {
	user, err := s.service.GetUserById(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, model.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		log.Error(ctx, "error_getting_user_by_id", log.ErrorParam(err))
		return nil, status.Error(codes.Internal, "could not get user")
	}

	return &gen.GetUserResponse{
		User: model.UserToGrpc(user),
	}, nil
}
