package web

import "github.com/faridlan/nostra-api-product/helper/mysql"

type ProductResponse struct {
	Id          string            `json:"id,omitempty"`
	Name        string            `json:"name,omitempty"`
	Price       int               `json:"price,omitempty"`
	Quantity    int               `json:"quantity,omitempty"`
	Description string            `json:"description,omitempty"`
	Image       *mysql.NullString `json:"image"`
	CategoryId  string            `json:"category_id,omitempty"`
	CreatedAt   int64             `json:"created_at,omitempty"`
	UpdatedAt   *mysql.NullInt    `json:"updated_at"`
}
