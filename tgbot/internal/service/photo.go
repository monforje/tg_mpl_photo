package service

import (
	"tgbot/internal/core/repo"
	"time"

	"github.com/google/uuid"
)

type PhotoService struct {
	photoRepo repo.PhotoRepo
}

func NewPhotoService(photoRepo repo.PhotoRepo) *PhotoService {
	return &PhotoService{
		photoRepo: photoRepo,
	}
}

func (p *PhotoService) UploadPhoto(
	userID int64,
	fileID string,
	uniqueID string,
	fileURL string,
) error {
	id := uuid.New()
	timeNow := time.Now()

	if err := p.photoRepo.CreatePhoto(
		id,
		userID,
		fileID,
		uniqueID,
		fileURL,
		timeNow,
		timeNow,
	); err != nil {
		return err
	}
	return nil
}
