package models

const SuppliersTableName = "suppliers"

type Supplier struct {
	ID           int    `db:"id"`
	CompanyName  string `db:"company_name"`
	ContactName  string `db:"contact_name"`
	ContactTitle string `db:"contact_title"`
	Address      string `db:"address"`
	City         string `db:"city"`
	Phone        string `db:"phone"`
	Email        string `db:"email"`
}

var SupplierColumns = []string{
	"company_name",
	"contact_name",
	"contact_title",
	"address",
	"city",
	"phone",
	"email",
}
