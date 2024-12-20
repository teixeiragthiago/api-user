package usercontroller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/teixeiragthiago/api-user/internal/dto"
	userservicemocks "github.com/teixeiragthiago/api-user/internal/service/mocks"
	testutils "github.com/teixeiragthiago/api-user/internal/test_utils"
)

func TestUserController_Login_MustReturnOk(t *testing.T) {
	//Arrange
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := userservicemocks.NewMockUserService(ctrl)
	service.EXPECT().Login(gomock.Any()).Return("token", nil)

	userController := NewUserController(service)

	body, _ := json.Marshal(testutils.MockValidUserLoginDto())
	req := httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	//Act
	userController.Login(ctx)

	//Assert

	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]string
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, "token", response["token"])
}

func TestUserController_Login_MustReturnUnauthorized(t *testing.T) {
	//Arrange
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := userservicemocks.NewMockUserService(ctrl)
	service.EXPECT().Login(gomock.Any()).Return("", errors.New("invalid credentials"))

	userController := NewUserController(service)

	body, _ := json.Marshal(testutils.MockInvalidUserLoginDto())
	req := httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	//Act
	userController.Login(ctx)

	//Assert

	assert.Equal(t, http.StatusUnauthorized, rec.Code)

	var response map[string]string
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, "invalid credentials", response["error"])
}

func TestUserController_RegisterUser_MustReturnOk(t *testing.T) {
	//Arrange
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := userservicemocks.NewMockUserService(ctrl)
	service.EXPECT().RegisterUser(gomock.Any()).Return("token", nil)

	userController := NewUserController(service)

	body, _ := json.Marshal(testutils.MockValidUserDto())
	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	//Act
	userController.RegisterUser(ctx)

	//Assert
	assert.Equal(t, http.StatusCreated, rec.Code)

	var response map[string]string
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, "token", response["token"])
}

func TestUserController_RegisterUser_MustReturnBadRequest(t *testing.T) {
	//Arrange
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := userservicemocks.NewMockUserService(ctrl)
	service.EXPECT().RegisterUser(gomock.Any()).Return("", errors.New("error creating user"))

	userController := NewUserController(service)

	body, _ := json.Marshal(testutils.MockInvalidUserDto())
	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	//Act
	userController.RegisterUser(ctx)

	//Assert
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]string
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, "error creating user", response["error"])
}

func TestUserController_Update_MustReturnOk(t *testing.T) {
	//Arrange
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := userservicemocks.NewMockUserService(ctrl)
	service.EXPECT().Update(gomock.Any()).Return("User updated successfully", nil)

	userController := NewUserController(service)

	body, _ := json.Marshal(testutils.MockValidUserDto())
	req := httptest.NewRequest(http.MethodPut, "/user", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	//Act
	userController.Update(ctx)

	//Assert
	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]string
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, "User updated successfully", response["data"])
}

func TestUserController_Update_MustReturnBadRequest(t *testing.T) {
	//Arrange
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := userservicemocks.NewMockUserService(ctrl)
	service.EXPECT().Update(gomock.Any()).Return("", errors.New("Error updating user"))

	userController := NewUserController(service)

	body, _ := json.Marshal(testutils.MockValidUserDto())
	req := httptest.NewRequest(http.MethodPut, "/user", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	//Act
	userController.Update(ctx)

	//Assert
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]string
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, "Error updating user", response["error"])
}

func TestUserController_Delete_MustReturnOk(t *testing.T) {
	//Arrange
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := userservicemocks.NewMockUserService(ctrl)
	service.EXPECT().Delete(gomock.Any()).Return("User deleted successfully", nil)

	userController := NewUserController(service)

	req := httptest.NewRequest(http.MethodDelete, "/user/1", nil)
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}

	//Act
	userController.Delete(ctx)

	//Assert
	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]string
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, "User deleted successfully", response["data"])
}

