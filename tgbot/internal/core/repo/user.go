package repo

import (
	"context"
	"tgbot/internal/core/model"
	"time"

	"github.com/google/uuid"
)

type UserRepo interface {
	CreateUser(
		ctx context.Context,
		id uuid.UUID,
		tgID int64,
		username string,
		createdAt time.Time,
	) error
	GetUserByTgID(ctx context.Context, tgID int64) (*model.User, error)
}
