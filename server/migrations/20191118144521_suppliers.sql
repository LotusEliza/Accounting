-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SEQUENCE IF NOT EXISTS items_seq_sub;

CREATE TABLE IF NOT EXISTS suppliers (
    id INT PRIMARY KEY DEFAULT nextval('items_seq_sub'),
    company_name STRING NOT NULL,
    contact_name STRING NOT NULL,
    contact_title STRING NOT NULL,
    address STRING NOT NULL,
    city STRING NOT NULL,
    phone STRING NOT NULL,
    email STRING NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- -- +goose StatementBegin
-- SELECT 'down SQL query';
-- DROP TABLE IF EXISTS suppliers;
-- DROP SEQUENCE IF EXISTS items_seq_sub;
-- -- +goose StatementEnd
