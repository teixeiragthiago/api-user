package service

import (
	"errors"

	"github.com/teixeiragthiago/api-user/internal/dto"
	"github.com/teixeiragthiago/api-user/internal/mapper"
	"github.com/teixeiragthiago/api-user/internal/repository"
	"github.com/teixeiragthiago/api-user/internal/util"
)

type UserService interface {
	RegisterUser(UserDTO *dto.UserDTO) (string, error)
	GetById(id uint) (*dto.UserResponseDto, error)
	Get(search string) ([]*dto.UserResponseDto, error)
	Delete(id uint) (string, error)
	Update(UserDTO *dto.UserDTO) (string, error)
	Login(userDTO *dto.UserLoginDto) (string, error)
}

type userService struct {
	userRepository repository.UserRepository
	jwtService     util.JwtGeneratorService
}

func NewUserService(userRepository repository.UserRepository, jwtService util.JwtGeneratorService) UserService {
	return &userService{userRepository, jwtService}
}

func (s *userService) Login(userDTO *dto.UserLoginDto) (string, error) {

	if err := userDTO.Validate(); err != nil {
		return "", err
	}

	userData, err := s.userRepository.GetByEmail(userDTO.Email)
	if err != nil {
		return "", err
	}

	if userData == nil {
		return "", errors.New("user could not be found")
	}

	_, err = util.MatchPassword(string(userData.Password), []byte(userDTO.Password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	token, err := s.jwtService.GenerateToken(&util.Claims{
		ID:       userData.ID,
		Nickname: userData.Nickname,
		Email:    userData.Email,
	})

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) RegisterUser(userDTO *dto.UserDTO) (string, error) {

	if err := userDTO.Validate(); err != nil {
		return "", err
	}

	if emailExists, err := s.userRepository.Exists("email", userDTO.Email); err != nil {
		return "", err
	} else if emailExists {
		return "", errors.New("e-mail already exists")
	}

	if nicknameExists, err := s.userRepository.Exists("nickname", userDTO.Nickname); err != nil {
		return "", err
	} else if nicknameExists {
		return "", errors.New("nickname already exists")
	}

	user := mapper.MapDtoToEntity(userDTO)

	err := s.userRepository.Save(user)
	if err != nil {
		return "", err
	}

	token, err := s.jwtService.GenerateToken(&util.Claims{
		ID:       user.ID,
		Nickname: user.Nickname,
		Email:    user.Email,
	})

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) Update(userDTO *dto.UserDTO) (string, error) {

	if err := userDTO.ValidateUpdate(); err != nil {
		return "", err
	}

	userEntity := mapper.MapDtoToEntity(userDTO)
	user, _ := s.userRepository.GetById(userEntity.ID)

	if user == nil {
		return "", errors.New("user could not be found to update")
	}

	emailInUseAnotherUser, _ := s.userRepository.InUseByAnotherUser(userDTO.ID, "email", userDTO.Email)
	if emailInUseAnotherUser {
		return "", errors.New("e-mail already in use by another user")
	}

	nicknameInUseAnotherUser, _ := s.userRepository.InUseByAnotherUser(userDTO.ID, "nickname", userDTO.Nickname)
	if nicknameInUseAnotherUser {
		return "", errors.New("nickname already in use by another user")
	}

	err := s.userRepository.Update(userEntity)
	if err != nil {
		return "", err
	}

	return "User updated successfully", nil
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

	user, _ := s.userRepository.GetById(id)

	if user == nil {
		return "", errors.New("user could not be found to delete")
	}

	err := s.userRepository.Delete(user)
	if err != nil {
		return "", err
	}

	return "User removed successfully", nil
}
