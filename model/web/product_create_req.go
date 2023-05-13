package web

type ProductCreateReq struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty" validate:"required,gte=5"`
	Price       int    `json:"price,omitempty" validate:"required"`
	Quantity    int    `json:"quantity,omitempty" validate:"required"`
	Description string `json:"description,omitempty" validate:"required,gte=10"`
	Image       string `json:"image"`
	CategoryId  string `json:"category_id,omitempty" validate:"required"`
	CreatedAt   int64  `json:"created_at,omitempty"`
}
