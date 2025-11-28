package bot

import tele "gopkg.in/telebot.v4"

type Bot struct {
	b *tele.Bot
}

func New(token string) (*Bot, error) {
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10},
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}

	return &Bot{b: b}, nil
}

func (b *Bot) Run() {
	b.b.Start()
}

func (b *Bot) Stop() {
	b.b.Stop()
}
