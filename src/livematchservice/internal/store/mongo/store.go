package mongo

import (
	"context"
	"errors"
	"time"

	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

const (
	db         = "liveMatch"
	collection = "liveMatch"
)

var limit = int64(20)

type store struct {
	client *mongo.Client
}

func New(ctx context.Context, client *mongo.Client) (*store, error) {
	if client == nil {
		return nil, errors.New("mongo_client_is_nil")
	}

	s := &store{
		client: client,
	}

	if err := s.Ping(ctx); err != nil {
		return nil, err
	}

	if err := s.ensureIndexes(ctx); err != nil {
		return nil, err
	}

	return s, nil
}

func (r *store) ensureIndexes(ctx context.Context) error {
	_, err := r.client.
		Database(db).
		Collection(collection).
		Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys: bsonx.Doc{
				{Key: "matchDate", Value: bsonx.Int64(1)},
			},
			Options: options.Index().
				SetName("matchDate_idx").
				SetUnique(false).
				SetSparse(true).
				SetBackground(true),
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *store) GetUpcomingMatches() ([]common.MatchFacts, error) {
	year, month, day := time.Now().Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location())

	return r.getMatches(
		bson.D{
			{
				Key: "matchDate",
				Value: bson.D{
					{
						Key:   "$gte",
						Value: today,
					},
				},
			},
		},
		&options.FindOptions{
			Limit: &limit,
			Sort: bson.D{
				bson.E{Key: "matchDate", Value: 1},
			},
		},
	)
}

func (r *store) getMatches(filter interface{}, opts *options.FindOptions) ([]common.MatchFacts, error) {
	ctx := context.Background()

	cur, err := r.client.
		Database(db).
		Collection(collection).
		Find(
			ctx,
			filter,
			opts,
		)

	if err != nil {
		return nil, err
	}

	matches := []common.MatchFacts{}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var m matchFactsEntity
		err := cur.Decode(&m)
		if err != nil {
			return nil, err
		}

		matches = append(matches, *toMatchFacts(&m))
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return matches, nil
}

func (r *store) GetMatchFacts(id string) (*common.MatchFacts, error) {
	var m matchFactsEntity

	err := r.client.
		Database(db).
		Collection(collection).
		FindOne(
			context.Background(),
			bson.M{
				"_id": id,
			},
		).
		Decode(&m)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, model.ErrMatchNotFound
		}

		return nil, err
	}

	return toMatchFacts(&m), nil
}

func (s *store) Ping(ctx context.Context) error {
	ctx, _ = context.WithTimeout(ctx, 2*time.Second)
	return s.client.Ping(ctx, nil)
}

func (s *store) Close(ctx context.Context) error {
	return s.client.Disconnect(ctx)
}
