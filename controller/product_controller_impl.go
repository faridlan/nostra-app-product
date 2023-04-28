package controller

import (
	"embed"
	"encoding/json"
	"net/http"

	"github.com/faridlan/nostra-api-product/helper"
	"github.com/faridlan/nostra-api-product/model/web"
	"github.com/faridlan/nostra-api-product/service"
	"github.com/julienschmidt/httprouter"
)

//go:embed json/products.json

var Json embed.FS

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productCreate := web.ProductCreateReq{}
	helper.ReadFromRequestBody(request, &productCreate)

	product := controller.ProductService.Create(request.Context(), productCreate)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   product,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productUpdate := web.ProductUpdateReq{}
	helper.ReadFromRequestBody(request, &productUpdate)

	id := params.ByName("productId")
	productUpdate.Id = id

	product := controller.ProductService.Update(request.Context(), productUpdate)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   product,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("productId")

	controller.ProductService.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("productId")

	product := controller.ProductService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   product,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	products := controller.ProductService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   products,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) SeederCreate(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	products, err := Json.ReadFile("json/products.json")
	helper.PanicIfError(err)

	productCreate := []web.ProductCreateReq{}
	json.Unmarshal(products, &productCreate)

	product := controller.ProductService.CreateMany(request.Context(), productCreate)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   product,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) SeederDelete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	controller.ProductService.DeleteAll(request.Context())

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
