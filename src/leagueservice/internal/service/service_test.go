package league_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/mock/store"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/mock/time"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/service"
)

var testErr = errors.New("error")

const (
	id1   = "🆔"
	id2   = "⚽️"
	id3   = "🤙️"
	pin1  = int64(12345)
	pin2  = int64(67890)
	name1 = "League of champions"
	name2 = "🏆🏆🏆🏆🏆🏆"

	count       = int64(1234)
	overallRank = int64(123)
)

var (
	users1 = []string{id1, id2}
	users2 = []string{id1, id3}
)

func TestService_GetUsersLeagueList(t *testing.T) {
	leagues := []*model.League{
		{
			Pin:   pin1,
			Name:  name1,
			Users: users1,
		},
		{
			Pin:   pin2,
			Name:  name2,
			Users: users2,
		},
	}

	testErr := errors.New("error")

	t.Run("returns error if there is a problem getting info for a league from db", func(t *testing.T) {
		ctrl, ctx := gomock.WithContext(context.Background(), t)
		defer ctrl.Finish()

		leagueStore := store_mock.NewMockLeagueStore(ctrl)
		userStore := store_mock.NewMockUserStore(ctrl)
		timer := time_mock.NewMockTimer(ctrl)

		service, err := league.New(leagueStore, userStore, timer)
		require.NoError(t, err)

		leagueStore.EXPECT().GetLeaguesByUserId(gomock.Any(), id1).Return(nil, testErr)
		userStore.EXPECT().Count(gomock.Any()).MaxTimes(1).Return(count, nil)
		userStore.EXPECT().Get(gomock.Any(), id1).MaxTimes(1).Return(model.LeagueUser{Rank: 123}, nil)

		result, err := service.GetUsersLeagueList(ctx, id1)
		require.Error(t, err)

		assert.Empty(t, result)
	})

	t.Run("returns error if there is a problem getting overall rank", func(t *testing.T) {
		ctrl, ctx := gomock.WithContext(context.Background(), t)
		defer ctrl.Finish()

		leagueStore := store_mock.NewMockLeagueStore(ctrl)
		userStore := store_mock.NewMockUserStore(ctrl)
		timer := time_mock.NewMockTimer(ctrl)

		service, err := league.New(leagueStore, userStore, timer)
		require.NoError(t, err)

		leagueStore.EXPECT().GetLeaguesByUserId(gomock.Any(), id1).MaxTimes(1).Return(leagues, nil)
		userStore.EXPECT().Count(gomock.Any()).MaxTimes(1).Return(count, nil)
		userStore.EXPECT().Get(gomock.Any(), id1).Return(model.LeagueUser{}, testErr)

		result, err := service.GetUsersLeagueList(ctx, id1)
		require.Error(t, err)

		assert.Empty(t, result)
	})

	t.Run("returns error if there is a problem getting overall user count", func(t *testing.T) {
		ctrl, ctx := gomock.WithContext(context.Background(), t)
		defer ctrl.Finish()

		leagueStore := store_mock.NewMockLeagueStore(ctrl)
		userStore := store_mock.NewMockUserStore(ctrl)
		timer := time_mock.NewMockTimer(ctrl)

		service, err := league.New(leagueStore, userStore, timer)
		require.NoError(t, err)

		leagueStore.EXPECT().GetLeaguesByUserId(gomock.Any(), id1).MaxTimes(1).Return(leagues, nil)
		userStore.EXPECT().Count(gomock.Any()).Return(int64(0), testErr)
		userStore.EXPECT().Get(gomock.Any(), id1).MaxTimes(1).Return(model.LeagueUser{Rank: 123}, nil)

		result, err := service.GetUsersLeagueList(ctx, id1)
		require.Error(t, err)

		assert.Empty(t, result)
	})

	t.Run("returns error if there is a problem getting rank for a league", func(t *testing.T) {
		ctrl, ctx := gomock.WithContext(context.Background(), t)
		defer ctrl.Finish()

		leagueStore := store_mock.NewMockLeagueStore(ctrl)
		userStore := store_mock.NewMockUserStore(ctrl)
		timer := time_mock.NewMockTimer(ctrl)

		service, err := league.New(leagueStore, userStore, timer)
		require.NoError(t, err)

		leagueStore.EXPECT().GetLeaguesByUserId(gomock.Any(), id1).MaxTimes(1).Return(leagues, nil)
		userStore.EXPECT().Count(gomock.Any()).MaxTimes(1).Return(count, nil)
		userStore.EXPECT().Get(gomock.Any(), id1).MaxTimes(1).Return(model.LeagueUser{Rank: 123}, nil)

		u1 := model.LeagueUser{
			Score: 10,
		}
		u2 := model.LeagueUser{
			Score: 30,
		}
		u3 := model.LeagueUser{
			ID:    id1,
			Score: 20,
		}

		userStore.EXPECT().List(gomock.Any(), users1).Return(nil, testErr)
		userStore.EXPECT().List(gomock.Any(), users2).MaxTimes(1).Return([]model.LeagueUser{u1, u2, u3}, nil)

		result, err := service.GetUsersLeagueList(ctx, id1)
		require.Error(t, err)

		assert.Empty(t, result)
	})

	t.Run("returns standings overview from the user", func(t *testing.T) {
		ctrl, ctx := gomock.WithContext(context.Background(), t)
		defer ctrl.Finish()

		leagueStore := store_mock.NewMockLeagueStore(ctrl)
		userStore := store_mock.NewMockUserStore(ctrl)
		timer := time_mock.NewMockTimer(ctrl)

		service, err := league.New(leagueStore, userStore, timer)
		require.NoError(t, err)

		leagueStore.EXPECT().GetLeaguesByUserId(gomock.Any(), id1).MaxTimes(1).Return(leagues, nil)
		userStore.EXPECT().Count(gomock.Any()).MaxTimes(1).Return(count, nil)
		userStore.EXPECT().Get(gomock.Any(), id1).MaxTimes(1).Return(model.LeagueUser{Rank: 123}, nil)

		u1 := model.LeagueUser{
			Score: 10,
		}
		u2 := model.LeagueUser{
			Score: 30,
		}
		u3 := model.LeagueUser{
			ID:    id1,
			Score: 20,
		}

		userStore.EXPECT().List(gomock.Any(), users1).Return([]model.LeagueUser{u2, u3}, nil)
		userStore.EXPECT().List(gomock.Any(), users2).Return([]model.LeagueUser{u1, u3}, nil)

		expectedResult := &model.StandingsOverview{
			OverallLeagueOverview: model.OverallLeagueOverview{
				Rank:      overallRank,
				UserCount: count,
			},
			UserLeagues: []model.LeagueOverview{
				{
					Rank:       2,
					LeagueName: name1,
					Pin:        pin1,
				},
				{
					Rank:       1,
					LeagueName: name2,
					Pin:        pin2,
				},
			},
		}

		result, err := service.GetUsersLeagueList(ctx, id1)
		require.NoError(t, err)

		assert.Equal(t, expectedResult.OverallLeagueOverview, result.OverallLeagueOverview)

		if result.UserLeagues[0].Pin == pin1 {
			assert.Equal(t, expectedResult.UserLeagues[0], result.UserLeagues[0])
			assert.Equal(t, expectedResult.UserLeagues[1], result.UserLeagues[1])
		} else {
			assert.Equal(t, expectedResult.UserLeagues[1], result.UserLeagues[0])
			assert.Equal(t, expectedResult.UserLeagues[0], result.UserLeagues[1])
		}
	})
}

