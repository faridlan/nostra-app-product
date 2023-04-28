package domain

import (
	"github.com/faridlan/nostra-api-product/helper/mysql"
)

type User struct {
	UserId    int
	Id        string
	Username  string
	Password  string
	Email     string
	Image     *mysql.NullString
	RoleId    string
	CreatedAt int64
	UpdatedAt *mysql.NullInt
}
