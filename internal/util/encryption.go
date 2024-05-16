package util

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) ([]byte, error) {

	if password == "" {
		return nil, errors.New("password cannot be empty")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Fatal(err)
	}

	return hashedPassword, nil
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
