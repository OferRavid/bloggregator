-- +goose Up
CREATE TABLE users(
    id UUID primary key,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null,
    name TEXT unique not null
);

-- +goose Down
DROP TABLE users;