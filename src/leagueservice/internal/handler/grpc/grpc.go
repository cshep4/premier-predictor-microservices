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
		JoinUserLeague(ctx context.Context, id string, pin int64) (*model.LeagueOverview, error)
		AddUserLeague(ctx context.Context, id, name string) (*model.League, error)
		LeaveUserLeague(ctx context.Context, id string, pin int64) error
		RenameUserLeague(ctx context.Context, pin int64, name string) error
		GetLeagueTable(ctx context.Context, pin int64) ([]model.LeagueUser, error)
		GetOverallLeagueTable(ctx context.Context) ([]model.LeagueUser, error)
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
		log.Error(ctx, "error getting user leagues", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not get user's leagues")
	}

	return &pb.ListLeaguesResponse{
		Leagues: s.leagueSummaryList(overview.UserLeagues),
	}, nil
}

func (s *server) GetOverview(ctx context.Context, req *pb.GetOverviewRequest) (*pb.GetOverviewResponse, error) {
	if req.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid userId")
	}

	overview, err := s.service.GetUsersLeagueList(ctx, req.GetId())
	if err != nil {
		log.Error(ctx, "error getting user leagues", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not get user's leagues")
	}

	return &pb.GetOverviewResponse{
		Rank:      overview.OverallLeagueOverview.Rank,
		Score:     int32(overview.OverallLeagueOverview.Score),
		UserCount: overview.OverallLeagueOverview.UserCount,
		Leagues:   s.leagueSummaryList(overview.UserLeagues),
	}, nil
}

func (s *server) AddLeague(ctx context.Context, req *pb.AddLeagueRequest) (*pb.AddLeagueResponse, error) {
	switch {
	case req.GetUserId() == "":
		return nil, status.Error(codes.InvalidArgument, "invalid userId")
	case req.GetName() == "":
		return nil, status.Error(codes.InvalidArgument, "invalid name")
	}

	league, err := s.service.AddUserLeague(ctx, req.GetUserId(), req.GetName())
	if err != nil {
		log.Error(ctx, "error adding league",
			zap.Error(err),
			zap.String("user_id", req.GetUserId()),
			zap.String("name", req.GetName()),
		)

		return nil, s.statusError(err, "could not add league")
	}

	return &pb.AddLeagueResponse{
		Pin:   league.Pin,
		Name:  league.Name,
		Users: league.Users,
	}, nil
}

func (s *server) JoinLeague(ctx context.Context, req *pb.JoinLeagueRequest) (*pb.JoinLeagueResponse, error) {
	switch {
	case req.GetUserId() == "":
		return nil, status.Error(codes.InvalidArgument, "invalid userId")
	case req.GetPin() <= 0:
		return nil, status.Errorf(codes.InvalidArgument, "invalid pin: %d", req.GetPin())
	}

	league, err := s.service.JoinUserLeague(ctx, req.GetUserId(), req.GetPin())
	if err != nil {
		log.Error(ctx, "error joining league",
			zap.Error(err),
			zap.Int64("pin", req.GetPin()),
			zap.String("user_id", req.GetUserId()),
		)

		return nil, s.statusError(err, "could not join league")
	}

	return &pb.JoinLeagueResponse{
		Pin:  league.Pin,
		Name: league.LeagueName,
		Rank: league.Rank,
	}, nil
}

func (s *server) LeaveLeague(ctx context.Context, req *pb.LeaveLeagueRequest) (*pb.LeaveLeagueResponse, error) {
	switch {
	case req.GetUserId() == "":
		return nil, status.Error(codes.InvalidArgument, "invalid userId")
	case req.GetPin() <= 0:
		return nil, status.Errorf(codes.InvalidArgument, "invalid pin: %d", req.GetPin())
	}

	if err := s.service.LeaveUserLeague(ctx, req.GetUserId(), req.GetPin()); err != nil {
		log.Error(ctx, "error leaving league",
			zap.Error(err),
			zap.Int64("pin", req.GetPin()),
			zap.String("user_id", req.GetUserId()),
		)

		return nil, s.statusError(err, "could not leave league")
	}

	return &pb.LeaveLeagueResponse{}, nil
}

func (s *server) RenameLeague(ctx context.Context, req *pb.RenameLeagueRequest) (*pb.RenameLeagueResponse, error) {
	switch {
	case req.GetPin() <= 0:
		return nil, status.Errorf(codes.InvalidArgument, "invalid pin: %d", req.GetPin())
	case req.GetName() == "":
		return nil, status.Error(codes.InvalidArgument, "invalid name")
	}

	err := s.service.RenameUserLeague(ctx, req.GetPin(), req.GetName())
	if err != nil {
		log.Error(ctx, "error renaming league",
			zap.Error(err),
			zap.Int64("pin", req.GetPin()),
			zap.String("name", req.GetName()),
		)

		return nil, s.statusError(err, "could not rename league")
	}

	return &pb.RenameLeagueResponse{}, nil
}

func (s *server) GetLeagueTable(ctx context.Context, req *pb.GetLeagueTableRequest) (*pb.GetLeagueTableResponse, error) {
	if req.GetPin() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid pin: %d", req.GetPin())
	}

	users, err := s.service.GetLeagueTable(ctx, req.GetPin())
	if err != nil {
		log.Error(ctx, "error getting league table",
			zap.Error(err),
			zap.Int64("pin", req.GetPin()),
		)

		return nil, s.statusError(err, "could not get league table")
	}

	return &pb.GetLeagueTableResponse{
		Users: s.leagueUsers(users),
	}, nil
}

func (s *server) GetOverallLeagueTable(ctx context.Context, _ *pb.GetOverallLeagueTableRequest) (*pb.GetOverallLeagueTableResponse, error) {
	users, err := s.service.GetOverallLeagueTable(ctx)
	if err != nil {
		log.Error(ctx, "error getting overall league table", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not get overall league table")
	}

	return &pb.GetOverallLeagueTableResponse{
		Users: s.leagueUsers(users),
	}, nil
}

func (s *server) leagueUsers(users []model.LeagueUser) []*pb.LeagueUser {
	leagueUsers := make([]*pb.LeagueUser, len(users))
	for i, u := range users {
		leagueUsers[i] = &pb.LeagueUser{
			Id:              u.ID,
			FirstName:       u.FirstName,
			Surname:         u.Surname,
			PredictedWinner: u.PredictedWinner,
			Rank:            u.Rank,
			Score:           int32(u.Score),
			CreatedAt:       timestamppb.New(u.CreatedAt),
			UpdatedAt:       timestamppb.New(u.UpdatedAt),
		}
	}
	return leagueUsers
}

func (s *server) leagueSummaryList(leagues []model.LeagueOverview) []*pb.LeagueSummary {
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

func (s *server) statusError(err error, genericMessage string) error {
	switch {
	case errors.As(err, &model.InvalidParameterError{}):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, model.ErrLeagueNotFound):
		return status.Error(codes.NotFound, "league not found")
	}

	return status.Error(codes.Internal, genericMessage)
}
