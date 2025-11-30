package handler

import (
	"errors"
	"tgbot/internal/service"
	"tgbot/pkg/errorx"
	"tgbot/pkg/logx"
	"tgbot/pkg/tg/message"
	"tgbot/pkg/tg/sticker"

	tele "gopkg.in/telebot.v4"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) HandleReg(c tele.Context) error {
	tgID := c.Sender().ID
	username := c.Sender().Username

	err := h.userService.Reg(tgID, username)

	if errors.Is(err, errorx.ErrAlreadyRegistered) {
		logx.Warn(
			"user already registered",
			"username", username,
			"tg_id", tgID,
		)
		sticker.SendSticker(c, sticker.StickerRegSuccess)
		return c.Send(message.MsgRegAlready)
	}

	if err != nil {
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
