package helper

import (
	"mime/multipart"
	"net/http"
)

func MultipartForm(formName string, request *http.Request) multipart.File {

	err := request.ParseMultipartForm(10 << 20)
	PanicIfError(err)

	file, _, err := request.FormFile(formName)
	PanicIfError(err)

	defer file.Close()

	return file

}
