package run

import (
	"context"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

func HandleShutdown(g *errgroup.Group, ctx context.Context, cancel context.CancelFunc, shutdown func(context.Context) error) func() error {
	return func() error {
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		select {
		case <-c:
			cancel()
			return nil
		case <-ctx.Done():
			g.Go(func() error { return shutdown(ctx) })
			return ctx.Err()
		}
	}
}
