package app

import (
	"context"
)

type Option func(*App)

func WithRunner(r Runner) Option {
	return func(a *App) {
		a.runners = append(a.runners, r)
	}
}

func WithStartupFunc(s func() error) Option {
	return func(a *App) {
		a.startups = append(a.startups, func(ctx context.Context) error { return s() })
	}
}

func WithStartupFuncContext(s Func) Option {
	return func(a *App) {
		a.startups = append(a.startups, s)
	}
}

func WithShutdownFunc(s func() error) Option {
	return func(a *App) {
		a.shutdowns = append(a.shutdowns, func(ctx context.Context) error { return s() })
	}
}

func WithShutdownFuncContext(s Func) Option {
	return func(a *App) {
		a.shutdowns = append(a.shutdowns, s)
	}
}
