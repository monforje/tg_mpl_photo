package app

import (
	"context"
	"tgbot/internal/env"
	"tgbot/internal/tg/bot"
	"tgbot/pkg/logx"
)

type App struct {
	b *bot.Bot
}

func New(ctx context.Context) (*App, error) {
	e, err := env.New()
	if err != nil {
		return nil, err
	}

	b, err := bot.New(e.TgToken)
	if err != nil {
		return nil, err
	}

	logx.Info("app initialized successfully")

	return &App{
		b: b,
	}, nil
}

func (a *App) Run(ctx context.Context) {
	logx.Info("app is starting")

	done := make(chan struct{})
	go func() {
		defer close(done)
		a.b.Run()
	}()

	<-ctx.Done()
	a.Stop()
	<-done
	logx.Info("app has stopped")
}

func (a *App) Stop() {
	a.b.Stop()
}
