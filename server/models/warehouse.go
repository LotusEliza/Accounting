package models

const WarehouseTableName = "warehouse"

type Warehouse struct {
	ID         int    `db:"id"`
	DateCreate string `db:"date_create"`
	ProductID  int    `db:"product_id"`
	//SupplierID int     `db:"supplier_id"`
	SupplyID int    `db:"supply_id"`
	Amount   int    `db:"amount"`
	Comment  string `db:"comment"`
}

type SelectWarehouse struct {
	ID          int     `db:"id"`
	DateCreate  string  `db:"date_create"`
	ProductName string  `db:"product_name"`
	CompanyName string  `db:"company_name"`
	BuyPrice    float64 `db:"buy_price"`
	Surcharge   int     `db:"surcharge"`
	Amount      int     `db:"amount"`
	Unit        string  `db:"unit"`
	Comment     string  `db:"comment"`
}

var WarehouseColumns = []string{
	"date_create",
	"product_id",
	"supply_id",
	//"buy_price",
	//"surcharge",
	"amount",
	//"unit",
	"comment",
}

//type Decommissioned struct {
//	ID          int     `db:"id"`
//	ProductID   int     `db:"product_id"`
//	SellPrice   float64 `db:"sell_price"`
//	Amount      int     `db:"amount"`
//	Comment     string  `db:"comment"`
//}
