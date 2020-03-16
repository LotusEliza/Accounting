package models

const DebitTableName = "debit"

type Debit struct {
	ID         int     `db:"id"`
	Debit      float64 `db:"debit"`
	DateCreate string  `db:"date_create"`
}

type LastId struct {
	ID int `db:"id"`
}

var DebitColumns = []string{
	"debit",
	"date_create",
}

type ReportDebit struct {
	ID    int     `db:"id"`
	Date  string  `db:"date_create"`
	Debit float64 `db:"debit"`
}

//
//type Report struct {
//	Date    string  `db:"date_create"`
//	Comment string  `db:"comment"`
//	Credit  float64 `db:"credit"`
//}
//
//type ReportDate struct {
//	From string
//	To   string
//}

//type Report struct {
//	ID          int     `db:"id"`
//	Credit      float64 `db:"credit"`
//	//Amount      int     `db:"amount"`
//	//UnitID      int     `db:"unit_id"`
//	//UnitName    string  `db:"unit_name"`
//	//Comment     string  `db:"comment"`
//}
