-- +migrate Up
-- +migrate StatementBegin

CREATE SEQUENCE categories_id_seq;

ALTER TABLE categories
ALTER COLUMN id SET DEFAULT nextval('categories_id_seq');

-- +migrate StatementEnd