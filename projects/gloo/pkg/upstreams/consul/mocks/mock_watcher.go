// Code generated by MockGen. DO NOT EDIT.
// Source: watcher.go

// Package mock_consul is a generated GoMock package.
package mock_consul

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/hashicorp/consul/api"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	consul "github.com/solo-io/gloo/projects/gloo/pkg/upstreams/consul"
)

// MockConsulWatcher is a mock of ConsulWatcher interface.
type MockConsulWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockConsulWatcherMockRecorder
}

// MockConsulWatcherMockRecorder is the mock recorder for MockConsulWatcher.
type MockConsulWatcherMockRecorder struct {
	mock *MockConsulWatcher
}

// NewMockConsulWatcher creates a new mock instance.
func NewMockConsulWatcher(ctrl *gomock.Controller) *MockConsulWatcher {
	mock := &MockConsulWatcher{ctrl: ctrl}
	mock.recorder = &MockConsulWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsulWatcher) EXPECT() *MockConsulWatcherMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockConsulWatcher) Connect(service, tag string, q *api.QueryOptions) ([]*api.CatalogService, *api.QueryMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", service, tag, q)
	ret0, _ := ret[0].([]*api.CatalogService)
	ret1, _ := ret[1].(*api.QueryMeta)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Connect indicates an expected call of Connect.
func (mr *MockConsulWatcherMockRecorder) Connect(service, tag, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockConsulWatcher)(nil).Connect), service, tag, q)
}

// DataCenters mocks base method.
func (m *MockConsulWatcher) DataCenters() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataCenters")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataCenters indicates an expected call of DataCenters.
func (mr *MockConsulWatcherMockRecorder) DataCenters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataCenters", reflect.TypeOf((*MockConsulWatcher)(nil).DataCenters))
}

// Service mocks base method.
func (m *MockConsulWatcher) Service(service, tag string, q *api.QueryOptions) ([]*api.CatalogService, *api.QueryMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Service", service, tag, q)
	ret0, _ := ret[0].([]*api.CatalogService)
	ret1, _ := ret[1].(*api.QueryMeta)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Service indicates an expected call of Service.
func (mr *MockConsulWatcherMockRecorder) Service(service, tag, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Service", reflect.TypeOf((*MockConsulWatcher)(nil).Service), service, tag, q)
}

// Services mocks base method.
func (m *MockConsulWatcher) Services(q *api.QueryOptions) (map[string][]string, *api.QueryMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Services", q)
	ret0, _ := ret[0].(map[string][]string)
	ret1, _ := ret[1].(*api.QueryMeta)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Services indicates an expected call of Services.
func (mr *MockConsulWatcherMockRecorder) Services(q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Services", reflect.TypeOf((*MockConsulWatcher)(nil).Services), q)
}

// WatchServices mocks base method.
func (m *MockConsulWatcher) WatchServices(ctx context.Context, dataCenters []string, cm v1.Settings_ConsulUpstreamDiscoveryConfiguration_ConsulConsistencyModes) (<-chan []*consul.ServiceMeta, <-chan error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchServices", ctx, dataCenters, cm)
	ret0, _ := ret[0].(<-chan []*consul.ServiceMeta)
	ret1, _ := ret[1].(<-chan error)
	return ret0, ret1
}

// WatchServices indicates an expected call of WatchServices.
func (mr *MockConsulWatcherMockRecorder) WatchServices(ctx, dataCenters, cm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchServices", reflect.TypeOf((*MockConsulWatcher)(nil).WatchServices), ctx, dataCenters, cm)
}
