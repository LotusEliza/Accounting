package models

const ProductsTableName = "products"

type Product struct {
	ID          int     `db:"id"`
	VendorCode  string  `db:"vendor_code"`
	ProductName string  `db:"product_name"`
	SupplierID  int     `db:"supplier_id"`
	CategoryID  int     `db:"category_id"`
	BuyPrice    float64 `db:"buy_price"`
	SellPrice   float64 `db:"sell_price"`
	Amount      int     `db:"amount"`
	DateCreate  string  `db:"date_create"`
}
type SelectProduct struct {
	ID           int     `db:"id"`
	VendorCode   string  `db:"vendor_code"`
	ProductName  string  `db:"product_name"`
	SupplierID   int     `db:"supplier_id"`
	CompanyName  string  `db:"company_name"`
	CategoryID   int     `db:"category_id"`
	CategoryName string  `db:"category_name"`
	BuyPrice     float64 `db:"buy_price"`
	SellPrice    float64 `db:"sell_price"`
	Amount       int     `db:"amount"`
	DateCreate   string  `db:"date_create"`
}

//type Decommissioned struct {
//	ID          int     `db:"id"`
//	ProductID   int     `db:"product_id"`
//	SellPrice   float64 `db:"sell_price"`
//	Amount      int     `db:"amount"`
//	Comment     string  `db:"comment"`
//}

var ProductColumns = []string{
	"vendor_code",
	"product_name",
	"supplier_id",
	"category_id",
	"buy_price",
	"sell_price",
	"amount",
	"date_create",
}
