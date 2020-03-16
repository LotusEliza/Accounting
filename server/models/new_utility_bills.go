package models

const NewUtilTableName = "new_utility_bills"

type NewUtil struct {
	ID         int     `db:"id"`
	Amount     float64 `db:"amount"`
	BillName   string  `db:"bill_name"`
	DateCreate string  `db:"date_create"`
}

//type ResponseUtil struct {
//	ID          int            `db:"id"`
//	Electricity float64        `db:"electricity"`
//	Water       float64        `db:"water"`
//	Gas         float64        `db:"gas"`
//	DateCreate  string         `db:"date_create"`
//	Credit      float64        `db:"credit"`
//	BillID      float64        `db:"bill_id"`
//	Comment     dbr.NullString `db:"comment"`
//}
//
type UtilName struct {
	BillName string `db:"bill_name"`
}

type ReportUtilBill struct {
	Date     string  `db:"date_create"`
	BillName string  `db:"bill_name"`
	Amount   float64 `db:"amount"`
}

var NewUtilityColumns = []string{
	"amount",
	"bill_name",
	"date_create",
}
