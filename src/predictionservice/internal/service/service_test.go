package prediction

import (
	"errors"
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/mocks/fixture"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/mocks/prediction"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

const (
	userId   = "1"
	matchId  = "1"
	matchId2 = "2"
	teamOne  = "Team 1"
	teamTwo  = "Team 2"
)

var (
	now = time.Now()
	e   = errors.New("error")
	one = 1
	two = 2
)

func TestService_GetFixturesWithPredictions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fixtureService := fixture_mocks.NewMockFixtureService(ctrl)
	store := prediction_mocks.NewMockStore(ctrl)

	service, err := New(store, fixtureService)
	require.NoError(t, err)

	t.Run("Gets fixtures and predictions and merges them", func(t *testing.T) {
		fixtures := []common.Fixture{
			{
				Id:        matchId,
				HomeTeam:  teamOne,
				AwayTeam:  teamTwo,
				HomeGoals: &one,
				AwayGoals: &one,
				Played:    1,
				DateTime:  now,
				Matchday:  1,
			},
			{
				Id:        matchId2,
				HomeTeam:  teamTwo,
				AwayTeam:  teamOne,
				HomeGoals: &one,
				AwayGoals: &two,
				Played:    1,
				DateTime:  now,
				Matchday:  2,
			},
		}

		predictions := []common.Prediction{
			{
				MatchId:   matchId,
				UserId:    userId,
				HomeGoals: 1,
				AwayGoals: 1,
			},
		}

		fixtureService.EXPECT().GetMatches().Return(fixtures, nil)
		store.EXPECT().GetPredictionsByUserId(userId).Return(predictions, nil)

		result, err := service.GetFixturesWithPredictions(userId)

		require.NoError(t, err)
		assert.Equal(t, matchId, result[0].Id)
		assert.Equal(t, userId, result[0].UserId)
		assert.Equal(t, teamOne, result[0].HomeTeam)
		assert.Equal(t, teamTwo, result[0].AwayTeam)
		assert.Equal(t, 1, *result[0].HomeGoals)
		assert.Equal(t, 1, *result[0].AwayGoals)
		assert.Equal(t, 1, *result[0].HomeResult)
		assert.Equal(t, 1, *result[0].AwayResult)
		assert.Equal(t, 1, result[0].Played)
		assert.Equal(t, now, result[0].DateTime)
		assert.Equal(t, 1, result[0].Matchday)
		assert.Equal(t, matchId2, result[1].Id)
		assert.Equal(t, userId, result[1].UserId)
		assert.Equal(t, teamTwo, result[1].HomeTeam)
		assert.Equal(t, teamOne, result[1].AwayTeam)
		assert.Nil(t, result[1].HomeGoals)
		assert.Nil(t, result[1].AwayGoals)
		assert.Equal(t, 1, *result[1].HomeResult)
		assert.Equal(t, 2, *result[1].AwayResult)
		assert.Equal(t, 1, result[1].Played)
		assert.Equal(t, now, result[1].DateTime)
		assert.Equal(t, 2, result[1].Matchday)
	})

	t.Run("Error if there is a problem getting fixtureService from FixtureService", func(t *testing.T) {
		var predictions []common.Prediction

		fixtureService.EXPECT().GetMatches().Return(nil, e)
		store.EXPECT().GetPredictionsByUserId(userId).Return(predictions, nil)

		result, err := service.GetFixturesWithPredictions(userId)

		require.Error(t, err)
		assert.Contains(t, err.Error(), e.Error())
		assert.Nil(t, result)
	})

	t.Run("Error if there is a problem getting predictions from db", func(t *testing.T) {
		var fixtures []common.Fixture

		fixtureService.EXPECT().GetMatches().Return(fixtures, nil)
		store.EXPECT().GetPredictionsByUserId(userId).Return(nil, e)

		result, err := service.GetFixturesWithPredictions(userId)

		require.Error(t, err)
		assert.Contains(t, err.Error(), e.Error())
		assert.Nil(t, result)
	})
}

