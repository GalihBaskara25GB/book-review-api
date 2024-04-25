-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE reviews (
  id BIGINT PRIMARY KEY,
  user_id BIGINT REFERENCES users (id),
  book_id BIGINT REFERENCES books (id),
  description VARCHAR(256) NULL,
  stars INT NOT NULL,
  created_at timestamp with time zone NULL,
  updated_at timestamp with time zone NULL
)

-- +migrate StatementEnd