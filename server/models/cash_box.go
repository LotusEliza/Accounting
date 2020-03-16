package models

const CashBoxTableName = "cashbox"

type CashBox struct {
	ID         int    `db:"id"`
	DateCreate string `db:"date_create"`
	ProductID  int    `db:"product_id"`
	Amount     int    `db:"amount"`
	SellPrice  int    `db:"sell_price"`
}

var CashBoxColumns = []string{
	"date_create",
	"product_id",
	"amount",
	"sell_price",
}
