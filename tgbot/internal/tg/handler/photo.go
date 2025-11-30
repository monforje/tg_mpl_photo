package handler

import (
	"tgbot/internal/service"
	"tgbot/pkg/errorx"
	"tgbot/pkg/logx"
	"tgbot/pkg/tg/message"
	"tgbot/pkg/tg/sticker"

	tele "gopkg.in/telebot.v4"
)

type PhotoHandler struct {
	photoService *service.PhotoService
}

func NewPhotoHandler(photoService *service.PhotoService) *PhotoHandler {
	return &PhotoHandler{
		photoService: photoService,
	}
}

func (p *PhotoHandler) HandleUpload(c tele.Context) error {
	userID := c.Sender().ID
	username := c.Sender().Username

	photo := c.Message().Photo
	if _, err := c.Bot().File(&photo.File); err != nil {
		return c.Send(message.MsgPhotoUploadFail)
	}

	err := p.photoService.UploadPhoto(
		userID,
		photo.FileID,
		photo.UniqueID,
		photo.FileURL,
	)

	if err == errorx.ErrPhotoDuplicate {
		logx.Warn(
			"photo already uploaded",
			"username", username,
			"tg_id", userID,
			"file_id", photo.FileID,
		)
		sticker.SendSticker(c, sticker.StickerPhotoUploadSuccess)
		return c.Send(message.MsgPhotoDuplicate)
	}

	if err != nil {
		logx.Error(
			"photo upload failed",
			"error", err,
			"username", username,
			"tg_id", userID,
			"file_id", photo.FileID,
		)
		sticker.SendSticker(c, sticker.StickerPhotoUploadFail)
		return c.Send(message.MsgPhotoUploadFail)
	}

	logx.Info(
		"photo uploaded successfully",
		"username", username,
		"tg_id", userID,
		"file_id", photo.FileID,
	)
	sticker.SendSticker(c, sticker.StickerPhotoUploadSuccess)
	return c.Send(message.MsgPhotoUploadSuccess)
}
