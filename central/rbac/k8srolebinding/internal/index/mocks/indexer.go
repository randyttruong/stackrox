// Code generated by MockGen. DO NOT EDIT.
// Source: indexer.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	search "github.com/stackrox/rox/pkg/search"
	blevesearch "github.com/stackrox/rox/pkg/search/blevesearch"
)

// MockIndexer is a mock of Indexer interface.
type MockIndexer struct {
	ctrl     *gomock.Controller
	recorder *MockIndexerMockRecorder
}

// MockIndexerMockRecorder is the mock recorder for MockIndexer.
type MockIndexerMockRecorder struct {
	mock *MockIndexer
}

// NewMockIndexer creates a new mock instance.
func NewMockIndexer(ctrl *gomock.Controller) *MockIndexer {
	mock := &MockIndexer{ctrl: ctrl}
	mock.recorder = &MockIndexerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIndexer) EXPECT() *MockIndexerMockRecorder {
	return m.recorder
}

// AddK8SRoleBinding mocks base method.
func (m *MockIndexer) AddK8SRoleBinding(k8srolebinding *storage.K8SRoleBinding) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddK8SRoleBinding", k8srolebinding)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddK8SRoleBinding indicates an expected call of AddK8SRoleBinding.
func (mr *MockIndexerMockRecorder) AddK8SRoleBinding(k8srolebinding interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddK8SRoleBinding", reflect.TypeOf((*MockIndexer)(nil).AddK8SRoleBinding), k8srolebinding)
}

// AddK8SRoleBindings mocks base method.
func (m *MockIndexer) AddK8SRoleBindings(k8srolebindings []*storage.K8SRoleBinding) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddK8SRoleBindings", k8srolebindings)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddK8SRoleBindings indicates an expected call of AddK8SRoleBindings.
func (mr *MockIndexerMockRecorder) AddK8SRoleBindings(k8srolebindings interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddK8SRoleBindings", reflect.TypeOf((*MockIndexer)(nil).AddK8SRoleBindings), k8srolebindings)
}

// Count mocks base method.
func (m *MockIndexer) Count(ctx context.Context, q *v1.Query, opts ...blevesearch.SearchOption) (int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, q}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Count", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockIndexerMockRecorder) Count(ctx, q interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, q}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockIndexer)(nil).Count), varargs...)
}

// DeleteK8SRoleBinding mocks base method.
func (m *MockIndexer) DeleteK8SRoleBinding(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteK8SRoleBinding", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteK8SRoleBinding indicates an expected call of DeleteK8SRoleBinding.
func (mr *MockIndexerMockRecorder) DeleteK8SRoleBinding(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteK8SRoleBinding", reflect.TypeOf((*MockIndexer)(nil).DeleteK8SRoleBinding), id)
}

// DeleteK8SRoleBindings mocks base method.
func (m *MockIndexer) DeleteK8SRoleBindings(ids []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteK8SRoleBindings", ids)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteK8SRoleBindings indicates an expected call of DeleteK8SRoleBindings.
func (mr *MockIndexerMockRecorder) DeleteK8SRoleBindings(ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteK8SRoleBindings", reflect.TypeOf((*MockIndexer)(nil).DeleteK8SRoleBindings), ids)
}

// MarkInitialIndexingComplete mocks base method.
func (m *MockIndexer) MarkInitialIndexingComplete() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkInitialIndexingComplete")
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkInitialIndexingComplete indicates an expected call of MarkInitialIndexingComplete.
func (mr *MockIndexerMockRecorder) MarkInitialIndexingComplete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkInitialIndexingComplete", reflect.TypeOf((*MockIndexer)(nil).MarkInitialIndexingComplete))
}

// NeedsInitialIndexing mocks base method.
func (m *MockIndexer) NeedsInitialIndexing() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NeedsInitialIndexing")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NeedsInitialIndexing indicates an expected call of NeedsInitialIndexing.
func (mr *MockIndexerMockRecorder) NeedsInitialIndexing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NeedsInitialIndexing", reflect.TypeOf((*MockIndexer)(nil).NeedsInitialIndexing))
}

// Search mocks base method.
func (m *MockIndexer) Search(ctx context.Context, q *v1.Query, opts ...blevesearch.SearchOption) ([]search.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, q}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Search", varargs...)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockIndexerMockRecorder) Search(ctx, q interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, q}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockIndexer)(nil).Search), varargs...)
}
