-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS photos (
    id UUID PRIMARY KEY,
    user_id BIGINT NOT NULL,
    file_id VARCHAR(255) NOT NULL,
    unique_id VARCHAR(255) NOT NULL,
    file_url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_photos_user
        FOREIGN KEY (user_id)
        REFERENCES users(tg_id)
        ON DELETE CASCADE
);

CREATE INDEX idx_photos_user_id ON photos(user_id);
CREATE INDEX idx_photos_file_id ON photos(file_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS photos;
-- +goose StatementEnd
