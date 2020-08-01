package prediction

import (
	"context"
	"errors"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/auth"
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type predictor struct {
	client gen.PredictionServiceClient
}

var (
	ErrPredictionNotFound = errors.New("prediction not found")
)

func New(client gen.PredictionServiceClient) (*predictor, error) {
	return &predictor{
		client: client,
	}, nil
}

func (p *predictor) GetPrediction(ctx context.Context, req model.PredictionRequest) (*common.Prediction, error) {
	metadata, err := auth.MetadataFromContext(ctx)
	if err != nil {
		return nil, err
	}

	prediction, err := p.client.GetPrediction(metadata, &gen.PredictionRequest{
		UserId:  req.UserId,
		MatchId: req.MatchId,
	})
	if err != nil {
		statusErr, ok := status.FromError(err)

		switch {
		case !ok:
			return nil, err
		case statusErr.Code() == codes.NotFound:
			return nil, ErrPredictionNotFound
		}

		return nil, err
	}

	return common.PredictionFromGrpc(prediction), nil
}

func (p *predictor) GetPredictionSummary(ctx context.Context, matchId string) (*common.MatchPredictionSummary, error) {
	metadata, err := auth.MetadataFromContext(ctx)
	if err != nil {
		return nil, err
	}

	predictionSummary, err := p.client.GetPredictionSummary(metadata, &gen.IdRequest{
		Id: matchId,
	})
	if err != nil {
		return nil, err
	}

	return common.MatchPredictionSummaryFromGrpc(predictionSummary), nil
}
