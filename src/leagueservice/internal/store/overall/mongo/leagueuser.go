package mongo

import (
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type leagueUserEntity struct {
	ID              primitive.ObjectID `bson:"_id"`
	Name            string             `bson:"name"`
	PredictedWinner string             `bson:"predictedWinner"`
	Rank            int64              `bson:"rank"`
	Score           int                `bson:"score"`
}

func (lu leagueUserEntity) toLeagueUser() model.LeagueUser {
	return model.LeagueUser{
		ID:              lu.ID.Hex(),
		Name:            lu.Name,
		PredictedWinner: lu.PredictedWinner,
		Rank:            lu.Rank,
		Score:           lu.Score,
	}
}
