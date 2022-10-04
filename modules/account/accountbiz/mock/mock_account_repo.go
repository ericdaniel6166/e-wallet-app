// Code generated by MockGen. DO NOT EDIT.
// Source: e-wallet-app/modules/account/accountbiz (interfaces: AccountRepo)

// Package mockaccountrepo is a generated GoMock package.
package mockaccountrepo

import (
	context "context"
	db "e-wallet-app/db/sqlc"
	accountmodel "e-wallet-app/modules/account/accountmodel"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAccountRepo is a mock of AccountRepo interface.
type MockAccountRepo struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRepoMockRecorder
}

// MockAccountRepoMockRecorder is the mock recorder for MockAccountRepo.
type MockAccountRepoMockRecorder struct {
	mock *MockAccountRepo
}

// NewMockAccountRepo creates a new mock instance.
func NewMockAccountRepo(ctrl *gomock.Controller) *MockAccountRepo {
	mock := &MockAccountRepo{ctrl: ctrl}
	mock.recorder = &MockAccountRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountRepo) EXPECT() *MockAccountRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAccountRepo) Create(arg0 context.Context, arg1 *accountmodel.CreateAccountRequest) (*db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAccountRepoMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccountRepo)(nil).Create), arg0, arg1)
}

// GetById mocks base method.
func (m *MockAccountRepo) GetById(arg0 context.Context, arg1 *accountmodel.GetAccountRequest) (*db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0, arg1)
	ret0, _ := ret[0].(*db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockAccountRepoMockRecorder) GetById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockAccountRepo)(nil).GetById), arg0, arg1)
}

// List mocks base method.
func (m *MockAccountRepo) List(arg0 context.Context, arg1 *accountmodel.ListAccountRequest) ([]db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAccountRepoMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAccountRepo)(nil).List), arg0, arg1)
}

// Transfer mocks base method.
func (m *MockAccountRepo) Transfer(arg0 context.Context, arg1 *accountmodel.TransferAccountRequest) (*accountmodel.TransferAccountResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transfer", arg0, arg1)
	ret0, _ := ret[0].(*accountmodel.TransferAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Transfer indicates an expected call of Transfer.
func (mr *MockAccountRepoMockRecorder) Transfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transfer", reflect.TypeOf((*MockAccountRepo)(nil).Transfer), arg0, arg1)
}
