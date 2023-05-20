package hash

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(hashedPwd string, plainPwd []byte) error {

	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return errors.New("username or password incorect")
	} else {
		return nil
	}
}
