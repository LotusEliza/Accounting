package models

const CategoriesTableName = "categories"

type Category struct {
	ID         		     int       `db:"id"`
	CategoryName         string    `db:"category_name"`
	CategoryDescription  string    `db:"category_description"`
}

var CategoryColumns = []string{
	"category_name",
	"category_description",
}


