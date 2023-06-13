package controller

import (
	"net/http"

	"github.com/faridlan/nostra-api-product/helper"
	"github.com/faridlan/nostra-api-product/model/web"
	"github.com/faridlan/nostra-api-product/service"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateReq := web.CategoryCreateReq{}
	helper.ReadFromRequestBody(request, &categoryCreateReq)

	category := controller.CategoryService.Create(request.Context(), categoryCreateReq)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   category,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateReq := web.CategoryUpdateReq{}
	helper.ReadFromRequestBody(request, &categoryUpdateReq)
	id := params.ByName("categoryId")

	categoryUpdateReq.CategoryId = id

	category := controller.CategoryService.Update(request.Context(), categoryUpdateReq)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   category,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("categoryId")

	controller.CategoryService.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("categoryId")

	category := controller.CategoryService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   category,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categories := controller.CategoryService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categories,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
