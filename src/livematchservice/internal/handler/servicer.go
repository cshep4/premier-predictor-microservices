package handler

import (
	"context"
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/model"
	"time"
)

type Servicer interface {
	GetMatchSummary(ctx context.Context, req model.PredictionRequest) (*model.MatchSummary, error)
	GetMatchFacts(ctx context.Context, id string) (*common.MatchFacts, error)
	SubscribeToMatch(ctx context.Context, id string, observer model.MatchObserver) error
	GetUpcomingMatches(ctx context.Context) (map[time.Time][]common.MatchFacts, error)
	GetTodaysMatches(ctx context.Context) ([]common.MatchFacts, error)
	SubscribeToTodaysMatches(ctx context.Context, observer model.MatchObserver) error
}
