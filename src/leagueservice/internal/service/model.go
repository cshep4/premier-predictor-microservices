package league

import (
	"context"
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

	UserStore interface {
		Get(ctx context.Context, id string) (model.LeagueUser, error)
		Count(ctx context.Context) (int64, error)
		List(ctx context.Context, ids []string) ([]model.LeagueUser, error)
		ListAll(ctx context.Context) ([]model.LeagueUser, error)
	}

	LeagueStore interface {
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
)
