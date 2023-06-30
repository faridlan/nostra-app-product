package multipart

import (
	"mime/multipart"
	"net/http"

	"github.com/faridlan/nostra-api-product/exception"
)

func MultipartForm(formName string, request *http.Request) multipart.File {

	err := request.ParseMultipartForm(10 << 20)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	file, _, err := request.FormFile(formName)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}
	defer file.Close()

	return file

}
