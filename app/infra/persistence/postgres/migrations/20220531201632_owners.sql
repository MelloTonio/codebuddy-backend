-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS owner
(
	id uuid primary key default gen_random_uuid(),
    user_id uuid references users(id),
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);


-- INSERT into owner (user_id) values ('34d7f715-f059-49fa-a916-bfc1b5687d10');



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS owner CASCADE;
-- +goose StatementEnd
