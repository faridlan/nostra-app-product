package controller

import (
	"net/http"
	"strings"

	"github.com/faridlan/nostra-api-product/helper"
	"github.com/faridlan/nostra-api-product/helper/auth"
	"github.com/faridlan/nostra-api-product/helper/multipart"
	"github.com/faridlan/nostra-api-product/model/web"
	"github.com/faridlan/nostra-api-product/service"
	"github.com/julienschmidt/httprouter"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
	Upload      service.UploadS3AWS
}

func NewAuthController(authService service.AuthService, upload service.UploadS3AWS) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
		Upload:      upload,
	}
}

func (controller *AuthControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreate := web.UserCreateReq{}
	helper.ReadFromRequestBody(request, &userCreate)

	user := controller.AuthService.Register(request.Context(), userCreate)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   user,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AuthControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdate := web.UserUpdateReq{}
	helper.ReadFromRequestBody(request, &userUpdate)

	Id := params.ByName("userId")
	userUpdate.UserId = Id

	user := controller.AuthService.Update(request.Context(), userUpdate)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   user,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *AuthControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	Id := params.ByName("userId")

	user := controller.AuthService.FindById(request.Context(), Id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   user,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AuthControllerImpl) Profile(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	Id := auth.GetIdProfile(request)

	user := controller.AuthService.FindById(request.Context(), Id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   user,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AuthControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	users := controller.AuthService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   users,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	loginReq := web.Login{}
	helper.ReadFromRequestBody(request, &loginReq)

	user := controller.AuthService.Login(request.Context(), loginReq)

	webRespone := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   user,
	}

	helper.WriteToResponseBody(writer, webRespone)
}

func (controller *AuthControllerImpl) UploadIamge(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	file := multipart.MultipartForm("userImage", request)

	uploadResponse := controller.Upload.Upload(file, "nostra-app/users")

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   uploadResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AuthControllerImpl) Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	auth := request.Header.Get("Authorization")
	authString := strings.Replace(auth, "Bearer ", "", -1)
	controller.AuthService.DeleteWL(request.Context(), authString)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
