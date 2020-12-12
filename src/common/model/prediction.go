package model

import "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"

type MatchPredictionSummary struct {
	HomeWin int `json:"homeWin,omitempty"`
	Draw    int `json:"draw,omitempty"`
	AwayWin int `json:"awayWin,omitempty"`
}

func MatchPredictionSummaryFromGrpc(predictionSummary *model.MatchPredictionSummary) *MatchPredictionSummary {
	if predictionSummary == nil {
		return &MatchPredictionSummary{}
	}
	return &MatchPredictionSummary{
		HomeWin: int(predictionSummary.HomeWin),
		Draw:    int(predictionSummary.Draw),
		AwayWin: int(predictionSummary.AwayWin),
	}
}

func MatchPredictionSummaryToGrpc(predictionSummary *MatchPredictionSummary) *model.MatchPredictionSummary {
	if predictionSummary == nil {
		return &model.MatchPredictionSummary{}
	}
	return &model.MatchPredictionSummary{
		HomeWin: int32(predictionSummary.HomeWin),
		Draw:    int32(predictionSummary.Draw),
		AwayWin: int32(predictionSummary.AwayWin),
	}
}

type Prediction struct {
	UserId    string `json:"userId"`
	MatchId   string `json:"matchId"`
	HomeGoals int    `json:"hGoals"`
	AwayGoals int    `json:"aGoals"`
}

func PredictionFromGrpc(prediction *model.Prediction) *Prediction {
	if prediction == nil {
		return nil
	}

	return &Prediction{
		UserId:    prediction.UserId,
		MatchId:   prediction.MatchId,
		HomeGoals: int(prediction.HGoals),
		AwayGoals: int(prediction.AGoals),
	}
}

func PredictionToGrpc(prediction *Prediction) *model.Prediction {
	if prediction == nil {
		return nil
	}

	return &model.Prediction{
		UserId:  prediction.UserId,
		MatchId: prediction.MatchId,
		HGoals:  int32(prediction.HomeGoals),
		AGoals:  int32(prediction.AwayGoals),
	}
}
