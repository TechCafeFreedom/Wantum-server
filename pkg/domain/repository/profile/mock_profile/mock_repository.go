// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/domain/repository/profile/repository.go

// Package mock_profile is a generated GoMock package.
package mock_profile

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	repository "wantum/pkg/domain/repository"
	model "wantum/pkg/infrastructure/mysql/model"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// InsertProfile mocks base method
func (m *MockRepository) InsertProfile(ctx context.Context, masterTx repository.MasterTx, profileEntity *model.ProfileModel) (*model.ProfileModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertProfile", ctx, masterTx, profileEntity)
	ret0, _ := ret[0].(*model.ProfileModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertProfile indicates an expected call of InsertProfile
func (mr *MockRepositoryMockRecorder) InsertProfile(ctx, masterTx, profileEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertProfile", reflect.TypeOf((*MockRepository)(nil).InsertProfile), ctx, masterTx, profileEntity)
}

// SelectByUserID mocks base method
func (m *MockRepository) SelectByUserID(ctx context.Context, masterTx repository.MasterTx, userID int) (*model.ProfileModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByUserID", ctx, masterTx, userID)
	ret0, _ := ret[0].(*model.ProfileModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByUserID indicates an expected call of SelectByUserID
func (mr *MockRepositoryMockRecorder) SelectByUserID(ctx, masterTx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByUserID", reflect.TypeOf((*MockRepository)(nil).SelectByUserID), ctx, masterTx, userID)
}

// SelectByUserIDs mocks base method
func (m *MockRepository) SelectByUserIDs(ctx context.Context, masterTx repository.MasterTx, userIDs []int) (model.ProfileModelSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByUserIDs", ctx, masterTx, userIDs)
	ret0, _ := ret[0].(model.ProfileModelSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByUserIDs indicates an expected call of SelectByUserIDs
func (mr *MockRepositoryMockRecorder) SelectByUserIDs(ctx, masterTx, userIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByUserIDs", reflect.TypeOf((*MockRepository)(nil).SelectByUserIDs), ctx, masterTx, userIDs)
}
