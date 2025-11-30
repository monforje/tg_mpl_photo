-- 00002_create_photos_table.sql
-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS photos (
    id UUID PRIMARY KEY,
    user_id BIGINT NOT NULL,
    file_id VARCHAR(255) NOT NULL UNIQUE, -- добавил UNIQUE
    unique_id VARCHAR(255) NOT NULL UNIQUE, -- добавил UNIQUE
    file_url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_photos_user
        FOREIGN KEY (user_id)
        REFERENCES users(tg_id)
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