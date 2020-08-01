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
	GetUpcomingMatches(ctx context.Context) (map[time.Time][]common.MatchFacts, error)
}
