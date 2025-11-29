package handler

import (
	"tgbot/internal/service"
	"tgbot/pkg/errorx"
	"tgbot/pkg/logx"
	"tgbot/pkg/tg/message"
	"tgbot/pkg/tg/sticker"

	tele "gopkg.in/telebot.v4"
)

type RegHandler struct {
	regService *service.RegService
}

func NewRegHandler(regService *service.RegService) *RegHandler {
	return &RegHandler{
		regService: regService,
	}
}

func (r *RegHandler) HandleReg(c tele.Context) error {
	tgID := c.Sender().ID
	username := c.Sender().Username

	if err := r.regService.Reg(tgID, username); err != nil {
		if err == errorx.ErrAlreadyRegistered {
			logx.Error(
				"user already registered",
				"username", username,
				"tg_id", tgID,
			)
			sticker.SendSticker(c, sticker.StickerRegSuccess)
			return c.Send(message.MsgRegAlready)
		}
		logx.Error(
			"registration failed",
			"error", err,
			"username", username,
			"tg_id", tgID,
		)
		sticker.SendSticker(c, sticker.StickerRegFail)
		return c.Send(message.MsgRegFail)
	}

	logx.Info(
		"user registered successfully",
		"username", username,
		"tg_id", tgID,
	)
	sticker.SendSticker(c, sticker.StickerRegSuccess)
	return c.Send(message.MsgRegSuccess)
}
