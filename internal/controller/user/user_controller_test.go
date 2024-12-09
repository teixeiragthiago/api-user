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

	body, _ := json.Marshal(testutils.MockValidUserLoginDto())
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

	body, _ := json.Marshal(testutils.MockValidUserLoginDto())
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
