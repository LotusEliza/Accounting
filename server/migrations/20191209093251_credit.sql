-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SEQUENCE IF NOT EXISTS items_seq_credit;

CREATE TABLE IF NOT EXISTS credit (
    id INT PRIMARY KEY DEFAULT nextval('items_seq_credit'),
    credit INT NOT NULL,
    supply_id INT,
    bill_id INT,
    FOREIGN KEY (supply_id) REFERENCES supply(id) ON DELETE CASCADE,
    FOREIGN KEY (bill_id) REFERENCES utility_bills(id) ON DELETE CASCADE
);

-- INSERT INTO units VALUES (1,'kg'), (2,'item'), (3,'liter');

-- +goose StatementEnd

-- +goose Down
-- -- +goose StatementBegin
-- SELECT 'down SQL query';
-- DROP TABLE IF EXISTS products;
-- DROP SEQUENCE IF EXISTS items_seq_prod;
-- -- +goose StatementEnd
