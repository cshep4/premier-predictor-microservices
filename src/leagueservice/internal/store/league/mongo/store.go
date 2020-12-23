package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

const (
	db         = "league"
	collection = "league"
)

type (
	store struct {
		client *mongo.Client
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
		client: client,
	}

	if err := s.ping(ctx); err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	if err := s.ensureIndexes(ctx); err != nil {
		return nil, fmt.Errorf("ensure_indexes: %w", err)
	}

	return &s, nil
}

func (s *store) ensureIndexes(ctx context.Context) error {
	idxs := []struct {
		name   string
		field  []string
		unique bool
	}{
		{
			name:   "users_idx",
			field:  []string{"users"},
			unique: false,
		},
	}

	for _, i := range idxs {
		var doc bsonx.Doc
		for _, f := range i.field {
			doc = append(doc, bsonx.Elem{Key: f, Value: bsonx.Int64(1)})
		}

		opts := options.Index().
			SetName(i.name).
			SetUnique(i.unique).
			SetSparse(false).
			SetBackground(true)

		_, err := s.client.
			Database(db).
			Collection(collection).
			Indexes().CreateOne(
			ctx,
			mongo.IndexModel{
				Keys:    doc,
				Options: opts,
			},
		)

		if err != nil {
			return fmt.Errorf("create_one: %w", err)
		}
	}

	return nil
}

func (s *store) GetLeagueByPin(ctx context.Context, pin int64) (*model.League, error) {
	var u leagueEntity

	err := s.client.
		Database(db).
		Collection(collection).
		FindOne(
			ctx,
			bson.D{
				{
					Key:   "_id",
					Value: pin,
				},
			},
		).
		Decode(&u)

	switch {
	case err == mongo.ErrNoDocuments:
		return nil, model.ErrLeagueNotFound
	case err != nil:
		return nil, fmt.Errorf("find_one: %w", err)
	}

	return toLeague(u), nil
}

func (s *store) GetLeaguesByUserId(ctx context.Context, id string) ([]*model.League, error) {
	users := []*model.League{}

	cur, err := s.client.
		Database(db).
		Collection(collection).
		Find(
			ctx,
			bson.D{
				{
					Key:   "users",
					Value: id,
				},
			},
		)
	if err != nil {
		return nil, fmt.Errorf("find: %w", err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var l leagueEntity

		err := cur.Decode(&l)
		if err != nil {
			return nil, fmt.Errorf("decode: %w", err)
		}

		users = append(users, toLeague(l))
	}

	return users, nil
}

func (s *store) AddLeague(ctx context.Context, league model.League) error {
	_, err := s.client.
		Database(db).
		Collection(collection).
		InsertOne(
			ctx,
			fromLeague(league),
		)
	if err != nil {
		return fmt.Errorf("insert_one: %w", err)
	}

	return nil
}

func (s *store) JoinLeague(ctx context.Context, pin int64, id string) error {
	return s.editLeague(ctx,
		pin,
		bson.D{
			{
				Key: "$addToSet",
				Value: bson.D{
					{
						Key:   "users",
						Value: id,
					},
				},
			},
		},
	)
}

func (s *store) LeaveLeague(ctx context.Context, pin int64, id string) error {
	return s.editLeague(ctx,
		pin,
		bson.D{
			{
				Key: "$pull",
				Value: bson.D{
					{
						Key:   "users",
						Value: id,
					},
				},
			},
		},
	)
}

func (s *store) RenameLeague(ctx context.Context, pin int64, name string) error {
	return s.editLeague(ctx,
		pin,
		bson.D{
			{
				Key: "$set",
				Value: bson.D{
					{
						Key:   "name",
						Value: name,
					},
				},
			},
		},
	)
}

func (s *store) editLeague(ctx context.Context, pin int64, update bson.D) error {
	res, err := s.client.
		Database(db).
		Collection(collection).
		UpdateOne(
			ctx,
			bson.D{
				{
					Key:   "_id",
					Value: pin,
				},
			},
			update,
		)

	if err != nil {
		return fmt.Errorf("update_one: %w", err)
	}

	if res.MatchedCount == 0 {
		return model.ErrLeagueNotFound
	}

	return nil
}

func (s *store) ping(ctx context.Context) error {
	ctx, _ = context.WithTimeout(ctx, 1*time.Minute)
	return s.client.Ping(ctx, nil)
}

func (s *store) Close(ctx context.Context) error {
	return s.client.Disconnect(ctx)
}
