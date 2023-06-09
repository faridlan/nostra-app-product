package web

import "github.com/faridlan/nostra-api-product/helper/mysql"

type CategoryResponse struct {
	CategoryId string         `json:"category_id,omitempty"`
	Name       string         `json:"name,omitempty"`
	CreatedAt  int64          `json:"created_at,omitempty"`
	UpdatedAt  *mysql.NullInt `json:"updated_at,omitempty"`
}
