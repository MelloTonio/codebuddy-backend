-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cards
(
    id uuid primary key default gen_random_uuid(),
	deck_holder uuid references decks(id),
	owner_id uuid references owner(id),
    question TEXT,
    answer TEXT,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cards CASCADE;
-- +goose StatementEnd
