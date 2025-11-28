package repo

import (
	"tgbot/internal/core/model"
	"time"
)

type UserRepo interface {
	CreateUser(tgID int64, username string, createdAt time.Time) (int64, error)
	GetUserByTgID(tgID int64) (*model.User, error)
}
