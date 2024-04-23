-- +migrate Up
-- +migrate StatementBegin

CREATE SEQUENCE reviews_id_seq;

ALTER TABLE reviews
ALTER COLUMN id SET DEFAULT nextval('reviews_id_seq');

-- +migrate StatementEnd