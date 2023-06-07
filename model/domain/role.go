package domain

import "github.com/faridlan/nostra-api-product/helper/mysql"

type Role struct {
	Id        int
	RoleId    string
	Name      string
	CreatedAt int64
	UpdatedAt *mysql.NullInt
}
