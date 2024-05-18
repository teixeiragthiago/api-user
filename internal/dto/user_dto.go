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
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type UserResponseDto struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Nickname  string    `json:"nickname"`
	Active    bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
}

func (u *UserDTO) Validate() error {
	if u.Name == "" {
		return errors.New("name cannot be empty")
	}

	if u.Nickname == "" {
		return errors.New("nickname cannot be empty")
	}

	if err := validatePassword(u.Password); err != nil {
		return err
	}

	if err := validateEmail(u.Email); err != nil {
		return err
	}

	return nil
}

func isValidEmail(email string) bool {

	regex := `^[\w.-]+@[a-zA-Z\d.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(regex)

	return re.MatchString(email)
}

func validateEmail(email string) error {

	if email == "" {
		return errors.New("e-mail cannot be empty")
	}

	if !isValidEmail(email) {
		return errors.New("invalid e-mail")
	}

	return nil
}

func validatePassword(password string) error {
	if password == "" {
		return errors.New("password cannot be empty")
	}

	if len(password) > 12 {
		return errors.New("password cannot have more than 12 characters")
	}

	return nil
}
