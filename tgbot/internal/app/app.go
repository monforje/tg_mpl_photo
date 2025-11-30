package app

import (
	"context"
	"tgbot/internal/env"
	"tgbot/internal/kafka/producer"
	"tgbot/internal/postgres"
	"tgbot/internal/postgres/repoimpl"
	"tgbot/internal/service"
	"tgbot/internal/tg/bot"
	"tgbot/internal/tg/handler"
	"tgbot/pkg/logx"
)

type App struct {
	b  *bot.Bot
	db *postgres.Postgres
}

func New(ctx context.Context) (*App, error) {
	e, err := env.New()
	if err != nil {
		return nil, err
	}

	// cfg, err := config.New()
	// if err != nil {
	// 	return nil, err
	// }

	db, err := postgres.New(ctx, e.PostgresDSN)
	if err != nil {
		return nil, err
	}

	// kafka
	// kafka.New() mb
	photoProducer := producer.NewPhotoUploadEventProducer()

	// user
	userRepo := repoimpl.NewUserRepoImpl(db.Pool)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// photo
	photoRepo := repoimpl.NewPhotoRepoImpl(db.Pool)
	photoService := service.NewPhotoService(
		photoRepo,
		userRepo,
		photoProducer,
	)
	photoHandler := handler.NewPhotoHandler(
		photoService,
		e.TgToken,
	)

	b, err := bot.New(
		e.TgToken,
		userHandler,
		photoHandler,
	)
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
