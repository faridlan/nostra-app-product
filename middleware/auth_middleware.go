package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/faridlan/nostra-api-product/helper"
	"github.com/faridlan/nostra-api-product/model/web"
	"github.com/golang-jwt/jwt/v4"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (authMiddleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	authorizationHeader := request.Header.Get("Authorization")

	if request.URL.Path == "/api/users/login" {
		authMiddleware.Handler.ServeHTTP(writer, request)
		return
	}

	if request.URL.Path == "/api/users" && request.Method == "POST" {
		authMiddleware.Handler.ServeHTTP(writer, request)
		return
	}

	if !strings.Contains(authorizationHeader, "Bearer") {

		writer.Header().Add("Content-Type", "application.json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		// logging.ProductLoggerError(webResponse, writer, request, "Auth Header Bearer Not Found")
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	var claim = &web.JWTClaim{}

	token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("signing method invalid")
		}
		return web.JwtSecret, nil
	})

	if err != nil {
		writer.Header().Add("Content-Type", "application.json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		// logging.ProductLoggerError(webResponse, writer, request, err)
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	if !token.Valid {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	} else {

		endpoints := helper.UserEndpoints(request)

		for _, enpoint := range endpoints {
			if request.URL.Path == enpoint.Url && request.Method == enpoint.Method && claim.RoleId == "d11cd32cfa4811edbc140242ac130002" {

				authMiddleware.Handler.ServeHTTP(writer, request)
				return

			} else if claim.RoleId != "d11cd32cfa4811edbc140242ac130002" {

				authMiddleware.Handler.ServeHTTP(writer, request)
				return

			}
		}

		writer.Header().Add("Content-Type", "application.json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return

	}

	// authMiddleware.Handler.ServeHTTP(writer, request)

}
