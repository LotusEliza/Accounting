package models

const DebitsProductsTableName = "debits_products"

type DebitsProducts struct {
	ID        int     `db:"id"`
	DebitID   int     `db:"debit_id"`
	ProductID int     `db:"product_id"`
	Amount    float64 `db:"amount"`
	Price     float64 `db:"price"`
}

var DebitsProductsColumns = []string{
	"debit_id",
	"product_id",
	"amount",
	"price",
}
