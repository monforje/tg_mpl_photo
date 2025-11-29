package app

import (
	"context"
	"tgbot/internal/env"
	"tgbot/internal/psql"
	"tgbot/internal/tg/bot"
	"tgbot/pkg/logx"
)

type App struct {
	b  *bot.Bot
	db *psql.Psql
}

func New(ctx context.Context) (*App, error) {
	e, err := env.New()
	if err != nil {
		return nil, err
	}

	db, err := psql.New(ctx, e.PostgresDSN)
	if err != nil {
		return nil, err
	}

	b, err := bot.New(e.TgToken, db.Pool)
	if err != nil {
		return nil, err
	}

	logx.Info("app initialized successfully")

	return &App{
		b:  b,
		db: db,
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
	a.db.Stop()
}
