package mongo

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/model"

	"github.com/cshep4/premier-predictor-microservices/src/common/store/mongo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	firstName = "a first name"
	surname   = "a surname"
	email     = "an email address"
	password  = "a new password"
)

var (
	id  = primitive.NewObjectID().Hex()
	id1 = primitive.NewObjectID()
	id2 = primitive.NewObjectID()
	id3 = primitive.NewObjectID()
	id4 = primitive.NewObjectID()
)

const (
	email1 = "📧1"
	email2 = "📧2"
	email3 = "📧3"
	email4 = "📧4"
)

func TestStore_GetUserById(t *testing.T) {
	store := newTestStore(t)
	defer store.Close(context.Background())

	t.Run("returns not found if user cannot be found", func(t *testing.T) {
		cleanupDb(store)

		user, err := store.GetUserById(context.Background(), id)
		require.Error(t, err)

		assert.Equal(t, model.ErrUserNotFound, err)
		assert.Nil(t, user)
	})

	t.Run("gets the specified user", func(t *testing.T) {
		defer cleanupDb(store)

		user := model.User{
			Id:        id,
			FirstName: "Chris",
			Surname:   "Shepherd",
		}

		entity, err := fromUser(user)
		require.NoError(t, err)

		createUser(entity, store, t)

		result, err := store.GetUserById(context.Background(), id)
		require.NoError(t, err)

		assert.Equal(t, &user, result)
	})
}

func TestStore_UpdateUserInfo(t *testing.T) {
	store := newTestStore(t)
	defer store.Close(context.Background())

	userInfo := model.UserInfo{
		Id:        id,
		FirstName: firstName,
		Surname:   surname,
		Email:     email,
	}

	t.Run("returns error if the user does not exist", func(t *testing.T) {
		cleanupDb(store)

		err := store.UpdateUserInfo(context.Background(), userInfo)
		require.Error(t, err)

		assert.Equal(t, model.ErrUserNotFound, err)
	})

	t.Run("updates the correct user's information", func(t *testing.T) {
		defer cleanupDb(store)

		user := model.User{
			Id: id,
		}

		entity, err := fromUser(user)
		require.NoError(t, err)

		createUser(entity, store, t)

		err = store.UpdateUserInfo(context.Background(), userInfo)
		require.NoError(t, err)

		result, err := store.GetUserById(context.Background(), id)
		require.NoError(t, err)

		assert.Equal(t, id, result.Id)
		assert.Equal(t, firstName, result.FirstName)
		assert.Equal(t, surname, result.Surname)
		assert.Equal(t, email, result.Email)
	})
}

func TestStore_UpdatePassword(t *testing.T) {
	store := newTestStore(t)
	defer store.Close(context.Background())

	t.Run("returns error if the user does not exist", func(t *testing.T) {
		cleanupDb(store)

		err := store.UpdatePassword(context.Background(), id, password)
		require.Error(t, err)

		assert.Equal(t, model.ErrUserNotFound, err)
	})

	t.Run("updates the correct user's password", func(t *testing.T) {
		defer cleanupDb(store)

		user := model.User{
			Id:        id,
			FirstName: firstName,
			Password:  "an old password",
		}

		entity, err := fromUser(user)
		require.NoError(t, err)

		createUser(entity, store, t)

		err = store.UpdatePassword(context.Background(), id, password)
		require.NoError(t, err)

		result, err := store.GetUserById(context.Background(), id)
		require.NoError(t, err)

		assert.Equal(t, id, result.Id)
		assert.Equal(t, password, result.Password)
	})
}

func TestStore_GetAllUsers(t *testing.T) {
	store := newTestStore(t)
	defer store.Close(context.Background())

	t.Run("returns empty slice if no users exist", func(t *testing.T) {
		cleanupDb(store)

		users, err := store.GetAllUsers(context.Background())
		require.NoError(t, err)

		assert.NotNil(t, users)
		assert.Equal(t, 0, len(users))
	})

	t.Run("returns all users from db", func(t *testing.T) {
		defer cleanupDb(store)

		id1 := primitive.NewObjectID()
		id2 := primitive.NewObjectID()
		id3 := primitive.NewObjectID()

		createUser(&userEntity{Id: &id1, Email: "1"}, store, t)
		createUser(&userEntity{Id: &id2, Email: "2"}, store, t)
		createUser(&userEntity{Id: &id3, Email: "3"}, store, t)

		users, err := store.GetAllUsers(context.Background())
		require.NoError(t, err)

		assert.NotNil(t, users)
		assert.Equal(t, 3, len(users))
		assert.Equal(t, id1.Hex(), users[0].Id)
		assert.Equal(t, id2.Hex(), users[1].Id)
		assert.Equal(t, id3.Hex(), users[2].Id)
	})
}

