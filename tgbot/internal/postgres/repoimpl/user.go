package repoimpl

import (
	"context"
	"errors"
	"tgbot/internal/core/model"
	"tgbot/pkg/errorx"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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
) error {
	existingUser, err := u.GetUserByTgID(tgID)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return err
		}
	}
	if existingUser != nil {
		return errorx.ErrAlreadyRegistered
	}

	query := `INSERT INTO users (id, tg_id, username, created_at) VALUES ($1, $2, $3, $4)`
	_, err = u.pool.Exec(context.Background(), query, id, tgID, username, createdAt)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepoImpl) GetUserByTgID(tgID int64) (*model.User, error) {
	query := `SELECT id, tg_id, username, created_at FROM users WHERE tg_id = $1`

	var user model.User
	err := u.pool.QueryRow(context.Background(), query, tgID).Scan(
		&user.ID,
		&user.TgID,
		&user.Username,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
