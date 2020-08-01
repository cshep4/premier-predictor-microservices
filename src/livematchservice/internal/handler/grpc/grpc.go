package grpc

import (
	"errors"
	"time"

	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/handler"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/model"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type server struct {
	service  handler.Servicer
	interval time.Duration
}

func New(service handler.Servicer, interval time.Duration) (*server, error) {
	if service == nil {
		return nil, errors.New("service_is_nil")
	}
	if interval == 0 {
		return nil, errors.New("interval_is_zero")
	}

	return &server{
		service:  service,
		interval: interval,
	}, nil
}

func (s *server) Register(g *grpc.Server) {
	gen.RegisterLiveMatchServiceServer(g, s)
}

func (s *server) GetUpcomingMatches(_ *empty.Empty, stream gen.LiveMatchService_GetUpcomingMatchesServer) error {
	matches, err := s.service.GetUpcomingMatches(stream.Context())
	if err != nil {
		return err
	}

	resp := model.ToUpcomingMatchesResponse(matches)

	if err := stream.Send(resp); err != nil {
		return err
	}

	ticker := time.NewTicker(s.interval)
	for {
		select {
		case <-ticker.C:
			matches, err := s.service.GetUpcomingMatches(stream.Context())
			if err != nil {
				return nil
			}

			resp := model.ToUpcomingMatchesResponse(matches)

			if err := stream.Send(resp); err != nil {
				return nil
			}
		}
	}
}

func (s *server) GetMatchSummary(req *gen.PredictionRequest, stream gen.LiveMatchService_GetMatchSummaryServer) error {
	r := model.PredictionRequest{
		UserId:  req.UserId,
		MatchId: req.MatchId,
	}

	matchSummary, err := s.service.GetMatchSummary(stream.Context(), r)
	if err != nil {
		return err
	}

	resp := model.MatchSummaryToGrpc(matchSummary)

	if err := stream.Send(resp); err != nil {
		return err
	}

	ticker := time.NewTicker(s.interval)
	for {
		select {
		case <-ticker.C:
			match, err := s.service.GetMatchFacts(stream.Context(), req.MatchId)
			if err != nil {
				return nil
			}

			resp.Match = common.MatchFactsToGrpc(match)

			if err := stream.Send(resp); err != nil {
				return nil
			}
		}
	}
}
