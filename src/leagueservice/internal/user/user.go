package user

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"

	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"

	pb "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
)

type (
	TokenGenerator interface {
		Generate(ctx context.Context, service string) (string, error)
	}

	service struct {
		generator TokenGenerator
		client    pb.UserServiceClient
	}

	// InvalidParameterError is returned when a required parameter passed to New is invalid.
	InvalidParameterError struct {
		Parameter string
	}
)

func (i InvalidParameterError) Error() string {
	return fmt.Sprintf("invalid parameter %s", i.Parameter)
}

func New(generator TokenGenerator, client pb.UserServiceClient) (*service, error) {
	switch {
	case generator == nil:
		return nil, InvalidParameterError{Parameter: "generator"}
	case client == nil:
		return nil, InvalidParameterError{Parameter: "client"}
	}

	return &service{
		generator: generator,
		client:    client,
	}, nil
}

func (s *service) GetAllUsers(ctx context.Context) ([]model.LeagueUser, error) {
	ctx, err := s.authorizeContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("authorize_context: %w", err)
	}

	resp, err := s.client.GetAllUsers(ctx, &pb.GetAllUsersRequest{})
	if err != nil {
		return nil, fmt.Errorf("get_all_users: %w", err)
	}

	var users []model.LeagueUser
	for _, u := range resp.Users {
		users = append(users, model.LeagueUserFromGrpc(u))
	}

	return users, nil
}

func (s *service) GetLeagueUsers(ctx context.Context, ids []string) ([]model.LeagueUser, error) {
	ctx, err := s.authorizeContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("authorize_context: %w", err)
	}

	resp, err := s.client.GetUsersByIds(ctx, &pb.GetUsersByIdsRequest{
		Ids: ids,
	})
	if err != nil {
		return nil, fmt.Errorf("get_users_by_ids: %w", err)
	}

	var users []model.LeagueUser
	for _, u := range resp.Users {
		users = append(users, model.LeagueUserFromGrpc(u))
	}

	return users, nil
}

func (s *service) GetOverallRank(ctx context.Context, id string) (int64, error) {
	ctx, err := s.authorizeContext(ctx)
	if err != nil {
		return 0, fmt.Errorf("authorize_context: %w", err)
	}

	resp, err := s.client.GetOverallRank(ctx, &pb.GetOverallRankRequest{
		Id: id,
	})
	if err != nil {
		return 0, fmt.Errorf("get_overall_rank: %w", err)
	}

	return resp.Rank, nil
}

func (s *service) GetLeagueRank(ctx context.Context, id string, ids []string) (int64, error) {
	ctx, err := s.authorizeContext(ctx)
	if err != nil {
		return 0, fmt.Errorf("authorize_context: %w", err)
	}

	resp, err := s.client.GetRankForGroup(ctx, &pb.GetRankForGroupRequest{
		Id:  id,
		Ids: ids,
	})
	if err != nil {
		return 0, fmt.Errorf("get_rank_for_group: %w", err)
	}

	return resp.Rank, nil
}

func (s *service) GetUserCount(ctx context.Context) (int64, error) {
	ctx, err := s.authorizeContext(ctx)
	if err != nil {
		return 0, fmt.Errorf("authorize_context: %w", err)
	}

	resp, err := s.client.GetUserCount(ctx, &pb.GetUserCountRequest{})
	if err != nil {
		return 0, fmt.Errorf("get_user_count: %w", err)
	}

	return resp.Count, nil
}

func (s *service) authorizeContext(ctx context.Context) (context.Context, error) {
	token, err := s.generator.Generate(ctx, "user")
	if err != nil {
		return nil, fmt.Errorf("generate: %w", err)
	}

	return metadata.AppendToOutgoingContext(ctx, "token", token), nil
}
