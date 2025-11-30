package service

import (
	"tgbot/internal/core/repo"
	"time"

	"github.com/google/uuid"
)

type PhotoService struct {
	photoRepo repo.PhotoRepo
	userRepo  repo.UserRepo
}

func NewPhotoService(
	photoRepo repo.PhotoRepo,
	userRepo repo.UserRepo,
) *PhotoService {
	return &PhotoService{
		photoRepo: photoRepo,
		userRepo:  userRepo,
	}
}

func (p *PhotoService) UploadPhoto(
	tgID int64,
	fileID string,
	uniqueID string,
	fileURL string,
) error {
	user, err := p.userRepo.GetUserByTgID(tgID)
	if err != nil {
		return err
	}

	id := uuid.New()
	timeNow := time.Now()

	if err := p.photoRepo.CreatePhoto(
		id,
		user.ID,
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
