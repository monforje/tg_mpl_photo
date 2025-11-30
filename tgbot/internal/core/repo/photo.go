package repo

import (
	"tgbot/internal/core/model"
	"time"

	"github.com/google/uuid"
)

type PhotoRepo interface {
	CreatePhoto(
		id uuid.UUID,
		userID uuid.UUID,
		fileID string,
		uniqueID string,
		fileURL string,
		createdAt time.Time,
		updatedAt time.Time,
	) error
	GetPhotoByFileID(fileID string) (*model.Photo, error)
}
