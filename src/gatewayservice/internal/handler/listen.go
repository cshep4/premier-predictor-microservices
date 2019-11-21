package handler

import (
	"context"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

func Listen(ctx context.Context, address string, h http.Handler) error {
	s := &http.Server{
		Addr:         address,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      h,
	}

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		<-ctx.Done()
		return s.Shutdown(context.Background())
	})

	g.Go(func() error {
		// server shutdown successfully so we don't want to return an error
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	return g.Wait()
}
