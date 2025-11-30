package bot

import (
	"tgbot/internal/postgres/repoimpl"
	"tgbot/internal/service"
	"tgbot/internal/tg/handler"
	"tgbot/internal/tg/router"
	"tgbot/pkg/logx"

	"github.com/jackc/pgx/v5/pgxpool"
	tele "gopkg.in/telebot.v4"
)

type Bot struct {
	b *tele.Bot
}

func New(token string, pool *pgxpool.Pool) (*Bot, error) {
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10},
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}

	// user
	userRepo := repoimpl.NewUserRepoImpl(pool)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// photo
	photoRepo := repoimpl.NewPhotoRepoImpl(pool)
	photoService := service.NewPhotoService(photoRepo)
	photoHandler := handler.NewPhotoHandler(photoService)

	router.New(
		b,
		userHandler,
		photoHandler,
	)

	logx.Info("bot initialized successfully")

	return &Bot{
		b: b,
	}, nil
}

func (b *Bot) Run() {
	logx.Info("bot is starting")
	b.b.Start()
}

func (b *Bot) Stop() {
	b.b.Stop()
	logx.Info("bot has stopped")
}
