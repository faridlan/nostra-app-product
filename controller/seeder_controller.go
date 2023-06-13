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

//go:embed json/roles.json
//go:embed json/users.json
//go:embed json/categories.json
//go:embed json/products.json

var JsonEmbed embed.FS

// var JsonUsers embed.FS
// var JsonCategories embed.FS
// var JsonProducts embed.FS

type SeederController interface {
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	SaveMany(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type SeederControllerImpl struct {
	SeederService service.SeederService
}

func NewSeederController(seederService service.SeederService) SeederController {
	return &SeederControllerImpl{
		SeederService: seederService,
	}
}

func (controller *SeederControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	controller.SeederService.Delete(request.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SeederControllerImpl) SaveMany(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	roles, err := JsonEmbed.ReadFile("json/roles.json")
	helper.PanicIfError(err)

	rolesCreate := []web.RoleCreateReq{}
	json.Unmarshal(roles, &rolesCreate)

	//
	users, err := JsonEmbed.ReadFile("json/users.json")
	helper.PanicIfError(err)

	usersCreate := []web.UserCreateReq{}
	json.Unmarshal(users, &usersCreate)

	//
	categories, err := JsonEmbed.ReadFile("json/categories.json")
	helper.PanicIfError(err)

	categoriesCreate := []web.CategoryCreateReq{}
	json.Unmarshal(categories, &categoriesCreate)

	//
	products, err := JsonEmbed.ReadFile("json/products.json")
	helper.PanicIfError(err)

	productsCreate := []web.ProductCreateReq{}
	json.Unmarshal(products, &productsCreate)

	controller.SeederService.SaveMany(request.Context(), rolesCreate, usersCreate, categoriesCreate, productsCreate)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
