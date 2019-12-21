// Code generated by MockGen. DO NOT EDIT.
// Source: client_manager.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

// MockClientForClusterHandler is a mock of ClientForClusterHandler interface
type MockClientForClusterHandler struct {
	ctrl     *gomock.Controller
	recorder *MockClientForClusterHandlerMockRecorder
}

// MockClientForClusterHandlerMockRecorder is the mock recorder for MockClientForClusterHandler
type MockClientForClusterHandlerMockRecorder struct {
	mock *MockClientForClusterHandler
}

// NewMockClientForClusterHandler creates a new mock instance
func NewMockClientForClusterHandler(ctrl *gomock.Controller) *MockClientForClusterHandler {
	mock := &MockClientForClusterHandler{ctrl: ctrl}
	mock.recorder = &MockClientForClusterHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientForClusterHandler) EXPECT() *MockClientForClusterHandlerMockRecorder {
	return m.recorder
}

// HandleNewClusterClient mocks base method
func (m *MockClientForClusterHandler) HandleNewClusterClient(cluster string, client clients.ResourceClient) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleNewClusterClient", cluster, client)
}

// HandleNewClusterClient indicates an expected call of HandleNewClusterClient
func (mr *MockClientForClusterHandlerMockRecorder) HandleNewClusterClient(cluster, client interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleNewClusterClient", reflect.TypeOf((*MockClientForClusterHandler)(nil).HandleNewClusterClient), cluster, client)
}

// HandleRemovedClusterClient mocks base method
func (m *MockClientForClusterHandler) HandleRemovedClusterClient(cluster string, client clients.ResourceClient) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleRemovedClusterClient", cluster, client)
}

// HandleRemovedClusterClient indicates an expected call of HandleRemovedClusterClient
func (mr *MockClientForClusterHandlerMockRecorder) HandleRemovedClusterClient(cluster, client interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleRemovedClusterClient", reflect.TypeOf((*MockClientForClusterHandler)(nil).HandleRemovedClusterClient), cluster, client)
}

// MockClusterClientGetter is a mock of ClusterClientGetter interface
type MockClusterClientGetter struct {
	ctrl     *gomock.Controller
	recorder *MockClusterClientGetterMockRecorder
}

// MockClusterClientGetterMockRecorder is the mock recorder for MockClusterClientGetter
type MockClusterClientGetterMockRecorder struct {
	mock *MockClusterClientGetter
}

// NewMockClusterClientGetter creates a new mock instance
func NewMockClusterClientGetter(ctrl *gomock.Controller) *MockClusterClientGetter {
	mock := &MockClusterClientGetter{ctrl: ctrl}
	mock.recorder = &MockClusterClientGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterClientGetter) EXPECT() *MockClusterClientGetterMockRecorder {
	return m.recorder
}

// ClientForCluster mocks base method
func (m *MockClusterClientGetter) ClientForCluster(cluster string) (clients.ResourceClient, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientForCluster", cluster)
	ret0, _ := ret[0].(clients.ResourceClient)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ClientForCluster indicates an expected call of ClientForCluster
func (mr *MockClusterClientGetterMockRecorder) ClientForCluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientForCluster", reflect.TypeOf((*MockClusterClientGetter)(nil).ClientForCluster), cluster)
}