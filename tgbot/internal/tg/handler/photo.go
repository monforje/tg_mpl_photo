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
	token        string
}

func NewPhotoHandler(
	photoService *service.PhotoService,
	token string,
) *PhotoHandler {

	return &PhotoHandler{
		photoService: photoService,
		token:        token,
	}
}

func (p *PhotoHandler) HandleUpload(c tele.Context) error {
	userID := c.Sender().ID
	username := c.Sender().Username

	photo := c.Message().Photo
	if _, err := c.Bot().File(&photo.File); err != nil {
		logx.Error(
			"photo upload failed",
			"error", err,
			"username", username,
			"tg_id", userID,
			"unique_id", photo.UniqueID,
		)
		return c.Send(message.MsgPhotoUploadFail)
	}
	err := p.photoService.UploadPhoto(
		userID,
		photo.FileID,
		photo.UniqueID,
		getURL(p.token, photo.FilePath),
	)

	if err == errorx.ErrPhotoDuplicate {
		logx.Warn(
			"photo already uploaded",
			"username", username,
			"tg_id", userID,
			"unique_id", photo.UniqueID,
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
			"unique_id", photo.UniqueID,
		)
		sticker.SendSticker(c, sticker.StickerPhotoUploadFail)
		return c.Send(message.MsgPhotoUploadFail)
	}

	logx.Info(
		"photo uploaded successfully",
		"username", username,
		"tg_id", userID,
		"unique_id", photo.UniqueID,
	)
	sticker.SendSticker(c, sticker.StickerPhotoUploadSuccess)
	return c.Send(message.MsgPhotoUploadSuccess)
}

func getURL(token string, filePath string) string {
	return "https://api.telegram.org/file/bot" + token + "/" + filePath
}
