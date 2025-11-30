package repoimpl

import (
	"tgbot/internal/core/model"
	"time"

	"github.com/google/uuid"
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
	updatedAt time.Time,
	createdAt time.Time,
) error {
	return nil
}

func (p *PhotoRepoImpl) GetPhotoByFileID(fileID string) (*model.Photo, error) {
	return nil, nil
}
