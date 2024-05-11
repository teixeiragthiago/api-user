package mapper

import (
	"log"

	"github.com/teixeiragthiago/api-user/internal/dto"
	"github.com/teixeiragthiago/api-user/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

func MapDtoToEntity(userDTO *dto.UserDTO) *entity.User {

	hashedPassword, _ := hashPassword(userDTO.Password)

	return &entity.User{
		ID:       userDTO.ID,
		Name:     userDTO.Name,
		Nick:     userDTO.Nick,
		Password: hashedPassword,
		Active:   true,
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Fatal(err)
	}

	return string(bytes), nil
}

func MapEntityToResponseDto(user *entity.User) *dto.UserResponseDto {
	return &dto.UserResponseDto{
		ID:        user.ID,
		Name:      user.Name,
		Nick:      user.Nick,
		Active:    user.Active,
		CreatedAt: user.CreatedAt,
	}
}
