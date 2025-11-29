package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `db:"id"`
	TgID      int64     `db:"tg_id"`
	Username  string    `db:"username"`
	CreatedAt time.Time `db:"created_at"`
}
