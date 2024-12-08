package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/teixeiragthiago/api-user/internal/entity"
	"github.com/teixeiragthiago/api-user/internal/repository/mocks"
	testutils "github.com/teixeiragthiago/api-user/internal/test_utils"
	utilmocks "github.com/teixeiragthiago/api-user/internal/util/mocks"
)

func TestUserService_Delete_MustReturnSuccess(t *testing.T) {
	//Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockJwtGenerator := utilmocks.NewMockJwtGeneratorService(ctrl)
	service := NewUserService(mockRepo, mockJwtGenerator)

	mockRepo.EXPECT().GetById(gomock.Eq(uint(1))).Return(&entity.User{}, nil)
	mockRepo.EXPECT().Delete(gomock.Any()).Return(nil)

	//Act
	result, err := service.Delete(1)

	//Assert
	assert.NoError(t, err)
	assert.Equal(t, "User removed successfully", result)
}

func TestUserService_Delete_ErrorWhenUserNotFound(t *testing.T) {
	//Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockJwtGenerator := utilmocks.NewMockJwtGeneratorService(ctrl)
	service := NewUserService(mockRepo, mockJwtGenerator)

	mockRepo.EXPECT().GetById(gomock.Eq(uint(1)))

	//Act
	result, err := service.Delete(1)

	//Assert
	assert.Error(t, err)
	assert.Empty(t, result)
	assert.Equal(t, "user could not be found to delete", err.Error())
}

func TestUserService_Delete_MustReturnErrorWhenFailedToDelete(t *testing.T) {
	//Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockJwtGenerator := utilmocks.NewMockJwtGeneratorService(ctrl)
	service := NewUserService(mockRepo, mockJwtGenerator)

	mockRepo.EXPECT().GetById(gomock.Eq(uint(1))).Return(&entity.User{}, nil)
	mockRepo.EXPECT().Delete(gomock.Any()).Return(errors.New("error deleting user"))

	//Act
	result, err := service.Delete(1)

	//Assert
	assert.Error(t, err)
	assert.Empty(t, result)
	assert.Equal(t, "error deleting user", err.Error())
}

func TestUserService_Get_MustReturnSuccess(t *testing.T) {
	//Arrange

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockJwtGenerator := utilmocks.NewMockJwtGeneratorService(ctrl)
	service := NewUserService(mockRepo, mockJwtGenerator)

	mockRepo.EXPECT().Get(gomock.Any()).Return([]*entity.User{
		testutils.MockValidUser(),
	}, nil)

	//Act
	result, err := service.Get("thiago")

	//Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestUserService_GetById_MustReturnWithSuccess(t *testing.T) {
	//Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockJwtGenerator := utilmocks.NewMockJwtGeneratorService(ctrl)
	service := NewUserService(mockRepo, mockJwtGenerator)

	mockRepo.EXPECT().GetById(gomock.Any()).Return(&entity.User{}, nil)

	//Act
	result, err := service.GetById(1)

	//Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestUserService_GetById_MustReturnError(t *testing.T) {
	//Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockJwtGenerator := utilmocks.NewMockJwtGeneratorService(ctrl)
	service := NewUserService(mockRepo, mockJwtGenerator)

	mockRepo.EXPECT().GetById(gomock.Any()).Return(nil, errors.New("error getting user by id"))

	//Act
	result, err := service.GetById(1)

	//Assert
	assert.Error(t, err)
	assert.Equal(t, "error getting user by id", err.Error())
	assert.Nil(t, result)
}

func TestUserService_Update_MustReturnSuccess(t *testing.T) {
	//Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockJwtGenerator := utilmocks.NewMockJwtGeneratorService(ctrl)
	service := NewUserService(mockRepo, mockJwtGenerator)

	mockRepo.EXPECT().GetById(gomock.Any()).Return(&entity.User{}, nil)
	mockRepo.EXPECT().InUseByAnotherUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, nil)
	mockRepo.EXPECT().InUseByAnotherUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, nil)
	mockRepo.EXPECT().Update(gomock.Any()).Return(nil)

	//Act
	result, err := service.Update(testutils.MockValidUserDto())

	//Assert
	assert.NoError(t, err)
	assert.Equal(t, "User updated successfully", result)
}

