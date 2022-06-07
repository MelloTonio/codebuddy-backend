-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS decks
(
	id uuid primary key default gen_random_uuid(),
	deck_name TEXT,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS decks CASCADE;
-- +goose StatementEnd
