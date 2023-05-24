package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/faridlan/nostra-api-product/helper"
	"github.com/faridlan/nostra-api-product/model/web"
	"github.com/golang-jwt/jwt/v4"
)

func GetIdProfile(request *http.Request) string {

	authorizationHeader := request.Header.Get("Authorization")

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	var claim = &web.JWTClaim{}

	_, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("signing method invalid")
		}
		return web.JwtSecret, nil
	})

	helper.PanicIfError(err)

	return claim.Id
}
