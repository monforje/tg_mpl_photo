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
	id uuid.UUID,
	userID int64,
	fileID string,
	uniqueID string,
	fileURL string,
	createdAt time.Time,
	updatedAt time.Time,
) error {

	existingPhoto, err := p.GetPhotoByFileID(fileID)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return err
		}
	}
	if existingPhoto != nil {
		return errorx.ErrPhotoDuplicate
	}

	query := `
        INSERT INTO photos (
            id, user_id, file_id, unique_id, file_url, created_at, updated_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7)
    `

	_, err = p.pool.Exec(
		context.Background(),
		query,
		id,
		userID,
		fileID,
		uniqueID,
		fileURL,
		createdAt,
		updatedAt,
	)

	return err
}

func (p *PhotoRepoImpl) GetPhotoByFileID(fileID string) (*model.Photo, error) {

	query := `
        SELECT 
            id, user_id, file_id, unique_id, file_url, created_at, updated_at
        FROM photos
        WHERE file_id = $1
    `

	var photo model.Photo

	err := p.pool.QueryRow(context.Background(), query, fileID).Scan(
		&photo.ID,
		&photo.UserID,
		&photo.FileID,
		&photo.UniqueID,
		&photo.FileURL,
		&photo.CreatedAt,
		&photo.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &photo, nil
}
