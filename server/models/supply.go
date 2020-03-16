package models

const SupplyTableName = "supply"

type Supply struct {
	ID         int     `db:"id"`
	DateCreate string  `db:"date_create"`
	ProductID  int     `db:"product_id"`
	SupplierID int     `db:"supplier_id"`
	BuyPrice   float64 `db:"buy_price"`
	SellPrice  float64 `db:"sell_price"`
	Surcharge  int     `db:"surcharge"`
	SupAmount  int     `db:"sup_amount"`
	Amount     int     `db:"amount"`
	UnitID     int     `db:"unit_id"`
	Comment    string  `db:"comment"`
}

var SupplyColumns = []string{
	"date_create",
	"product_id",
	"supplier_id",
	"buy_price",
	"sell_price",
	"surcharge",
	"sup_amount",
	"amount",
	"unit_id",
	"comment",
}

type SelectSupply struct {
	ID          int     `db:"id"`
	DateCreate  string  `db:"date_create"`
	ProductName string  `db:"product_name"`
	ProductID   int     `db:"product_id"`
	CompanyName string  `db:"company_name"`
	SupplierID  int     `db:"supplier_id"`
	BuyPrice    float64 `db:"buy_price"`
	SellPrice   float64 `db:"sell_price"`
	Surcharge   int     `db:"surcharge"`
	SupAmount   int     `db:"sup_amount"`
	Amount      int     `db:"amount"`
	UnitID      int     `db:"unit_id"`
	UnitName    string  `db:"unit_name"`
	Comment     string  `db:"comment"`
}

type IdSelect struct {
	ID int `db:"id"`
}

//type Decommissioned struct {
//	ID          int     `db:"id"`
//	ProductID   int     `db:"product_id"`
//	SellPrice   float64 `db:"sell_price"`
//	Amount      int     `db:"amount"`
//	Comment     string  `db:"comment"`
//}
