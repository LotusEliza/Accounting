-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SEQUENCE IF NOT EXISTS items_seq_prod;

CREATE TABLE IF NOT EXISTS products (
    id INT PRIMARY KEY DEFAULT nextval('items_seq_prod'),
    vendor_code VARCHAR(255) NOT NULL,
    product_name STRING NOT NULL,
    supplier_id INT NOT NULL,
    category_id INT NOT NULL,
    buy_price DECIMAL(6,2)  NOT NULL,
    sell_price DECIMAL(6,2)  NOT NULL,
    amount INT NOT NULL,
    date_create DATE NOT NULL,
    FOREIGN KEY (supplier_id) REFERENCES suppliers(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- -- +goose StatementBegin
-- SELECT 'down SQL query';
-- DROP TABLE IF EXISTS products;
-- DROP SEQUENCE IF EXISTS items_seq_prod;
-- -- +goose StatementEnd
