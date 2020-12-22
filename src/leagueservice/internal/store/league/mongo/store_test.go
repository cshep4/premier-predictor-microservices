// +build integration

package mongo

import (
	"context"
	"os"
	"testing"

	"github.com/cshep4/premier-predictor-microservices/src/common/store/mongo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"
)

const (
	pin     = int64(12345)
	name1   = "League of champions"
	name2   = "🏆🏆🏆🏆🏆🏆"
	userId1 = "1"
	userId2 = "2"
)

func TestStore_GetLeagueByPin(t *testing.T) {
	ctx := context.Background()

	store := newTestStore(t)
	defer store.Close(ctx)

	t.Run("returns not found if league cannot be found", func(t *testing.T) {
		cleanupDb(store)

		league, err := store.GetLeagueByPin(ctx, pin)
		require.Error(t, err)

		assert.Equal(t, model.ErrLeagueNotFound, err)
		assert.Nil(t, league)
	})

	t.Run("gets the specified league", func(t *testing.T) {
		defer cleanupDb(store)

		league := model.League{
			Pin:   pin,
			Name:  name1,
			Users: []string{userId1, userId2},
		}

		createLeague(fromLeague(league), store, t)

		result, err := store.GetLeagueByPin(ctx, pin)
		require.NoError(t, err)

		assert.Equal(t, &league, result)
	})
}

func TestStore_GetLeaguesByUserId(t *testing.T) {
	ctx := context.Background()

	store := newTestStore(t)
	defer store.Close(ctx)

	t.Run("returns a slice of the leagues the user has joined", func(t *testing.T) {
		defer cleanupDb(store)

		league1 := model.League{
			Pin:   1,
			Name:  name1,
			Users: []string{userId1, userId2},
		}

		createLeague(fromLeague(league1), store, t)

		league2 := model.League{
			Pin:   2,
			Name:  name2,
			Users: []string{userId1},
		}

		createLeague(fromLeague(league2), store, t)

		result, err := store.GetLeaguesByUserId(ctx, userId1)
		require.NoError(t, err)

		assert.Equal(t, &league1, result[0])
		assert.Equal(t, &league2, result[1])
	})

	t.Run("returns an empty slice if the user has no leagues", func(t *testing.T) {
		defer cleanupDb(store)

		league1 := model.League{
			Pin:   pin,
			Name:  name1,
			Users: []string{userId2},
		}

		createLeague(fromLeague(league1), store, t)

		result, err := store.GetLeaguesByUserId(ctx, userId1)
		require.NoError(t, err)

		assert.Equal(t, 0, len(result))
	})
}

func TestStore_AddLeague(t *testing.T) {
	ctx := context.Background()

	store := newTestStore(t)
	defer store.Close(ctx)

	t.Run("adds league to db", func(t *testing.T) {
		defer cleanupDb(store)

		league := model.League{
			Pin:   pin,
			Name:  name1,
			Users: []string{userId1},
		}

		err := store.AddLeague(ctx, league)
		require.NoError(t, err)

		result, err := store.GetLeagueByPin(ctx, pin)
		require.NoError(t, err)

		assert.Equal(t, &league, result)
	})
}

func TestStore_JoinLeague(t *testing.T) {
	ctx := context.Background()

	store := newTestStore(t)
	defer store.Close(ctx)

	t.Run("adds user to existing league", func(t *testing.T) {
		defer cleanupDb(store)

		createLeague(&leagueEntity{
			Pin:   pin,
			Name:  name1,
			Users: []string{userId2},
		}, store, t)

		err := store.JoinLeague(ctx, pin, userId1)
		require.NoError(t, err)

		league, err := store.GetLeagueByPin(ctx, pin)
		require.NoError(t, err)

		assert.Equal(t, pin, league.Pin)
		assert.Equal(t, userId2, league.Users[0])
		assert.Equal(t, userId1, league.Users[1])
	})

	t.Run("returns error if league does not exist", func(t *testing.T) {
		cleanupDb(store)

		err := store.JoinLeague(ctx, pin, userId1)
		require.Error(t, err)

		assert.Equal(t, model.ErrLeagueNotFound, err)
	})

	t.Run("does not re-add user or return error if user is already in the league", func(t *testing.T) {
		defer cleanupDb(store)

		createLeague(&leagueEntity{
			Pin:   pin,
			Name:  name1,
			Users: []string{userId2},
		}, store, t)

		err := store.JoinLeague(ctx, pin, userId1)
		require.NoError(t, err)

		err = store.JoinLeague(ctx, pin, userId1)
		require.NoError(t, err)

		league, err := store.GetLeagueByPin(ctx, pin)
		require.NoError(t, err)

		assert.Equal(t, pin, league.Pin)
		assert.Equal(t, userId2, league.Users[0])
		assert.Equal(t, userId1, league.Users[1])
		assert.Equal(t, 2, len(league.Users))
	})
}

