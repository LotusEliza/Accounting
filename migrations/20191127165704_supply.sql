-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SEQUENCE IF NOT EXISTS items_seq_sup;

CREATE TABLE IF NOT EXISTS supply (
    id INT PRIMARY KEY DEFAULT nextval('items_seq_sup'),
    product_id INT NOT NULL,
    amount INT NOT NULL,
    sell_price DECIMAL(6,2)  NOT NULL,
    comment STRING NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- -- +goose StatementBegin
-- SELECT 'down SQL query';
-- DROP TABLE IF EXISTS products;
-- DROP SEQUENCE IF EXISTS items_seq_prod;
-- -- +goose StatementEnd
