package controller

import (
	"embed"
	"encoding/json"
	"net/http"

	"github.com/faridlan/nostra-api-product/helper"
	"github.com/faridlan/nostra-api-product/helper/auth"
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

//go:embed json/users.json

var JsonUsers embed.FS

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

func (controller *AuthControllerImpl) CreateMany(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	users, err := JsonUsers.ReadFile("json/users.json")
	helper.PanicIfError(err)

	usersCreate := []web.UserCreateReq{}
	json.Unmarshal(users, &usersCreate)

	userResponses := controller.AuthService.SaveMany(request.Context(), usersCreate)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AuthControllerImpl) DeleteAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	controller.AuthService.DeleteAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *AuthControllerImpl) UploadIamge(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	file := helper.MultipartForm("userImage", request)

	uploadResponse := controller.Upload.Upload(file, "users")

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   uploadResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
