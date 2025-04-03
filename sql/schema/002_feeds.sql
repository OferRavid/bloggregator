-- +goose Up
CREATE TABLE feeds(
    id UUID primary key,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null,
    name TEXT unique not null,
    url TEXT unique not null,
    user_id UUID unique not null,
    CONSTRAINT fk_user_id
    FOREIGN KEY (user_id)
    REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;