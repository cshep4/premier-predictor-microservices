package prediction

import (
	"context"
	"errors"
	"testing"

	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/mocks/prediction"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

const (
	userId  = "1"
	matchId = "2"
	hGoals  = 2
	aGoals  = 1
	hWin    = 100
	draw    = 12
	aWin    = 400
)

func TestPredictor_GetPrediction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	predictionClient := prediction_mocks.NewMockPredictionServiceClient(ctrl)

	predictor, err := New(predictionClient)
	require.NoError(t, err)

	req := &gen.PredictionRequest{
		UserId:  userId,
		MatchId: matchId,
	}

	tokenMap := map[string][]string{
		"token": {"token"},
	}

	ctx := metadata.NewIncomingContext(context.Background(), tokenMap)

	t.Run("Get the prediction from the predictionService", func(t *testing.T) {
		prediction := &gen.Prediction{
			UserId:  userId,
			MatchId: matchId,
			HGoals:  hGoals,
			AGoals:  aGoals,
		}
		predictionClient.EXPECT().GetPrediction(ctx, req).Return(prediction, nil)

		r := model.PredictionRequest{
			UserId:  userId,
			MatchId: matchId,
		}

		result, err := predictor.GetPrediction(ctx, r)

		expectedResult := common.PredictionFromGrpc(prediction)

		require.NoError(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Returns an error if there is a problem", func(t *testing.T) {
		e := errors.New("")
		predictionClient.EXPECT().GetPrediction(gomock.Any(), req).Return(nil, e)

		r := model.PredictionRequest{
			UserId:  userId,
			MatchId: matchId,
		}

		result, err := predictor.GetPrediction(ctx, r)
		require.Error(t, err)
		assert.Equal(t, e, err)
		assert.Nil(t, result)
	})
}

func TestPredictor_GetPredictionSummary(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	predictionClient := prediction_mocks.NewMockPredictionServiceClient(ctrl)

	predictor, err := New(predictionClient)
	require.NoError(t, err)

	req := &gen.IdRequest{
		Id: matchId,
	}

	tokenMap := map[string][]string{
		"token": {"token"},
	}

	ctx := metadata.NewIncomingContext(context.Background(), tokenMap)

	t.Run("Get the prediction summary from the predictionService", func(t *testing.T) {
		predictionSummary := &gen.MatchPredictionSummary{
			HomeWin: hWin,
			Draw:    draw,
			AwayWin: aWin,
		}
		predictionClient.EXPECT().GetPredictionSummary(gomock.Any(), req).Return(predictionSummary, nil)

		result, err := predictor.GetPredictionSummary(ctx, matchId)

		expectedResult := common.MatchPredictionSummaryFromGrpc(predictionSummary)

		require.NoError(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Returns an error if there is a problem", func(t *testing.T) {
		e := errors.New("")
		predictionClient.EXPECT().GetPredictionSummary(gomock.Any(), req).Return(nil, e)

		result, err := predictor.GetPredictionSummary(ctx, matchId)
		require.Error(t, err)
		assert.Equal(t, e, err)
		assert.Nil(t, result)
	})
}
