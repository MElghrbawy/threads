-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS category
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS set_timestamp ON category;
DROP FUNCTION IF EXISTS update_updated_at_column;
DROP TABLE IF EXISTS category;
-- +goose StatementEnd
