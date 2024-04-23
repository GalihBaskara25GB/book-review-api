-- +migrate Up
-- +migrate StatementBegin

CREATE SEQUENCE users_id_seq;

ALTER TABLE users
ALTER COLUMN id SET DEFAULT nextval('users_id_seq');

-- +migrate StatementEnd