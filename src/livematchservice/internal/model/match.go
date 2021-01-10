package model

import (
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/model"
	_ "github.com/golang/protobuf/proto"
	"time"
)

type (
	MatchObserver interface {
		Update(matchFacts *model.MatchFacts) error
	}
	MatchObservable struct {
		observers []MatchObserver
	}
	MatchSummary struct {
		Match             *model.MatchFacts             `json:"match"`
		PredictionSummary *model.MatchPredictionSummary `json:"predictionSummary"`
		Prediction        *model.Prediction             `json:"prediction"`
	}
)

func (m *MatchObservable) AddObserver(observer MatchObserver) {
	m.observers = append(m.observers, observer)
}

func (m *MatchObservable) Notify(matchFacts *model.MatchFacts) error {
	for _, o := range m.observers {
		if err := o.Update(matchFacts); err != nil {
			return err
		}
	}

	return nil
}

func ToUpcomingMatchesResponse(upcomingMatches map[time.Time][]model.MatchFacts) *gen.UpcomingMatchesResponse {
	matches := make(map[string]*gen.MatchFactsList)

	for k, v := range upcomingMatches {
		date := k.Format("02-01-2006")
		matches[date] = &gen.MatchFactsList{}

		var matchFacts []*gen.MatchFacts
		for i := range v {
			matchFacts = append(matchFacts, model.MatchFactsToGrpc(&v[i]))
		}

		matches[date].Matches = matchFacts
	}

	return &gen.UpcomingMatchesResponse{
		Matches: matches,
	}
}
