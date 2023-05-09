package web

type UserUpdateReq struct {
	UserId    int    `json:"user_id,omitempty"`
	Id        string `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	Image     string `json:"image,omitempty"`
	RoleId    string `json:"role_id,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}
