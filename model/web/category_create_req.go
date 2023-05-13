package web

type CategoryCreateReq struct {
	Name      string `json:"name,omitempty" validate:"required,gte=3"`
	CreatedAt int64  `json:"created_at,omitempty"`
}
