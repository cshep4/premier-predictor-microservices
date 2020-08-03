package mongo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

const (
	db         = "user"
	collection = "user"
)

var (
	ErrCannotCreateObjectId = errors.New("cannot create objectId")
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

	s := &store{
		client: client,
	}

	if err := s.Ping(ctx); err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	if err := s.ensureIndexes(ctx); err != nil {
		return nil, fmt.Errorf("ensure_indexes: %w", err)
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
			name:   "email_idx",
			field:  []string{"email"},
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
			return fmt.Errorf("create_one: %w", err)
		}
	}

	return nil
}

func (s *store) GetUserById(ctx context.Context, id string) (*model.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrCannotCreateObjectId
	}

	var u userEntity

	err = s.client.
		Database(db).
		Collection(collection).
		FindOne(
			ctx,
			bson.D{
				{
					Key:   "_id",
					Value: objectId,
				},
			},
		).
		Decode(&u)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, model.ErrUserNotFound
		}

		return nil, fmt.Errorf("find_one: %w", err)
	}

	return toUser(u), nil
}

func (s *store) UpdateUserInfo(ctx context.Context, userInfo model.UserInfo) error {
	return s.updateUser(ctx, userInfo.Id, bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{
					Key:   "firstName",
					Value: userInfo.FirstName,
				},
				{
					Key:   "surname",
					Value: userInfo.Surname,
				},
				{
					Key:   "email",
					Value: userInfo.Email,
				},
			},
		},
	})
}

func (s *store) UpdatePassword(ctx context.Context, id, password string) error {
	return s.updateUser(ctx, id, bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{
					Key:   "password",
					Value: password,
				},
			},
		},
	})
}

func (s *store) UpdateSignature(ctx context.Context, id, signature string) error {
	return s.updateUser(ctx, id, bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{
					Key:   "signature",
					Value: signature,
				},
			},
		},
	})
}

func (s *store) updateUser(ctx context.Context, id string, update bson.D) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrCannotCreateObjectId
	}

	result, err := s.client.
		Database(db).
		Collection(collection).
		UpdateOne(
			ctx,
			bson.D{
				{
					Key:   "_id",
					Value: objectId,
				},
			},
			update,
		)
	if err != nil {
		return fmt.Errorf("update_one: %w", err)
	}

	if result.ModifiedCount == 0 {
		return model.ErrUserNotFound
	}

	return nil
}

func (s *store) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	return s.findUsers(ctx, bson.D{})
}

func (s *store) GetAllUsersByIds(ctx context.Context, ids []string) ([]*model.User, error) {
	var objectIds []primitive.ObjectID
	for _, id := range ids {
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return []*model.User{}, ErrCannotCreateObjectId
		}

		objectIds = append(objectIds, objectId)
	}

	return s.findUsers(
		ctx,
		bson.D{
			{
				Key: "_id",
				Value: bson.D{
					{
						Key:   "$in",
						Value: objectIds,
					},
				},
			},
		},
	)
}

func (s *store) findUsers(ctx context.Context, filter bson.D) ([]*model.User, error) {
	users := []*model.User{}

	cur, err := s.client.
		Database(db).
		Collection(collection).
		Find(
			ctx,
			filter,
		)
	if err != nil {
		return users, fmt.Errorf("find: %w", err)
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var u userEntity

		err := cur.Decode(&u)
		if err != nil {
			return nil, fmt.Errorf("decode: %w", err)
		}

		users = append(users, toUser(u))
	}

	return users, nil
}

func (s *store) IsEmailTakenByADifferentUser(ctx context.Context, id, email string) bool {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return true
	}

	var u userEntity

	err = s.client.
		Database(db).
		Collection(collection).
		FindOne(
			ctx,
			bson.D{
				{
					Key:   "email",
					Value: email,
				},
				{
					Key: "_id",
					Value: bson.D{
						{
							Key:   "$ne",
							Value: objectId,
						},
					},
				},
			},
		).
		Decode(&u)

	if err != nil && err == mongo.ErrNoDocuments {
		return false
	}

	return true
}

func (s *store) GetOverallRank(ctx context.Context, id string) (int64, error) {
	return s.getRank(ctx, id, bson.D{})
}

func (s *store) GetRankForGroup(ctx context.Context, id string, ids []string) (int64, error) {
	if !contains(ids, id) {
		return 0, errors.New("user not in group")
	}

	objectIds, err := toObjectIds(ids)
	if err != nil {
		return 0, err
	}

	return s.getRank(
		ctx,
		id,
		bson.D{
			{
				Key: "_id",
				Value: bson.D{
					{
						Key:   "$in",
						Value: objectIds,
					},
				},
			},
		},
	)
}

func (s *store) getRank(ctx context.Context, id string, filter bson.D) (int64, error) {
	cur, err := s.client.
		Database(db).
		Collection(collection).
		Find(
			ctx,
			filter,
			&options.FindOptions{
				Sort: bson.D{
					bson.E{Key: "score", Value: -1},
				},
			},
		)
	if err != nil {
		return 0, fmt.Errorf("find: %w", err)
	}

	defer cur.Close(ctx)

	previousScore := -1
	rank := int64(0)
	jointRanking := int64(1)
	for cur.Next(ctx) {
		var u userEntity

		err := cur.Decode(&u)
		if err != nil {
			return 0, fmt.Errorf("decode: %w", err)
		}

		if previousScore != u.Score {
			rank += jointRanking
			jointRanking = 1
		} else {
			jointRanking++
		}
		previousScore = u.Score

		if u.Id.Hex() == id {
			return rank, nil
		}
	}

	return 0, model.ErrUserNotFound
}

func (s *store) GetUserCount(ctx context.Context) (int64, error) {
	count, err := s.client.
		Database(db).
		Collection(collection).
		CountDocuments(
			ctx,
			bson.M{},
		)
	if err != nil {
		return 0, fmt.Errorf("count_documents: %w", err)
	}

	return count, nil
}

func (s *store) StoreUser(ctx context.Context, user model.User) (string, error) {
	user.Id = primitive.NewObjectID().Hex()

	entity, err := fromUser(user)
	if err != nil {
		return "", fmt.Errorf("from_user: %w", err)
	}

	_, err = s.client.
		Database(db).
		Collection(collection).
		InsertOne(ctx, entity)
	if err != nil {
		return "", fmt.Errorf("insert_one: %w", err)
	}

	return user.Id, nil
}

func (s *store) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var u userEntity

	err := s.client.
		Database(db).
		Collection(collection).
		FindOne(
			ctx,
			bson.D{
				{
					Key:   "email",
					Value: email,
				},
			},
		).
		Decode(&u)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, model.ErrUserNotFound
		}

		return nil, fmt.Errorf("find_one: %w", err)
	}

	return toUser(u), nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func toObjectIds(ids []string) (objectIds []primitive.ObjectID, err error) {
	for _, id := range ids {
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, ErrCannotCreateObjectId
		}

		objectIds = append(objectIds, objectId)
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
