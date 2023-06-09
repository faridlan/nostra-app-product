package domain

import "github.com/faridlan/nostra-api-product/helper/mysql"

type Category struct {
	Id         int
	CategoryId string
	Name       string
	CreatedAt  int64
	UpdatedAt  *mysql.NullInt
}
