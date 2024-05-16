package dto

import (
	"errors"
	"regexp"
	"time"
)

type UserDTO struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Nick     string `json:"nick"`
	Password string `json:"password"`
}

type UserResponseDto struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Nick      string    `json:"nick"`
	Active    bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
}

func (u *UserDTO) Validate() error {
	if u.Name == "" {
		return errors.New("name cannot be empty")
	}

	if u.Nick == "" {
		return errors.New("nick cannot be empty")
	}

	if u.Email == "" {
		return errors.New("e-mail cannot be empty")
	}

	emailMatch, _ := regexp.MatchString("/^[a-z0-9.]+@[a-z0-9]+\\.[a-z]+\\.([a-z]+)?$/i", u.Email)
	if !emailMatch {
		return errors.New("invalid e-mail")
	}

	if u.Password == "" {
		return errors.New("password cannot be empty")
	}

	if len(u.Password) > 12 {
		return errors.New("password cannot have more than 12 characters")
	}

	return nil
}
