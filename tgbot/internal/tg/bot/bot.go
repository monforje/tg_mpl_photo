package bot

import (
	"tgbot/internal/tg/handler"
	"tgbot/internal/tg/router"
	"tgbot/pkg/logx"

	tele "gopkg.in/telebot.v4"
)

type Bot struct {
	b *tele.Bot
}

func New(
	token string,
	userHandler *handler.UserHandler,
	photoHandler *handler.PhotoHandler,
) (*Bot, error) {

	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10},
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}

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
