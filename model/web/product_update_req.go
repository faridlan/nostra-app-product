package web

type ProductUpdateReq struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Price       int    `json:"price,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image"`
	CategoryId  string `json:"category_id,omitempty"`
	UpdatedAt   int64  `json:"updated_at,omitempty"`
}
