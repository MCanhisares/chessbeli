-- +goose Up
CREATE EXTENSION pgcrypto;

CREATE TABLE users (
  id UUID PRIMARY KEY,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  username TEXT NOT NULL
);

-- +goose Down
DROP EXTENSION pgcrypto;
DROP TABLE users;
