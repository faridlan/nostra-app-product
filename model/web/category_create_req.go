package web

type CategoryCreateReq struct {
	Name      string `json:"name,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
}
