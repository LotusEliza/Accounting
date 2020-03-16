-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SEQUENCE IF NOT EXISTS items_seq_s;

CREATE TABLE IF NOT EXISTS supply (
    id INT PRIMARY KEY DEFAULT nextval('items_seq_s'),
    date_create STRING NOT NULL,
    product_id INT NOT NULL,
    supplier_id INT NOT NULL,
    buy_price DECIMAL(6,2)  NOT NULL,
    sell_price DECIMAL(6,2)  NOT NULL,
    surcharge INT NOT NULL,
    sup_amount INT NOT NULL,
    amount INT NOT NULL,
    unit_id INT NOT NULL,
    comment STRING NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
    FOREIGN KEY (supplier_id) REFERENCES suppliers(id) ON DELETE CASCADE,
    FOREIGN KEY (unit_id) REFERENCES units(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- -- +goose StatementBegin
-- SELECT 'down SQL query';
-- DROP TABLE IF EXISTS products;
-- DROP SEQUENCE IF EXISTS items_seq_prod;
-- -- +goose StatementEnd
