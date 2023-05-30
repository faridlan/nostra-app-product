package helper

import (
	"fmt"
	"net/http"
	"strings"
)

type Endpoint struct {
	Url    string
	Method string
}

func UserEndpoints(request *http.Request) []Endpoint {

	prefix := 0
	if strings.HasPrefix(request.URL.Path, "/api/products/") {
		prefix = len("/api/products/")
	}

	return []Endpoint{
		{
			Url:    "/api/categories",
			Method: "GET",
		},
		{
			Url:    "/api/products",
			Method: "GET",
		},
		{
			Url:    fmt.Sprintf("/api/products/%s", request.URL.Path[prefix:]),
			Method: "GET",
		},
		{
			Url:    "/api/users/profile",
			Method: "GET",
		},
	}
}
