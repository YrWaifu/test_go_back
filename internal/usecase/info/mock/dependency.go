// Code generated by MockGen. DO NOT EDIT.
// Source: dependency.go
//
// Generated by this command:
//
//	mockgen -source dependency.go -destination mock/dependency.go
//

// Package mock_info is a generated GoMock package.
package mock_info

import (
	context "context"
	reflect "reflect"

	service "github.com/YrWaifu/test_go_back/internal/domain/purchase/service"
	service0 "github.com/YrWaifu/test_go_back/internal/domain/transaction/service"
	service1 "github.com/YrWaifu/test_go_back/internal/domain/user/service"
	gomock "go.uber.org/mock/gomock"
)

// MockPurchaseService is a mock of PurchaseService interface.
type MockPurchaseService struct {
	ctrl     *gomock.Controller
	recorder *MockPurchaseServiceMockRecorder
	isgomock struct{}
}

// MockPurchaseServiceMockRecorder is the mock recorder for MockPurchaseService.
type MockPurchaseServiceMockRecorder struct {
	mock *MockPurchaseService
}

// NewMockPurchaseService creates a new mock instance.
func NewMockPurchaseService(ctrl *gomock.Controller) *MockPurchaseService {
	mock := &MockPurchaseService{ctrl: ctrl}
	mock.recorder = &MockPurchaseServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPurchaseService) EXPECT() *MockPurchaseServiceMockRecorder {
	return m.recorder
}

// ListByUserID mocks base method.
func (m *MockPurchaseService) ListByUserID(ctx context.Context, r service.ListByUserIDRequest) (service.ListByUserIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByUserID", ctx, r)
	ret0, _ := ret[0].(service.ListByUserIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByUserID indicates an expected call of ListByUserID.
func (mr *MockPurchaseServiceMockRecorder) ListByUserID(ctx, r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByUserID", reflect.TypeOf((*MockPurchaseService)(nil).ListByUserID), ctx, r)
}

// MockTransactionService is a mock of TransactionService interface.
type MockTransactionService struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionServiceMockRecorder
	isgomock struct{}
}

// MockTransactionServiceMockRecorder is the mock recorder for MockTransactionService.
type MockTransactionServiceMockRecorder struct {
	mock *MockTransactionService
}

// NewMockTransactionService creates a new mock instance.
func NewMockTransactionService(ctrl *gomock.Controller) *MockTransactionService {
	mock := &MockTransactionService{ctrl: ctrl}
	mock.recorder = &MockTransactionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionService) EXPECT() *MockTransactionServiceMockRecorder {
	return m.recorder
}

// ListByUserID mocks base method.
func (m *MockTransactionService) ListByUserID(ctx context.Context, r service0.ListByUserIDRequest) (service0.ListByUserIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByUserID", ctx, r)
	ret0, _ := ret[0].(service0.ListByUserIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByUserID indicates an expected call of ListByUserID.
func (mr *MockTransactionServiceMockRecorder) ListByUserID(ctx, r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByUserID", reflect.TypeOf((*MockTransactionService)(nil).ListByUserID), ctx, r)
}

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
	isgomock struct{}
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// GetByID mocks base method.
func (m *MockUserService) GetByID(ctx context.Context, req service1.GetByIDRequest) (service1.GetByIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, req)
	ret0, _ := ret[0].(service1.GetByIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockUserServiceMockRecorder) GetByID(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUserService)(nil).GetByID), ctx, req)
}
