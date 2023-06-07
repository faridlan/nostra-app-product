package web

type UserUpdateReq struct {
	Id        int    `json:"id,omitempty"`
	UserId    string `json:"user_id,omitempty"`
	Username  string `json:"username,omitempty" validate:"required,gte=5"`
	Email     string `json:"email,omitempty" validate:"required,email"`
	Image     string `json:"image,omitempty"`
	RoleId    string `json:"role_id,omitempty" validate:"required"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}
