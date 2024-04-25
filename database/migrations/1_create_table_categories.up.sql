-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE categories (
  id BIGINT PRIMARY KEY,
  name VARCHAR(256),
  created_at timestamp with time zone NULL,
  updated_at timestamp with time zone NULL
)

-- +migrate StatementEnd