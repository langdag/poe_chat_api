-- +goose Up
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN connection_type;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN connection_type VARCHAR(20) NOT NULL DEFAULT 'pending';
UPDATE users SET connection_type = 'pending' WHERE connection_type IS NULL;
ALTER TABLE users ADD CONSTRAINT chk_connection_type CHECK (connection_type IN ('telegram', 'tiktok', 'pending'));
-- +goose StatementEnd
