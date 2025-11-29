package router

import (
	"tgbot/internal/tg/handler"

	tele "gopkg.in/telebot.v4"
)

type Router struct {
	bot          *tele.Bot
	regHandler   *handler.RegHandler
	photoHandler *handler.PhotoHandler
}

func New(
	bot *tele.Bot,
	regHandler *handler.RegHandler,
	photoHandler *handler.PhotoHandler,
) *Router {
	r := &Router{
		bot:          bot,
		regHandler:   regHandler,
		photoHandler: photoHandler,
	}
	r.Setup()

	return r
}

func (r *Router) Setup() {
	r.bot.Handle("/start", r.regHandler.HandleReg)
	// r.photoHandler.HandleUpload
}
