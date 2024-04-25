-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE users (
  id BIGINT PRIMARY KEY,
  username VARCHAR(256) NOT NULL,
  password VARCHAR(256) NOT NULL,
  role VARCHAR(256) NOT NULL,
  created_at timestamp with time zone NULL,
  updated_at timestamp with time zone NULL
)

-- +migrate StatementEnd