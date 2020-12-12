package model

import (
	pb "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
)

type LeagueUser struct {
	Id              string `json:"id"`
	FirstName       string `json:"firstName"`
	Surname         string `json:"surname"`
	PredictedWinner string `json:"predictedWinner"`
	Score           int    `json:"score"`
}

func LeagueUserFromGrpc(user *pb.User) LeagueUser {
	return LeagueUser{
		Id:              user.Id,
		FirstName:       user.FirstName,
		Surname:         user.Surname,
		PredictedWinner: user.PredictedWinner,
		Score:           int(user.Score),
	}
}

type LeagueUserSlice []LeagueUser

func (l LeagueUserSlice) Len() int {
	return len(l)
}

func (l LeagueUserSlice) Less(i, j int) bool {
	return l[i].Score > l[j].Score
}

func (l LeagueUserSlice) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

type League struct {
	Pin   int64    `json:"pin"`
	Name  string   `json:"name"`
	Users []string `json:"users"`
}

type LeagueOverview struct {
	LeagueName string `json:"leagueName"`
	Pin        int64  `json:"pin"`
	Rank       int64  `json:"rank"`
}

type OverallLeagueOverview struct {
	Rank      int64 `json:"rank"`
	UserCount int64 `json:"userCount"`
}

type StandingsOverview struct {
	OverallLeagueOverview OverallLeagueOverview `json:"overallLeagueOverview"`
	UserLeagues           []LeagueOverview      `json:"userLeagues"`
}
