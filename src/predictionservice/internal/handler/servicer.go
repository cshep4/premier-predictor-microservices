package handler

import (
	"context"
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/model"
)

type Servicer interface {
	GetFixturesWithPredictions(ctx context.Context, id string) ([]model.FixturePrediction, error)
	GetPredictorData(ctx context.Context, id string) (*model.PredictorData, error)
	GetUsersPastPredictions(ctx context.Context, id string) (*model.PredictionSummary, error)
	UpdatePredictions(ctx context.Context, predictions []common.Prediction) error
	GetPrediction(ctx context.Context, userId, matchId string) (*common.Prediction, error)
	GetMatchPredictionSummary(ctx context.Context, id string) (*common.MatchPredictionSummary, error)
}
