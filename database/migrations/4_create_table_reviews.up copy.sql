-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE reviews (
  id BIGINT NOT NULL,
  user_id BIGINT NOT NULL,
  book_id BIGINT NOT NULL,
  description VARCHAR(256) NULL,
  stars INT NOT NULL,
  created_at timestamp with time zone NULL,
  updated_at timestamp with time zone NULL,
  CONSTRAINT reviews_pk PRIMARY KEY (id)
)

-- +migrate StatementEnd