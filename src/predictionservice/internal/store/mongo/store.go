package mongo

import (
	"context"
	"errors"
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"time"
)

const (
	db         = "prediction"
	collection = "prediction"
)

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

func (s *store) ensureIndexes(ctx context.Context) error {
	idxs := []struct {
		name   string
		field  []string
		unique bool
	}{
		{
			name:   "userId_idx",
			field:  []string{"userId"},
			unique: false,
		},
		{
			name:   "match_idx",
			field:  []string{"matchId"},
			unique: false,
		},
		{
			name:   "prediction_idx",
			field:  []string{"userId", "matchId"},
			unique: true,
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
			Indexes().
			CreateOne(
				ctx,
				mongo.IndexModel{
					Keys:    doc,
					Options: opts,
				},
			)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *store) GetPrediction(ctx context.Context, userId, matchId string) (*common.Prediction, error) {
	var p predictionEntity

	err := s.client.
		Database(db).
		Collection(collection).
		FindOne(
			ctx,
			bson.M{
				"userId":  userId,
				"matchId": matchId,
			},
		).
		Decode(&p)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, model.ErrPredictionNotFound
		}

		return nil, err
	}

	return toPrediction(&p), nil
}

func (s *store) GetPredictionsByUserId(ctx context.Context, id string) ([]common.Prediction, error) {
	cur, err := s.client.
		Database(db).
		Collection(collection).
		Find(
			ctx,
			bson.M{
				"userId": id,
			},
		)

	if err != nil {
		return nil, err
	}

	predictions := []common.Prediction{}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var m predictionEntity
		err := cur.Decode(&m)
		if err != nil {
			return nil, err
		}

		predictions = append(predictions, *toPrediction(&m))
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return predictions, nil
}

func (s *store) UpdatePredictions(ctx context.Context, predictions []common.Prediction) error {
	opts := options.FindOneAndReplaceOptions{}
	opts.SetUpsert(true)

	for _, p := range predictions {
		err := s.client.
			Database(db).
			Collection(collection).
			FindOneAndReplace(
				ctx,
				bson.M{
					"userId":  p.UserId,
					"matchId": p.MatchId,
				},
				fromPrediction(&p),
				&opts,
			)
		if err.Err() != nil && err.Err() != mongo.ErrNoDocuments {
			return err.Err()
		}
	}

	return nil
}

func (s *store) GetMatchPredictionSummary(ctx context.Context, id string) (homeWins int, draw int, awayWins int, err error) {
	cur, err := s.client.
		Database(db).
		Collection(collection).
		Find(
			ctx,
			bson.M{
				"matchId": id,
			},
		)

	if err != nil {
		return
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var m predictionEntity
		err = cur.Decode(&m)
		if err != nil {
			return
		}

		if m.HomeGoals > m.AwayGoals {
			homeWins++
		} else if m.HomeGoals < m.AwayGoals {
			awayWins++
		} else {
			draw++
		}
	}

	if err = cur.Err(); err != nil {
		return
	}

	return
}

func (s *store) Ping(ctx context.Context) error {
	ctx, _ = context.WithTimeout(ctx, 2*time.Second)
	return s.client.Ping(ctx, nil)
}

func (s *store) Close(ctx context.Context) error {
	return s.client.Disconnect(ctx)
}
