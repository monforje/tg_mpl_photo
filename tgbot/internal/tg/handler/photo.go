package handler

import "tgbot/internal/service"

type PhotoHandler struct {
	photoService *service.PhotoService
}

func NewPhotoHandler(photoService *service.PhotoService) *PhotoHandler {
	return &PhotoHandler{
		photoService: photoService,
	}
}

func (p *PhotoHandler) HandleUpload() {
	/*
		обрабатывает загрузку фото
		вызывает сервисы
	*/
}