func TestStore_GetAllUsersByIds(t *testing.T) {
	store := newTestStore(t)
	defer store.Close(context.Background())

	t.Run("returns empty slice if no users exist", func(t *testing.T) {
		cleanupDb(store)

		users, err := store.GetAllUsers(context.Background())
		require.NoError(t, err)

		assert.NotNil(t, users)
		assert.Equal(t, 0, len(users))
	})

	t.Run("returns all users with specified id", func(t *testing.T) {
		defer cleanupDb(store)

		id1 := primitive.NewObjectID()
		id2 := primitive.NewObjectID()
		id3 := primitive.NewObjectID()

		createUser(&userEntity{Id: &id1, Email: "1"}, store, t)
		createUser(&userEntity{Id: &id2, Email: "2"}, store, t)
		createUser(&userEntity{Id: &id3, Email: "3"}, store, t)

		users, err := store.GetAllUsersByIds(context.Background(), []string{id1.Hex(), id3.Hex()})
		require.NoError(t, err)

		assert.NotNil(t, users)
		assert.Equal(t, 2, len(users))
		assert.Equal(t, id1.Hex(), users[0].Id)
		assert.Equal(t, id3.Hex(), users[1].Id)
	})
}

func TestStore_IsEmailTakenByADifferentUser(t *testing.T) {
	store := newTestStore(t)
	defer store.Close(context.Background())

	t.Run("returns true if a different user already has the same email", func(t *testing.T) {
		defer cleanupDb(store)

		otherId := primitive.NewObjectID()

		createUser(&userEntity{Id: &otherId, Email: email}, store, t)

		taken := store.IsEmailTakenByADifferentUser(context.Background(), id, email)

		assert.True(t, taken)
	})

	t.Run("returns false if the specified user already has the same email", func(t *testing.T) {
		defer cleanupDb(store)

		id := primitive.NewObjectID()

		createUser(&userEntity{Id: &id, Email: email}, store, t)

		taken := store.IsEmailTakenByADifferentUser(context.Background(), id.Hex(), email)

		assert.False(t, taken)
	})

	t.Run("returns false if the email address is not taken", func(t *testing.T) {
		cleanupDb(store)

		taken := store.IsEmailTakenByADifferentUser(context.Background(), id, email)

		assert.False(t, taken)
	})
}

func TestStore_GetOverallRank(t *testing.T) {
	store := newTestStore(t)
	defer store.Close(context.Background())

	t.Run("will get rank for a group of users", func(t *testing.T) {
		defer cleanupDb(store)

		createUser(&userEntity{Id: &id1, Email: email1, Score: 1}, store, t)
		createUser(&userEntity{Id: &id2, Email: email2, Score: 2}, store, t)
		createUser(&userEntity{Id: &id3, Email: email3, Score: 3}, store, t)
		createUser(&userEntity{Id: &id4, Email: email4, Score: 4}, store, t)

		rank, err := store.GetOverallRank(context.Background(), id1.Hex())
		require.NoError(t, err)

		assert.Equal(t, int64(4), rank)

		rank, err = store.GetOverallRank(context.Background(), id2.Hex())
		require.NoError(t, err)

		assert.Equal(t, int64(3), rank)

		rank, err = store.GetOverallRank(context.Background(), id3.Hex())
		require.NoError(t, err)

		assert.Equal(t, int64(2), rank)

		rank, err = store.GetOverallRank(context.Background(), id4.Hex())
		require.NoError(t, err)

		assert.Equal(t, int64(1), rank)
	})

	t.Run("if users have the same amount of points their rank will be the same", func(t *testing.T) {
		defer cleanupDb(store)

		createUser(&userEntity{Id: &id1, Email: email1, Score: 1}, store, t)
		createUser(&userEntity{Id: &id2, Email: email2, Score: 2}, store, t)
		createUser(&userEntity{Id: &id3, Email: email3, Score: 2}, store, t)
		createUser(&userEntity{Id: &id4, Email: email4, Score: 4}, store, t)

		rank, err := store.GetOverallRank(context.Background(), id1.Hex())
		require.NoError(t, err)

		assert.Equal(t, int64(4), rank)

		rank, err = store.GetOverallRank(context.Background(), id2.Hex())
		require.NoError(t, err)

		assert.Equal(t, int64(2), rank)

		rank, err = store.GetOverallRank(context.Background(), id3.Hex())
		require.NoError(t, err)

		assert.Equal(t, int64(2), rank)

		rank, err = store.GetOverallRank(context.Background(), id4.Hex())
		require.NoError(t, err)

		assert.Equal(t, int64(1), rank)
	})

	t.Run("returns error if user cannot be found", func(t *testing.T) {
		defer cleanupDb(store)

		createUser(&userEntity{Id: &id1, Email: email1, Score: 1}, store, t)
		createUser(&userEntity{Id: &id2, Email: email2, Score: 2}, store, t)

		rank, err := store.GetOverallRank(context.Background(), id3.Hex())
		require.Error(t, err)

		assert.Equal(t, model.ErrUserNotFound, err)
		assert.Empty(t, rank)
	})
}

