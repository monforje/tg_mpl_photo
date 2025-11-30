package service

import (
	"tgbot/internal/core/event"
	"tgbot/internal/core/repo"
	"tgbot/pkg/logx"
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

	photoEvent := &event.PhotoUploadEvent{
		ID:        id,
		UserID:    user.ID,
		FileID:    fileID,
		UniqueID:  uniqueID,
		FileURL:   fileURL,
		CreatedAt: timeNow,
	}

	if err := p.photoProducer.Produce(photoEvent); err != nil {
		logx.Error(
			"failed to produce photo upload event",
			"error", err,
			"photo_id", id,
			"user_id", user.ID,
		)
	}

	return nil
}
