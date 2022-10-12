// Code generated by MockGen. DO NOT EDIT.
// Source: e-wallet-app/modules/account/accountbiz (interfaces: AccountRepo)

// Package mockaccountrepo is a generated GoMock package.
package mockaccountrepo

import (
	context "context"
	common "e-wallet-app/common"
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

// GetByAccountNumber mocks base method.
func (m *MockAccountRepo) GetByAccountNumber(arg0 context.Context, arg1 string) (*db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByAccountNumber", arg0, arg1)
	ret0, _ := ret[0].(*db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByAccountNumber indicates an expected call of GetByAccountNumber.
func (mr *MockAccountRepoMockRecorder) GetByAccountNumber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAccountNumber", reflect.TypeOf((*MockAccountRepo)(nil).GetByAccountNumber), arg0, arg1)
}

// GetByID mocks base method.
func (m *MockAccountRepo) GetByID(arg0 context.Context, arg1 int64) (*db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(*db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockAccountRepoMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockAccountRepo)(nil).GetByID), arg0, arg1)
}

// List mocks base method.
func (m *MockAccountRepo) List(arg0 context.Context, arg1 *accountmodel.AccountFilter, arg2 *common.Paging, arg3 *common.Sorting, arg4 common.Requester) ([]db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAccountRepoMockRecorder) List(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAccountRepo)(nil).List), arg0, arg1, arg2, arg3, arg4)
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
