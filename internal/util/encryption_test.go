package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptPassword_ShouldEncryptWithSuccess(t *testing.T) {
	//Arrange

	assert := assert.New(t)

	var plainPassword string = "123456789"

	//Act
	result, err := EncryptPassword(plainPassword)

	//Assert
	assert.NotEqual(plainPassword, result)
	assert.Empty(err)
}

func TestEncryptPassword_ShouldReturnErrorWhenInputIsEmpty(t *testing.T) {
	//Arrange
	assert := assert.New(t)

	//Act
	result, err := EncryptPassword("")

	//Assert
	assert.Empty(result)
	assert.Error(err)
	assert.Equal("password cannot be empty", err.Error())
}
