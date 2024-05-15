// Code generated by MockGen. DO NOT EDIT.
// Source: google.golang.org/grpc/credentials (interfaces: PerRPCCredentials)
//
// Generated by this command:
//
//	mockgen -package grpcmock google.golang.org/grpc/credentials PerRPCCredentials
//

// Package grpcmock is a generated GoMock package.
package grpcmock

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPerRPCCredentials is a mock of PerRPCCredentials interface.
type MockPerRPCCredentials struct {
	ctrl     *gomock.Controller
	recorder *MockPerRPCCredentialsMockRecorder
}

// MockPerRPCCredentialsMockRecorder is the mock recorder for MockPerRPCCredentials.
type MockPerRPCCredentialsMockRecorder struct {
	mock *MockPerRPCCredentials
}

// NewMockPerRPCCredentials creates a new mock instance.
func NewMockPerRPCCredentials(ctrl *gomock.Controller) *MockPerRPCCredentials {
	mock := &MockPerRPCCredentials{ctrl: ctrl}
	mock.recorder = &MockPerRPCCredentialsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPerRPCCredentials) EXPECT() *MockPerRPCCredentialsMockRecorder {
	return m.recorder
}

// GetRequestMetadata mocks base method.
func (m *MockPerRPCCredentials) GetRequestMetadata(arg0 context.Context, arg1 ...string) (map[string]string, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRequestMetadata", varargs...)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequestMetadata indicates an expected call of GetRequestMetadata.
func (mr *MockPerRPCCredentialsMockRecorder) GetRequestMetadata(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestMetadata", reflect.TypeOf((*MockPerRPCCredentials)(nil).GetRequestMetadata), varargs...)
}

// RequireTransportSecurity mocks base method.
func (m *MockPerRPCCredentials) RequireTransportSecurity() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequireTransportSecurity")
	ret0, _ := ret[0].(bool)
	return ret0
}

// RequireTransportSecurity indicates an expected call of RequireTransportSecurity.
func (mr *MockPerRPCCredentialsMockRecorder) RequireTransportSecurity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequireTransportSecurity", reflect.TypeOf((*MockPerRPCCredentials)(nil).RequireTransportSecurity))
}