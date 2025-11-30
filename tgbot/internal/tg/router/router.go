package router

import (
	"tgbot/internal/tg/handler"

	tele "gopkg.in/telebot.v4"
)

type Router struct {
	bot          *tele.Bot
	userHandler  *handler.UserHandler
	photoHandler *handler.PhotoHandler
}

func New(
	bot *tele.Bot,
	userHandler *handler.UserHandler,
	photoHandler *handler.PhotoHandler,
) *Router {
	r := &Router{
		bot:          bot,
		userHandler:  userHandler,
		photoHandler: photoHandler,
	}
	r.Setup()

	return r
}

func (r *Router) Setup() {
	r.bot.Handle("/start", r.userHandler.HandleReg)
	r.bot.Handle(tele.OnPhoto, r.photoHandler.HandleUpload)
}
