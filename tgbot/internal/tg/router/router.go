package router

import (
	"tgbot/internal/tg/handler"

	tele "gopkg.in/telebot.v4"
)

type Router struct {
	bot        *tele.Bot
	regHandler *handler.RegHandler
}

func New(
	bot *tele.Bot,
	regHandler *handler.RegHandler,
) *Router {
	r := &Router{
		bot:        bot,
		regHandler: regHandler,
	}
	r.Setup()

	return r
}

func (r *Router) Setup() {
	r.bot.Handle("/start", r.regHandler.HandleReg)
}
