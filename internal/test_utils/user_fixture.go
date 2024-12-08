package testutils

import (
	"time"

	"github.com/teixeiragthiago/api-user/internal/dto"
	"github.com/teixeiragthiago/api-user/internal/entity"
)

func MockValidUser() *entity.User {
	return &entity.User{
		ID:        1,
		Name:      "Thiago",
		Email:     "thiago@teste.com",
		Nickname:  "thiago_teste",
		Password:  []byte("hashedpassword"),
		CreatedAt: time.Now().Truncate(time.Millisecond),
		Active:    true,
	}
}

func MockValidUserDto() *dto.UserDTO {
	return &dto.UserDTO{
		ID:       1,
		Name:     "Thiago",
		Email:    "thiago@teste.com",
		Nickname: "thiago_teste",
		Password: "SenH@123!al",
	}
}

func MockInvalidUserDto() *dto.UserDTO {
	return &dto.UserDTO{
		Name:     "",
		Email:    "",
		Nickname: "",
		Password: "123",
	}
}
