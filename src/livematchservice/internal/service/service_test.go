package live

import (
	"context"
	"errors"
	"testing"
	"time"
	
	"github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/mocks/live"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/mocks/prediction"
	. "github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	userId  = "1"
	matchId = "2"
)

var (
	ctx = context.Background()
)

func TestService_GetMatchFacts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := live_mocks.NewMockStore(ctrl)
	predictor := prediction_mocks.NewMockPredictor(ctrl)

	service, err := New(store, predictor)
	require.NoError(t, err)

	t.Run("Retrieves match from db", func(t *testing.T) {
		matchFacts := &model.MatchFacts{}

		store.EXPECT().GetMatchFacts(ctx, userId).Return(matchFacts, nil)

		result, err := service.GetMatchFacts(ctx, userId)
		require.NoError(t, err)
		assert.Equal(t, matchFacts, result)
	})

	t.Run("Returns error if there is a problem", func(t *testing.T) {
		e := errors.New("")

		store.EXPECT().GetMatchFacts(ctx, userId).Return(nil, e)

		result, err := service.GetMatchFacts(ctx, userId)
		require.Error(t, err)
		assert.Equal(t, e, err)
		assert.Nil(t, result)
	})
}

func TestService_GetMatchSummary(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := live_mocks.NewMockStore(ctrl)
	predictor := prediction_mocks.NewMockPredictor(ctrl)

	service, err := New(store, predictor)
	require.NoError(t, err)

	t.Run("Retrieves match and prediction summary", func(t *testing.T) {
		matchFacts := &model.MatchFacts{}
		predictionSummary := &model.MatchPredictionSummary{}
		prediction := &model.Prediction{}

		req := PredictionRequest{
			UserId:  userId,
			MatchId: matchId,
		}

		store.EXPECT().GetMatchFacts(ctx, matchId).Return(matchFacts, nil)
		predictor.EXPECT().GetPredictionSummary(ctx, matchId).Return(predictionSummary, nil)
		predictor.EXPECT().GetPrediction(ctx, req).Return(prediction, nil)

		result, err := service.GetMatchSummary(ctx, req)
		require.NoError(t, err)
		assert.Equal(t, prediction, result.Prediction)
		assert.Equal(t, predictionSummary, result.PredictionSummary)
		assert.Equal(t, matchFacts, result.Match)
	})

	t.Run("Returns error if there is a problem getting match from db", func(t *testing.T) {
		predictionSummary := &model.MatchPredictionSummary{}
		prediction := &model.Prediction{}
		e := errors.New("")

		req := PredictionRequest{
			UserId:  userId,
			MatchId: matchId,
		}

		store.EXPECT().GetMatchFacts(ctx, matchId).Return(nil, e)
		predictor.EXPECT().GetPredictionSummary(ctx, matchId).Return(predictionSummary, nil)
		predictor.EXPECT().GetPrediction(ctx, req).Return(prediction, nil)

		result, err := service.GetMatchSummary(ctx, req)
		require.Error(t, err)
		assert.Equal(t, e, err)
		assert.Nil(t, result)
	})

	t.Run("Returns error if there is a problem getting prediction summary", func(t *testing.T) {
		matchFacts := &model.MatchFacts{}
		e := errors.New("")
		prediction := &model.Prediction{}

		req := PredictionRequest{
			UserId:  userId,
			MatchId: matchId,
		}

		store.EXPECT().GetMatchFacts(ctx, matchId).Return(matchFacts, nil)
		predictor.EXPECT().GetPredictionSummary(ctx, matchId).Return(nil, e)
		predictor.EXPECT().GetPrediction(ctx, req).Return(prediction, nil)

		result, err := service.GetMatchSummary(ctx, req)
		require.Error(t, err)
		assert.Equal(t, e, err)
		assert.Nil(t, result)
	})

	t.Run("Returns error if there is a problem getting prediction", func(t *testing.T) {
		matchFacts := &model.MatchFacts{}
		predictionSummary := &model.MatchPredictionSummary{}
		e := errors.New("")

		req := PredictionRequest{
			UserId:  userId,
			MatchId: matchId,
		}

		store.EXPECT().GetMatchFacts(ctx, matchId).Return(matchFacts, nil)
		predictor.EXPECT().GetPredictionSummary(ctx, matchId).Return(predictionSummary, nil)
		predictor.EXPECT().GetPrediction(ctx, req).Return(nil, e)

		result, err := service.GetMatchSummary(ctx, req)
		require.Error(t, err)
		assert.Equal(t, e, err)
		assert.Nil(t, result)
	})
}

func TestService_GetUpcomingMatches(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := live_mocks.NewMockStore(ctrl)
	predictor := prediction_mocks.NewMockPredictor(ctrl)

	service, err := New(store, predictor)
	require.NoError(t, err)

	t.Run("Gets upcoming game from db and groups them by date", func(t *testing.T) {
		today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location())
		tomorrow := today.AddDate(0, 0, 1)

		id1 := "1"
		m1 := model.MatchFacts{
			Id:            id1,
			MatchDate:     tomorrow,
			FormattedDate: tomorrow.Format("02.01.2006"),
			Time:          "15:00",
		}
		id2 := "2"
		m2 := model.MatchFacts{
			Id:            id2,
			MatchDate:     today,
			FormattedDate: today.Format("02.01.2006"),
			Time:          "12:00",
		}
		id3 := "3"
		m3 := model.MatchFacts{
			Id:            id3,
			MatchDate:     tomorrow,
			FormattedDate: tomorrow.Format("02.01.2006"),
			Time:          "12:00",
		}
		store.EXPECT().GetUpcomingMatches(ctx).Return([]model.MatchFacts{m1, m2, m3}, nil)

		expectedResult := map[time.Time][]model.MatchFacts{
			tomorrow: {m3, m1},
			today:    {m2},
		}

		result, err := service.GetUpcomingMatches(ctx)
		require.NoError(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Returns error if there is a problem getting from db", func(t *testing.T) {
		e := errors.New("")

		store.EXPECT().GetUpcomingMatches(ctx).Return(nil, e)

		result, err := service.GetUpcomingMatches(ctx)
		require.Error(t, err)
		assert.Equal(t, e, err)
		assert.Nil(t, result)
	})
}
