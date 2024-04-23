-- +migrate Up
-- +migrate StatementBegin

CREATE SEQUENCE books_id_seq;

ALTER TABLE books
ALTER COLUMN id SET DEFAULT nextval('books_id_seq');

-- +migrate StatementEnd