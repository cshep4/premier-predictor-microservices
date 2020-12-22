package model

import (
	"sort"
)

type LeagueUser struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	PredictedWinner string `json:"predictedWinner"`
	Rank            int64  `json:"rank"`
	Score           int    `json:"score"`
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

func (l LeagueUserSlice) Rank(id string) (int64, bool) {
	sort.Sort(l)

	rank := int64(0)
	previousScore := -1
	usersOnScore := int64(1)

	for _, u := range l {
		if previousScore != u.Score {
			rank += usersOnScore
			usersOnScore = 1
		} else {
			usersOnScore++
		}
		previousScore = u.Score

		if id == u.ID {
			return rank, true
		}
	}

	return 0, false
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
