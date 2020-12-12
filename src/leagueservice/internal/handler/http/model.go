package http

import (
	"context"
	"fmt"

	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"
)

type (
	addLeagueRequest struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}

	leagueRequest struct {
		Id  string `json:"id"`
		Pin int64  `json:"pin"`
	}

	renameRequest struct {
		Pin  int64  `json:"pin"`
		Name string `json:"name"`
	}

	serverError struct {
		Message string `json:"message"`
	}

	Service interface {
		RebuildOverallLeagueTable(ctx context.Context) error
		GetUsersLeagueList(ctx context.Context, id string) (*model.StandingsOverview, error)
		JoinUserLeague(ctx context.Context, id string, pin int64) (*model.LeagueOverview, error)
		AddUserLeague(ctx context.Context, id, name string) (*model.League, error)
		LeaveUserLeague(ctx context.Context, id string, pin int64) error
		RenameUserLeague(ctx context.Context, pin int64, name string) error
		GetLeagueTable(ctx context.Context, pin int64) ([]model.LeagueUser, error)
		GetOverallLeagueTable(ctx context.Context) ([]model.LeagueUser, error)
	}

	// InvalidParameterError is returned when a required parameter passed to New is invalid.
	InvalidParameterError struct {
		Parameter string
	}
)

func (i InvalidParameterError) Error() string {
	return fmt.Sprintf("invalid parameter %s", i.Parameter)
}
