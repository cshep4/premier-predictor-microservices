//go:generate impl 's *server' io.ReadWriteCloser -output server.go

package grpc

import (
	"context"
	"errors"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/handler"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type server struct {
	service handler.Servicer
}

func New(service handler.Servicer) (*server, error) {
	if service == nil {
		return nil, errors.New("service_is_nil")
	}

	return &server{
		service: service,
	}, nil
}

func (s *server) Register(g *grpc.Server) {
	gen.RegisterPredictionServiceServer(g, s)
}

func (s *server) GetPrediction(ctx context.Context, req *gen.PredictionRequest) (*gen.Prediction, error) {
	prediction, err := s.service.GetPrediction(req.UserId, req.MatchId)
	if err != nil {
		if errors.Is(err, model.ErrPredictionNotFound) {
			return nil, status.Error(codes.NotFound, model.ErrPredictionNotFound.Error())
		}
		log.Printf("error_getting_prediction: %v", err)
		return nil, err
	}

	return common.PredictionToGrpc(prediction), nil
}

func (s *server) GetPredictionSummary(ctx context.Context, req *gen.IdRequest) (*gen.MatchPredictionSummary, error) {
	predictionSummary, err := s.service.GetMatchPredictionSummary(req.Id)
	if err != nil {
		log.Printf("error_getting_prediction_summary: %v", err)
		return nil, err
	}

	return common.MatchPredictionSummaryToGrpc(predictionSummary), nil
}
