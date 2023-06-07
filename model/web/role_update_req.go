package web

type RoleUpdateReq struct {
	RoleId    string `json:"role_id,omitempty"`
	Name      string `json:"name,omitempty" validate:"required,gte=3"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}
