package mongo

import (
	"context"
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/common/store/mongo"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"testing"
)

const (
	userId   = "userId"
	userId2  = "userId2"
	matchId  = "matchId"
	matchId2 = "matchId2"
)

func Test_Store(t *testing.T) {
	ctx := context.Background()

	err := os.Setenv("MONGO_PORT", "27017")
	require.NoError(t, err)
	err = os.Setenv("MONGO_HOST", "localhost")
	require.NoError(t, err)
	err = os.Setenv("MONGO_SCHEME", "mongodb")
	require.NoError(t, err)

	client, err := mongo.New(ctx)
	require.NoError(t, err)

	store, err := New(ctx, client)
	require.NoError(t, err)

	createPrediction := func(p *predictionEntity) {
		_, err = store.
			client.
			Database(db).
			Collection(collection).
			InsertOne(
				ctx,
				p,
			)

		require.NoError(t, err)
	}

	cleanupDb := func() {
		_, _ = store.
			client.
			Database(db).
			Collection(collection).
			DeleteMany(
				ctx,
				bson.M{},
			)
	}

	t.Run("GetPrediction", func(t *testing.T) {
		t.Run("gets prediction by userId and matchId", func(t *testing.T) {
			p := &predictionEntity{
				UserId:    userId,
				MatchId:   matchId,
				HomeGoals: 1,
				AwayGoals: 1,
			}

			defer cleanupDb()
			createPrediction(p)

			prediction, err := store.GetPrediction(ctx, userId, matchId)
			require.NoError(t, err)

			expectedResult := toPrediction(p)

			assert.Equal(t, expectedResult, prediction)
		})

		t.Run("returns error if not found", func(t *testing.T) {
			prediction, err := store.GetPrediction(ctx, userId, matchId)
			require.Error(t, err)

			assert.Nil(t, prediction)
			assert.Equal(t, model.ErrPredictionNotFound, err)
		})
	})

	t.Run("GetPredictionsByUserId", func(t *testing.T) {
		t.Run("gets all predictions by userId", func(t *testing.T) {
			p1 := &predictionEntity{
				UserId:    userId,
				MatchId:   matchId,
				HomeGoals: 1,
				AwayGoals: 1,
			}
			p2 := &predictionEntity{
				UserId:    userId,
				MatchId:   matchId2,
				HomeGoals: 1,
				AwayGoals: 1,
			}
			p3 := &predictionEntity{
				UserId:    userId2,
				MatchId:   matchId,
				HomeGoals: 1,
				AwayGoals: 1,
			}

			defer cleanupDb()
			createPrediction(p1)
			createPrediction(p2)
			createPrediction(p3)

			expectedResult := []common.Prediction{
				*toPrediction(p1),
				*toPrediction(p2),
			}

			predictions, err := store.GetPredictionsByUserId(ctx, userId)
			require.NoError(t, err)

			assert.Equal(t, expectedResult, predictions)
		})
	})

	t.Run("UpdatePredictions", func(t *testing.T) {
		t.Run("inserts new predictions", func(t *testing.T) {
			predictions := []common.Prediction{
				{
					UserId:    userId,
					MatchId:   matchId,
					HomeGoals: 2,
					AwayGoals: 1,
				},
				{
					UserId:    userId,
					MatchId:   matchId2,
					HomeGoals: 2,
					AwayGoals: 1,
				},
			}

			defer cleanupDb()

			err := store.UpdatePredictions(ctx, predictions)
			require.NoError(t, err)

			result, err := store.GetPredictionsByUserId(ctx, userId)
			require.NoError(t, err)

			assert.Equal(t, predictions, result)
		})

		t.Run("updates prediction if already exists", func(t *testing.T) {
			defer cleanupDb()

			err := store.UpdatePredictions(ctx, []common.Prediction{
				{
					UserId:    userId,
					MatchId:   matchId,
					HomeGoals: 2,
					AwayGoals: 1,
				},
			})
			require.NoError(t, err)

			predictions := []common.Prediction{
				{
					UserId:    userId,
					MatchId:   matchId,
					HomeGoals: 3,
					AwayGoals: 3,
				},
				{
					UserId:    userId,
					MatchId:   matchId2,
					HomeGoals: 2,
					AwayGoals: 1,
				},
			}

			err = store.UpdatePredictions(ctx, predictions)
			require.NoError(t, err)

			result, err := store.GetPredictionsByUserId(ctx, userId)
			require.NoError(t, err)

			assert.Equal(t, predictions, result)
		})
	})

	t.Run("GetMatchPredictionSummary", func(t *testing.T) {
		t.Run("gets a count of each prediction's result for a specified match	", func(t *testing.T) {
			predictions := []common.Prediction{
				{
					UserId:    "1",
					MatchId:   matchId,
					HomeGoals: 1,
					AwayGoals: 0,
				},
				{
					UserId:    "2",
					MatchId:   matchId2,
					HomeGoals: 1,
					AwayGoals: 1,
				},
				{
					UserId:    "3",
					MatchId:   matchId,
					HomeGoals: 3,
					AwayGoals: 1,
				},
				{
					UserId:    "4",
					MatchId:   matchId,
					HomeGoals: 0,
					AwayGoals: 1,
				},
			}

			err = store.UpdatePredictions(ctx, predictions)
			require.NoError(t, err)

			homeWins, draw, awayWins, err := store.GetMatchPredictionSummary(ctx, matchId)
			require.NoError(t, err)

			assert.Equal(t, 2, homeWins)
			assert.Equal(t, 0, draw)
			assert.Equal(t, 1, awayWins)
		})
	})
}
