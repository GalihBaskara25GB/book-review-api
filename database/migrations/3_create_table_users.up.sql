-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE users (
  id BIGINT NOT NULL,
  username VARCHAR(256) NOT NULL,
  password VARCHAR(256) NOT NULL,
  role VARCHAR(256) NOT NULL,
  created_at timestamp with time zone NULL,
  updated_at timestamp with time zone NULL,
  CONSTRAINT users_pk PRIMARY KEY (id)
)

-- +migrate StatementEnd