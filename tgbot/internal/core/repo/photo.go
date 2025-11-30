package repo

import (
	"tgbot/internal/core/model"
	"time"

	"github.com/google/uuid"
)

type PhotoRepo interface {
	CreatePhoto(
		id uuid.UUID,
		createdAt time.Time,
	) error
	GetPhotoByTgID() (*model.Photo, error)
}
