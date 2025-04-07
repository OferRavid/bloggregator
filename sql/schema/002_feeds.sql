-- +goose Up
CREATE TABLE feeds(
    id UUID primary key,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null,
    last_fetched_at TIMESTAMP,
    name TEXT unique not null,
    url TEXT unique not null,
    user_id UUID not null REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;