-- +goose Up
CREATE TABLE posts(
    id UUID primary key,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null,
    title TEXT not null,
    url TEXT unique not null,
    description TEXT,
    published_at TIMESTAMP,
    feed_id UUID not null REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;