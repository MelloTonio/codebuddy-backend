-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
	id uuid primary key default gen_random_uuid(),
    nickname TEXT,
    email TEXT,
    password_hash TEXT,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT into users (nickname, email, password_hash) values ('mello', 'mello@email.com','teste');


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users CASCADE;
-- +goose StatementEnd