func TestService_JoinUserLeague(t *testing.T) {
	t.Run("Returns error if it cannot get league info", func(t *testing.T) {
		ctrl, ctx := gomock.WithContext(context.Background(), t)
		defer ctrl.Finish()

		leagueStore := store_mock.NewMockLeagueStore(ctrl)
		userStore := store_mock.NewMockUserStore(ctrl)
		timer := time_mock.NewMockTimer(ctrl)

		service, err := league.New(leagueStore, userStore, timer)
		require.NoError(t, err)

		leagueStore.EXPECT().GetLeagueByPin(ctx, pin1).Return(nil, testErr)

		result, err := service.JoinUserLeague(ctx, id1, pin1)
		require.Error(t, err)

		assert.True(t, errors.Is(err, testErr))
		assert.Empty(t, result)
	})

	t.Run("Returns error if it cannot get league rank", func(t *testing.T) {
		ctrl, ctx := gomock.WithContext(context.Background(), t)
		defer ctrl.Finish()

		leagueStore := store_mock.NewMockLeagueStore(ctrl)
		userStore := store_mock.NewMockUserStore(ctrl)
		timer := time_mock.NewMockTimer(ctrl)

		service, err := league.New(leagueStore, userStore, timer)
		require.NoError(t, err)

		league := &model.League{
			Pin:   pin1,
			Name:  name1,
			Users: []string{id2, id3},
		}

		users := []string{id2, id3, id1}

		leagueStore.EXPECT().GetLeagueByPin(ctx, pin1).Return(league, nil)
		userStore.EXPECT().List(ctx, users).Return(nil, testErr)

		result, err := service.JoinUserLeague(ctx, id1, pin1)
		require.Error(t, err)

		assert.True(t, errors.Is(err, testErr))
		assert.Empty(t, result)
	})

	t.Run("Returns error if user cannot join league", func(t *testing.T) {
		ctrl, ctx := gomock.WithContext(context.Background(), t)
		defer ctrl.Finish()

		leagueStore := store_mock.NewMockLeagueStore(ctrl)
		userStore := store_mock.NewMockUserStore(ctrl)
		timer := time_mock.NewMockTimer(ctrl)

		service, err := league.New(leagueStore, userStore, timer)
		require.NoError(t, err)

		u := model.LeagueUser{
			ID: id1,
		}

		league := &model.League{
			Pin:   pin1,
			Name:  name1,
			Users: []string{id2, id3},
		}

		users := []string{id2, id3, id1}

		leagueStore.EXPECT().GetLeagueByPin(ctx, pin1).Return(league, nil)
		userStore.EXPECT().List(ctx, users).Return([]model.LeagueUser{u}, nil)
		leagueStore.EXPECT().JoinLeague(ctx, pin1, id1).Return(testErr)

		result, err := service.JoinUserLeague(ctx, id1, pin1)
		require.Error(t, err)

		assert.True(t, errors.Is(err, testErr))
		assert.Empty(t, result)
	})

	t.Run("Gets league info and joins", func(t *testing.T) {
		ctrl, ctx := gomock.WithContext(context.Background(), t)
		defer ctrl.Finish()

		leagueStore := store_mock.NewMockLeagueStore(ctrl)
		userStore := store_mock.NewMockUserStore(ctrl)
		timer := time_mock.NewMockTimer(ctrl)

		service, err := league.New(leagueStore, userStore, timer)
		require.NoError(t, err)

		leagueOverview := &model.LeagueOverview{
			Rank:       1,
			LeagueName: name1,
			Pin:        pin1,
		}

		leagueStore.EXPECT().GetLeagueByPin(ctx, pin1).Return(&model.League{
			Pin:   pin1,
			Name:  name1,
			Users: []string{id2, id3},
		}, nil)
		userStore.EXPECT().List(ctx, []string{id2, id3, id1}).Return([]model.LeagueUser{{ID: id1}}, nil)
		leagueStore.EXPECT().JoinLeague(ctx, pin1, id1).Return(nil)

		result, err := service.JoinUserLeague(ctx, id1, pin1)
		require.NoError(t, err)

		assert.Equal(t, leagueOverview, result)
	})
}

