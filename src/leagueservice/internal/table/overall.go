package table

import (
	"context"
	"fmt"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/rank"
)

type (
	UserService interface {
		GetAllUsers(ctx context.Context) ([]model.LeagueUser, error)
	}

	overallTable struct {
		userService UserService
		leagueTable
	}
)

func NewOverallTable(ctx context.Context, userService UserService) (*overallTable, error) {
	if userService == nil {
		return nil, InvalidParameterError{Parameter: "userService"}
	}

	l := &overallTable{
		userService: userService,
	}

	return l, l.Build(ctx)
}

func (l *overallTable) Build(ctx context.Context) error {
	users, err := l.userService.GetAllUsers(ctx)
	if err != nil {
		return fmt.Errorf("get_all_users: %w", err)
	}

	lt, err := NewLeagueTable(users, &rank.Ranker{})
	if err != nil {
		return fmt.Errorf("new_league_table: %w", err)
	}
	l.leagueTable = *lt

	return nil
}