func TestStore_LeaveLeague(t *testing.T) {
	ctx := context.Background()

	store := newTestStore(t)
	defer store.Close(ctx)

	t.Run("removes user from league", func(t *testing.T) {
		defer cleanupDb(store)

		createLeague(&leagueEntity{
			Pin:   pin,
			Name:  name1,
			Users: []string{userId1, userId2},
		}, store, t)

		err := store.LeaveLeague(ctx, pin, userId1)
		require.NoError(t, err)

		league, err := store.GetLeagueByPin(ctx, pin)
		require.NoError(t, err)

		assert.Equal(t, pin, league.Pin)
		assert.Equal(t, userId2, league.Users[0])
		assert.Equal(t, 1, len(league.Users))
	})

	t.Run("returns error if league does not exist", func(t *testing.T) {
		cleanupDb(store)

		err := store.LeaveLeague(ctx, pin, userId1)
		require.Error(t, err)

		assert.Equal(t, model.ErrLeagueNotFound, err)
	})

	t.Run("does not remove anyone from league and does not return error if user is not in the league", func(t *testing.T) {
		defer cleanupDb(store)

		createLeague(&leagueEntity{
			Pin:   pin,
			Name:  name1,
			Users: []string{userId2},
		}, store, t)

		err := store.LeaveLeague(ctx, pin, userId1)
		require.NoError(t, err)

		league, err := store.GetLeagueByPin(ctx, pin)
		require.NoError(t, err)

		assert.Equal(t, pin, league.Pin)
		assert.Equal(t, userId2, league.Users[0])
		assert.Equal(t, 1, len(league.Users))
	})
}

func TestStore_RenameLeague(t *testing.T) {
	ctx := context.Background()

	store := newTestStore(t)
	defer store.Close(ctx)

	t.Run("renames league", func(t *testing.T) {
		defer cleanupDb(store)

		createLeague(&leagueEntity{
			Pin:   pin,
			Name:  name1,
			Users: []string{userId2},
		}, store, t)

		err := store.RenameLeague(ctx, pin, name2)
		require.NoError(t, err)

		league, err := store.GetLeagueByPin(ctx, pin)
		require.NoError(t, err)

		assert.Equal(t, pin, league.Pin)
		assert.Equal(t, name2, league.Name)
	})

	t.Run("returns error if league does not exist", func(t *testing.T) {
		cleanupDb(store)

		err := store.RenameLeague(ctx, pin, name2)
		require.Error(t, err)

		assert.Equal(t, model.ErrLeagueNotFound, err)
	})

	t.Run("does not return error if the new name is the same as the old name", func(t *testing.T) {
		defer cleanupDb(store)

		createLeague(&leagueEntity{
			Pin:   pin,
			Name:  name1,
			Users: []string{userId2},
		}, store, t)

		err := store.RenameLeague(ctx, pin, name1)
		require.NoError(t, err)

		league, err := store.GetLeagueByPin(ctx, pin)
		require.NoError(t, err)

		assert.Equal(t, pin, league.Pin)
		assert.Equal(t, name1, league.Name)
	})
}

func newTestStore(t *testing.T) *store {
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

	return store
}

func createLeague(l *leagueEntity, store *store, t *testing.T) {
	_, err := store.
		client.
		Database(db).
		Collection(collection).
		InsertOne(
			context.Background(),
			l,
		)

	require.NoError(t, err)
}

func cleanupDb(store *store) {
	_, _ = store.
		client.
		Database(db).
		Collection(collection).
		DeleteMany(
			context.Background(),
			bson.M{},
		)
}