func TestService_AddUserLeague(t *testing.T) {
	t.Run("returns error if league cannot be added", func(t *testing.T) {
		ctrl, ctx := gomock.WithContext(context.Background(), t)
		defer ctrl.Finish()

		leagueStore := store_mock.NewMockLeagueStore(ctrl)
		userStore := store_mock.NewMockUserStore(ctrl)
		timer := time_mock.NewMockTimer(ctrl)

		service, err := league.New(leagueStore, userStore, timer)
		require.NoError(t, err)

		const (
			currentTime = 1512345678912
			pin         = int64(12345678912)
		)

		timer.EXPECT().Now().AnyTimes().Return(time.Unix(0, currentTime*int64(time.Millisecond)))
		leagueStore.EXPECT().AddLeague(ctx, model.League{
			Pin:   pin,
			Name:  name1,
			Users: []string{id1},
		}).Return(testErr)

		result, err := service.AddUserLeague(ctx, id1, name1)
		require.Error(t, err)

		assert.True(t, errors.Is(err, testErr))
		assert.Empty(t, result)
	})

	t.Run("returns league once added", func(t *testing.T) {
		ctrl, ctx := gomock.WithContext(context.Background(), t)
		defer ctrl.Finish()

		leagueStore := store_mock.NewMockLeagueStore(ctrl)
		userStore := store_mock.NewMockUserStore(ctrl)
		timer := time_mock.NewMockTimer(ctrl)

		service, err := league.New(leagueStore, userStore, timer)
		require.NoError(t, err)

		const (
			currentTime = 1512345678912
			pin         = int64(12345678912)
		)

		league := model.League{
			Pin:   pin,
			Name:  name1,
			Users: []string{id1},
		}

		timer.EXPECT().Now().AnyTimes().Return(time.Unix(0, currentTime*int64(time.Millisecond)))
		leagueStore.EXPECT().AddLeague(ctx, league).Return(nil)

		result, err := service.AddUserLeague(ctx, id1, name1)
		require.NoError(t, err)

		assert.Equal(t, &league, result)
	})
}

