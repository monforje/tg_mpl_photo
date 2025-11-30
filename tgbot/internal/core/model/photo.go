package model

import (
	"time"

	"github.com/google/uuid"
)

type Photo struct {
	ID        uuid.UUID `db:"id"`
	UserID    int64     `db:"user_id"`
	FileID    string    `db:"file_id"`
	UniqueID  string    `db:"unique_id"`
	FileURL   string    `db:"file_url"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
