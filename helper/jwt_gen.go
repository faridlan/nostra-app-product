package helper

import (
	"time"

	"github.com/faridlan/nostra-api-product/model/domain"
	"github.com/faridlan/nostra-api-product/model/web"
	"github.com/golang-jwt/jwt/v4"
)

func JwtGen(user domain.User) string {

	strRandom := RandStringRunes(20)

	expirationtime := time.Now().Add(5 * time.Minute)

	claim := &web.JWTClaim{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		RoleId:   user.Role.Id,
		Token:    strRandom,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationtime),
		},
	}

	token := jwt.NewWithClaims(web.JwtSigningMEethod, claim)
	tokenString, err := token.SignedString(web.JwtSecret)
	PanicIfError(err)

	return tokenString

}
