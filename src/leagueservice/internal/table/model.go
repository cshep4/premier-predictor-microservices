package table

import (
	"fmt"

	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"
)

type (
	leagueUserSlice []model.LeagueUser

	Ranker interface {
		Insert(val int)
		Rank(val int) (int, bool)
		Clear()
	}

	// InvalidParameterError is returned when a required parameter passed to New is invalid.
	InvalidParameterError struct {
		Parameter string
	}
)

func (i InvalidParameterError) Error() string {
	return fmt.Sprintf("invalid parameter %s", i.Parameter)
}

func (l leagueUserSlice) Len() int {
	return len(l)
}

func (l leagueUserSlice) Less(i, j int) bool {
	return l[i].Score > l[j].Score
}

func (l leagueUserSlice) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
