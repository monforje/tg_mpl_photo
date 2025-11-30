package repo

import (
	"context" // ✅ Добавили
	"tgbot/internal/core/model"
	"time"

	"github.com/google/uuid"
)

type PhotoRepo interface {
	CreatePhoto(
		ctx context.Context,
		id uuid.UUID,
		userID uuid.UUID,
		fileID string,
		uniqueID string,
		fileURL string,
		createdAt time.Time,
		updatedAt time.Time,
	) error
	GetPhotoByFileID(ctx context.Context, fileID string) (*model.Photo, error)
}
