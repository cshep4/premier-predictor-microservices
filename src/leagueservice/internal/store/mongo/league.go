package mongo

import (
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"
)

type leagueEntity struct {
	Pin   int64    `bson:"_id"`
	Name  string   `bson:"name"`
	Users []string `bson:"users"`
}

func fromLeague(league model.League) *leagueEntity {
	return &leagueEntity{
		Pin:   league.Pin,
		Name:  league.Name,
		Users: league.Users,
	}
}

func toLeague(league leagueEntity) *model.League {
	return &model.League{
		Pin:   league.Pin,
		Name:  league.Name,
		Users: league.Users,
	}
}
