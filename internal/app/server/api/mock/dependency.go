// Code generated by MockGen. DO NOT EDIT.
// Source: dependency.go
//
// Generated by this command:
//
//	mockgen -source dependency.go -destination mock/dependency.go
//

// Package mock_api is a generated GoMock package.
package mock_api

import (
	context "context"
	reflect "reflect"

	auth "github.com/YrWaifu/test_go_back/internal/usecase/auth"
	info "github.com/YrWaifu/test_go_back/internal/usecase/info"
	purchase "github.com/YrWaifu/test_go_back/internal/usecase/purchase"
	transaction "github.com/YrWaifu/test_go_back/internal/usecase/transaction"
	gomock "go.uber.org/mock/gomock"
)

// MockAuthUsecase is a mock of AuthUsecase interface.
type MockAuthUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockAuthUsecaseMockRecorder
	isgomock struct{}
}

// MockAuthUsecaseMockRecorder is the mock recorder for MockAuthUsecase.
type MockAuthUsecaseMockRecorder struct {
	mock *MockAuthUsecase
}

// NewMockAuthUsecase creates a new mock instance.
func NewMockAuthUsecase(ctrl *gomock.Controller) *MockAuthUsecase {
	mock := &MockAuthUsecase{ctrl: ctrl}
	mock.recorder = &MockAuthUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthUsecase) EXPECT() *MockAuthUsecaseMockRecorder {
	return m.recorder
}

// SignIn mocks base method.
func (m *MockAuthUsecase) SignIn(ctx context.Context, req auth.SignInRequest) (auth.SignInResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", ctx, req)
	ret0, _ := ret[0].(auth.SignInResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockAuthUsecaseMockRecorder) SignIn(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockAuthUsecase)(nil).SignIn), ctx, req)
}

// MockPurchaseUsecase is a mock of PurchaseUsecase interface.
type MockPurchaseUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockPurchaseUsecaseMockRecorder
	isgomock struct{}
}

// MockPurchaseUsecaseMockRecorder is the mock recorder for MockPurchaseUsecase.
type MockPurchaseUsecaseMockRecorder struct {
	mock *MockPurchaseUsecase
}

// NewMockPurchaseUsecase creates a new mock instance.
func NewMockPurchaseUsecase(ctrl *gomock.Controller) *MockPurchaseUsecase {
	mock := &MockPurchaseUsecase{ctrl: ctrl}
	mock.recorder = &MockPurchaseUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPurchaseUsecase) EXPECT() *MockPurchaseUsecaseMockRecorder {
	return m.recorder
}

// BuyMerch mocks base method.
func (m *MockPurchaseUsecase) BuyMerch(ctx context.Context, req purchase.BuyMerchRequest) (purchase.BuyMerchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuyMerch", ctx, req)
	ret0, _ := ret[0].(purchase.BuyMerchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuyMerch indicates an expected call of BuyMerch.
func (mr *MockPurchaseUsecaseMockRecorder) BuyMerch(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuyMerch", reflect.TypeOf((*MockPurchaseUsecase)(nil).BuyMerch), ctx, req)
}

// MockTransactionUsecase is a mock of TransactionUsecase interface.
type MockTransactionUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionUsecaseMockRecorder
	isgomock struct{}
}

// MockTransactionUsecaseMockRecorder is the mock recorder for MockTransactionUsecase.
type MockTransactionUsecaseMockRecorder struct {
	mock *MockTransactionUsecase
}

// NewMockTransactionUsecase creates a new mock instance.
func NewMockTransactionUsecase(ctrl *gomock.Controller) *MockTransactionUsecase {
	mock := &MockTransactionUsecase{ctrl: ctrl}
	mock.recorder = &MockTransactionUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionUsecase) EXPECT() *MockTransactionUsecaseMockRecorder {
	return m.recorder
}

// Transfer mocks base method.
func (m *MockTransactionUsecase) Transfer(ctx context.Context, req transaction.TransferRequest) (transaction.TransferResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transfer", ctx, req)
	ret0, _ := ret[0].(transaction.TransferResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Transfer indicates an expected call of Transfer.
func (mr *MockTransactionUsecaseMockRecorder) Transfer(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transfer", reflect.TypeOf((*MockTransactionUsecase)(nil).Transfer), ctx, req)
}

// MockInfoUsecase is a mock of InfoUsecase interface.
type MockInfoUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockInfoUsecaseMockRecorder
	isgomock struct{}
}

// MockInfoUsecaseMockRecorder is the mock recorder for MockInfoUsecase.
type MockInfoUsecaseMockRecorder struct {
	mock *MockInfoUsecase
}

// NewMockInfoUsecase creates a new mock instance.
func NewMockInfoUsecase(ctrl *gomock.Controller) *MockInfoUsecase {
	mock := &MockInfoUsecase{ctrl: ctrl}
	mock.recorder = &MockInfoUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInfoUsecase) EXPECT() *MockInfoUsecaseMockRecorder {
	return m.recorder
}

// Info mocks base method.
func (m *MockInfoUsecase) Info(ctx context.Context, req info.InfoRequest) (info.InfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", ctx, req)
	ret0, _ := ret[0].(info.InfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info.
func (mr *MockInfoUsecaseMockRecorder) Info(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockInfoUsecase)(nil).Info), ctx, req)
}
