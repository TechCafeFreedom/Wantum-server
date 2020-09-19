// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/domain/repository/wishcardtag/repository.go

// Package mock_wish_card_tag is a generated GoMock package.
package mock_wish_card_tag

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
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
func (m *MockRepository) Insert(ctx context.Context, masterTx repository.MasterTx, wishCardID, tagID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, masterTx, wishCardID, tagID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockRepositoryMockRecorder) Insert(ctx, masterTx, wishCardID, tagID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRepository)(nil).Insert), ctx, masterTx, wishCardID, tagID)
}

// BulkInsert mocks base method
func (m *MockRepository) BulkInsert(ctx context.Context, masterTx repository.MasterTx, wishCardID int, tagIDs []int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkInsert", ctx, masterTx, wishCardID, tagIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkInsert indicates an expected call of BulkInsert
func (mr *MockRepositoryMockRecorder) BulkInsert(ctx, masterTx, wishCardID, tagIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkInsert", reflect.TypeOf((*MockRepository)(nil).BulkInsert), ctx, masterTx, wishCardID, tagIDs)
}

// Delete mocks base method
func (m *MockRepository) Delete(ctx context.Context, masterTx repository.MasterTx, wishCardID, tagID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, masterTx, wishCardID, tagID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockRepositoryMockRecorder) Delete(ctx, masterTx, wishCardID, tagID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), ctx, masterTx, wishCardID, tagID)
}

// DeleteByWishCardID mocks base method
func (m *MockRepository) DeleteByWishCardID(ctx context.Context, masterTx repository.MasterTx, wishCardID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByWishCardID", ctx, masterTx, wishCardID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByWishCardID indicates an expected call of DeleteByWishCardID
func (mr *MockRepositoryMockRecorder) DeleteByWishCardID(ctx, masterTx, wishCardID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByWishCardID", reflect.TypeOf((*MockRepository)(nil).DeleteByWishCardID), ctx, masterTx, wishCardID)
}

// DeleteByIDs mocks base method
func (m *MockRepository) DeleteByIDs(ctx context.Context, masterTx repository.MasterTx, wishCardID int, tagIDs []int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByIDs", ctx, masterTx, wishCardID, tagIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByIDs indicates an expected call of DeleteByIDs
func (mr *MockRepositoryMockRecorder) DeleteByIDs(ctx, masterTx, wishCardID, tagIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByIDs", reflect.TypeOf((*MockRepository)(nil).DeleteByIDs), ctx, masterTx, wishCardID, tagIDs)
}