func TestUserService_Update_MustReturnErrorWhenEmailIsInUseByAnotherUser(t *testing.T) {
	//Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockJwtGenerator := utilmocks.NewMockJwtGeneratorService(ctrl)
	service := NewUserService(mockRepo, mockJwtGenerator)

	mockRepo.EXPECT().GetById(gomock.Any()).Return(&entity.User{}, nil)
	mockRepo.EXPECT().InUseByAnotherUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, errors.New("e-mail already in use by another user"))

	//Act
	result, err := service.Update(testutils.MockValidUserDto())

	//Assert
	assert.Empty(t, result)
	assert.Error(t, err)
	assert.Equal(t, "e-mail already in use by another user", err.Error())
}

func TestUserService_Update_MustReturnErrorWhenNickIsInUseByAnotherUser(t *testing.T) {
	//Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockJwtGenerator := utilmocks.NewMockJwtGeneratorService(ctrl)
	service := NewUserService(mockRepo, mockJwtGenerator)

	mockRepo.EXPECT().GetById(gomock.Any()).Return(&entity.User{}, nil)
	mockRepo.EXPECT().InUseByAnotherUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, nil)
	mockRepo.EXPECT().InUseByAnotherUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, errors.New("nickname already in use by another user"))

	//Act
	result, err := service.Update(testutils.MockValidUserDto())

	//Assert
	assert.Empty(t, result)
	assert.Error(t, err)
	assert.Equal(t, "nickname already in use by another user", err.Error())
}

func TestUserService_Update_MustReturnErrorWhenUpdateFailed(t *testing.T) {
	//Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockJwtGenerator := utilmocks.NewMockJwtGeneratorService(ctrl)
	service := NewUserService(mockRepo, mockJwtGenerator)

	mockRepo.EXPECT().GetById(gomock.Any()).Return(&entity.User{}, nil)
	mockRepo.EXPECT().InUseByAnotherUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, nil)
	mockRepo.EXPECT().InUseByAnotherUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, nil)
	mockRepo.EXPECT().Update(gomock.Any()).Return(errors.New("error updating user"))
	//Act
	result, err := service.Update(testutils.MockValidUserDto())

	//Assert
	assert.Empty(t, result)
	assert.Error(t, err)
	assert.Equal(t, "error updating user", err.Error())
}

func TestUserService_RegisterUser_MustReturnSuccessWhenUserIsValid(t *testing.T) {
	//Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockJwtGenerator := utilmocks.NewMockJwtGeneratorService(ctrl)
	service := NewUserService(mockRepo, mockJwtGenerator)

	mockRepo.EXPECT().Exists(gomock.Any(), gomock.Any()).Return(false, nil)
	mockRepo.EXPECT().Exists(gomock.Any(), gomock.Any()).Return(false, nil)
	mockJwtGenerator.EXPECT().GenerateToken(gomock.Any()).Return("token", nil)
	mockRepo.EXPECT().Save(gomock.Any()).Return(nil)

	//Act
	result, err := service.RegisterUser(testutils.MockValidUserDto())

	//Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}

func TestUserService_RegisterUser_MustReturnErrorWhenUserIsInvalid(t *testing.T) {
	//Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockJwtGenerator := utilmocks.NewMockJwtGeneratorService(ctrl)
	service := NewUserService(mockRepo, mockJwtGenerator)

	//Act
	_, err := service.RegisterUser(testutils.MockInvalidUserDto())

	//Assert
	assert.Error(t, err)
}

func TestUserService_RegisterUser_MustReturnErrorWhenEmailExists(t *testing.T) {
	//Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockJwtGenerator := utilmocks.NewMockJwtGeneratorService(ctrl)
	service := NewUserService(mockRepo, mockJwtGenerator)

	mockRepo.EXPECT().Exists(gomock.Any(), gomock.Any()).Return(true, errors.New("e-mail already exists"))

	//Act
	result, err := service.RegisterUser(testutils.MockValidUserDto())

	//Assert
	assert.Error(t, err)
	assert.Equal(t, "e-mail already exists", err.Error())
	assert.Empty(t, result)
}

func TestUserService_RegisterUser_MustReturnErrorWhenNicknameExists(t *testing.T) {
	//Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockJwtGenerator := utilmocks.NewMockJwtGeneratorService(ctrl)
	service := NewUserService(mockRepo, mockJwtGenerator)

	mockRepo.EXPECT().Exists(gomock.Any(), gomock.Any()).Return(false, nil)
	mockRepo.EXPECT().Exists(gomock.Any(), gomock.Any()).Return(true, errors.New("nickname already exists"))

	//Act
	result, err := service.RegisterUser(testutils.MockValidUserDto())

	//Assert
	assert.Error(t, err)
	assert.Equal(t, "nickname already exists", err.Error())
	assert.Empty(t, result)
}
