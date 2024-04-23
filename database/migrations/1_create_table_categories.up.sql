-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE categories (
  id BIGINT NOT NULL,
  name VARCHAR(256),
  created_at timestamp with time zone NULL,
  updated_at timestamp with time zone NULL,
  CONSTRAINT categories_pk PRIMARY KEY (id)
)

-- +migrate StatementEnd