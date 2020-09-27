// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/domain/repository/tag/repository.go

// Package mock_tag is a generated GoMock package.
package mock_tag

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
	tag "wantum/pkg/domain/entity/tag"
	repository "wantum/pkg/domain/repository"
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

// Insert mocks base method
func (m *MockRepository) Insert(ctx context.Context, masterTx repository.MasterTx, tag *tag.Entity) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, masterTx, tag)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockRepositoryMockRecorder) Insert(ctx, masterTx, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRepository)(nil).Insert), ctx, masterTx, tag)
}

// UpDeleteFlag mocks base method
func (m *MockRepository) UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, tagID int, updatedAt, deletedAt *time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpDeleteFlag", ctx, masterTx, tagID, updatedAt, deletedAt)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpDeleteFlag indicates an expected call of UpDeleteFlag
func (mr *MockRepositoryMockRecorder) UpDeleteFlag(ctx, masterTx, tagID, updatedAt, deletedAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpDeleteFlag", reflect.TypeOf((*MockRepository)(nil).UpDeleteFlag), ctx, masterTx, tagID, updatedAt, deletedAt)
}

// DownDeleteFlag mocks base method
func (m *MockRepository) DownDeleteFlag(ctx context.Context, masterTx repository.MasterTx, tagID int, updatedAt *time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownDeleteFlag", ctx, masterTx, tagID, updatedAt)
	ret0, _ := ret[0].(error)
	return ret0
}

// DownDeleteFlag indicates an expected call of DownDeleteFlag
func (mr *MockRepositoryMockRecorder) DownDeleteFlag(ctx, masterTx, tagID, updatedAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownDeleteFlag", reflect.TypeOf((*MockRepository)(nil).DownDeleteFlag), ctx, masterTx, tagID, updatedAt)
}

// Delete mocks base method
func (m *MockRepository) Delete(ctx context.Context, masterTx repository.MasterTx, tag *tag.Entity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, masterTx, tag)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockRepositoryMockRecorder) Delete(ctx, masterTx, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), ctx, masterTx, tag)
}

// SelectByID mocks base method
func (m *MockRepository) SelectByID(ctx context.Context, masterTx repository.MasterTx, tagID int) (*tag.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByID", ctx, masterTx, tagID)
	ret0, _ := ret[0].(*tag.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByID indicates an expected call of SelectByID
func (mr *MockRepositoryMockRecorder) SelectByID(ctx, masterTx, tagID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByID", reflect.TypeOf((*MockRepository)(nil).SelectByID), ctx, masterTx, tagID)
}

// SelectByIDs mocks base method
func (m *MockRepository) SelectByIDs(ctx context.Context, masterTx repository.MasterTx, tagIDs []int) (tag.EntitySlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByIDs", ctx, masterTx, tagIDs)
	ret0, _ := ret[0].(tag.EntitySlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByIDs indicates an expected call of SelectByIDs
func (mr *MockRepositoryMockRecorder) SelectByIDs(ctx, masterTx, tagIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByIDs", reflect.TypeOf((*MockRepository)(nil).SelectByIDs), ctx, masterTx, tagIDs)
}

// SelectByName mocks base method
func (m *MockRepository) SelectByName(ctx context.Context, masterTx repository.MasterTx, name string) (*tag.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByName", ctx, masterTx, name)
	ret0, _ := ret[0].(*tag.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByName indicates an expected call of SelectByName
func (mr *MockRepositoryMockRecorder) SelectByName(ctx, masterTx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByName", reflect.TypeOf((*MockRepository)(nil).SelectByName), ctx, masterTx, name)
}

// SelectByWishCardID mocks base method
func (m *MockRepository) SelectByWishCardID(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (tag.EntitySlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByWishCardID", ctx, masterTx, wishCardID)
	ret0, _ := ret[0].(tag.EntitySlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByWishCardID indicates an expected call of SelectByWishCardID
func (mr *MockRepositoryMockRecorder) SelectByWishCardID(ctx, masterTx, wishCardID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByWishCardID", reflect.TypeOf((*MockRepository)(nil).SelectByWishCardID), ctx, masterTx, wishCardID)
}

// SelectByMemoryID mocks base method
func (m *MockRepository) SelectByMemoryID(ctx context.Context, masterTx repository.MasterTx, memoryID int) (tag.EntitySlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByMemoryID", ctx, masterTx, memoryID)
	ret0, _ := ret[0].(tag.EntitySlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByMemoryID indicates an expected call of SelectByMemoryID
func (mr *MockRepositoryMockRecorder) SelectByMemoryID(ctx, masterTx, memoryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByMemoryID", reflect.TypeOf((*MockRepository)(nil).SelectByMemoryID), ctx, masterTx, memoryID)
}
