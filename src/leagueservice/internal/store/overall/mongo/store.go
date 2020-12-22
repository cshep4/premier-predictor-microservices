package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"
)

const (
	db         = "league"
	collection = "overall"
)

type (
	store struct {
		client     *mongo.Client
		collection *mongo.Collection
	}

	// InvalidParameterError is returned when a required parameter passed to New is invalid.
	InvalidParameterError struct {
		Parameter string
	}
)

func (i InvalidParameterError) Error() string {
	return fmt.Sprintf("invalid parameter %s", i.Parameter)
}

func New(ctx context.Context, client *mongo.Client) (*store, error) {
	if client == nil {
		return nil, InvalidParameterError{Parameter: "client"}
	}

	s := store{
		client:     client,
		collection: client.Database(db).Collection(collection),
	}

	if err := s.ping(ctx); err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	return &s, nil
}

func (s *store) Get(ctx context.Context, id string) (model.LeagueUser, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.LeagueUser{}, fmt.Errorf("cannot create objectID for %s: %w", id, err)
	}

	var lu leagueUserEntity

	err = s.collection.FindOne(ctx, bson.D{
		{
			Key:   "_id",
			Value: objectId,
		},
	}).Decode(&lu)

	switch {
	case err == mongo.ErrNoDocuments:
		return model.LeagueUser{}, model.ErrLeagueUserNotFound
	case err != nil:
		return model.LeagueUser{}, fmt.Errorf("cannot find one for %s: %w", id, err)
	}

	return lu.toLeagueUser(), nil
}

func (s *store) Count(ctx context.Context) (int64, error) {
	count, err := s.collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		return 0, fmt.Errorf("cannot count documents: %w", err)
	}

	return count, nil
}

func (s *store) List(ctx context.Context, ids []string) ([]model.LeagueUser, error) {
	var objectIDs []primitive.ObjectID
	for _, id := range ids {
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, fmt.Errorf("cannot create objectID for %s: %w", id, err)
		}

		objectIDs = append(objectIDs, oid)
	}

	leagueUsers, err := s.list(ctx, bson.D{
		{
			Key: "_id",
			Value: bson.D{
				{
					Key:   "$in",
					Value: objectIDs,
				},
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("cannot list league users for %v: %w", ids, err)
	}

	return leagueUsers, nil
}

func (s *store) ListAll(ctx context.Context) ([]model.LeagueUser, error) {
	leagueUsers, err := s.list(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("cannot list all league users: %w", err)
	}

	return leagueUsers, nil
}

func (s *store) list(ctx context.Context, filter bson.D) ([]model.LeagueUser, error) {
	leagueUsers := []model.LeagueUser{}

	opts := options.Find()
	opts.SetSort(bson.D{{"score", -1}})

	cur, err := s.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("find: %w", err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var lu leagueUserEntity

		if err := cur.Decode(&lu); err != nil {
			return nil, fmt.Errorf("decode: %w", err)
		}

		leagueUsers = append(leagueUsers, lu.toLeagueUser())
	}

	return leagueUsers, nil
}

func (s *store) ping(ctx context.Context) error {
	ctx, _ = context.WithTimeout(ctx, 2*time.Second)
	return s.client.Ping(ctx, nil)
}

func (s *store) Close(ctx context.Context) error {
	return s.client.Disconnect(ctx)
}