func TestStore_GetOverallRank_Performance(t *testing.T) {
	store := newTestStore(t)
	defer store.Close(context.Background())
	defer cleanupDb(store)

	const documentNumber = 10000

	for i := 0; i < documentNumber; i++ {
		id := primitive.NewObjectID()
		createUser(&userEntity{Id: &id, Email: fmt.Sprintf("📨%d", i), Score: i}, store, t)
	}

	t.Run("performance test", func(t *testing.T) {
		createUser(&userEntity{Id: &id1, Email: email1, Score: 0}, store, t)

		rank, err := store.GetOverallRank(context.Background(), id1.Hex())
		require.NoError(t, err)

		assert.Equal(t, int64(documentNumber), rank)
	})
}

func TestStore_GetRankForGroup(t *testing.T) {
	store := newTestStore(t)
	defer store.Close(context.Background())

	t.Run("will get rank for a group of users", func(t *testing.T) {
		defer cleanupDb(store)

		createUser(&userEntity{Id: &id1, Email: email1, Score: 1}, store, t)
		createUser(&userEntity{Id: &id2, Email: email2, Score: 2}, store, t)
		createUser(&userEntity{Id: &id3, Email: email3, Score: 3}, store, t)
		createUser(&userEntity{Id: &id4, Email: email4, Score: 4}, store, t)

		ids := []string{id1.Hex(), id2.Hex(), id3.Hex(), id4.Hex()}

		rank, err := store.GetRankForGroup(context.Background(), id1.Hex(), ids)
		require.NoError(t, err)

		assert.Equal(t, int64(4), rank)

		rank, err = store.GetRankForGroup(context.Background(), id2.Hex(), ids)
		require.NoError(t, err)

		assert.Equal(t, int64(3), rank)

		rank, err = store.GetRankForGroup(context.Background(), id3.Hex(), ids)
		require.NoError(t, err)

		assert.Equal(t, int64(2), rank)

		rank, err = store.GetRankForGroup(context.Background(), id4.Hex(), ids)
		require.NoError(t, err)

		assert.Equal(t, int64(1), rank)
	})

	t.Run("returns error if user not in group", func(t *testing.T) {
		ids := []string{id1.Hex(), id2.Hex()}

		rank, err := store.GetRankForGroup(context.Background(), id3.Hex(), ids)
		require.Error(t, err)

		assert.Equal(t, "user not in group", err.Error())
		assert.Empty(t, rank)
	})

	t.Run("if users have the same amount of points their rank will be the same", func(t *testing.T) {
		defer cleanupDb(store)

		createUser(&userEntity{Id: &id1, Email: email1, Score: 1}, store, t)
		createUser(&userEntity{Id: &id2, Email: email2, Score: 2}, store, t)
		createUser(&userEntity{Id: &id3, Email: email3, Score: 2}, store, t)
		createUser(&userEntity{Id: &id4, Email: email4, Score: 4}, store, t)

		ids := []string{id1.Hex(), id2.Hex(), id3.Hex(), id4.Hex()}

		rank, err := store.GetRankForGroup(context.Background(), id1.Hex(), ids)
		require.NoError(t, err)

		assert.Equal(t, int64(4), rank)

		rank, err = store.GetRankForGroup(context.Background(), id2.Hex(), ids)
		require.NoError(t, err)

		assert.Equal(t, int64(2), rank)

		rank, err = store.GetRankForGroup(context.Background(), id3.Hex(), ids)
		require.NoError(t, err)

		assert.Equal(t, int64(2), rank)

		rank, err = store.GetRankForGroup(context.Background(), id4.Hex(), ids)
		require.NoError(t, err)

		assert.Equal(t, int64(1), rank)
	})

	t.Run("returns error if user cannot be found", func(t *testing.T) {
		defer cleanupDb(store)

		createUser(&userEntity{Id: &id1, Email: email1, Score: 1}, store, t)
		createUser(&userEntity{Id: &id2, Email: email2, Score: 2}, store, t)

		ids := []string{id1.Hex(), id2.Hex(), id3.Hex()}

		rank, err := store.GetRankForGroup(context.Background(), id3.Hex(), ids)
		require.Error(t, err)

		assert.Equal(t, model.ErrUserNotFound, err)
		assert.Empty(t, rank)
	})
}

