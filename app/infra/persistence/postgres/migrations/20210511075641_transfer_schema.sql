-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS transfers
(
	id uuid PRIMARY KEY,
	account_origin_id uuid REFERENCES accounts,
	account_destination_id uuid REFERENCES accounts,
	amount int NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transfers CASCADE;
-- +goose StatementEnd
