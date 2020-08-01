package user

import (
	"errors"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type userEntity struct {
	Id              *primitive.ObjectID `bson:"_id"`
	FirstName       string              `bson:"firstName"`
	Surname         string              `bson:"surname"`
	Email           string              `bson:"email"`
	Password        string              `bson:"password"`
	PredictedWinner string              `bson:"predictedWinner"`
	Signature       string              `bson:"signature"`
	Score           int                 `bson:"score"`
	Joined          time.Time           `bson:"joined"`
	Admin           bool                `bson:"admin"`
	AdFree          bool                `bson:"adFree"`
}

func fromUser(user model.User) (*userEntity, error) {
	id, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return nil, errors.New("cannot create objectId")
	}

	return &userEntity{
		Id:              &id,
		FirstName:       user.FirstName,
		Surname:         user.Surname,
		Email:           user.Email,
		Password:        user.Password,
		PredictedWinner: user.PredictedWinner,
		Signature: 		 user.Signature,
		Score:           user.Score,
		Joined:          user.Joined,
		Admin:           user.Admin,
		AdFree:          user.AdFree,
	}, nil
}

func toUser(user userEntity) *model.User {
	return &model.User{
		Id:              user.Id.Hex(),
		FirstName:       user.FirstName,
		Surname:         user.Surname,
		Email:           user.Email,
		Password:        user.Password,
		PredictedWinner: user.PredictedWinner,
		Signature: 		 user.Signature,
		Score:           user.Score,
		Joined:          user.Joined,
		Admin:           user.Admin,
		AdFree:          user.AdFree,
	}
}
