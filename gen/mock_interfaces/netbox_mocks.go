// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/netbox/interfaces/netbox.go
//
// Generated by this command:
//
//	mockgen -destination gen/mock_interfaces/netbox_mocks.go -source=pkg/netbox/interfaces/netbox.go
//

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	reflect "reflect"

	runtime "github.com/go-openapi/runtime"
	extras "github.com/netbox-community/go-netbox/v3/netbox/client/extras"
	ipam "github.com/netbox-community/go-netbox/v3/netbox/client/ipam"
	tenancy "github.com/netbox-community/go-netbox/v3/netbox/client/tenancy"
	gomock "go.uber.org/mock/gomock"
)

// MockIpamInterface is a mock of IpamInterface interface.
type MockIpamInterface struct {
	ctrl     *gomock.Controller
	recorder *MockIpamInterfaceMockRecorder
}

// MockIpamInterfaceMockRecorder is the mock recorder for MockIpamInterface.
type MockIpamInterfaceMockRecorder struct {
	mock *MockIpamInterface
}

// NewMockIpamInterface creates a new mock instance.
func NewMockIpamInterface(ctrl *gomock.Controller) *MockIpamInterface {
	mock := &MockIpamInterface{ctrl: ctrl}
	mock.recorder = &MockIpamInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIpamInterface) EXPECT() *MockIpamInterfaceMockRecorder {
	return m.recorder
}

