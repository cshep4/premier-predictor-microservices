package prediction

import (
	"context"
	"errors"
	"fmt"
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/model"
	"golang.org/x/sync/errgroup"
	"time"
)

type (
	FixtureService interface {
		GetMatches(ctx context.Context) ([]common.Fixture, error)
		GetTeamForm(ctx context.Context) (map[string]model.TeamForm, error)
		GetFutureFixtures(ctx context.Context) (map[string]string, error)
	}
	Store interface {
		GetPrediction(ctx context.Context, userId, matchId string) (*common.Prediction, error)
		GetPredictionsByUserId(ctx context.Context, id string) ([]common.Prediction, error)
		UpdatePredictions(ctx context.Context, predictions []common.Prediction) error
		GetMatchPredictionSummary(ctx context.Context, id string) (int, int, int, error)
	}

	service struct {
		store          Store
		fixtureService FixtureService
	}
)

func New(store Store, fixtureService FixtureService) (*service, error) {
	if store == nil {
		return nil, errors.New("store_is_nil")
	}
	if fixtureService == nil {
		return nil, errors.New("fixture_service_is_nil")
	}

	return &service{
		store:          store,
		fixtureService: fixtureService,
	}, nil
}

func (s *service) GetFixturesWithPredictions(ctx context.Context, id string) ([]model.FixturePrediction, error) {
	fixturesChan := make(chan []common.Fixture)
	predictionsChan := make(chan []common.Prediction)

	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		f, err := s.fixtureService.GetMatches(ctx)
		fixturesChan <- f
		if err != nil {
			return fmt.Errorf("get_fixtures: %w", err)
		}

		return nil
	})
	g.Go(func() error {
		p, err := s.store.GetPredictionsByUserId(ctx, id)
		predictionsChan <- p
		if err != nil {
			return fmt.Errorf("get_predictions: %w", err)
		}
		return nil
	})

	fixtures := <-fixturesChan
	predictions := <-predictionsChan

	if err := g.Wait(); err != nil {
		return nil, err
	}

	var fixturePredictions []model.FixturePrediction
	for _, f := range fixtures {
		fp := model.FixturePrediction{
			UserId:     id,
			Id:         f.Id,
			HomeTeam:   f.HomeTeam,
			AwayTeam:   f.AwayTeam,
			HomeResult: f.HomeGoals,
			AwayResult: f.AwayGoals,
			Played:     f.Played,
			DateTime:   f.DateTime,
			Matchday:   f.Matchday,
		}

		for _, p := range predictions {
			if p.MatchId == f.Id {
				fp.HomeGoals = &p.HomeGoals
				fp.AwayGoals = &p.AwayGoals
				break
			}
		}

		fixturePredictions = append(fixturePredictions, fp)
	}

	return fixturePredictions, nil
}

func (s *service) GetPredictorData(ctx context.Context, id string) (*model.PredictorData, error) {
	fixturePredictionsChan := make(chan []model.FixturePrediction)
	formChan := make(chan map[string]model.TeamForm)

	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		f, err := s.GetFixturesWithPredictions(ctx, id)
		fixturePredictionsChan <- f
		if err != nil {
			return fmt.Errorf("get_fixtures_with_predictions: %w", err)
		}

		return nil
	})
	g.Go(func() error {
		p, err := s.fixtureService.GetTeamForm(ctx)
		formChan <- p
		if err != nil {
			return fmt.Errorf("get_team_form: %w", err)
		}
		return nil
	})

	fixturePredictions := <-fixturePredictionsChan
	forms := <-formChan

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &model.PredictorData{
		Predictions: fixturePredictions,
		Forms:       forms,
	}, nil
}

func (s *service) GetUsersPastPredictions(ctx context.Context, id string) (*model.PredictionSummary, error) {
	fixturePredictions, err := s.GetFixturesWithPredictions(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get_fixtures_with_predictions: %w", err)
	}

	var fp []model.FixturePrediction
	for _, f := range fixturePredictions {
		if time.Now().UTC().After(f.DateTime) {
			fp = append(fp, f)
		}
	}

	return &model.PredictionSummary{
		Matches: fp,
	}, nil
}

func (s *service) UpdatePredictions(ctx context.Context, predictions []common.Prediction) error {
	futureFixtures, err := s.fixtureService.GetFutureFixtures(ctx)
	if err != nil {
		return fmt.Errorf("get_future_fixtures: %w", err)
	}

	var validPredictions []common.Prediction
	for _, p := range predictions {
		if _, ok := futureFixtures[p.MatchId]; ok {
			validPredictions = append(validPredictions, p)
		}
	}

	err = s.store.UpdatePredictions(ctx, validPredictions)
	if err != nil {
		return fmt.Errorf("update_prediction: %w", err)
	}

	return nil
}

func (s *service) GetPrediction(ctx context.Context, userId, matchId string) (*common.Prediction, error) {
	prediction, err := s.store.GetPrediction(ctx, userId, matchId)
	if err != nil {
		return nil, fmt.Errorf("get_prediction: %w", err)
	}
	return prediction, nil
}

func (s *service) GetMatchPredictionSummary(ctx context.Context, id string) (*common.MatchPredictionSummary, error) {
	homeWins, draw, awayWins, err := s.store.GetMatchPredictionSummary(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get_match_prediction_summary: %w", err)
	}

	return &common.MatchPredictionSummary{
		HomeWin: homeWins,
		Draw:    draw,
		AwayWin: awayWins,
	}, nil
}
