-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SEQUENCE IF NOT EXISTS items_seq_prod;

CREATE TABLE IF NOT EXISTS products (
    id INT PRIMARY KEY DEFAULT nextval('items_seq_prod'),
    vendor_code VARCHAR(255) NOT NULL,
    product_name STRING NOT NULL,
    category_id INT NOT NULL,
    img_link STRING NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- -- +goose StatementBegin
-- SELECT 'down SQL query';
-- DROP TABLE IF EXISTS products;
-- DROP SEQUENCE IF EXISTS items_seq_prod;
-- -- +goose StatementEnd
