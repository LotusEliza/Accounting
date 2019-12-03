package models

const DecommissionedTableName = "decommissioned"

type Decommissioned struct {
	ID         int     `db:"id"`
	ProductID  int     `db:"product_id"`
	SellPrice  float64 `db:"sell_price"`
	Amount     int     `db:"amount"`
	Comment    string  `db:"comment"`
	DateCreate string  `db:"date_create"`
}

type SelectDecommissioned struct {
	ID          int     `db:"id"`
	ProductID   int     `db:"product_id"`
	SellPrice   float64 `db:"sell_price"`
	AmountDecom int     `db:"amount"`
	Amount      int
	Comment     string `db:"comment"`
	DateCreate  string `db:"date_create"`
}

var DecommissionedColumns = []string{
	"product_id",
	"sell_price",
	"amount",
	"comment",
	"date_create",
}
