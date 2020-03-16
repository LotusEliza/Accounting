-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SEQUENCE IF NOT EXISTS items_seq_cat;

CREATE TABLE IF NOT EXISTS categories (
    id INT PRIMARY KEY DEFAULT nextval('items_seq_cat'),
    category_name STRING NOT NULL,
    category_description STRING NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- -- +goose StatementBegin
-- SELECT 'down SQL query';
-- DROP TABLE IF EXISTS categories;
-- DROP SEQUENCE IF EXISTS items_seq_cat;
-- -- +goose StatementEnd
