-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SEQUENCE IF NOT EXISTS items_seq_debit;

CREATE TABLE IF NOT EXISTS debit (
    id INT PRIMARY KEY DEFAULT nextval('items_seq_debit'),
    debit FLOAT NOT NULL,
    date_create STRING NOT NULL
);

-- INSERT INTO units VALUES (1,'kg'), (2,'item'), (3,'liter');

-- +goose StatementEnd

-- +goose Down
-- -- +goose StatementBegin
-- SELECT 'down SQL query';
-- DROP TABLE IF EXISTS products;
-- DROP SEQUENCE IF EXISTS items_seq_prod;
-- -- +goose StatementEnd
