package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Fatal(err)
	}

	return string(bytes), nil
}