func TestUserController_Delete_MustReturnBadRequestWhenIdIsInvalid(t *testing.T) {
	//Arrange
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := userservicemocks.NewMockUserService(ctrl)

	req := httptest.NewRequest(http.MethodDelete, "/user/invalid", nil)
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req
	ctx.Params = []gin.Param{{Key: "id", Value: "invalid id"}}

	userController := NewUserController(service)

	//Act
	userController.Delete(ctx)

	//Assert
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]string
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, "Error converting parameter `id`", response["error"])
}

func TestUserController_Delete_MustReturnBadRequest(t *testing.T) {
	//Arrange
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := userservicemocks.NewMockUserService(ctrl)
	service.EXPECT().Delete(gomock.Any()).Return("", errors.New("Error deleting user"))

	userController := NewUserController(service)

	req := httptest.NewRequest(http.MethodDelete, "/user/1", nil)
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}

	//Act
	userController.Delete(ctx)

	//Assert
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]string
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, "Error deleting user", response["error"])
}

func TestUserController_GetById_MustReturnOk(t *testing.T) {
	//Arrange
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedUserResponse := testutils.MockUserResponseDto()

	service := userservicemocks.NewMockUserService(ctrl)
	service.EXPECT().GetById(gomock.Any()).Return(expectedUserResponse, nil)

	userController := NewUserController(service)

	req := httptest.NewRequest(http.MethodGet, "/user/1", nil)
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}

	//Act
	userController.GetById(ctx)

	//Assert
	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]*dto.UserResponseDto
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, expectedUserResponse, response["data"])
}

func TestUserController_GetById_MustReturnBadRequestWhenIdIsInvalid(t *testing.T) {
	//Arrange
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := userservicemocks.NewMockUserService(ctrl)

	userController := NewUserController(service)

	req := httptest.NewRequest(http.MethodGet, "/user/invalid", nil)
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req
	ctx.Params = []gin.Param{{Key: "id", Value: "invalid id"}}

	//Act
	userController.GetById(ctx)

	//Assert
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]string
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, "Invalid user ID", response["error"])
}

func TestUserController_GetById_MustReturnInternalServerErrorWhenFailed(t *testing.T) {
	//Arrange
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := userservicemocks.NewMockUserService(ctrl)
	service.EXPECT().GetById(gomock.Any()).Return(nil, errors.New("Error fetching user"))

	userController := NewUserController(service)

	req := httptest.NewRequest(http.MethodGet, "/user/1", nil)
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}

	//Act
	userController.GetById(ctx)

	//Assert
	assert.Equal(t, http.StatusInternalServerError, rec.Code)

	var response map[string]string
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, "Error fetching user", response["error"])
}

func TestUserController_Get_MustReturnOk(t *testing.T) {
	//Arrange
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := userservicemocks.NewMockUserService(ctrl)

	mockUsers := []*dto.UserResponseDto{
		{ID: 1, Name: "John Doe", Email: "john.doe@example.com"},
		{ID: 2, Name: "Johnny Smith", Email: "johnny.smith@example.com"},
	}

	service.EXPECT().Get(gomock.Any()).Return(mockUsers, nil)

	userController := NewUserController(service)

	req := httptest.NewRequest(http.MethodGet, "/user?search=teste", nil)
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	//Act
	userController.Get(ctx)

	//Assert
	var response map[string]any
	_ = json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotNil(t, response["data"])
	assert.Len(t, mockUsers, 2)
}

func TestUserController_Get_MustReturnBadRequestWhenIsFailed(t *testing.T) {
	//Arrange
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := userservicemocks.NewMockUserService(ctrl)

	service.EXPECT().Get(gomock.Any()).Return(nil, errors.New("error fetching users"))

	userController := NewUserController(service)

	req := httptest.NewRequest(http.MethodGet, "/user?search=teste", nil)
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	//Act
	userController.Get(ctx)

	//Assert
	var response map[string]any
	_ = json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "error fetching users", response["error"])
}
