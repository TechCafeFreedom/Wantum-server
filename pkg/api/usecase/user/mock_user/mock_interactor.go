// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/api/usecase/user/interactor.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	context "context"
	reflect "reflect"
	entity "wantum/pkg/domain/entity"

	gomock "github.com/golang/mock/gomock"
)

// MockInteractor is a mock of Interactor interface
type MockInteractor struct {
	ctrl     *gomock.Controller
	recorder *MockInteractorMockRecorder
}

// MockInteractorMockRecorder is the mock recorder for MockInteractor
type MockInteractorMockRecorder struct {
	mock *MockInteractor
}

// NewMockInteractor creates a new mock instance
func NewMockInteractor(ctrl *gomock.Controller) *MockInteractor {
	mock := &MockInteractor{ctrl: ctrl}
	mock.recorder = &MockInteractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInteractor) EXPECT() *MockInteractorMockRecorder {
	return m.recorder
}

// CreateNewUser mocks base method
func (m *MockInteractor) CreateNewUser(ctx context.Context, authID, userName, mail, name, thumbnail, bio, phone, place, birth string, gender int) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewUser", ctx, authID, userName, mail, name, thumbnail, bio, phone, place, birth, gender)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewUser indicates an expected call of CreateNewUser
func (mr *MockInteractorMockRecorder) CreateNewUser(ctx, authID, userName, mail, name, thumbnail, bio, phone, place, birth, gender interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewUser", reflect.TypeOf((*MockInteractor)(nil).CreateNewUser), ctx, authID, userName, mail, name, thumbnail, bio, phone, place, birth, gender)
}

// GetAuthorizedUser mocks base method
func (m *MockInteractor) GetAuthorizedUser(ctx context.Context, authID string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorizedUser", ctx, authID)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorizedUser indicates an expected call of GetAuthorizedUser
func (mr *MockInteractorMockRecorder) GetAuthorizedUser(ctx, authID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorizedUser", reflect.TypeOf((*MockInteractor)(nil).GetAuthorizedUser), ctx, authID)
}

// GetAll mocks base method
func (m *MockInteractor) GetAll(ctx context.Context) (entity.UserSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].(entity.UserSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockInteractorMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockInteractor)(nil).GetAll), ctx)
}
