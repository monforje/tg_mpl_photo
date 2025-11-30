-- 00002_create_photos_table.sql
-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS photos (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL, -- изменил на UUID
    file_id VARCHAR(255) NOT NULL UNIQUE,
    unique_id VARCHAR(255) NOT NULL UNIQUE,
    file_url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_photos_user
        FOREIGN KEY (user_id)
        REFERENCES users(id) -- теперь ссылается на users.id
        ON DELETE CASCADE
);

CREATE INDEX idx_photos_user_id ON photos(user_id);
CREATE INDEX idx_photos_file_id ON photos(file_id);
CREATE INDEX idx_photos_unique_id ON photos(unique_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_photos_unique_id;
DROP INDEX IF EXISTS idx_photos_file_id;
DROP INDEX IF EXISTS idx_photos_user_id;
DROP TABLE IF EXISTS photos;
-- +goose StatementEnd