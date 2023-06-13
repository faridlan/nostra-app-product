package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthController interface {
	Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Profile(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UploadIamge(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
