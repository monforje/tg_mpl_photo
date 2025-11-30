package service

import (
	"tgbot/internal/core/event"
	"tgbot/internal/core/repo"
	"time"

	"github.com/google/uuid"
)

type PhotoService struct {
	photoRepo     repo.PhotoRepo
	userRepo      repo.UserRepo
	photoProducer event.PhotoUploadEventProducer
}

func NewPhotoService(
	photoRepo repo.PhotoRepo,
	userRepo repo.UserRepo,
	photoProducer event.PhotoUploadEventProducer,
) *PhotoService {
	return &PhotoService{
		photoRepo:     photoRepo,
		userRepo:      userRepo,
		photoProducer: photoProducer,
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

	// create event

	// produce event

	return nil
}
