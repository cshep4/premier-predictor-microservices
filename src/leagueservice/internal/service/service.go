package league

import (
	"context"
	"errors"
	"fmt"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/rank"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/table"
	"sort"
	"sync"
	"time"

	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"

	"golang.org/x/sync/errgroup"
)

type service struct {
	store        Store
	userService  UserService
	overallTable LeagueTable
	time         Timer
}

func New(store Store, userService UserService, overallTable LeagueTable, time Timer) (*service, error) {
	switch {
	case store == nil:
		return nil, InvalidParameterError{Parameter: "store"}
	case userService == nil:
		return nil, InvalidParameterError{Parameter: "userService"}
	case overallTable == nil:
		return nil, InvalidParameterError{Parameter: "overallTable"}
	}

	return &service{
		store:        store,
		userService:  userService,
		overallTable: overallTable,
		time:         time,
	}, nil
}

func (s *service) RebuildOverallLeagueTable(ctx context.Context) error {
	overallTable, err := table.NewOverallTable(ctx, s.userService)
	if err != nil {
		// if there is an error building the league table,
		// restart service to trigger the rebuild again
		panic(err)
	}
	s.overallTable = overallTable

	return nil
}

func (s *service) GetUsersLeagueList(ctx context.Context, id string) (*model.StandingsOverview, error) {
	if id == "" {
		return nil, model.InvalidParameterError{Parameter: "id"}
	}

	var userLeagues []*model.League
	var userCount, overallRank int64

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		l, err := s.store.GetLeaguesByUserId(gCtx, id)
		if err != nil {
			return fmt.Errorf("get_leagues_by_user_id: %w", err)
		}
		userLeagues = l
		return nil
	})

	g.Go(func() error {
		r, ok := s.overallTable.Rank(id)
		if !ok {
			return errors.New("user not in league table")
		}
		overallRank = int64(r)
		return nil
	})

	g.Go(func() error {
		c, err := s.userService.GetUserCount(gCtx)
		if err != nil {
			return fmt.Errorf("get_user_count: %w", err)
		}
		userCount = c
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	g, ctx = errgroup.WithContext(ctx)

	leagueOverviews := make([]model.LeagueOverview, 0, len(userLeagues))
	var mu sync.Mutex

	for _, l := range userLeagues {
		league := l
		g.Go(func() error {
			rank, err := s.getLeagueRank(ctx, id, league.Users)
			if err != nil {
				return fmt.Errorf("get_league_rank: %w", err)
			}

			mu.Lock()
			defer mu.Unlock()

			leagueOverviews = append(leagueOverviews, model.LeagueOverview{
				Pin:        league.Pin,
				LeagueName: league.Name,
				Rank:       int64(rank),
			})

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &model.StandingsOverview{
		UserLeagues: leagueOverviews,
		OverallLeagueOverview: model.OverallLeagueOverview{
			Rank:      overallRank,
			UserCount: userCount,
		},
	}, nil
}

func (s *service) JoinUserLeague(ctx context.Context, id string, pin int64) (*model.LeagueOverview, error) {
	switch {
	case id == "":
		return nil, model.InvalidParameterError{Parameter: "id"}
	case pin == 0:
		return nil, model.InvalidParameterError{Parameter: "pin"}
	}

	league, err := s.store.GetLeagueByPin(ctx, pin)
	if err != nil {
		return nil, fmt.Errorf("get_league_by_pin: %w", err)
	}

	ids := append(league.Users, id)

	rank, err := s.getLeagueRank(ctx, id, ids)
	if err != nil {
		return nil, fmt.Errorf("get_league_rank: %w", err)
	}

	err = s.store.JoinLeague(ctx, pin, id)
	if err != nil {
		return nil, fmt.Errorf("join_league: %w", err)
	}

	return &model.LeagueOverview{
		LeagueName: league.Name,
		Pin:        league.Pin,
		Rank:       int64(rank),
	}, nil
}

func (s *service) getLeagueRank(ctx context.Context, id string, ids []string) (int, error) {
	users, err := s.userService.GetLeagueUsers(ctx, ids)
	if err != nil {
		return 0, fmt.Errorf("get_league_users: %w", err)
	}

	leagueTable, err := table.NewLeagueTable(users, &rank.Ranker{})
	if err != nil {
		return 0, fmt.Errorf("new_league_table: %w", err)
	}

	rank, ok := leagueTable.Rank(id)
	if !ok {
		return 0, errors.New("user not in league table")
	}

	return rank, nil
}

func (s *service) AddUserLeague(ctx context.Context, id, name string) (*model.League, error) {
	switch {
	case id == "":
		return nil, model.InvalidParameterError{Parameter: "id"}
	case name == "":
		return nil, model.InvalidParameterError{Parameter: "name"}
	}

	league := model.League{
		Pin:   s.time.Now().UnixNano()/int64(time.Millisecond) - timeSubtractor,
		Name:  name,
		Users: []string{id},
	}

	err := s.store.AddLeague(ctx, league)
	if err != nil {
		return nil, fmt.Errorf("add_league: %w", err)
	}

	return &league, nil
}

func (s *service) LeaveUserLeague(ctx context.Context, id string, pin int64) error {
	switch {
	case id == "":
		return model.InvalidParameterError{Parameter: "id"}
	case pin == 0:
		return model.InvalidParameterError{Parameter: "pin"}
	}

	err := s.store.LeaveLeague(ctx, pin, id)
	if err != nil {
		return fmt.Errorf("leave_league: %w", err)
	}

	return nil
}

func (s *service) RenameUserLeague(ctx context.Context, pin int64, name string) error {
	switch {
	case name == "":
		return model.InvalidParameterError{Parameter: "name"}
	case pin == 0:
		return model.InvalidParameterError{Parameter: "pin"}
	}

	err := s.store.RenameLeague(ctx, pin, name)
	if err != nil {
		return fmt.Errorf("rename_league: %w", err)
	}

	return nil
}

func (s *service) GetLeagueTable(ctx context.Context, pin int64) ([]model.LeagueUser, error) {
	if pin == 0 {
		return nil, model.InvalidParameterError{Parameter: "pin"}
	}

	league, err := s.store.GetLeagueByPin(ctx, pin)
	if err != nil {
		return nil, fmt.Errorf("get_league_by_pin: %w", err)
	}

	users, err := s.userService.GetLeagueUsers(ctx, league.Users)
	if err != nil {
		return nil, fmt.Errorf("get_league_users: %w", err)
	}

	leagueUsers := model.LeagueUserSlice(users)
	sort.Sort(leagueUsers)

	return leagueUsers, nil
}

func (s *service) GetOverallLeagueTable(context.Context) ([]model.LeagueUser, error) {
	return s.overallTable.LeagueTable(), nil
}
