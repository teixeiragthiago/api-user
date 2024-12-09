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

func TestUserController_Login_MustReturnSuccess(t *testing.T) {
	//Arrange
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := userservicemocks.NewMockUserService(ctrl)
	service.EXPECT().Login(gomock.Any()).Return("token", nil)

	userController := NewUserController(service)

	body, _ := json.Marshal(testutils.MockValidUserLoginDto())
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
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
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
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
