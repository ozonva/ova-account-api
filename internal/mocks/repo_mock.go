// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozonva/ova-account-api/internal/repo (interfaces: Repo)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/ozonva/ova-account-api/internal/entity"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddAccounts mocks base method.
func (m *MockRepo) AddAccounts(arg0 context.Context, arg1 []entity.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAccounts", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAccounts indicates an expected call of AddAccounts.
func (mr *MockRepoMockRecorder) AddAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAccounts", reflect.TypeOf((*MockRepo)(nil).AddAccounts), arg0, arg1)
}

// DescribeAccount mocks base method.
func (m *MockRepo) DescribeAccount(arg0 context.Context, arg1 string) (*entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeAccount", arg0, arg1)
	ret0, _ := ret[0].(*entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAccount indicates an expected call of DescribeAccount.
func (mr *MockRepoMockRecorder) DescribeAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAccount", reflect.TypeOf((*MockRepo)(nil).DescribeAccount), arg0, arg1)
}

// ListAccounts mocks base method.
func (m *MockRepo) ListAccounts(arg0 context.Context, arg1, arg2 uint64) ([]entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccounts", arg0, arg1, arg2)
	ret0, _ := ret[0].([]entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts.
func (mr *MockRepoMockRecorder) ListAccounts(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockRepo)(nil).ListAccounts), arg0, arg1, arg2)
}

// RemoveAccount mocks base method.
func (m *MockRepo) RemoveAccount(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAccount indicates an expected call of RemoveAccount.
func (mr *MockRepoMockRecorder) RemoveAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAccount", reflect.TypeOf((*MockRepo)(nil).RemoveAccount), arg0, arg1)
}

// UpdateAccount mocks base method.
func (m *MockRepo) UpdateAccount(arg0 context.Context, arg1 entity.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockRepoMockRecorder) UpdateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockRepo)(nil).UpdateAccount), arg0, arg1)
}
