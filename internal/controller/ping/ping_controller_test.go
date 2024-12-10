package pingcontroller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestPingController_Ping_MustReturnOk(t *testing.T) {
	//Arrange
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	pingController := NewPingController()

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	//Act
	pingController.Ping(ctx)

	//Assert
	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]string
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, "pong", response["message"])
}
