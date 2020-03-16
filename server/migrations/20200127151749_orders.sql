-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SEQUENCE IF NOT EXISTS items_seq_ord;

CREATE TABLE IF NOT EXISTS orders (
    id INT PRIMARY KEY DEFAULT nextval('items_seq_ord'),
    customer_id INT NOT NULL,
    product_id INT NOT NULL,
    amount INT NOT NULL,
    date_create STRING NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- -- +goose StatementBegin
-- SELECT 'down SQL query';
-- DROP TABLE IF EXISTS categories;
-- DROP SEQUENCE IF EXISTS items_seq_cat;
-- -- +goose StatementEnd
