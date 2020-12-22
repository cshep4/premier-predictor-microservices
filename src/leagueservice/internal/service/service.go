package league

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"

	"golang.org/x/sync/errgroup"
)

type service struct {
	leagueStore LeagueStore
	userStore   UserStore
	time        Timer
}

func New(
	leagueStore LeagueStore,
	userStore UserStore,
	time Timer,
) (*service, error) {
	switch {
	case leagueStore == nil:
		return nil, model.InvalidParameterError{Parameter: "leagueStore"}
	case userStore == nil:
		return nil, model.InvalidParameterError{Parameter: "userStore"}
	case time == nil:
		return nil, model.InvalidParameterError{Parameter: "time"}
	}

	return &service{
		leagueStore: leagueStore,
		userStore:   userStore,
		time:        time,
	}, nil
}

func (s *service) GetUsersLeagueList(ctx context.Context, id string) (*model.StandingsOverview, error) {
	if id == "" {
		return nil, model.InvalidParameterError{Parameter: "id"}
	}

	var userLeagues []*model.League
	var userCount, overallRank int64

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		l, err := s.leagueStore.GetLeaguesByUserId(gCtx, id)
		if err != nil {
			return fmt.Errorf("could not get user leagues: %w", err)
		}
		userLeagues = l
		return nil
	})

	g.Go(func() error {
		lu, err := s.userStore.Get(gCtx, id)
		if err != nil {
			return fmt.Errorf("could not get user: %w", err)
		}
		overallRank = lu.Rank
		return nil
	})

	g.Go(func() error {
		c, err := s.userStore.Count(gCtx)
		if err != nil {
			return fmt.Errorf("could not get overall count: %w", err)
		}
		userCount = c
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	leagueOverviews, err := s.buildLeagueOverview(ctx, id, userLeagues)
	if err != nil {
		return nil, fmt.Errorf("cannot build league overview: %w", err)
	}

	return &model.StandingsOverview{
		UserLeagues: leagueOverviews,
		OverallLeagueOverview: model.OverallLeagueOverview{
			Rank:      overallRank,
			UserCount: userCount,
		},
	}, nil
}

func (s *service) buildLeagueOverview(ctx context.Context, id string, userLeagues []*model.League) ([]model.LeagueOverview, error) {
	g, ctx := errgroup.WithContext(ctx)

	leagueOverviews := make([]model.LeagueOverview, 0, len(userLeagues))
	var mu sync.Mutex

	for _, l := range userLeagues {
		league := l
		g.Go(func() error {
			rank, err := s.leagueRank(ctx, id, league.Users)
			if err != nil {
				return fmt.Errorf("cannot get league rank: %w", err)
			}

			mu.Lock()
			defer mu.Unlock()

			leagueOverviews = append(leagueOverviews, model.LeagueOverview{
				Pin:        league.Pin,
				LeagueName: league.Name,
				Rank:       rank,
			})

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return leagueOverviews, nil
}

func (s *service) leagueRank(ctx context.Context, id string, users []string) (int64, error) {
	lu, err := s.userStore.List(ctx, users)
	if err != nil {
		return 0, fmt.Errorf("could not get league users: %w", err)
	}

	rank, ok := model.LeagueUserSlice(lu).Rank(id)
	if !ok {
		return 0, errors.New("user not in league")
	}

	return rank, nil
}

func (s *service) JoinUserLeague(ctx context.Context, id string, pin int64) (*model.LeagueOverview, error) {
	switch {
	case id == "":
		return nil, model.InvalidParameterError{Parameter: "id"}
	case pin == 0:
		return nil, model.InvalidParameterError{Parameter: "pin"}
	}

	league, err := s.leagueStore.GetLeagueByPin(ctx, pin)
	if err != nil {
		return nil, fmt.Errorf("cannot get league by pin: %w", err)
	}

	ids := append(league.Users, id)

	rank, err := s.leagueRank(ctx, id, ids)
	if err != nil {
		return nil, fmt.Errorf("cannot get league rank: %w", err)
	}

	err = s.leagueStore.JoinLeague(ctx, pin, id)
	if err != nil {
		return nil, fmt.Errorf("cannot join league: %w", err)
	}

	return &model.LeagueOverview{
		LeagueName: league.Name,
		Pin:        league.Pin,
		Rank:       rank,
	}, nil
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

	err := s.leagueStore.AddLeague(ctx, league)
	if err != nil {
		return nil, fmt.Errorf("cannot add league: %w", err)
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

	err := s.leagueStore.LeaveLeague(ctx, pin, id)
	if err != nil {
		return fmt.Errorf("cannot leave league: %w", err)
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

	err := s.leagueStore.RenameLeague(ctx, pin, name)
	if err != nil {
		return fmt.Errorf("cannot rename league: %w", err)
	}

	return nil
}

func (s *service) GetLeagueTable(ctx context.Context, pin int64) ([]model.LeagueUser, error) {
	if pin == 0 {
		return nil, model.InvalidParameterError{Parameter: "pin"}
	}

	league, err := s.leagueStore.GetLeagueByPin(ctx, pin)
	if err != nil {
		return nil, fmt.Errorf("cannot get league by pin: %w", err)
	}

	users, err := s.userStore.List(ctx, league.Users)
	if err != nil {
		return nil, fmt.Errorf("cannot get league users: %w", err)
	}

	return users, nil
}

func (s *service) GetOverallLeagueTable(ctx context.Context) ([]model.LeagueUser, error) {
	users, err := s.userStore.ListAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot get all league users: %w", err)
	}

	return users, nil
}
