package web

type CategoryUpdateReq struct {
	CategoryId string `json:"category_id,omitempty"`
	Name       string `json:"name,omitempty" validate:"required,gte=5"`
	UpdatedAt  int64  `json:"updated_at,omitempty"`
}
