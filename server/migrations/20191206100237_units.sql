-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SEQUENCE IF NOT EXISTS items_seq_unit;

CREATE TABLE IF NOT EXISTS units (
    id INT PRIMARY KEY DEFAULT nextval('items_seq_unit'),
    unit_name STRING NOT NULL
);

-- INSERT INTO units VALUES (1,'kg'), (2,'item'), (3,'liter');

-- +goose StatementEnd

-- +goose Down
-- -- +goose StatementBegin
-- SELECT 'down SQL query';
-- DROP TABLE IF EXISTS products;
-- DROP SEQUENCE IF EXISTS items_seq_prod;
-- -- +goose StatementEnd
