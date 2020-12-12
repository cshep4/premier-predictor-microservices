package app

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"

	"github.com/cshep4/premier-predictor-microservices/src/common/run"
)

type Runner interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type Func func(context.Context) error

type App struct {
	runners   []Runner
	startups  []Func
	shutdowns []Func
}

func New(opts ...Option) App {
	a := App{}

	for _, opt := range opts {
		opt(&a)
	}

	return a
}

func (a *App) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error { return a.run(ctx) })
	g.Go(run.HandleShutdown(g, ctx, cancel, a.shutdown))

	return g.Wait()
}

func (a *App) run(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	for _, s := range a.startups {
		if err := s(ctx); err != nil {
			return fmt.Errorf("startup_func: %w", err)
		}
	}

	for _, r := range a.runners {
		runner := r
		g.Go(func() error { return runner.Start(ctx) })
	}

	g.Go(func() error {
		<-ctx.Done()
		return ctx.Err()
	})

	return g.Wait()
}

func (a *App) shutdown(ctx context.Context) error {
	for _, s := range a.shutdowns {
		if err := s(ctx); err != nil {
			return fmt.Errorf("shutdown_func: %w", err)
		}
	}

	for _, r := range a.runners {
		if err := r.Stop(ctx); err != nil {
			return fmt.Errorf("shutdown_server: %w", err)
		}
	}

	return nil
}
