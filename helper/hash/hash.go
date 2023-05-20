package hash

import (
	"github.com/faridlan/nostra-api-product/helper"
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	helper.PanicIfError(err)

	return string(hash)

}
