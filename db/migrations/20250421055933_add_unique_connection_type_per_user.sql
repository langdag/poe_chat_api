-- +goose Up
-- +goose StatementBegin
ALTER TABLE connections
ADD CONSTRAINT unique_user_connection_type UNIQUE (user_id, connection_type);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE connections
DROP CONSTRAINT unique_user_connection_type;
-- +goose StatementEnd
