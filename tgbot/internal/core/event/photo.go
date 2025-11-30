package event

import (
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

type PhotoUploadEventProducer interface {
	Produce(event *PhotoUploadEvent) error
}
