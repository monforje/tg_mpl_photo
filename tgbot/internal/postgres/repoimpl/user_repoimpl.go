package repoimpl

import (
	"tgbot/internal/core/model"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepoImpl struct {
	pool *pgxpool.Pool
}

func NewUserRepoImpl(pool *pgxpool.Pool) *UserRepoImpl {
	return &UserRepoImpl{
		pool: pool,
	}
}

func (u *UserRepoImpl) CreateUser(
	id uuid.UUID,
	tgID int64,
	username string,
	createdAt time.Time,
) (int64, error) {
	// TODO: implement
	return 0, nil
}

func (u *UserRepoImpl) GetUserByTgID(tgID int64) (*model.User, error) {
	// TODO: implement
	return nil, nil
}
