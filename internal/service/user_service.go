package service

import (
	"github.com/teixeiragthiago/api-user/internal/dto"
	"github.com/teixeiragthiago/api-user/internal/mapper"
	"github.com/teixeiragthiago/api-user/internal/repository"
)

type UserService interface {
	RegisterUser(UserDTO *dto.UserDTO) error
	GetById(id uint) (*dto.UserResponseDto, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) RegisterUser(userDTO *dto.UserDTO) error {

	if err := userDTO.Validate(); err != nil {
		return err
	}

	user := mapper.MapDtoToEntity(userDTO)

	err := s.userRepository.Save(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) GetById(id uint) (*dto.UserResponseDto, error) {
	user, err := s.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return mapper.MapEntityToResponseDto(user), nil
}
