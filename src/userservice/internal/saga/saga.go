package saga

import (
	"context"

	"github.com/cshep4/data-structures/saga"
	"github.com/cshep4/premier-predictor-microservices/src/common/log"
)

type (
	ErrorHandler    struct{}
	RollbackHandler struct{}
)

func (e ErrorHandler) Handle(ctx context.Context, process saga.Process, err error) {
	log.Error(ctx, "saga_process_error",
		log.SafeParam("saga_process", process.Name()),
		log.ErrorParam(err),
	)
}

func (r RollbackHandler) Handle(ctx context.Context, process saga.Process) {
	log.Error(ctx, "saga_process_rollback_success",
		log.SafeParam("saga_process", process.Name()),
	)
}
