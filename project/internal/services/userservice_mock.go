// Code generated by MockGen. DO NOT EDIT.
// Source: userService.go
//
// Generated by this command:
//
//	mockgen -source=userService.go -destination=userservice_mock.go -package=services
//
// Package services is a generated GoMock package.
package services

import (
	model "project/internal/model"
	reflect "reflect"

	jwt "github.com/golang-jwt/jwt/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockUsersService is a mock of UsersService interface.
type MockUsersService struct {
	ctrl     *gomock.Controller
	recorder *MockUsersServiceMockRecorder
}

// MockUsersServiceMockRecorder is the mock recorder for MockUsersService.
type MockUsersServiceMockRecorder struct {
	mock *MockUsersService
}

// NewMockUsersService creates a new mock instance.
func NewMockUsersService(ctrl *gomock.Controller) *MockUsersService {
	mock := &MockUsersService{ctrl: ctrl}
	mock.recorder = &MockUsersServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersService) EXPECT() *MockUsersServiceMockRecorder {
	return m.recorder
}

// UserForgetPassword mocks base method.
func (m *MockUsersService) UserForgetPassword(uf model.UserForgetPassword) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserForgetPassword", uf)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserForgetPassword indicates an expected call of UserForgetPassword.
func (mr *MockUsersServiceMockRecorder) UserForgetPassword(uf any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserForgetPassword", reflect.TypeOf((*MockUsersService)(nil).UserForgetPassword), uf)
}

// UserSignup mocks base method.
func (m *MockUsersService) UserSignup(nu model.UserSignup) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserSignup", nu)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserSignup indicates an expected call of UserSignup.
func (mr *MockUsersServiceMockRecorder) UserSignup(nu any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserSignup", reflect.TypeOf((*MockUsersService)(nil).UserSignup), nu)
}

// UserUpdatePassword mocks base method.
func (m *MockUsersService) UserUpdatePassword(up model.UserUpdatePassword) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserUpdatePassword", up)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserUpdatePassword indicates an expected call of UserUpdatePassword.
func (mr *MockUsersServiceMockRecorder) UserUpdatePassword(up any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserUpdatePassword", reflect.TypeOf((*MockUsersService)(nil).UserUpdatePassword), up)
}

// Userlogin mocks base method.
func (m *MockUsersService) Userlogin(l model.UserLogin) (jwt.RegisteredClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Userlogin", l)
	ret0, _ := ret[0].(jwt.RegisteredClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Userlogin indicates an expected call of Userlogin.
func (mr *MockUsersServiceMockRecorder) Userlogin(l any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Userlogin", reflect.TypeOf((*MockUsersService)(nil).Userlogin), l)
}
