package models

import (
	"github.com/gocraft/dbr"
)

const CreditTableName = "credit"

type Credit struct {
	ID       int     `db:"id"`
	Credit   float64 `db:"credit"`
	SupplyID int     `db:"supply_id"`
	BillID   int     `db:"bill_id"`
}

var CreditColumns = []string{
	"credit",
	"supply_id",
	"bill_id",
}

type Report struct {
	Date    dbr.NullString `db:"date_create"`
	Comment dbr.NullString `db:"comment"`
	Credit  float64        `db:"credit"`
}

type ReportDate struct {
	From string
	To   string
}

//type Report struct {
//	ID          int     `db:"id"`
//	Credit      float64 `db:"credit"`
//	//Amount      int     `db:"amount"`
//	//UnitID      int     `db:"unit_id"`
//	//UnitName    string  `db:"unit_name"`
//	//Comment     string  `db:"comment"`
//}
