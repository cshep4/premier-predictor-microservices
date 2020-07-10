package fixture

import (
	"context"
	"errors"
	"fmt"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/model"
	"github.com/golang/protobuf/ptypes/empty"
)

type fixtureService struct {
	fixtureClient gen.FixtureServiceClient
}

func New(fixtureClient gen.FixtureServiceClient) (*fixtureService, error) {
	if fixtureClient == nil {
		return nil, errors.New("fixture_client_is_nil")
	}

	return &fixtureService{
		fixtureClient: fixtureClient,
	}, nil
}

func (f *fixtureService) GetMatches(ctx context.Context) ([]common.Fixture, error) {
	resp, err := f.fixtureClient.GetMatches(ctx, &empty.Empty{})
	if err != nil {
		return nil, fmt.Errorf("get_matches: %w", err)
	}

	var fixtures []common.Fixture
	for _, m := range resp.Matches {
		fixtures = append(fixtures, common.FixtureFromGrpc(m))
	}

	return fixtures, nil
}

func (f *fixtureService) GetTeamForm(ctx context.Context) (map[string]model.TeamForm, error) {
	resp, err := f.fixtureClient.GetTeamForm(ctx, &empty.Empty{})
	if err != nil {
		return nil, fmt.Errorf("get_team_form: %w", err)
	}

	forms, err := model.TeamFormFromGrpc(resp)
	if err != nil {
		return nil, fmt.Errorf("convert_team_form: %w", err)
	}

	return forms, nil
}

func (f *fixtureService) GetFutureFixtures(ctx context.Context) (map[string]string, error) {
	resp, err := f.fixtureClient.GetFutureFixtures(ctx, &empty.Empty{})
	if err != nil {
		return nil, fmt.Errorf("get_future_fixtures: %w", err)
	}

	return resp.Matches, nil
}
