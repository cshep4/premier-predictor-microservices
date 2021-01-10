package grpc

import (
	"context"
	"errors"
	"time"

	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/handler"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/model"

	pb "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/log"
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	server struct {
		service  handler.Servicer
		interval time.Duration
	}
	observer struct {
		update func(matchFacts *common.MatchFacts) error
	}
)

func (o observer) Update(matchFacts *common.MatchFacts) error {
	return o.update(matchFacts)
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
	pb.RegisterLiveMatchServiceServer(g, s)
}

func (s *server) GetUpcomingMatches(_ *empty.Empty, stream pb.LiveMatchService_GetUpcomingMatchesServer) error {
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

func (s *server) GetMatchSummary(req *pb.PredictionRequest, stream pb.LiveMatchService_GetMatchSummaryServer) error {
	matchFacts, err := s.service.GetMatchFacts(stream.Context(), req.GetMatchId())
	if err != nil {
		log.Error(stream.Context(), "error_getting_match_summary",
			log.ErrorParam(err),
			log.SafeParam("matchId", req.GetMatchId()),
			//log.SafeParam("userId", req.GetUserId()),
		)
		return status.Error(codes.Internal, "could not get match summary")
	}

	if err := stream.Send(&pb.MatchSummary{Match: common.MatchFactsToGrpc(matchFacts)}); err != nil {
		log.Error(stream.Context(), "error_sending_response",
			log.ErrorParam(err),
			log.SafeParam("matchId", req.GetMatchId()),
			//log.SafeParam("userId", req.GetUserId()),
		)
		return status.Error(codes.Internal, "could not send response")
	}

	obvs := observer{update: func(matchFacts *common.MatchFacts) error {
		if err := stream.Send(&pb.MatchSummary{Match: common.MatchFactsToGrpc(matchFacts)}); err != nil {
			log.Error(stream.Context(), "error_sending_response",
				log.ErrorParam(err),
				log.SafeParam("matchId", req.GetMatchId()),
				//log.SafeParam("userId", req.GetUserId()),
			)
			return status.Error(codes.Internal, "could not send response")
		}

		return nil
	}}

	err = s.service.SubscribeToMatch(stream.Context(), req.GetMatchId(), obvs)
	if err != nil {
		log.Error(stream.Context(), "error_subscribing_to_live_match",
			log.ErrorParam(err),
			log.SafeParam("matchId", req.GetMatchId()),
			//log.SafeParam("userId", req.GetUserId()),
		)
		return status.Error(codes.Internal, "could not subscribe to live match")
	}

	return nil
}

func (s *server) GetTodaysLiveMatches(_ *pb.GetTodaysLiveMatchesRequest, stream pb.LiveMatchService_GetTodaysLiveMatchesServer) error {
	obvs := observer{update: func(matchFacts *common.MatchFacts) error {
		if err := stream.Send(&pb.GetTodaysLiveMatchesResponse{Match: common.MatchFactsToGrpc(matchFacts)}); err != nil {
			log.Error(stream.Context(), "error_sending_response", log.ErrorParam(err))
			return status.Error(codes.Internal, "could not send response")
		}

		return nil
	}}

	err := s.service.SubscribeToTodaysMatches(stream.Context(), obvs)
	if err != nil {
		log.Error(stream.Context(), "error_subscribing_to_todays_matches", log.ErrorParam(err))
		return status.Error(codes.Internal, "could not subscribe to todays matches")
	}

	return nil
}

func (s *server) GetLiveMatch(ctx context.Context, req *pb.GetLiveMatchRequest) (*pb.GetLiveMatchResponse, error) {
	match, err := s.service.GetMatchFacts(ctx, req.GetId())
	if err != nil {
		log.Error(ctx, "error_getting_live_match", log.ErrorParam(err))
		return nil, status.Error(codes.Internal, "could not get live match")
	}

	return &pb.GetLiveMatchResponse{
		Match: common.MatchFactsToGrpc(match),
	}, nil
}

func (s *server) ListTodaysMatches(ctx context.Context, _ *pb.ListTodaysMatchesRequest) (*pb.ListTodaysMatchesResponse, error) {
	matches, err := s.service.GetTodaysMatches(ctx)
	if err != nil {
		log.Error(ctx, "error_getting_todays_matches", log.ErrorParam(err))
		return nil, status.Error(codes.Internal, "could not list today's matches")
	}

	var matchFacts []*pb.MatchFacts
	for _, m := range matches {
		matchFacts = append(matchFacts, common.MatchFactsToGrpc(&m))
	}

	return &pb.ListTodaysMatchesResponse{
		Matches: matchFacts,
	}, nil
}
