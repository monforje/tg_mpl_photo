package handler

import (
	"tgbot/internal/service"

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
	// TODO: implement
	return nil
}
