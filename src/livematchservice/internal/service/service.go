package live

import (
	"context"
	"errors"
	"github.com/ahl5esoft/golang-underscore"
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/model"
	predictor "github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/prediction"
	"sort"
	"time"
)

type (
	Predictor interface {
		GetPrediction(ctx context.Context, req model.PredictionRequest) (*common.Prediction, error)
		GetPredictionSummary(ctx context.Context, matchId string) (*common.MatchPredictionSummary, error)
	}
	Store interface {
		GetUpcomingMatches() ([]common.MatchFacts, error)
		GetMatchFacts(id string) (*common.MatchFacts, error)
	}

	predictionResult struct {
		result *common.Prediction
		err    error
	}

	matchPredictionSummaryResult struct {
		result *common.MatchPredictionSummary
		err    error
	}

	matchFactsResult struct {
		result *common.MatchFacts
		err    error
	}

	service struct {
		store     Store
		predictor Predictor
	}
)

func New(store Store, predictor Predictor) (*service, error) {
	if store == nil {
		return nil, errors.New("store_is_nil")
	}
	if predictor == nil {
		return nil, errors.New("predictor_is_nil")
	}

	return &service{
		store:     store,
		predictor: predictor,
	}, nil
}

func (s *service) GetMatchSummary(ctx context.Context, req model.PredictionRequest) (*model.MatchSummary, error) {
	predictionChan := make(chan predictionResult)
	matchPredictionSummaryChan := make(chan matchPredictionSummaryResult)
	matchFactsChan := make(chan matchFactsResult)

	go func() {
		res, err := s.predictor.GetPrediction(ctx, req)
		predictionChan <- predictionResult{result: res, err: err}
	}()

	go func() {
		res, err := s.predictor.GetPredictionSummary(ctx, req.MatchId)
		matchPredictionSummaryChan <- matchPredictionSummaryResult{result: res, err: err}
	}()

	go func() {
		res, err := s.store.GetMatchFacts(req.MatchId)
		matchFactsChan <- matchFactsResult{result: res, err: err}
	}()

	prediction := <-predictionChan
	if prediction.err != nil && prediction.err != predictor.ErrPredictionNotFound {
		return nil, prediction.err
	}

	matchPredictionSummary := <-matchPredictionSummaryChan
	if matchPredictionSummary.err != nil {
		return nil, matchPredictionSummary.err
	}

	matchFacts := <-matchFactsChan
	if matchFacts.err != nil {
		return nil, matchFacts.err
	}

	return &model.MatchSummary{
		Match:             matchFacts.result,
		PredictionSummary: matchPredictionSummary.result,
		Prediction:        prediction.result,
	}, nil
}

func (s *service) GetMatchFacts(id string) (*common.MatchFacts, error) {
	return s.store.GetMatchFacts(id)
}

func (s *service) GetUpcomingMatches() (map[time.Time][]common.MatchFacts, error) {
	matches, err := s.store.GetUpcomingMatches()
	if err != nil {
		return nil, err
	}

	sort.Sort(common.MatchFactsSlice(matches))

	upcomingMatches := make(map[time.Time][]common.MatchFacts)

	underscore.Chain(matches).
		Group(s.groupByMatchDate).
		Value(&upcomingMatches)

	if err := recover(); err != nil {
		return nil, errors.New("could not map matches")
	}

	return upcomingMatches, nil
}

func (s *service) groupByMatchDate(m common.MatchFacts, _ int) time.Time {
	return m.MatchDate
}
