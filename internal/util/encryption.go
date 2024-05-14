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

func MatchPassword(password string, salt []byte) (bool, error) {

	if password == "" {
		return false, errors.New("password cannot be empty")
	}

	err := bcrypt.CompareHashAndPassword([]byte(password), salt)
	if err != nil {
		return false, err
	}

	return true, nil
}
