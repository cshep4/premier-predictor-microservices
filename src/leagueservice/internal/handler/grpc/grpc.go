package grpc

import (
	"context"

	pb "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"
)

type (
	Service interface {
		GetUsersLeagueList(ctx context.Context, id string) (*model.StandingsOverview, error)
	}

	server struct {
		service Service
	}
)

func New(service Service) (*server, error) {
	if service == nil {
		return nil, model.InvalidParameterError{Parameter: "service"}
	}

	return &server{
		service: service,
	}, nil
}

func (s *server) Register(g *grpc.Server) {
	pb.RegisterLeagueServiceServer(g, s)
}

func (s *server) ListLeagues(ctx context.Context, req *pb.ListLeaguesRequest) (*pb.ListLeaguesResponse, error) {
	if req.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid userId")
	}

	overview, err := s.service.GetUsersLeagueList(ctx, req.GetId())
	if err != nil {
		log.Error(ctx, "error_getting_user_leagues", log.ErrorParam(err))
		return nil, status.Error(codes.Internal, "could not get user's leagues")
	}

	return &pb.ListLeaguesResponse{
		Leagues: s.toLeagueSummaryList(overview.UserLeagues),
	}, nil
}

func (s *server) GetOverview(ctx context.Context, req *pb.GetOverviewRequest) (*pb.GetOverviewResponse, error) {
	if req.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid userId")
	}

	overview, err := s.service.GetUsersLeagueList(ctx, req.GetId())
	if err != nil {
		log.Error(ctx, "error_getting_user_leagues", log.ErrorParam(err))
		return nil, status.Error(codes.Internal, "could not get user's leagues")
	}

	return &pb.GetOverviewResponse{
		Rank:      overview.OverallLeagueOverview.Rank,
		UserCount: overview.OverallLeagueOverview.UserCount,
		Leagues:   s.toLeagueSummaryList(overview.UserLeagues),
	}, nil
}

func (s *server) toLeagueSummaryList(leagues []model.LeagueOverview) []*pb.LeagueSummary {
	var ls []*pb.LeagueSummary
	for _, l := range leagues {
		ls = append(ls, &pb.LeagueSummary{
			LeagueName: l.LeagueName,
			Pin:        l.Pin,
			Rank:       l.Rank,
		})
	}
	return ls
}
