package domain

import (
	"github.com/faridlan/nostra-api-product/helper/mysql"
)

type Product struct {
	ProductId   int
	Id          string
	Name        string
	Price       int
	Quantity    int
	Description string
	Image       *mysql.NullString
	Category    Category
	CreatedAt   int64
	UpdatedAt   *mysql.NullInt
}