// IpamIPAddressesCreate mocks base method.
func (m *MockIpamInterface) IpamIPAddressesCreate(params *ipam.IpamIPAddressesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ipam.ClientOption) (*ipam.IpamIPAddressesCreateCreated, error) {
	m.ctrl.T.Helper()
	varargs := []any{params, authInfo}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IpamIPAddressesCreate", varargs...)
	ret0, _ := ret[0].(*ipam.IpamIPAddressesCreateCreated)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpamIPAddressesCreate indicates an expected call of IpamIPAddressesCreate.
func (mr *MockIpamInterfaceMockRecorder) IpamIPAddressesCreate(params, authInfo any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{params, authInfo}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpamIPAddressesCreate", reflect.TypeOf((*MockIpamInterface)(nil).IpamIPAddressesCreate), varargs...)
}

// IpamIPAddressesDelete mocks base method.
func (m *MockIpamInterface) IpamIPAddressesDelete(params *ipam.IpamIPAddressesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ipam.ClientOption) (*ipam.IpamIPAddressesDeleteNoContent, error) {
	m.ctrl.T.Helper()
	varargs := []any{params, authInfo}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IpamIPAddressesDelete", varargs...)
	ret0, _ := ret[0].(*ipam.IpamIPAddressesDeleteNoContent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpamIPAddressesDelete indicates an expected call of IpamIPAddressesDelete.
func (mr *MockIpamInterfaceMockRecorder) IpamIPAddressesDelete(params, authInfo any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{params, authInfo}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpamIPAddressesDelete", reflect.TypeOf((*MockIpamInterface)(nil).IpamIPAddressesDelete), varargs...)
}

// IpamIPAddressesList mocks base method.
func (m *MockIpamInterface) IpamIPAddressesList(params *ipam.IpamIPAddressesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ipam.ClientOption) (*ipam.IpamIPAddressesListOK, error) {
	m.ctrl.T.Helper()
	varargs := []any{params, authInfo}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IpamIPAddressesList", varargs...)
	ret0, _ := ret[0].(*ipam.IpamIPAddressesListOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpamIPAddressesList indicates an expected call of IpamIPAddressesList.
func (mr *MockIpamInterfaceMockRecorder) IpamIPAddressesList(params, authInfo any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{params, authInfo}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpamIPAddressesList", reflect.TypeOf((*MockIpamInterface)(nil).IpamIPAddressesList), varargs...)
}

// IpamIPAddressesUpdate mocks base method.
func (m *MockIpamInterface) IpamIPAddressesUpdate(params *ipam.IpamIPAddressesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ipam.ClientOption) (*ipam.IpamIPAddressesUpdateOK, error) {
	m.ctrl.T.Helper()
	varargs := []any{params, authInfo}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IpamIPAddressesUpdate", varargs...)
	ret0, _ := ret[0].(*ipam.IpamIPAddressesUpdateOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpamIPAddressesUpdate indicates an expected call of IpamIPAddressesUpdate.
func (mr *MockIpamInterfaceMockRecorder) IpamIPAddressesUpdate(params, authInfo any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{params, authInfo}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpamIPAddressesUpdate", reflect.TypeOf((*MockIpamInterface)(nil).IpamIPAddressesUpdate), varargs...)
}

// IpamPrefixesAvailableIpsList mocks base method.
func (m *MockIpamInterface) IpamPrefixesAvailableIpsList(params *ipam.IpamPrefixesAvailableIpsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ipam.ClientOption) (*ipam.IpamPrefixesAvailableIpsListOK, error) {
	m.ctrl.T.Helper()
	varargs := []any{params, authInfo}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IpamPrefixesAvailableIpsList", varargs...)
	ret0, _ := ret[0].(*ipam.IpamPrefixesAvailableIpsListOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpamPrefixesAvailableIpsList indicates an expected call of IpamPrefixesAvailableIpsList.
func (mr *MockIpamInterfaceMockRecorder) IpamPrefixesAvailableIpsList(params, authInfo any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{params, authInfo}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpamPrefixesAvailableIpsList", reflect.TypeOf((*MockIpamInterface)(nil).IpamPrefixesAvailableIpsList), varargs...)
}

// IpamPrefixesAvailablePrefixesList mocks base method.
func (m *MockIpamInterface) IpamPrefixesAvailablePrefixesList(params *ipam.IpamPrefixesAvailablePrefixesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ipam.ClientOption) (*ipam.IpamPrefixesAvailablePrefixesListOK, error) {
	m.ctrl.T.Helper()
	varargs := []any{params, authInfo}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IpamPrefixesAvailablePrefixesList", varargs...)
	ret0, _ := ret[0].(*ipam.IpamPrefixesAvailablePrefixesListOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpamPrefixesAvailablePrefixesList indicates an expected call of IpamPrefixesAvailablePrefixesList.
func (mr *MockIpamInterfaceMockRecorder) IpamPrefixesAvailablePrefixesList(params, authInfo any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{params, authInfo}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpamPrefixesAvailablePrefixesList", reflect.TypeOf((*MockIpamInterface)(nil).IpamPrefixesAvailablePrefixesList), varargs...)
}

// IpamPrefixesCreate mocks base method.
func (m *MockIpamInterface) IpamPrefixesCreate(params *ipam.IpamPrefixesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ipam.ClientOption) (*ipam.IpamPrefixesCreateCreated, error) {
	m.ctrl.T.Helper()
	varargs := []any{params, authInfo}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IpamPrefixesCreate", varargs...)
	ret0, _ := ret[0].(*ipam.IpamPrefixesCreateCreated)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpamPrefixesCreate indicates an expected call of IpamPrefixesCreate.
func (mr *MockIpamInterfaceMockRecorder) IpamPrefixesCreate(params, authInfo any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{params, authInfo}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpamPrefixesCreate", reflect.TypeOf((*MockIpamInterface)(nil).IpamPrefixesCreate), varargs...)
}

// IpamPrefixesDelete mocks base method.
func (m *MockIpamInterface) IpamPrefixesDelete(params *ipam.IpamPrefixesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ipam.ClientOption) (*ipam.IpamPrefixesDeleteNoContent, error) {
	m.ctrl.T.Helper()
	varargs := []any{params, authInfo}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IpamPrefixesDelete", varargs...)
	ret0, _ := ret[0].(*ipam.IpamPrefixesDeleteNoContent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpamPrefixesDelete indicates an expected call of IpamPrefixesDelete.
func (mr *MockIpamInterfaceMockRecorder) IpamPrefixesDelete(params, authInfo any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{params, authInfo}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpamPrefixesDelete", reflect.TypeOf((*MockIpamInterface)(nil).IpamPrefixesDelete), varargs...)
}

// IpamPrefixesList mocks base method.
func (m *MockIpamInterface) IpamPrefixesList(params *ipam.IpamPrefixesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ipam.ClientOption) (*ipam.IpamPrefixesListOK, error) {
	m.ctrl.T.Helper()
	varargs := []any{params, authInfo}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IpamPrefixesList", varargs...)
	ret0, _ := ret[0].(*ipam.IpamPrefixesListOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpamPrefixesList indicates an expected call of IpamPrefixesList.
func (mr *MockIpamInterfaceMockRecorder) IpamPrefixesList(params, authInfo any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{params, authInfo}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpamPrefixesList", reflect.TypeOf((*MockIpamInterface)(nil).IpamPrefixesList), varargs...)
}

// IpamPrefixesUpdate mocks base method.
func (m *MockIpamInterface) IpamPrefixesUpdate(params *ipam.IpamPrefixesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ipam.ClientOption) (*ipam.IpamPrefixesUpdateOK, error) {
	m.ctrl.T.Helper()
	varargs := []any{params, authInfo}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IpamPrefixesUpdate", varargs...)
	ret0, _ := ret[0].(*ipam.IpamPrefixesUpdateOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpamPrefixesUpdate indicates an expected call of IpamPrefixesUpdate.
func (mr *MockIpamInterfaceMockRecorder) IpamPrefixesUpdate(params, authInfo any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{params, authInfo}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpamPrefixesUpdate", reflect.TypeOf((*MockIpamInterface)(nil).IpamPrefixesUpdate), varargs...)
}

// MockTenancyInterface is a mock of TenancyInterface interface.
type MockTenancyInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTenancyInterfaceMockRecorder
}

// MockTenancyInterfaceMockRecorder is the mock recorder for MockTenancyInterface.
type MockTenancyInterfaceMockRecorder struct {
	mock *MockTenancyInterface
}

// NewMockTenancyInterface creates a new mock instance.
func NewMockTenancyInterface(ctrl *gomock.Controller) *MockTenancyInterface {
	mock := &MockTenancyInterface{ctrl: ctrl}
	mock.recorder = &MockTenancyInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTenancyInterface) EXPECT() *MockTenancyInterfaceMockRecorder {
	return m.recorder
}

// TenancyTenantsList mocks base method.
func (m *MockTenancyInterface) TenancyTenantsList(params *tenancy.TenancyTenantsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...tenancy.ClientOption) (*tenancy.TenancyTenantsListOK, error) {
	m.ctrl.T.Helper()
	varargs := []any{params, authInfo}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TenancyTenantsList", varargs...)
	ret0, _ := ret[0].(*tenancy.TenancyTenantsListOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TenancyTenantsList indicates an expected call of TenancyTenantsList.
func (mr *MockTenancyInterfaceMockRecorder) TenancyTenantsList(params, authInfo any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{params, authInfo}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenancyTenantsList", reflect.TypeOf((*MockTenancyInterface)(nil).TenancyTenantsList), varargs...)
}

// MockExtrasInterface is a mock of ExtrasInterface interface.
type MockExtrasInterface struct {
	ctrl     *gomock.Controller
	recorder *MockExtrasInterfaceMockRecorder
}

// MockExtrasInterfaceMockRecorder is the mock recorder for MockExtrasInterface.
type MockExtrasInterfaceMockRecorder struct {
	mock *MockExtrasInterface
}

// NewMockExtrasInterface creates a new mock instance.
func NewMockExtrasInterface(ctrl *gomock.Controller) *MockExtrasInterface {
	mock := &MockExtrasInterface{ctrl: ctrl}
	mock.recorder = &MockExtrasInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExtrasInterface) EXPECT() *MockExtrasInterfaceMockRecorder {
	return m.recorder
}

// ExtrasCustomFieldsList mocks base method.
func (m *MockExtrasInterface) ExtrasCustomFieldsList(params *extras.ExtrasCustomFieldsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...extras.ClientOption) (*extras.ExtrasCustomFieldsListOK, error) {
	m.ctrl.T.Helper()
	varargs := []any{params, authInfo}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExtrasCustomFieldsList", varargs...)
	ret0, _ := ret[0].(*extras.ExtrasCustomFieldsListOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExtrasCustomFieldsList indicates an expected call of ExtrasCustomFieldsList.
func (mr *MockExtrasInterfaceMockRecorder) ExtrasCustomFieldsList(params, authInfo any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{params, authInfo}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtrasCustomFieldsList", reflect.TypeOf((*MockExtrasInterface)(nil).ExtrasCustomFieldsList), varargs...)
}