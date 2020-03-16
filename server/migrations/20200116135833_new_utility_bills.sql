-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SEQUENCE IF NOT EXISTS items_seq_bills;

CREATE TABLE IF NOT EXISTS new_utility_bills (
    id INT PRIMARY KEY DEFAULT nextval('items_seq_bills'),
    amount INT NOT NULL,
    bill_name STRING NOT NULL,
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
