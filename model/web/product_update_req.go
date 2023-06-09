package web

type ProductUpdateReq struct {
	ProductId   string `json:"product_id,omitempty"`
	Name        string `json:"name,omitempty" validate:"required,gte=5"`
	Price       int    `json:"price,omitempty" validate:"required"`
	Quantity    int    `json:"quantity,omitempty" validate:"required"`
	Description string `json:"description,omitempty" validate:"required,gte=10"`
	Image       string `json:"image,omitempty"`
	CategoryId  string `json:"category_id,omitempty" validate:"required"`
	UpdatedAt   int64  `json:"updated_at,omitempty"`
}
