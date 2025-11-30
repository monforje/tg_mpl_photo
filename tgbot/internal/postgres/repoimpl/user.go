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
	ctx context.Context,
	id uuid.UUID,
	tgID int64,
	username string,
	createdAt time.Time,
) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
        INSERT INTO users (id, tg_id, username, created_at) 
        VALUES ($1, $2, $3, $4)
        ON CONFLICT (tg_id) DO NOTHING
        RETURNING id
    `

	var insertedID uuid.UUID
	err := u.pool.QueryRow(
		ctx,
		query,
		id,
		tgID,
		username,
		createdAt,
	).Scan(&insertedID)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errorx.ErrAlreadyRegistered
		}
		return err
	}

	return nil
}

func (u *UserRepoImpl) GetUserByTgID(ctx context.Context, tgID int64) (*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `SELECT id, tg_id, username, created_at FROM users WHERE tg_id = $1`

	var user model.User

	err := u.pool.QueryRow(ctx, query, tgID).Scan(
		&user.ID,
		&user.TgID,
		&user.Username,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errorx.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}
