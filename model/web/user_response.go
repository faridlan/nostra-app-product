package web

import "github.com/faridlan/nostra-api-product/helper/mysql"

type UserResponse struct {
	UserId    string            `json:"user_id,omitempty"`
	Username  string            `json:"username,omitempty"`
	Email     string            `json:"email,omitempty"`
	Image     *mysql.NullString `json:"image"`
	Role      *RoleResponse     `json:"role,omitempty"`
	CreatedAt int64             `json:"created_at,omitempty"`
	UpdatedAt *mysql.NullInt    `json:"updated_at"`
}

type LoginResponse struct {
	User  *UserResponse `json:"user,omitempty"`
	Token string        `json:"token,omitempty"`
}
