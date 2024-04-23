-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE books (
  id BIGINT NOT NULL,
  title VARCHAR(256),
  description VARCHAR(256),
  image_url VARCHAR(256),
  release_year INT,
  price INT,
  total_page INT,
  author VARCHAR(256),
  created_at timestamp with time zone NULL,
  updated_at timestamp with time zone NULL,
  category_id BIGINT NOT NULL,
  CONSTRAINT books_pk PRIMARY KEY (id)
)

-- +migrate StatementEnd