func TestService_GetPredictorData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fixtureService := fixture_mocks.NewMockFixtureService(ctrl)
	store := prediction_mocks.NewMockStore(ctrl)

	service, err := New(store, fixtureService)
	require.NoError(t, err)

	t.Run("Gets fixtures and predictions and merges them, then gets team forms", func(t *testing.T) {
		fixtures := []common.Fixture{
			{
				Id:        matchId,
				HomeTeam:  teamOne,
				AwayTeam:  teamTwo,
				HomeGoals: &one,
				AwayGoals: &one,
				Played:    1,
				DateTime:  now,
				Matchday:  1,
			},
			{
				Id:        matchId2,
				HomeTeam:  teamTwo,
				AwayTeam:  teamOne,
				HomeGoals: &one,
				AwayGoals: &two,
				Played:    1,
				DateTime:  now,
				Matchday:  2,
			},
		}

		predictions := []common.Prediction{
			{
				MatchId:   matchId,
				UserId:    userId,
				HomeGoals: 1,
				AwayGoals: 1,
			},
		}

		var forms map[string]model.TeamForm

		fixtureService.EXPECT().GetMatches().Return(fixtures, nil)
		store.EXPECT().GetPredictionsByUserId(userId).Return(predictions, nil)
		fixtureService.EXPECT().GetTeamForm().Return(forms, nil)

		result, err := service.GetPredictorData(userId)

		require.NoError(t, err)
		assert.Equal(t, matchId, result.Predictions[0].Id)
		assert.Equal(t, userId, result.Predictions[0].UserId)
		assert.Equal(t, teamOne, result.Predictions[0].HomeTeam)
		assert.Equal(t, teamTwo, result.Predictions[0].AwayTeam)
		assert.Equal(t, 1, *result.Predictions[0].HomeGoals)
		assert.Equal(t, 1, *result.Predictions[0].AwayGoals)
		assert.Equal(t, 1, *result.Predictions[0].HomeResult)
		assert.Equal(t, 1, *result.Predictions[0].AwayResult)
		assert.Equal(t, 1, result.Predictions[0].Played)
		assert.Equal(t, now, result.Predictions[0].DateTime)
		assert.Equal(t, 1, result.Predictions[0].Matchday)
		assert.Equal(t, matchId2, result.Predictions[1].Id)
		assert.Equal(t, userId, result.Predictions[1].UserId)
		assert.Equal(t, teamTwo, result.Predictions[1].HomeTeam)
		assert.Equal(t, teamOne, result.Predictions[1].AwayTeam)
		assert.Nil(t, result.Predictions[1].HomeGoals)
		assert.Nil(t, result.Predictions[1].AwayGoals)
		assert.Equal(t, 1, *result.Predictions[1].HomeResult)
		assert.Equal(t, 2, *result.Predictions[1].AwayResult)
		assert.Equal(t, 1, result.Predictions[1].Played)
		assert.Equal(t, now, result.Predictions[1].DateTime)
		assert.Equal(t, 2, result.Predictions[1].Matchday)
		assert.Equal(t, forms, result.Forms)
	})

	t.Run("Error if there is a problem getting fixtureService from FixtureService", func(t *testing.T) {
		var predictions []common.Prediction
		var forms map[string]model.TeamForm

		fixtureService.EXPECT().GetMatches().Return(nil, e)
		store.EXPECT().GetPredictionsByUserId(userId).Return(predictions, nil)
		fixtureService.EXPECT().GetTeamForm().Return(forms, nil)

		result, err := service.GetPredictorData(userId)

		require.Error(t, err)
		assert.Contains(t, err.Error(), e.Error())
		assert.Nil(t, result)
	})

	t.Run("Error if there is a problem getting predictions from db", func(t *testing.T) {
		var fixtures []common.Fixture
		var forms map[string]model.TeamForm

		fixtureService.EXPECT().GetMatches().Return(fixtures, nil)
		store.EXPECT().GetPredictionsByUserId(userId).Return(nil, e)
		fixtureService.EXPECT().GetTeamForm().Return(forms, nil)

		result, err := service.GetPredictorData(userId)

		require.Error(t, err)
		assert.Contains(t, err.Error(), e.Error())
		assert.Nil(t, result)
	})

	t.Run("Error if there is a problem getting forms from fixtureService", func(t *testing.T) {
		var fixtures []common.Fixture
		var predictions []common.Prediction

		fixtureService.EXPECT().GetMatches().Return(fixtures, nil)
		store.EXPECT().GetPredictionsByUserId(userId).Return(predictions, nil)
		fixtureService.EXPECT().GetTeamForm().Return(nil, e)

		result, err := service.GetPredictorData(userId)

		require.Error(t, err)
		assert.Contains(t, err.Error(), e.Error())
		assert.Nil(t, result)
	})
}

