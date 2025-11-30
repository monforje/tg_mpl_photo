package router

import (
	"tgbot/internal/tg/handler"

	tele "gopkg.in/telebot.v4"
)

func Setup(
	bot *tele.Bot,
	userHandler *handler.UserHandler,
	photoHandler *handler.PhotoHandler,
) {
	bot.Handle("/start", userHandler.HandleReg)
	bot.Handle(tele.OnPhoto, photoHandler.HandleUpload)
}
