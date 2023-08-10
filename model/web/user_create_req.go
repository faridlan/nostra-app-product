package web

type UserCreateReq struct {
	UserId    string `json:"user_id,omitempty"`
	Username  string `json:"username,omitempty" validate:"required,gte=5"`
	Password  string `json:"password,omitempty" validate:"required,gte=8"`
	Email     string `json:"email,omitempty" validate:"required,email"`
	Image     string `json:"image,omitempty"`
	RoleId    string `json:"role_id,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
}