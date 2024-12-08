package testutils

import (
	"time"

	"github.com/teixeiragthiago/api-user/internal/dto"
	"github.com/teixeiragthiago/api-user/internal/entity"
	"github.com/teixeiragthiago/api-user/internal/util"
)

func MockValidUser() *entity.User {

	hashedPassword, _ := util.EncryptPassword("SenH@123!al")
	return &entity.User{
		ID:        1,
		Name:      "Thiago",
		Email:     "thiago@teste.com",
		Nickname:  "thiago_teste",
		Password:  hashedPassword,
		CreatedAt: time.Now().Truncate(time.Millisecond),
		Active:    true,
	}
}

func MockInvalidUser() *entity.User {

	hashedPassword, _ := util.EncryptPassword("123")
	return &entity.User{
		ID:        0,
		Name:      "",
		Email:     "thiago.com",
		Nickname:  "",
		Password:  hashedPassword,
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

func MockValidUserLoginDto() *dto.UserLoginDto {
	return &dto.UserLoginDto{
		Email:    "thiago@email.com",
		Password: "SenH@123!al",
	}
}

func MockInvalidUserLoginDto() *dto.UserLoginDto {
	return &dto.UserLoginDto{
		Email:    "thiago@email.com",
		Password: "inválida",
	}
}
