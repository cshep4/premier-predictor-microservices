package table

import (
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"
	"sort"
)

type leagueTable struct {
	users  map[string]model.LeagueUser
	table  []model.LeagueUser
	ranker Ranker
}

func NewLeagueTable(users []model.LeagueUser, ranker Ranker) (*leagueTable, error) {
	switch {
	case users == nil:
		return nil, InvalidParameterError{Parameter: "users"}
	case len(users) == 0:
		return nil, InvalidParameterError{Parameter: "users"}
	case ranker == nil:
		return nil, InvalidParameterError{Parameter: "ranker"}
	}

	l := &leagueTable{
		ranker: ranker,
	}
	l.build(users)

	return l, nil
}

func (l *leagueTable) build(users []model.LeagueUser) {
	leagueTable := leagueUserSlice(users)
	sort.Sort(leagueTable)

	l.table = leagueTable

	l.ranker.Clear()
	l.users = make(map[string]model.LeagueUser)

	for _, u := range l.table {
		l.users[u.Id] = u
		l.ranker.Insert(u.Score)
	}
}

func (l *leagueTable) AddUser(user model.LeagueUser) {
	l.table = append(l.table, user)
	l.build(l.table)
}

func (l *leagueTable) Rank(id string) (int, bool) {
	u, ok := l.users[id]
	if !ok {
		return l.lastPlace()
	}

	return l.ranker.Rank(u.Score)
}

// lastPlace returns the last position in the table, this will either be:
// the same as the current last user if their score is 0
// if their score isn't zero, it will be one place below
func (l *leagueTable) lastPlace() (int, bool) {
	lastUser := l.table[len(l.users)-1]

	if lastUser.Score == 0 {
		return l.ranker.Rank(lastUser.Score)
	}

	return len(l.users) + 1, true
}

func (l *leagueTable) LeagueUser(id string) (*model.LeagueUser, bool) {
	u, ok := l.users[id]
	if !ok {
		return nil, false
	}

	return &u, true
}

func (l *leagueTable) LeagueTable() []model.LeagueUser {
	return l.table
}
