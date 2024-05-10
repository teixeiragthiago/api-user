package mapper

import (
	"github.com/teixeiragthiago/api-user/internal/dto"
	"github.com/teixeiragthiago/api-user/internal/entity"
)

func MapDtoToEntity(userDTO *dto.UserDTO) *entity.User {
	return &entity.User{
		ID:       userDTO.ID,
		Name:     userDTO.Name,
		Nick:     userDTO.Nick,
		Password: userDTO.Password,
		Active:   true,
	}
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
