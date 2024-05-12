package service

import (
	"errors"

	"github.com/teixeiragthiago/api-user/internal/dto"
	"github.com/teixeiragthiago/api-user/internal/mapper"
	"github.com/teixeiragthiago/api-user/internal/repository"
)

type UserService interface {
	RegisterUser(UserDTO *dto.UserDTO) error
	GetById(id uint) (*dto.UserResponseDto, error)
	Get(search string) ([]*dto.UserResponseDto, error)
	Delete(id uint) (string, error)
	Update(UserDTO *dto.UserDTO) error
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
func (s *userService) Get(search string) ([]*dto.UserResponseDto, error) {

	users, err := s.userRepository.Get(search)
	if err != nil {
		return nil, err
	}

	return mapper.MapEntitiesToResponseDto(users), nil
}

func (s *userService) Delete(id uint) (string, error) {
	panic("uninplemented")
}

func (s *userService) Update(userDTO *dto.UserDTO) error {

	userEntity := mapper.MapDtoToEntity(userDTO)
	user, err := s.userRepository.GetById(userEntity.ID)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user could not be found!")
	}

	err = s.userRepository.Update(userEntity)
	if err != nil {
		return err
	}

	return nil
}
