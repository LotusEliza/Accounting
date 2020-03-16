-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SEQUENCE IF NOT EXISTS items_seq_deb_prod;

CREATE TABLE IF NOT EXISTS debits_products (
    id INT PRIMARY KEY DEFAULT nextval('items_seq_deb_prod'),
    debit_id INT NOT NULL,
    product_id INT NOT NULL,
    amount FLOAT NOT NULL,
    price DECIMAL(6,2) NOT NULL,
    FOREIGN KEY (debit_id) REFERENCES debit(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
);

-- INSERT INTO units VALUES (1,'kg'), (2,'item'), (3,'liter');

-- +goose StatementEnd

-- +goose Down
-- -- +goose StatementBegin
-- SELECT 'down SQL query';
-- DROP TABLE IF EXISTS products;
-- DROP SEQUENCE IF EXISTS items_seq_prod;
-- -- +goose StatementEnd
