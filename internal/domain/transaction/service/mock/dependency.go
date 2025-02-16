// Code generated by MockGen. DO NOT EDIT.
// Source: dependency.go
//
// Generated by this command:
//
//	mockgen -source dependency.go -destination mock/dependency.go
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	transaction "github.com/YrWaifu/test_go_back/internal/domain/transaction"
	user "github.com/YrWaifu/test_go_back/internal/domain/user"
	storage "github.com/YrWaifu/test_go_back/internal/domain/user/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockTransactionStorage is a mock of TransactionStorage interface.
type MockTransactionStorage struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionStorageMockRecorder
	isgomock struct{}
}

// MockTransactionStorageMockRecorder is the mock recorder for MockTransactionStorage.
type MockTransactionStorageMockRecorder struct {
	mock *MockTransactionStorage
}

// NewMockTransactionStorage creates a new mock instance.
func NewMockTransactionStorage(ctrl *gomock.Controller) *MockTransactionStorage {
	mock := &MockTransactionStorage{ctrl: ctrl}
	mock.recorder = &MockTransactionStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionStorage) EXPECT() *MockTransactionStorageMockRecorder {
	return m.recorder
}

// BeginTransaction mocks base method.
func (m *MockTransactionStorage) BeginTransaction(ctx context.Context, fn func(context.Context) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTransaction", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeginTransaction indicates an expected call of BeginTransaction.
func (mr *MockTransactionStorageMockRecorder) BeginTransaction(ctx, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTransaction", reflect.TypeOf((*MockTransactionStorage)(nil).BeginTransaction), ctx, fn)
}

// CreateTransaction mocks base method.
func (m *MockTransactionStorage) CreateTransaction(ctx context.Context, senderID, receiverID string, amount int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", ctx, senderID, receiverID, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockTransactionStorageMockRecorder) CreateTransaction(ctx, senderID, receiverID, amount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockTransactionStorage)(nil).CreateTransaction), ctx, senderID, receiverID, amount)
}

// ListByUserID mocks base method.
func (m *MockTransactionStorage) ListByUserID(ctx context.Context, userID string) ([]transaction.Transaction, []transaction.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByUserID", ctx, userID)
	ret0, _ := ret[0].([]transaction.Transaction)
	ret1, _ := ret[1].([]transaction.Transaction)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListByUserID indicates an expected call of ListByUserID.
func (mr *MockTransactionStorageMockRecorder) ListByUserID(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByUserID", reflect.TypeOf((*MockTransactionStorage)(nil).ListByUserID), ctx, userID)
}

// MockUserStorage is a mock of UserStorage interface.
type MockUserStorage struct {
	ctrl     *gomock.Controller
	recorder *MockUserStorageMockRecorder
	isgomock struct{}
}

// MockUserStorageMockRecorder is the mock recorder for MockUserStorage.
type MockUserStorageMockRecorder struct {
	mock *MockUserStorage
}

// NewMockUserStorage creates a new mock instance.
func NewMockUserStorage(ctrl *gomock.Controller) *MockUserStorage {
	mock := &MockUserStorage{ctrl: ctrl}
	mock.recorder = &MockUserStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserStorage) EXPECT() *MockUserStorageMockRecorder {
	return m.recorder
}

// GetById mocks base method.
func (m *MockUserStorage) GetById(ctx context.Context, id string, opts storage.GetOptions) (user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id, opts)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockUserStorageMockRecorder) GetById(ctx, id, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockUserStorage)(nil).GetById), ctx, id, opts)
}

// GetByUsername mocks base method.
func (m *MockUserStorage) GetByUsername(ctx context.Context, username string, opts storage.GetOptions) (user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUsername", ctx, username, opts)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUsername indicates an expected call of GetByUsername.
func (mr *MockUserStorageMockRecorder) GetByUsername(ctx, username, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUsername", reflect.TypeOf((*MockUserStorage)(nil).GetByUsername), ctx, username, opts)
}

// IncrementBalance mocks base method.
func (m *MockUserStorage) IncrementBalance(ctx context.Context, username string, inc int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrementBalance", ctx, username, inc)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrementBalance indicates an expected call of IncrementBalance.
func (mr *MockUserStorageMockRecorder) IncrementBalance(ctx, username, inc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementBalance", reflect.TypeOf((*MockUserStorage)(nil).IncrementBalance), ctx, username, inc)
}
