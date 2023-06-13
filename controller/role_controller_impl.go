package controller

import (
	"net/http"

	"github.com/faridlan/nostra-api-product/helper"
	"github.com/faridlan/nostra-api-product/model/web"
	"github.com/faridlan/nostra-api-product/service"
	"github.com/julienschmidt/httprouter"
)

type RoleControllerImpl struct {
	RoleService service.RoleService
}

func NewRoleController(roleSerivce service.RoleService) RoleController {
	return &RoleControllerImpl{
		RoleService: roleSerivce,
	}
}

func (controller *RoleControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleCreate := web.RoleCreateReq{}
	helper.ReadFromRequestBody(request, &roleCreate)

	roleResponse := controller.RoleService.Create(request.Context(), roleCreate)

	WebResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roleResponse,
	}

	helper.WriteToResponseBody(writer, WebResponse)
}

func (controller *RoleControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleUpdate := web.RoleUpdateReq{}
	helper.ReadFromRequestBody(request, &roleUpdate)

	id := params.ByName("roleId")

	roleUpdate.RoleId = id

	roleResponse := controller.RoleService.Update(request.Context(), roleUpdate)

	WebResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roleResponse,
	}

	helper.WriteToResponseBody(writer, WebResponse)
}

func (controller *RoleControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleId := params.ByName("roleId")

	roleResponse := controller.RoleService.FindById(request.Context(), roleId)

	WebResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roleResponse,
	}

	helper.WriteToResponseBody(writer, WebResponse)
}

func (controller *RoleControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleResponses := controller.RoleService.FindAll(request.Context())

	WebResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roleResponses,
	}

	helper.WriteToResponseBody(writer, WebResponse)
}
