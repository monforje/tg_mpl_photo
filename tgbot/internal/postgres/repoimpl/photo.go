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

type PhotoRepoImpl struct {
	pool *pgxpool.Pool
}

func NewPhotoRepoImpl(pool *pgxpool.Pool) *PhotoRepoImpl {
	return &PhotoRepoImpl{
		pool: pool,
	}
}

func (p *PhotoRepoImpl) CreatePhoto(
	ctx context.Context,
	id uuid.UUID,
	userID uuid.UUID,
	fileID string,
	uniqueID string,
	fileURL string,
	createdAt time.Time,
	updatedAt time.Time,
) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
        INSERT INTO photos (
            id, user_id, file_id, unique_id, file_url, created_at, updated_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7)
        ON CONFLICT (file_id) DO NOTHING
        RETURNING id
    `

	var insertedID uuid.UUID
	err := p.pool.QueryRow(
		ctx,
		query,
		id,
		userID,
		fileID,
		uniqueID,
		fileURL,
		createdAt,
		updatedAt,
	).Scan(&insertedID)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errorx.ErrPhotoDuplicate
		}
		return err
	}

	return nil
}

func (p *PhotoRepoImpl) GetPhotoByFileID(ctx context.Context, fileID string) (*model.Photo, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
        SELECT 
            id, user_id, file_id, unique_id, file_url, created_at, updated_at
        FROM photos
        WHERE file_id = $1
    `

	var photo model.Photo
	err := p.pool.QueryRow(ctx, query, fileID).Scan(
		&photo.ID,
		&photo.UserID,
		&photo.FileID,
		&photo.UniqueID,
		&photo.FileURL,
		&photo.CreatedAt,
		&photo.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errorx.ErrPhotoNotFound
		}
		return nil, err
	}

	return &photo, nil
}
