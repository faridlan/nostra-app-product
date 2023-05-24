package middleware

import (
	"fmt"
	"log"
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

	// enpoints := exception.EndpointsGlobal()

	// for _, Enpoint := range enpoints {
	// 	if request.URL.Path != Enpoint.Url && request.Method != Enpoint.Method {
	// 		authMiddleware.Handler.ServeHTTP(writer, request)
	// 		return
	// 	}
	// }

	// if request.URL.Path != "/api/seeder/products" && request.URL.Path != "api/log" {
	// 	authMiddleware.Handler.ServeHTTP(writer, request)
	// 	return
	// }

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
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if request.URL.Path == "/api/users/profile" && request.Method == "GET" && claim.RoleId != "7f03c5c7f97711ed9f900242ac130002" {
		writer.Header().Add("Content-Type", "application.json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	log.Println(claim.RoleId)
	authMiddleware.Handler.ServeHTTP(writer, request)

}