func TestService_LeaveUserLeague(t *testing.T) {
	t.Run("Returns error if there is a problem", func(t *testing.T) {
		ctrl, ctx := gomock.WithContext(context.Background(), t)
		defer ctrl.Finish()

		leagueStore := store_mock.NewMockLeagueStore(ctrl)
		userStore := store_mock.NewMockUserStore(ctrl)
		timer := time_mock.NewMockTimer(ctrl)

		service, err := league.New(leagueStore, userStore, timer)
		require.NoError(t, err)

		leagueStore.EXPECT().LeaveLeague(ctx, pin1, id1).Return(testErr)

		err = service.LeaveUserLeague(ctx, id1, pin1)
		require.Error(t, err)

		assert.True(t, errors.Is(err, testErr))
	})

	t.Run("Adds user to league", func(t *testing.T) {
		ctrl, ctx := gomock.WithContext(context.Background(), t)
		defer ctrl.Finish()

		leagueStore := store_mock.NewMockLeagueStore(ctrl)
		userStore := store_mock.NewMockUserStore(ctrl)
		timer := time_mock.NewMockTimer(ctrl)

		service, err := league.New(leagueStore, userStore, timer)
		require.NoError(t, err)

		leagueStore.EXPECT().LeaveLeague(ctx, pin1, id1).Return(nil)

		err = service.LeaveUserLeague(ctx, id1, pin1)
		require.NoError(t, err)
	})
}

func TestService_RenameUserLeague(t *testing.T) {
	t.Run("Returns error if there is a problem", func(t *testing.T) {
		ctrl, ctx := gomock.WithContext(context.Background(), t)
		defer ctrl.Finish()

		leagueStore := store_mock.NewMockLeagueStore(ctrl)
		userStore := store_mock.NewMockUserStore(ctrl)
		timer := time_mock.NewMockTimer(ctrl)

		service, err := league.New(leagueStore, userStore, timer)
		require.NoError(t, err)

		leagueStore.EXPECT().RenameLeague(ctx, pin1, name2).Return(testErr)

		err = service.RenameUserLeague(ctx, pin1, name2)
		require.Error(t, err)

		assert.True(t, errors.Is(err, testErr))
	})

	t.Run("Renames league", func(t *testing.T) {
		ctrl, ctx := gomock.WithContext(context.Background(), t)
		defer ctrl.Finish()

		leagueStore := store_mock.NewMockLeagueStore(ctrl)
		userStore := store_mock.NewMockUserStore(ctrl)
		timer := time_mock.NewMockTimer(ctrl)

		service, err := league.New(leagueStore, userStore, timer)
		require.NoError(t, err)

		leagueStore.EXPECT().RenameLeague(ctx, pin1, name2).Return(nil)

		err = service.RenameUserLeague(ctx, pin1, name2)
		require.NoError(t, err)
	})
}

