-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SEQUENCE IF NOT EXISTS items_seq_cust;

CREATE TABLE IF NOT EXISTS customers (
    id INT PRIMARY KEY DEFAULT nextval('items_seq_cust'),
    customer_name STRING NOT NULL,
    phone STRING NOT NULL,
    chat_id INT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- -- +goose StatementBegin
-- SELECT 'down SQL query';
-- DROP TABLE IF EXISTS categories;
-- DROP SEQUENCE IF EXISTS items_seq_cat;
-- -- +goose StatementEnd