func TestService_GetUsersPastPredictions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fixtureService := fixture_mocks.NewMockFixtureService(ctrl)
	store := prediction_mocks.NewMockStore(ctrl)

	service, err := New(store, fixtureService)
	require.NoError(t, err)

	t.Run("Gets fixtures and predictions and merges them, then strips out future games", func(t *testing.T) {
		fixtures := []common.Fixture{
			{
				Id:        matchId,
				HomeTeam:  teamOne,
				AwayTeam:  teamTwo,
				HomeGoals: &one,
				AwayGoals: &one,
				Played:    1,
				DateTime:  now,
				Matchday:  1,
			},
			{
				Id:        matchId2,
				HomeTeam:  teamTwo,
				AwayTeam:  teamOne,
				HomeGoals: &one,
				AwayGoals: &two,
				Played:    1,
				DateTime:  now.AddDate(0, 0, 1),
				Matchday:  2,
			},
		}

		predictions := []common.Prediction{
			{
				MatchId:   matchId,
				UserId:    userId,
				HomeGoals: 1,
				AwayGoals: 1,
			},
		}

		fixtureService.EXPECT().GetMatches().Return(fixtures, nil)
		store.EXPECT().GetPredictionsByUserId(userId).Return(predictions, nil)

		predictionSummary, err := service.GetUsersPastPredictions(userId)

		result := predictionSummary.Matches

		require.NoError(t, err)
		assert.Equal(t, matchId, result[0].Id)
		assert.Equal(t, userId, result[0].UserId)
		assert.Equal(t, teamOne, result[0].HomeTeam)
		assert.Equal(t, teamTwo, result[0].AwayTeam)
		assert.Equal(t, 1, *result[0].HomeGoals)
		assert.Equal(t, 1, *result[0].AwayGoals)
		assert.Equal(t, 1, *result[0].HomeResult)
		assert.Equal(t, 1, *result[0].AwayResult)
		assert.Equal(t, 1, result[0].Played)
		assert.Equal(t, now, result[0].DateTime)
		assert.Equal(t, 1, result[0].Matchday)
		assert.Equal(t, 1, len(result))
	})

	t.Run("Error if there is a problem getting fixtureService from FixtureService", func(t *testing.T) {
		var predictions []common.Prediction

		fixtureService.EXPECT().GetMatches().Return(nil, e)
		store.EXPECT().GetPredictionsByUserId(userId).Return(predictions, nil)

		result, err := service.GetUsersPastPredictions(userId)

		require.Error(t, err)
		assert.Contains(t, err.Error(), e.Error())
		assert.Nil(t, result)
	})

	t.Run("Error if there is a problem getting predictions from db", func(t *testing.T) {
		var fixtures []common.Fixture

		fixtureService.EXPECT().GetMatches().Return(fixtures, nil)
		store.EXPECT().GetPredictionsByUserId(userId).Return(nil, e)

		result, err := service.GetUsersPastPredictions(userId)

		require.Error(t, err)
		assert.Contains(t, err.Error(), e.Error())
		assert.Nil(t, result)
	})
}

func TestService_UpdatePredictions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fixtureService := fixture_mocks.NewMockFixtureService(ctrl)
	store := prediction_mocks.NewMockStore(ctrl)

	service, err := New(store, fixtureService)
	require.NoError(t, err)

	t.Run("Stores predictions in db", func(t *testing.T) {
		var predictions []common.Prediction

		store.EXPECT().UpdatePredictions(predictions).Return(nil)

		err := service.UpdatePredictions(predictions)

		require.NoError(t, err)
	})

	t.Run("Returns error if there is a problem", func(t *testing.T) {
		var predictions []common.Prediction

		store.EXPECT().UpdatePredictions(predictions).Return(e)

		err := service.UpdatePredictions(predictions)

		require.Error(t, err)
		assert.Contains(t, err.Error(), e.Error())
	})
}

func TestService_GetPrediction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fixtureService := fixture_mocks.NewMockFixtureService(ctrl)
	store := prediction_mocks.NewMockStore(ctrl)

	service, err := New(store, fixtureService)
	require.NoError(t, err)

	t.Run("Gets prediction from db", func(t *testing.T) {
		prediction := &common.Prediction{
			UserId:  userId,
			MatchId: matchId,
		}

		store.EXPECT().GetPrediction(userId, matchId).Return(prediction, nil)

		result, err := service.GetPrediction(userId, matchId)

		require.NoError(t, err)
		assert.Equal(t, prediction, result)
	})

	t.Run("Returns error if there is a problem", func(t *testing.T) {
		store.EXPECT().GetPrediction(userId, matchId).Return(nil, e)

		result, err := service.GetPrediction(userId, matchId)

		require.Error(t, err)
		assert.Contains(t, err.Error(), e.Error())
		assert.Nil(t, result)
	})
}

func TestService_GetMatchPredictionSummary(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fixtureService := fixture_mocks.NewMockFixtureService(ctrl)
	store := prediction_mocks.NewMockStore(ctrl)

	service, err := New(store, fixtureService)
	require.NoError(t, err)

	t.Run("gets match prediction summary from db", func(t *testing.T) {
		homeWins := 24
		draw := 12
		awayWins := 4

		store.EXPECT().GetMatchPredictionSummary(matchId).Return(homeWins, draw, awayWins, nil)

		matchPredictionSummary, err := service.GetMatchPredictionSummary(matchId)
		require.NoError(t, err)

		assert.Equal(t, homeWins, matchPredictionSummary.HomeWin)
		assert.Equal(t, draw, matchPredictionSummary.Draw)
		assert.Equal(t, awayWins, matchPredictionSummary.AwayWin)
	})

	t.Run("returns error if there was a problem", func(t *testing.T) {
		e := errors.New("some error")

		store.EXPECT().GetMatchPredictionSummary(matchId).Return(0, 0, 0, e)

		matchPredictionSummary, err := service.GetMatchPredictionSummary(matchId)
		require.Error(t, err)

		assert.Contains(t, err.Error(), e.Error())
		assert.Nil(t, matchPredictionSummary)
	})
}
