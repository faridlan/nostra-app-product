package web

type RoleCreateReq struct {
	RoleId    string `json:"role_id,omitempty"`
	Name      string `json:"name,omitempty" validate:"required,gte=3"`
	CreatedAt int64  `json:"created_at,omitempty"`
}
