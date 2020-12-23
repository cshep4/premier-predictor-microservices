package mongo

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
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
	collection = "2020-2021"
)

var limit = int64(20)

type store struct {
	client *mongo.Client
}

func New(ctx context.Context, client *mongo.Client) (*store, error) {
	if client == nil {
		return nil, errors.New("mongo client is nil")
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

func (s *store) ensureIndexes(ctx context.Context) error {
	_, err := s.client.
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

func (s *store) GetUpcomingMatches(ctx context.Context) ([]common.MatchFacts, error) {
	year, month, day := time.Now().Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location())

	return s.getMatches(
		ctx,
		bson.D{
			{
				Key:   "matchDate",
				Value: bson.D{{Key: "$gte", Value: today}},
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

func (s *store) GetTodaysMatches(ctx context.Context) ([]common.MatchFacts, error) {
	year, month, day := time.Now().Date()
	today := time.Date(year, month, day-1, 0, 0, 0, 0, time.Now().Location())
	tomorrow := time.Date(year, month, day+1, 0, 0, 0, 0, time.Now().Location())

	return s.getMatches(
		ctx,
		bson.D{
			{
				Key: "matchDate",
				Value: bson.D{
					{Key: "$gte", Value: today},
					{Key: "$lt", Value: tomorrow},
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

func (s *store) getMatches(ctx context.Context, filter interface{}, opts *options.FindOptions) ([]common.MatchFacts, error) {
	cur, err := s.client.
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

func (s *store) GetMatchFacts(ctx context.Context, id string) (*common.MatchFacts, error) {
	var m matchFactsEntity

	err := s.client.
		Database(db).
		Collection(collection).
		FindOne(
			ctx,
			bson.M{"_id": id},
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

func (s *store) SubscribeToMatch(ctx context.Context, id string, observer model.MatchObserver) error {
	cs, err := s.client.
		Database(db).
		Collection(collection).
		Watch(ctx, mongo.Pipeline{
			bson.D{{
				Key: "$match",
				Value: bson.D{
					{Key: "$and", Value: bson.A{
						bson.D{{Key: "$or", Value: bson.A{
							bson.D{{Key: "operationType", Value: "replace"}},
							bson.D{{Key: "operationType", Value: "update"}},
						}}},
						bson.D{{Key: "documentKey._id", Value: id}},
					}},
				},
			}},
		})
	if err != nil {
		return fmt.Errorf("watch_match: %w", err)
	}

	observable := model.MatchObservable{}
	observable.AddObserver(observer)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		for cs.Next(ctx) {
			m, err := s.GetMatchFacts(ctx, id)
			if err != nil {
				return fmt.Errorf("get_match_facts: %w", err)
			}

			if err := observable.Notify(m); err != nil {
				return fmt.Errorf("observable_notify: %w", err)
			}
		}
		return nil
	})

	return g.Wait()
}

func (s *store) SubscribeToMatches(ctx context.Context, ids []string, observer model.MatchObserver) error {
	cs, err := s.client.
		Database(db).
		Collection(collection).
		Watch(ctx, mongo.Pipeline{
			bson.D{{
				Key: "$match",
				Value: bson.D{
					{Key: "$and", Value: bson.A{
						bson.D{{Key: "$or", Value: bson.A{
							bson.D{{Key: "operationType", Value: "replace"}},
							bson.D{{Key: "operationType", Value: "update"}},
						}}},
						bson.D{{Key: "documentKey._id", Value: bson.D{{Key: "$in", Value: ids}}}},
					}},
				},
			}},
		})
	if err != nil {
		return fmt.Errorf("watch_matches: %w", err)
	}

	observable := model.MatchObservable{}
	observable.AddObserver(observer)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		for cs.Next(ctx) {
			var r map[string]interface{}
			err := cs.Decode(&r)
			if err != nil {
				return fmt.Errorf("cs_decode: %w", err)
			}

			dd, ok := r["documentKey"].(map[string]interface{})
			if !ok {
				return fmt.Errorf("invalid documentKey: %v", dd)
			}

			id, ok := dd["_id"].(string)
			if !ok {
				return fmt.Errorf("invalid _id: %v", id)
			}

			m, err := s.GetMatchFacts(ctx, id)
			if err != nil {
				return fmt.Errorf("get_match_facts: %w", err)
			}

			if err := observable.Notify(m); err != nil {
				return fmt.Errorf("observable_notify: %w", err)
			}
		}
		return nil
	})

	return g.Wait()
}

func (s *store) Ping(ctx context.Context) error {
	ctx, _ = context.WithTimeout(ctx, 1*time.Minute)
	return s.client.Ping(ctx, nil)
}

func (s *store) Close(ctx context.Context) error {
	return s.client.Disconnect(ctx)
}
