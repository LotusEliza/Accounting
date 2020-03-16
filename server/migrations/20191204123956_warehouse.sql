-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SEQUENCE IF NOT EXISTS items_seq_war;

CREATE TABLE IF NOT EXISTS warehouse (
    id INT PRIMARY KEY DEFAULT nextval('items_seq_war'),
    product_id INT NOT NULL,
    supply_id INT NOT NULL,
    amount INT NOT NULL,
    comment STRING NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
    FOREIGN KEY (supply_id) REFERENCES supply(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- -- +goose StatementBegin
-- SELECT 'down SQL query';
-- DROP TABLE IF EXISTS products;
-- DROP SEQUENCE IF EXISTS items_seq_prod;
-- -- +goose StatementEnd
