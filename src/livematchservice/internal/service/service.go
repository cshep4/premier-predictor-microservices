package live

import (
	"context"
	"errors"
	"fmt"
	"github.com/ahl5esoft/golang-underscore"
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/model"
	predictor "github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/prediction"
	"golang.org/x/sync/errgroup"
	"sort"
	"time"
)

type (
	Predictor interface {
		GetPrediction(ctx context.Context, req model.PredictionRequest) (*common.Prediction, error)
		GetPredictionSummary(ctx context.Context, matchId string) (*common.MatchPredictionSummary, error)
	}
	Store interface {
		GetUpcomingMatches(ctx context.Context) ([]common.MatchFacts, error)
		GetTodaysMatches(ctx context.Context) ([]common.MatchFacts, error)
		GetMatchFacts(ctx context.Context, id string) (*common.MatchFacts, error)
		SubscribeToMatch(ctx context.Context, id string, observer model.MatchObserver) error
		SubscribeToMatches(ctx context.Context, ids []string, observer model.MatchObserver) error
	}

	service struct {
		store     Store
		predictor Predictor
	}
)

func New(store Store, predictor Predictor) (*service, error) {
	if store == nil {
		return nil, errors.New("store is nil")
	}
	if predictor == nil {
		return nil, errors.New("predictor is nil")
	}

	return &service{
		store:     store,
		predictor: predictor,
	}, nil
}

func (s *service) GetMatchSummary(ctx context.Context, req model.PredictionRequest) (*model.MatchSummary, error) {
	predictionChan := make(chan *common.Prediction)
	matchPredictionSummaryChan := make(chan *common.MatchPredictionSummary)
	matchFactsChan := make(chan *common.MatchFacts)

	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		p, err := s.predictor.GetPrediction(ctx, req)
		predictionChan <- p
		if err != nil {
			return fmt.Errorf("get prediction: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		p, err := s.predictor.GetPredictionSummary(ctx, req.MatchId)
		matchPredictionSummaryChan <- p
		if err != nil {
			return fmt.Errorf("get prediction summary: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		m, err := s.store.GetMatchFacts(ctx, req.MatchId)
		matchFactsChan <- m
		if err != nil {
			return fmt.Errorf("get match facts: %w", err)
		}
		return nil
	})

	prediction := <-predictionChan
	matchPredictionSummary := <-matchPredictionSummaryChan
	matchFacts := <-matchFactsChan

	err := g.Wait()
	if err != nil && !errors.Is(err, predictor.ErrPredictionNotFound) {
		return nil, err
	}

	return &model.MatchSummary{
		Match:             matchFacts,
		PredictionSummary: matchPredictionSummary,
		Prediction:        prediction,
	}, nil
}

func (s *service) GetMatchFacts(ctx context.Context, id string) (*common.MatchFacts, error) {
	matchFacts, err := s.store.GetMatchFacts(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get match facts: %w", err)
	}

	return matchFacts, nil
}

func (s *service) GetUpcomingMatches(ctx context.Context) (map[time.Time][]common.MatchFacts, error) {
	matches, err := s.store.GetUpcomingMatches(ctx)
	if err != nil {
		return nil, fmt.Errorf("get upcoming matches: %w", err)
	}

	sort.Sort(common.MatchFactsSlice(matches))

	upcomingMatches := make(map[time.Time][]common.MatchFacts)

	underscore.Chain(matches).
		Group(func(m common.MatchFacts, _ int) time.Time { return m.MatchDate }).
		Value(&upcomingMatches)

	if err := recover(); err != nil {
		return nil, fmt.Errorf("group matches: %s", err)
	}

	return upcomingMatches, nil
}

func (s *service) GetTodaysMatches(ctx context.Context) ([]common.MatchFacts, error) {
	matches, err := s.store.GetTodaysMatches(ctx)
	if err != nil {
		return nil, fmt.Errorf("get todays matches: %w", err)
	}

	return matches, nil
}

func (s *service) SubscribeToMatch(ctx context.Context, id string, observer model.MatchObserver) error {
	err := s.store.SubscribeToMatch(ctx, id, observer)
	if err != nil {
		return fmt.Errorf("subscribe_to_match_%s: %w", id, err)
	}

	return nil
}

func (s *service) SubscribeToTodaysMatches(ctx context.Context, observer model.MatchObserver) error {
	matches, err := s.store.GetTodaysMatches(ctx)
	if err != nil {
		return fmt.Errorf("get_todays_matches: %w", err)
	}

	observable := model.MatchObservable{}
	observable.AddObserver(observer)

	var ids []string
	for _, m := range matches {
		if err := observable.Notify(&m); err != nil {
			return fmt.Errorf("observable_notify_for_match_%s: %w", m.Id, err)
		}
		ids = append(ids, m.Id)
	}

	err = s.store.SubscribeToMatches(ctx, ids, observer)
	if err != nil {
		return fmt.Errorf("subscribe_to_matches_%v: %w", ids, err)
	}

	return nil
}
