package grpc

import (
	"context"
	"errors"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/log"
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/handler"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	service handler.Servicer
}

func New(service handler.Servicer) (*server, error) {
	if service == nil {
		return nil, errors.New("service is nil")
	}

	return &server{
		service: service,
	}, nil
}

func (s *server) Register(g *grpc.Server) {
	gen.RegisterPredictionServiceServer(g, s)
}

func (s *server) GetPrediction(ctx context.Context, req *gen.PredictionRequest) (*gen.Prediction, error) {
	prediction, err := s.service.GetPrediction(ctx, req.UserId, req.MatchId)
	if err != nil {
		if errors.Is(err, model.ErrPredictionNotFound) {
			return nil, status.Error(codes.NotFound, "prediction not found")
		}
		log.Error(ctx, "error_getting_prediction", log.ErrorParam(err))
		return nil, err
	}

	return common.PredictionToGrpc(prediction), nil
}

func (s *server) GetUserPredictions(ctx context.Context, req *gen.GetUserPredictionsRequest) (*gen.GetUserPredictionsResponse, error) {
	predictions, err := s.service.GetUsersPredictions(ctx, req.GetUserId())
	if err != nil {
		log.Error(ctx, "error_getting_prediction", log.ErrorParam(err))
		return nil, status.Error(codes.Internal, "could not get predictions")
	}

	preds := make([]*gen.Prediction, 0, len(predictions))
	for _, p := range predictions {
		preds = append(preds, common.PredictionToGrpc(&p))
	}

	return &gen.GetUserPredictionsResponse{
		Predictions: preds,
	}, nil
}

func (s *server) GetPredictionSummary(ctx context.Context, req *gen.IdRequest) (*gen.MatchPredictionSummary, error) {
	predictionSummary, err := s.service.GetMatchPredictionSummary(ctx, req.Id)
	if err != nil {
		log.Error(ctx, "error_getting_prediction_summary", log.ErrorParam(err))
		return nil, err
	}

	return common.MatchPredictionSummaryToGrpc(predictionSummary), nil
}
