package fixture

import (
	"errors"
	"github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/mocks/fixture"
	. "github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	id      = "id"
	teamOne = "Team 1"
	teamTwo = "Team 2"
	score   = "2-1"
)

func TestFixtures_GetMatches(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fixtureClient := fixture_mocks.NewMockFixtureServiceClient(ctrl)

	fixtureService, err := New(fixtureClient)
	require.NoError(t, err)

	t.Run("Gets the matches from FixtureService", func(t *testing.T) {
		matches := []*model.Match{
			{
				Id:    id,
				HTeam: teamOne,
				ATeam: teamTwo,
			},
		}

		resp := &model.Matches{
			Matches: matches,
		}

		fixtureClient.EXPECT().GetMatches(gomock.Any(), gomock.Any()).Return(resp, nil)

		result, err := fixtureService.GetMatches()
		require.NoError(t, err)

		assert.Equal(t, id, result[0].Id)
		assert.Equal(t, teamOne, result[0].HomeTeam)
		assert.Equal(t, teamTwo, result[0].AwayTeam)
	})

	t.Run("Returns error if there is a problem calling FixtureService", func(t *testing.T) {
		fixtureClient.EXPECT().GetMatches(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))

		result, err := fixtureService.GetMatches()
		require.Error(t, err)

		assert.Nil(t, result)
	})
}

func TestFixtures_GetTeamForm(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fixtureClient := fixture_mocks.NewMockFixtureServiceClient(ctrl)

	fixtureService, err := New(fixtureClient)
	require.NoError(t, err)

	t.Run("Gets team forms from FixtureService", func(t *testing.T) {
		teamForms := map[string]*model.TeamForm{
			teamOne: {
				Forms: []*model.TeamMatchResult{
					{
						Result:   model.TeamMatchResult_WIN,
						Score:    score,
						Opponent: teamTwo,
						Location: model.TeamMatchResult_HOME,
					},
				},
			},
		}

		resp := &model.Forms{
			Teams: teamForms,
		}

		fixtureClient.EXPECT().GetTeamForm(gomock.Any(), gomock.Any()).Return(resp, nil)

		result, err := fixtureService.GetTeamForm()
		require.NoError(t, err)

		assert.Equal(t, WIN, result[teamOne].Forms[0].Result)
		assert.Equal(t, score, result[teamOne].Forms[0].Score)
		assert.Equal(t, teamTwo, result[teamOne].Forms[0].Opponent)
		assert.Equal(t, HOME, result[teamOne].Forms[0].Location)
	})

	t.Run("Returns error if there is a problem calling FixtureService", func(t *testing.T) {
		fixtureClient.EXPECT().GetTeamForm(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))

		result, err := fixtureService.GetTeamForm()
		require.Error(t, err)

		assert.Nil(t, result)
	})
}
