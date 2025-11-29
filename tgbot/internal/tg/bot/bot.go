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

	userRepo := repoimpl.NewUserRepoImpl(pool)

	regService := service.NewRegService(userRepo)

	regHandler := handler.NewRegHandler(regService)

	router.New(
		b,
		regHandler,
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
