package mongo

import (
	"context"
	"github.com/cshep4/premier-predictor-microservices/src/common/store/mongo"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	matchId = "matchId"
)

func Test_Store(t *testing.T) {
	ctx := context.Background()

	err := os.Setenv("MONGO_PORT", "27017")
	assert.NoError(t, err)
	err = os.Setenv("MONGO_HOST", "localhost")
	assert.NoError(t, err)
	err = os.Setenv("MONGO_SCHEME", "mongodb")
	assert.NoError(t, err)

	client, err := mongo.New(ctx)
	require.NoError(t, err)

	store, err := New(ctx, client)
	require.NoError(t, err)

	createMatch := func(m *matchFactsEntity) {
		_, err = store.
			client.
			Database(db).
			Collection(collection).
			InsertOne(
				context.Background(),
				m,
			)

		require.NoError(t, err)
	}

	cleanupDb := func() {
		_, _ = store.
			client.
			Database(db).
			Collection(collection).
			DeleteMany(
				context.Background(),
				bson.M{},
			)
	}

	t.Run("GetUpcomingMatches", func(t *testing.T) {
		t.Run("gets the next 20 matches, either in the future or today", func(t *testing.T) {
			const numMatches = 50
			const limit = 20

			defer cleanupDb()

			for i := 0; i < numMatches; i++ {
				m := &matchFactsEntity{
					Id:        matchId + strconv.Itoa(i),
					MatchDate: time.Now().AddDate(0, 0, numMatches-1-i),
				}
				createMatch(m)
			}

			matches, err := store.GetUpcomingMatches()
			require.NoError(t, err)

			assert.Equal(t, limit, len(matches))

			for i := range matches {
				id := matchId + strconv.Itoa(numMatches-i-1)
				assert.Equal(t, id, matches[i].Id)
			}
		})

		t.Run("gets today's matches even if they have passed", func(t *testing.T) {
			defer cleanupDb()

			m := &matchFactsEntity{
				Id:        matchId,
				MatchDate: time.Now().Add(-time.Hour).Round(time.Second).UTC(),
			}
			createMatch(m)

			matches, err := store.GetUpcomingMatches()
			require.NoError(t, err)

			assert.Equal(t, toMatchFacts(m), &matches[0])
		})

		t.Run("will not return matches in the past", func(t *testing.T) {
			const numMatches = 50

			defer cleanupDb()

			for i := 0; i < numMatches; i++ {
				m := &matchFactsEntity{
					Id:        matchId + strconv.Itoa(i),
					MatchDate: time.Now().AddDate(0, 0, -1-i),
				}
				createMatch(m)
			}

			matches, err := store.GetUpcomingMatches()
			require.NoError(t, err)

			assert.Equal(t, 0, len(matches))
		})
	})

	t.Run("GetMatchFacts", func(t *testing.T) {
		t.Run("gets the specified match", func(t *testing.T) {
			m := &matchFactsEntity{
				Id:         matchId,
				Commentary: &commentary{},
				MatchDate:  time.Now().Round(time.Second).UTC(),
			}

			createMatch(m)
			defer cleanupDb()

			match, err := store.GetMatchFacts(matchId)
			require.NoError(t, err)

			expectedResult := toMatchFacts(m)

			assert.Equal(t, expectedResult, match)
		})

		t.Run("Returns error if not found", func(t *testing.T) {
			match, err := store.GetMatchFacts(matchId)
			require.Error(t, err)

			assert.Nil(t, match)
			assert.Equal(t, model.ErrMatchNotFound, err)
		})
	})
}
