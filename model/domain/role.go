package domain

import "github.com/faridlan/nostra-api-product/helper/mysql"

type Role struct {
	RoleId    int
	Id        string
	Name      string
	CreatedAt int64
	UpdatedAt *mysql.NullInt
}
