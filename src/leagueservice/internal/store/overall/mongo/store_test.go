// +build integration

package mongo_test

import (
	"context"
	"os"
	"testing"

	commonmongo "github.com/cshep4/premier-predictor-microservices/src/common/store/mongo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongodb "go.mongodb.org/mongo-driver/mongo"

	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/store/overall/mongo"
)

const (
	name1 = "👋"
	name2 = "👋👋"
)

var (
	userId1 = primitive.NewObjectID()
	userId2 = primitive.NewObjectID()
)

func TestStore_Get(t *testing.T) {
	t.Run("returns not found if league user cannot be found", func(t *testing.T) {
		ctx := context.Background()

		client := mongoClient(ctx, t)

		store, err := mongo.New(ctx, client)
		require.NoError(t, err)

		defer store.Close(ctx)

		cleanupDb(ctx, client, t)

		league, err := store.Get(ctx, userId1.Hex())
		require.Error(t, err)

		assert.Equal(t, model.ErrLeagueUserNotFound, err)
		assert.Empty(t, league)
	})

	t.Run("gets the specified league user", func(t *testing.T) {
		ctx := context.Background()

		client := mongoClient(ctx, t)

		store, err := mongo.New(ctx, client)
		require.NoError(t, err)

		defer func() {
			cleanupDb(ctx, client, t)
			require.NoError(t, store.Close(ctx))
		}()

		entity := mongo.LeagueUserEntity{
			ID:   userId1,
			Name: name1,
		}

		leagueUser := model.LeagueUser{
			ID:   userId1.Hex(),
			Name: name1,
		}

		createLeagueUser(ctx, entity, client, t)

		result, err := store.Get(ctx, userId1.Hex())
		require.NoError(t, err)

		assert.Equal(t, leagueUser, result)
	})
}

func TestStore_Count(t *testing.T) {
	t.Run("returns count of all league users", func(t *testing.T) {
		ctx := context.Background()

		client := mongoClient(ctx, t)

		store, err := mongo.New(ctx, client)
		require.NoError(t, err)

		defer func() {
			cleanupDb(ctx, client, t)
			require.NoError(t, store.Close(ctx))
		}()

		createLeagueUser(ctx, mongo.LeagueUserEntity{
			ID:   userId1,
			Name: name1,
		}, client, t)
		createLeagueUser(ctx, mongo.LeagueUserEntity{
			ID:   userId2,
			Name: name2,
		}, client, t)

		count, err := store.Count(ctx)
		require.NoError(t, err)

		assert.Equal(t, 2, count)
	})
}

func mongoClient(ctx context.Context, t *testing.T) *mongodb.Client {
	err := os.Setenv("MONGO_PORT", "27017")
	require.NoError(t, err)
	err = os.Setenv("MONGO_HOST", "localhost")
	require.NoError(t, err)
	err = os.Setenv("MONGO_SCHEME", "mongodb")
	require.NoError(t, err)

	client, err := commonmongo.New(ctx)
	require.NoError(t, err)

	return client
}

func createLeagueUser(ctx context.Context, lu mongo.LeagueUserEntity, client *mongodb.Client, t *testing.T) {
	_, err := client.
		Database(mongo.DB).
		Collection(mongo.Collection).
		InsertOne(ctx, lu)

	require.NoError(t, err)
}

func cleanupDb(ctx context.Context, client *mongodb.Client, t *testing.T) {
	_, err := client.
		Database(mongo.DB).
		Collection(mongo.Collection).
		DeleteMany(ctx, bson.M{})

	require.NoError(t, err)
}
