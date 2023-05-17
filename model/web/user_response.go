package web

import "github.com/faridlan/nostra-api-product/helper/mysql"

type UserResponse struct {
	Id        string            `json:"id,omitempty"`
	Username  string            `json:"username,omitempty"`
	Email     string            `json:"email,omitempty"`
	Image     *mysql.NullString `json:"image,omitempty"`
	RoleId    string            `json:"role_id,omitempty"`
	CreatedAt int64             `json:"created_at,omitempty"`
	UpdatedAt *mysql.NullInt    `json:"updated_at,omitempty"`
}

type LoginResponse struct {
	User  *UserResponse `json:"user,omitempty"`
	Token string        `json:"token,omitempty"`
}
