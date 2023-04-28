package web

type ProductCreateReq struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Price       int    `json:"price,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image"`
	CategoryId  string `json:"category_id,omitempty"`
	CreatedAt   int64  `json:"created_at,omitempty"`
}
