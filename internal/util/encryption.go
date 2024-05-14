package util

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {

	if password == "" {
		return "", errors.New("password cannot be empty")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Fatal(err)
	}

	return string(bytes), nil
}
