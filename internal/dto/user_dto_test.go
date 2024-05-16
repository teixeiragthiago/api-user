package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate_MustNotReturnError_WhenDtoIsValid(t *testing.T) {

	//Arrange
	userDTO := GenerateValidUserDTO()

	assert := assert.New(t)

	//Act
	err := userDTO.Validate()

	//Assert
	assert.Empty(err)
}

func TestValidate_MustReturnError_WhenNameIsEmpty(t *testing.T) {
	//Arrange
	userDTO := GenerateInvalidUserDTO("", "teste", "teste")

	assert := assert.New(t)

	//Act
	err := userDTO.Validate()

	//Assert
	assert.Error(err)
	assert.NotEmpty(err.Error())
	assert.Equal("name cannot be empty", err.Error())
}

func TestValidate_MustReturnError_WhenNickIsEmpty(t *testing.T) {
	//Arrange
	userDTO := GenerateInvalidUserDTO("teste", "", "teste")

	assert := assert.New(t)

	//Act
	err := userDTO.Validate()

	//Assert
	assert.Error(err)
	assert.NotEmpty(err.Error())
	assert.Equal("nick cannot be empty", err.Error())
}

func TestValidate_MustReturnError_WhenPasswordIsEmpty(t *testing.T) {
	//Arrange
	userDTO := GenerateInvalidUserDTO("teste", "teste", "")

	assert := assert.New(t)

	//Act
	err := userDTO.Validate()

	//Assert
	assert.Error(err)
	assert.NotEmpty(err.Error())
	assert.Equal("password cannot be empty", err.Error())
}

func TestValidate_MustReturnError_WhenPasswordHasMoreThenTwelveCharacters(t *testing.T) {
	//Arrange
	userDTO := GenerateInvalidUserDTO("teste", "teste", "123456789101112")

	assert := assert.New(t)

	//Act
	err := userDTO.Validate()

	//Assert
	assert.Error(err)
	assert.NotEmpty(err.Error())
	assert.Equal("password cannot have more than 12 characters", err.Error())
	assert.True(len(userDTO.Password) > 12)
}

func GenerateValidUserDTO() *UserDTO {
	return &UserDTO{
		ID:       1,
		Name:     "Thiago",
		Nick:     "thiago_test",
		Password: "123456789",
	}
}

func GenerateInvalidUserDTO(name string, nick string, password string) *UserDTO {
	return &UserDTO{
		Name:     name,
		Nick:     nick,
		Password: password,
	}
}
