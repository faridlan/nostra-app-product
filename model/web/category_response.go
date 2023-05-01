package web

import "github.com/faridlan/nostra-api-product/helper/mysql"

type CategoryResponse struct {
	Id        string
	Name      string
	CreatedAt int64
	UpdatedAt *mysql.NullInt
}