func TestService_GetLeagueTable(t *testing.T) {
	t.Run("Returns error if there is a problem getting league", func(t *testing.T) {
		ctrl, ctx := gomock.WithContext(context.Background(), t)
		defer ctrl.Finish()

		leagueStore := store_mock.NewMockLeagueStore(ctrl)
		userStore := store_mock.NewMockUserStore(ctrl)
		timer := time_mock.NewMockTimer(ctrl)

		service, err := league.New(leagueStore, userStore, timer)
		require.NoError(t, err)

		leagueStore.EXPECT().GetLeagueByPin(ctx, pin1).Return(nil, testErr)

		result, err := service.GetLeagueTable(ctx, pin1)
		require.Error(t, err)

		assert.True(t, errors.Is(err, testErr))
		assert.Empty(t, result)
	})

	t.Run("Returns error if there is a problem getting users", func(t *testing.T) {
		ctrl, ctx := gomock.WithContext(context.Background(), t)
		defer ctrl.Finish()

		leagueStore := store_mock.NewMockLeagueStore(ctrl)
		userStore := store_mock.NewMockUserStore(ctrl)
		timer := time_mock.NewMockTimer(ctrl)

		service, err := league.New(leagueStore, userStore, timer)
		require.NoError(t, err)

		leagueStore.EXPECT().GetLeagueByPin(ctx, pin1).Return(&model.League{
			Pin:   pin1,
			Users: users1,
		}, nil)
		userStore.EXPECT().List(ctx, users1).Return(nil, testErr)

		result, err := service.GetLeagueTable(ctx, pin1)
		require.Error(t, err)

		assert.True(t, errors.Is(err, testErr))
		assert.Empty(t, result)
	})

	t.Run("Gets users and returns them sorted by points", func(t *testing.T) {
		ctrl, ctx := gomock.WithContext(context.Background(), t)
		defer ctrl.Finish()

		var (
			u1 = model.LeagueUser{Score: 30}
			u2 = model.LeagueUser{Score: 20}
			u3 = model.LeagueUser{Score: 10}
		)

		leagueStore := store_mock.NewMockLeagueStore(ctrl)
		userStore := store_mock.NewMockUserStore(ctrl)
		timer := time_mock.NewMockTimer(ctrl)

		service, err := league.New(leagueStore, userStore, timer)
		require.NoError(t, err)

		leagueStore.EXPECT().GetLeagueByPin(ctx, pin1).Return(&model.League{
			Pin:   pin1,
			Users: users1,
		}, nil)
		userStore.EXPECT().List(ctx, users1).Return([]model.LeagueUser{u1, u2, u3}, nil)

		result, err := service.GetLeagueTable(ctx, pin1)
		require.NoError(t, err)

		assert.Equal(t, []model.LeagueUser{u1, u2, u3}, result)
	})
}

func TestService_GetOverallLeagueTable(t *testing.T) {
	t.Run("Gets users and returns them sorted by points", func(t *testing.T) {
		ctrl, ctx := gomock.WithContext(context.Background(), t)
		defer ctrl.Finish()

		var (
			u1 = model.LeagueUser{Score: 10}
			u2 = model.LeagueUser{Score: 30}
			u3 = model.LeagueUser{Score: 20}
		)

		leagueStore := store_mock.NewMockLeagueStore(ctrl)
		userStore := store_mock.NewMockUserStore(ctrl)
		timer := time_mock.NewMockTimer(ctrl)

		service, err := league.New(leagueStore, userStore, timer)
		require.NoError(t, err)

		leagueTable := []model.LeagueUser{u2, u3, u1}

		userStore.EXPECT().ListAll(ctx).Return(leagueTable, nil)

		result, err := service.GetOverallLeagueTable(ctx)
		require.NoError(t, err)

		assert.Equal(t, leagueTable, result)
	})
}
