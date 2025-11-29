package repo

import (
	"tgbot/internal/core/model"
	"time"

	"github.com/google/uuid"
)

type UserRepo interface {
	CreateUser(
		id uuid.UUID,
		tgID int64,
		username string,
		createdAt time.Time,
	) (int64, error)
	GetUserByTgID(tgID int64) (*model.User, error)
}
