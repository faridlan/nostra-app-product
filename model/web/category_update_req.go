package web

type CategoryUpdateReq struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}
