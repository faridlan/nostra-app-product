package web

type UserCreateReq struct {
	Id        string `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	Email     string `json:"email,omitempty"`
	Image     string `json:"image,omitempty"`
	RoleId    string `json:"role_id,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
}
