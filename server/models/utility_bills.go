package models

import (
	"github.com/gocraft/dbr"
)

const UtilTableName = "utility_bills"

type Util struct {
	ID          int     `db:"id"`
	Electricity float64 `db:"electricity"`
	Water       float64 `db:"water"`
	Gas         float64 `db:"gas"`
	DateCreate  string  `db:"date_create"`
	Comment     string  `db:"comment"`
}

type ResponseUtil struct {
	ID          int            `db:"id"`
	Electricity float64        `db:"electricity"`
	Water       float64        `db:"water"`
	Gas         float64        `db:"gas"`
	DateCreate  string         `db:"date_create"`
	Credit      float64        `db:"credit"`
	BillID      float64        `db:"bill_id"`
	Comment     dbr.NullString `db:"comment"`
}

type ReportUtil struct {
	Date   string  `db:"date_create"`
	Credit float64 `db:"credit"`
}

var UtilityColumns = []string{
	"electricity",
	"water",
	"gas",
	"date_create",
	"comment",
}
