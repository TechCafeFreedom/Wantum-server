// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/domain/service/profile/service.go

// Package mock_profile is a generated GoMock package.
package mock_profile

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	entity "wantum/pkg/domain/entity"
	repository "wantum/pkg/domain/repository"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateNewProfile mocks base method
func (m *MockService) CreateNewProfile(ctx context.Context, masterTx repository.MasterTx, userID int, name, thumbnail, bio, phone, place, birth string, gender int) (*entity.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewProfile", ctx, masterTx, userID, name, thumbnail, bio, phone, place, birth, gender)
	ret0, _ := ret[0].(*entity.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewProfile indicates an expected call of CreateNewProfile
func (mr *MockServiceMockRecorder) CreateNewProfile(ctx, masterTx, userID, name, thumbnail, bio, phone, place, birth, gender interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewProfile", reflect.TypeOf((*MockService)(nil).CreateNewProfile), ctx, masterTx, userID, name, thumbnail, bio, phone, place, birth, gender)
}
