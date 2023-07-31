package web

import "github.com/faridlan/nostra-api-product/helper/mysql"

type ProductResponse struct {
	ProductId   string `json:"product_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Price       int    `json:"price,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
	Description string `json:"description,omitempty"`
	// Image       *mysql.NullString `json:"image"`
	Image     []string          `json:"image"`
	Category  *CategoryResponse `json:"category,omitempty"`
	CreatedAt int64             `json:"created_at,omitempty"`
	UpdatedAt *mysql.NullInt    `json:"updated_at"`
}
