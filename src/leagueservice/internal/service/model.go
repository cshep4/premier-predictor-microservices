package league

import (
	"context"
	"fmt"
	"time"

	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"
)

const (
	timeSubtractor = 1500000000000
)

type (
	LeagueTable interface {
		AddUser(user model.LeagueUser)
		Rank(id string) (int, bool)
		LeagueUser(id string) (*model.LeagueUser, bool)
		LeagueTable() []model.LeagueUser
	}

	UserService interface {
		GetAllUsers(ctx context.Context) ([]model.LeagueUser, error)
		GetLeagueUsers(ctx context.Context, ids []string) ([]model.LeagueUser, error)
		GetOverallRank(ctx context.Context, id string) (int64, error)
		GetLeagueRank(ctx context.Context, id string, ids []string) (int64, error)
		GetUserCount(ctx context.Context) (int64, error)
	}

	Store interface {
		GetLeagueByPin(ctx context.Context, pin int64) (*model.League, error)
		GetLeaguesByUserId(ctx context.Context, id string) ([]*model.League, error)
		AddLeague(ctx context.Context, league model.League) error
		JoinLeague(ctx context.Context, pin int64, id string) error
		LeaveLeague(ctx context.Context, pin int64, id string) error
		RenameLeague(ctx context.Context, pin int64, name string) error
	}

	Timer interface {
		Now() time.Time
	}

	// InvalidParameterError is returned when a required parameter passed to New is invalid.
	InvalidParameterError struct {
		Parameter string
	}
)

func (i InvalidParameterError) Error() string {
	return fmt.Sprintf("invalid parameter %s", i.Parameter)
}
