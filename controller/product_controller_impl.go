package controller

import (
	"net/http"

	"github.com/faridlan/nostra-api-product/exception"
	"github.com/faridlan/nostra-api-product/helper"
	"github.com/faridlan/nostra-api-product/helper/multipart"
	"github.com/faridlan/nostra-api-product/model/web"
	"github.com/faridlan/nostra-api-product/service"
	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
	UploadService  service.UploadS3AWS
}

func NewProductController(productService service.ProductService, uploadService service.UploadS3AWS) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
		UploadService:  uploadService,
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

func (controller *ProductControllerImpl) CreateMany(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productsCreate := []web.ProductCreateReq{}
	helper.ReadFromRequestBody(request, &productsCreate)

	products := controller.ProductService.CreateMany(request.Context(), productsCreate)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   products,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productUpdate := web.ProductUpdateReq{}
	helper.ReadFromRequestBody(request, &productUpdate)

	id := params.ByName("productId")
	productUpdate.ProductId = id

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

func (controller *ProductControllerImpl) UploadImage(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	file := multipart.MultipartForm("productImage", request)
	defer file.Close()

	upload := controller.UploadService.Upload(file, "nostra-app/products")
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   upload,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) UploadImageBatch(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	err := request.ParseMultipartForm(10 << 20)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	images := request.MultipartForm.File["productImage"]
	if len(images) == 0 {
		panic(exception.NewBadRequestError("No Such File"))
	}
	uploads := []web.UploadResponse{}

	for _, fh := range images {
		file, err := fh.Open()
		if err != nil {
			panic(exception.NewBadRequestError(err.Error()))
		}

		defer file.Close()

		url := controller.UploadService.Upload(file, "nostra-app/products")

		uploads = append(uploads, url)
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   uploads,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
