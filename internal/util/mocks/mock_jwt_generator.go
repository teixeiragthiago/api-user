// Code generated by MockGen. DO NOT EDIT.
// Source: jwt_generator.go

// Package utilmocks is a generated GoMock package.
package utilmocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	util "github.com/teixeiragthiago/api-user/internal/util"
)

// MockJwtGeneratorService is a mock of JwtGeneratorService interface.
type MockJwtGeneratorService struct {
	ctrl     *gomock.Controller
	recorder *MockJwtGeneratorServiceMockRecorder
}

// MockJwtGeneratorServiceMockRecorder is the mock recorder for MockJwtGeneratorService.
type MockJwtGeneratorServiceMockRecorder struct {
	mock *MockJwtGeneratorService
}

// NewMockJwtGeneratorService creates a new mock instance.
func NewMockJwtGeneratorService(ctrl *gomock.Controller) *MockJwtGeneratorService {
	mock := &MockJwtGeneratorService{ctrl: ctrl}
	mock.recorder = &MockJwtGeneratorServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJwtGeneratorService) EXPECT() *MockJwtGeneratorServiceMockRecorder {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockJwtGeneratorService) GenerateToken(claimsData *util.Claims) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", claimsData)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockJwtGeneratorServiceMockRecorder) GenerateToken(claimsData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockJwtGeneratorService)(nil).GenerateToken), claimsData)
}

// ValidateToken mocks base method.
func (m *MockJwtGeneratorService) ValidateToken(tokenString string) (*util.Claims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateToken", tokenString)
	ret0, _ := ret[0].(*util.Claims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateToken indicates an expected call of ValidateToken.
func (mr *MockJwtGeneratorServiceMockRecorder) ValidateToken(tokenString interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToken", reflect.TypeOf((*MockJwtGeneratorService)(nil).ValidateToken), tokenString)
}
