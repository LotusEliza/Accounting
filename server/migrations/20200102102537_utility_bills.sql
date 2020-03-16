-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SEQUENCE IF NOT EXISTS items_seq_bills;

CREATE TABLE IF NOT EXISTS utility_bills (
    id INT PRIMARY KEY DEFAULT nextval('items_seq_bills'),
    electricity INT NOT NULL,
    water INT NOT NULL,
    gas INT NOT NULL,
    date_create STRING NOT NULL,
    comment STRING NOT NULL
);

-- INSERT INTO units VALUES (1,'kg'), (2,'item'), (3,'liter');

-- +goose StatementEnd

-- +goose Down
-- -- +goose StatementBegin
-- SELECT 'down SQL query';
-- DROP TABLE IF EXISTS products;
-- DROP SEQUENCE IF EXISTS items_seq_prod;
-- -- +goose StatementEnd
