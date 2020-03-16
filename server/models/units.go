package models

const UnitsTableName = "units"

type Unit struct {
	ID       int    `db:"id"`
	UnitName string `db:"unit_name"`
}

var UnitColumns = []string{
	"unit_name",
}
