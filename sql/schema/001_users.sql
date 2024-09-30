-- +goose Up
CREATE TABLE users ( id UUID PRIMARY KEY, created_at TIMESTAMP, updated_at TIMESTAMP, name varchar(255) UNIQUE NOT NULL);

-- +goose Down
DROP TABLE users;
