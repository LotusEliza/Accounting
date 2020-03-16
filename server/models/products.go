package models

const ProductsTableName = "products"

type Product struct {
	ID          int    `db:"id"`
	VendorCode  string `db:"vendor_code"`
	ProductName string `db:"product_name"`
	CategoryID  int    `db:"category_id"`
	ImgLink     string `db:"img_link"`
}

type SelectProduct struct {
	ID           int    `db:"id"`
	VendorCode   string `db:"vendor_code"`
	ProductName  string `db:"product_name"`
	CategoryID   int    `db:"category_id"`
	CategoryName string `db:"category_name"`
	ImgLink      string `db:"img_link"`
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
	"category_id",
	"img_link",
}
