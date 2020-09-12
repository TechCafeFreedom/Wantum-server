// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/domain/service/wish_card/service.go

// Package mock_wish_card is a generated GoMock package.
package mock_wish_card

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
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

// Create mocks base method
func (m *MockService) Create(ctx context.Context, masterTx repository.MasterTx, activity, description string, date *time.Time, userID, categoryID, placeID int) (*entity.WishCard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, masterTx, activity, description, date, userID, categoryID, placeID)
	ret0, _ := ret[0].(*entity.WishCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockServiceMockRecorder) Create(ctx, masterTx, activity, description, date, userID, categoryID, placeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockService)(nil).Create), ctx, masterTx, activity, description, date, userID, categoryID, placeID)
}

// Update mocks base method
func (m *MockService) Update(ctx context.Context, masterTx repository.MasterTx, wishCardID int, activity, description string, date, doneAt *time.Time, userID, categoryID, placeID int) (*entity.WishCard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, masterTx, wishCardID, activity, description, date, doneAt, userID, categoryID, placeID)
	ret0, _ := ret[0].(*entity.WishCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockServiceMockRecorder) Update(ctx, masterTx, wishCardID, activity, description, date, doneAt, userID, categoryID, placeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockService)(nil).Update), ctx, masterTx, wishCardID, activity, description, date, doneAt, userID, categoryID, placeID)
}

// Delete mocks base method
func (m *MockService) Delete(ctx context.Context, masterTx repository.MasterTx, wishCardID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, masterTx, wishCardID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockServiceMockRecorder) Delete(ctx, masterTx, wishCardID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockService)(nil).Delete), ctx, masterTx, wishCardID)
}

// UpDeleteFlag mocks base method
func (m *MockService) UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (*entity.WishCard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpDeleteFlag", ctx, masterTx, wishCardID)
	ret0, _ := ret[0].(*entity.WishCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpDeleteFlag indicates an expected call of UpDeleteFlag
func (mr *MockServiceMockRecorder) UpDeleteFlag(ctx, masterTx, wishCardID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpDeleteFlag", reflect.TypeOf((*MockService)(nil).UpDeleteFlag), ctx, masterTx, wishCardID)
}

// DownDeleteFlag mocks base method
func (m *MockService) DownDeleteFlag(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (*entity.WishCard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownDeleteFlag", ctx, masterTx, wishCardID)
	ret0, _ := ret[0].(*entity.WishCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownDeleteFlag indicates an expected call of DownDeleteFlag
func (mr *MockServiceMockRecorder) DownDeleteFlag(ctx, masterTx, wishCardID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownDeleteFlag", reflect.TypeOf((*MockService)(nil).DownDeleteFlag), ctx, masterTx, wishCardID)
}

// GetByID mocks base method
func (m *MockService) GetByID(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (*entity.WishCard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, masterTx, wishCardID)
	ret0, _ := ret[0].(*entity.WishCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockServiceMockRecorder) GetByID(ctx, masterTx, wishCardID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockService)(nil).GetByID), ctx, masterTx, wishCardID)
}

// GetByIDs mocks base method
func (m *MockService) GetByIDs(ctx context.Context, masterTx repository.MasterTx, wishCardIDs []int) (entity.WishCardSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDs", ctx, masterTx, wishCardIDs)
	ret0, _ := ret[0].(entity.WishCardSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDs indicates an expected call of GetByIDs
func (mr *MockServiceMockRecorder) GetByIDs(ctx, masterTx, wishCardIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDs", reflect.TypeOf((*MockService)(nil).GetByIDs), ctx, masterTx, wishCardIDs)
}

// GetByCategoryID mocks base method
func (m *MockService) GetByCategoryID(ctx context.Context, masterTx repository.MasterTx, categoryID int) (entity.WishCardSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCategoryID", ctx, masterTx, categoryID)
	ret0, _ := ret[0].(entity.WishCardSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCategoryID indicates an expected call of GetByCategoryID
func (mr *MockServiceMockRecorder) GetByCategoryID(ctx, masterTx, categoryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCategoryID", reflect.TypeOf((*MockService)(nil).GetByCategoryID), ctx, masterTx, categoryID)
}