package app

import (
	"context"
	"tgbot/internal/env"
)

type App struct {
}

func New(ctx context.Context) (*App, error) {
	e, err := env.New()
	if err != nil {
		return nil, err
	}
	_ = e // use env

	return &App{}, nil
}

func (a *App) Run(ctx context.Context) {
	done := make(chan struct{})

	go func() {
		defer close(done)
		// start bot here
	}()

	<-ctx.Done()
	a.Stop()
	<-done
}

func (a *App) Stop() {
	// stop bot here
}