func TestStore_GetRankForGroup_Performance(t *testing.T) {
	store := newTestStore(t)
	defer store.Close(context.Background())
	defer cleanupDb(store)

	for i := 0; i < 10000; i++ {
		id := primitive.NewObjectID()
		createUser(&userEntity{Id: &id, Email: fmt.Sprintf("📨%d", i), Score: i}, store, t)
	}

	t.Run("performance test", func(t *testing.T) {
		createUser(&userEntity{Id: &id1, Email: email1, Score: 1}, store, t)
		createUser(&userEntity{Id: &id2, Email: email2, Score: 2}, store, t)
		createUser(&userEntity{Id: &id3, Email: email3, Score: 2}, store, t)
		createUser(&userEntity{Id: &id4, Email: email4, Score: 4}, store, t)

		ids := []string{id1.Hex(), id2.Hex(), id3.Hex(), id4.Hex()}

		rank, err := store.GetRankForGroup(context.Background(), id1.Hex(), ids)
		require.NoError(t, err)

		assert.Equal(t, int64(4), rank)

		rank, err = store.GetRankForGroup(context.Background(), id2.Hex(), ids)
		require.NoError(t, err)

		assert.Equal(t, int64(2), rank)

		rank, err = store.GetRankForGroup(context.Background(), id3.Hex(), ids)
		require.NoError(t, err)

		assert.Equal(t, int64(2), rank)

		rank, err = store.GetRankForGroup(context.Background(), id4.Hex(), ids)
		require.NoError(t, err)

		assert.Equal(t, int64(1), rank)
	})
}

func TestStore_GetUserCount(t *testing.T) {
	store := newTestStore(t)
	defer store.Close(context.Background())
	defer cleanupDb(store)

	userCount := 1234

	for i := 0; i < userCount; i++ {
		id := primitive.NewObjectID()
		createUser(&userEntity{Id: &id, Email: fmt.Sprintf("📨%d", i), Score: i}, store, t)
	}

	t.Run("gets total number of stored users", func(t *testing.T) {
		count, err := store.GetUserCount(context.Background())
		require.NoError(t, err)

		assert.Equal(t, int64(userCount), count)
	})
}

func TestStore_GetUserByEmail(t *testing.T) {
	store := newTestStore(t)
	defer store.Close(context.Background())

	t.Run("returns not found if user cannot be found", func(t *testing.T) {
		cleanupDb(store)

		user, err := store.GetUserByEmail(context.Background(), email1)
		require.Error(t, err)

		assert.Equal(t, model.ErrUserNotFound, err)
		assert.Nil(t, user)
	})

	t.Run("gets the specified user", func(t *testing.T) {
		defer cleanupDb(store)

		user := model.User{
			Id:        id,
			Email:     email1,
			FirstName: "Chris",
			Surname:   "Shepherd",
		}

		entity, err := fromUser(user)
		require.NoError(t, err)

		createUser(entity, store, t)

		result, err := store.GetUserByEmail(context.Background(), email1)
		require.NoError(t, err)

		assert.Equal(t, &user, result)
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

func createUser(u *userEntity, store *store, t *testing.T) {
	_, err := store.
		client.
		Database(db).
		Collection(collection).
		InsertOne(
			context.Background(),
			u,
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
