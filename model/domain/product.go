package domain

import (
	"github.com/faridlan/nostra-api-product/helper/mysql"
)

type Product struct {
	Id          int
	ProductId   string
	Name        string
	Price       int
	Quantity    int
	Description string
	ImageSingle *mysql.NullString
	Image       []string
	Category    Category
	CreatedAt   int64
	UpdatedAt   *mysql.NullInt
}
