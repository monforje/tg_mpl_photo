package event

import (
	"context"
	"tgbot/pkg/errorx"
	"time"

	"github.com/google/uuid"
)

type PhotoUploadEvent struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	FileID    string    `json:"file_id"`
	UniqueID  string    `json:"unique_id"`
	FileURL   string    `json:"file_url"`
	CreatedAt time.Time `json:"created_at"`
}

func (e *PhotoUploadEvent) Validate() error {
	if e.ID == uuid.Nil {
		return errorx.ErrPhotoRequired
	}
	if e.UserID == uuid.Nil {
		return errorx.ErrUserRequired
	}
	if e.FileID == "" {
		return errorx.ErrFileIDRequired
	}
	if e.UniqueID == "" {
		return errorx.ErrUniqueIDRequired
	}
	if e.FileURL == "" {
		return errorx.ErrFileURLRequired
	}
	if e.CreatedAt.IsZero() {
		return errorx.ErrCreatedAtRequired
	}
	return nil
}

type PhotoProducer interface {
	PublishPhoto(ctx context.Context, event *PhotoUploadEvent) error
}
