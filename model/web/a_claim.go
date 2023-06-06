package web

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JwtSigningMEethod = jwt.SigningMethodHS256
var JwtSecret = []byte("Anj1ngAd4l4hH3w4nuuAAii8sdA73DFed7")
var ExpiredTime = time.Now().Add(time.Minute * 1)

type JWTClaim struct {
	Id       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	RoleId   string `json:"role_id,omitempty"`
	RoleName string `json:"role_name,omitempty"`
	Token    string `json:"token,omitempty"`
	jwt.RegisteredClaims
}
