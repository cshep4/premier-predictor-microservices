package handler

import (
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/model"
)

type Servicer interface {
	GetFixturesWithPredictions(id string) ([]model.FixturePrediction, error)
	GetPredictorData(id string) (*model.PredictorData, error)
	GetUsersPastPredictions(id string) (*model.PredictionSummary, error)
	UpdatePredictions(predictions []common.Prediction) error
	GetPrediction(userId, matchId string) (*common.Prediction, error)
	GetMatchPredictionSummary(id string) (*common.MatchPredictionSummary, error)
}
