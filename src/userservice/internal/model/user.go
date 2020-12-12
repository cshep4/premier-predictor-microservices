package model

import (
	"time"

	"github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
)

type User struct {
	Id              string    `json:"id"`
	FirstName       string    `json:"firstName"`
	Surname         string    `json:"surname"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	PredictedWinner string    `json:"predictedWinner"`
	Signature       string    `json:"signature"`
	Score           int       `json:"score"`
	Joined          time.Time `json:"joined"`
	Admin           bool      `json:"admin"`
	AdFree          bool      `json:"adFree"`
}

func UserToGrpc(user *User) *model.User {
	return &model.User{
		Id:              user.Id,
		FirstName:       user.FirstName,
		Surname:         user.Surname,
		Email:           user.Email,
		PredictedWinner: user.PredictedWinner,
		Password:        user.Password,
		Signature:       user.Signature,
		Score:           int32(user.Score),
	}
}

func UserFromCreateReq(req model.CreateRequest) User {
	return User{
		FirstName:       req.FirstName,
		Surname:         req.Surname,
		Email:           req.Email,
		Password:        req.Password,
		PredictedWinner: req.PredictedWinner,
		Joined:          time.Now(),
	}
}

type UserInfo struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
}

type UpdatePassword struct {
	Id              string `json:"id"`
	OldPassword     string `json:"oldPassword"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

type UserScore struct {
	Score int `json:"score"`
}
