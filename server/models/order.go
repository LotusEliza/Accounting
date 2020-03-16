package models

const OrdersTableName = "orders"
const CustomersTableName = "customers"

type Order struct {
	ID         int    `db:"id"`
	CustomerID int    `db:"customer_id"`
	ProductID  int    `db:"product_id"`
	Amount     int    `db:"amount"`
	DateCreate string `db:"date_create"`
}

type RespOrder struct {
	ID           int    `db:"id"`
	CustomerName string `db:"customer_name"`
	DateCreate   string `db:"date_create"`
	Telephone    string `db:"phone"`
	Product      string `db:"product_name"`
	Amount       int    `db:"amount"`
	CustomerID   int    `db:"customer_id"`
	ProductID    int    `db:"product_id"`
}

type Customer struct {
	ID           int    `db:"id"`
	CustomerName string `db:"customer_name"`
	Phone        string `db:"phone"`
	ChatID       int    `db:"chat_id"`
}

var OrderColumns = []string{
	"customer_id",
	"product_id",
	"amount",
	"date_create",
}

var CustomerColumns = []string{
	"customer_name",
	"phone",
	"chat_id",
}

type LastCustomerId struct {
	ID int `db:"id"`
}
