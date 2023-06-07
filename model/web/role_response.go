package web

import "github.com/faridlan/nostra-api-product/helper/mysql"

type RoleResponse struct {
	RoleId    string         `json:"role_id,omitempty"`
	Name      string         `json:"name,omitempty"`
	CreatedAt int64          `json:"created_at,omitempty"`
	UpdatedAt *mysql.NullInt `json:"updated_at,omitempty"`
}
