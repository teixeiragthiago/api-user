package mapper

import (
	"github.com/teixeiragthiago/api-user/internal/dto"
	"github.com/teixeiragthiago/api-user/internal/entity"
	"github.com/teixeiragthiago/api-user/internal/util"
)

func MapDtoToEntity(userDTO *dto.UserDTO) *entity.User {

	hashedPassword, _ := util.EncryptPassword(userDTO.Password)

	return &entity.User{
		ID:       userDTO.ID,
		Name:     userDTO.Name,
		Nickname: userDTO.Nickname,
		Password: hashedPassword,
		Active:   true,
	}
}

func MapEntityToResponseDto(user *entity.User) *dto.UserResponseDto {
	return &dto.UserResponseDto{
		ID:        user.ID,
		Name:      user.Name,
		Nickname:  user.Nickname,
		Active:    user.Active,
		CreatedAt: user.CreatedAt,
	}
}

func MapEntitiesToResponseDto(users []*entity.User) []*dto.UserResponseDto {

	if len(users) == 0 {
		return nil
	}

	var usersDTO []*dto.UserResponseDto
	for _, user := range users {
		userDTO := MapEntityToResponseDto(user)
		usersDTO = append(usersDTO, userDTO)
	}

	return usersDTO
}
