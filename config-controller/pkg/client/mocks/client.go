// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/config-controller/pkg/client (interfaces: CachedPolicyClient,PolicyClient)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/client.go github.com/stackrox/rox/config-controller/pkg/client CachedPolicyClient,PolicyClient
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	storage "github.com/stackrox/rox/generated/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockCachedPolicyClient is a mock of CachedPolicyClient interface.
type MockCachedPolicyClient struct {
	ctrl     *gomock.Controller
	recorder *MockCachedPolicyClientMockRecorder
	isgomock struct{}
}

// MockCachedPolicyClientMockRecorder is the mock recorder for MockCachedPolicyClient.
type MockCachedPolicyClientMockRecorder struct {
	mock *MockCachedPolicyClient
}

// NewMockCachedPolicyClient creates a new mock instance.
func NewMockCachedPolicyClient(ctrl *gomock.Controller) *MockCachedPolicyClient {
	mock := &MockCachedPolicyClient{ctrl: ctrl}
	mock.recorder = &MockCachedPolicyClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCachedPolicyClient) EXPECT() *MockCachedPolicyClientMockRecorder {
	return m.recorder
}

// CreatePolicy mocks base method.
func (m *MockCachedPolicyClient) CreatePolicy(ctx context.Context, policy *storage.Policy) (*storage.Policy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePolicy", ctx, policy)
	ret0, _ := ret[0].(*storage.Policy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePolicy indicates an expected call of CreatePolicy.
func (mr *MockCachedPolicyClientMockRecorder) CreatePolicy(ctx, policy any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePolicy", reflect.TypeOf((*MockCachedPolicyClient)(nil).CreatePolicy), ctx, policy)
}

// DeletePolicy mocks base method.
func (m *MockCachedPolicyClient) DeletePolicy(ctx context.Context, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePolicy", ctx, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePolicy indicates an expected call of DeletePolicy.
func (mr *MockCachedPolicyClientMockRecorder) DeletePolicy(ctx, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePolicy", reflect.TypeOf((*MockCachedPolicyClient)(nil).DeletePolicy), ctx, name)
}

// EnsureFresh mocks base method.
func (m *MockCachedPolicyClient) EnsureFresh(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureFresh", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureFresh indicates an expected call of EnsureFresh.
func (mr *MockCachedPolicyClientMockRecorder) EnsureFresh(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureFresh", reflect.TypeOf((*MockCachedPolicyClient)(nil).EnsureFresh), ctx)
}

// FlushCache mocks base method.
func (m *MockCachedPolicyClient) FlushCache(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushCache", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// FlushCache indicates an expected call of FlushCache.
func (mr *MockCachedPolicyClientMockRecorder) FlushCache(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushCache", reflect.TypeOf((*MockCachedPolicyClient)(nil).FlushCache), ctx)
}

// GetPolicy mocks base method.
func (m *MockCachedPolicyClient) GetPolicy(ctx context.Context, name string) (*storage.Policy, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicy", ctx, name)
	ret0, _ := ret[0].(*storage.Policy)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPolicy indicates an expected call of GetPolicy.
func (mr *MockCachedPolicyClientMockRecorder) GetPolicy(ctx, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicy", reflect.TypeOf((*MockCachedPolicyClient)(nil).GetPolicy), ctx, name)
}

// ListPolicies mocks base method.
func (m *MockCachedPolicyClient) ListPolicies(ctx context.Context) ([]*storage.Policy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPolicies", ctx)
	ret0, _ := ret[0].([]*storage.Policy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPolicies indicates an expected call of ListPolicies.
func (mr *MockCachedPolicyClientMockRecorder) ListPolicies(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPolicies", reflect.TypeOf((*MockCachedPolicyClient)(nil).ListPolicies), ctx)
}

// UpdatePolicy mocks base method.
func (m *MockCachedPolicyClient) UpdatePolicy(ctx context.Context, policy *storage.Policy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePolicy", ctx, policy)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePolicy indicates an expected call of UpdatePolicy.
func (mr *MockCachedPolicyClientMockRecorder) UpdatePolicy(ctx, policy any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePolicy", reflect.TypeOf((*MockCachedPolicyClient)(nil).UpdatePolicy), ctx, policy)
}

// MockPolicyClient is a mock of PolicyClient interface.
type MockPolicyClient struct {
	ctrl     *gomock.Controller
	recorder *MockPolicyClientMockRecorder
	isgomock struct{}
}

// MockPolicyClientMockRecorder is the mock recorder for MockPolicyClient.
type MockPolicyClientMockRecorder struct {
	mock *MockPolicyClient
}

// NewMockPolicyClient creates a new mock instance.
func NewMockPolicyClient(ctrl *gomock.Controller) *MockPolicyClient {
	mock := &MockPolicyClient{ctrl: ctrl}
	mock.recorder = &MockPolicyClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPolicyClient) EXPECT() *MockPolicyClientMockRecorder {
	return m.recorder
}

// DeletePolicy mocks base method.
func (m *MockPolicyClient) DeletePolicy(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePolicy", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePolicy indicates an expected call of DeletePolicy.
func (mr *MockPolicyClientMockRecorder) DeletePolicy(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePolicy", reflect.TypeOf((*MockPolicyClient)(nil).DeletePolicy), ctx, id)
}

// GetPolicy mocks base method.
func (m *MockPolicyClient) GetPolicy(ctx context.Context, id string) (*storage.Policy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicy", ctx, id)
	ret0, _ := ret[0].(*storage.Policy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPolicy indicates an expected call of GetPolicy.
func (mr *MockPolicyClientMockRecorder) GetPolicy(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicy", reflect.TypeOf((*MockPolicyClient)(nil).GetPolicy), ctx, id)
}

// ListPolicies mocks base method.
func (m *MockPolicyClient) ListPolicies(arg0 context.Context) ([]*storage.ListPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPolicies", arg0)
	ret0, _ := ret[0].([]*storage.ListPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPolicies indicates an expected call of ListPolicies.
func (mr *MockPolicyClientMockRecorder) ListPolicies(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPolicies", reflect.TypeOf((*MockPolicyClient)(nil).ListPolicies), arg0)
}

// PostPolicy mocks base method.
func (m *MockPolicyClient) PostPolicy(arg0 context.Context, arg1 *storage.Policy) (*storage.Policy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostPolicy", arg0, arg1)
	ret0, _ := ret[0].(*storage.Policy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostPolicy indicates an expected call of PostPolicy.
func (mr *MockPolicyClientMockRecorder) PostPolicy(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostPolicy", reflect.TypeOf((*MockPolicyClient)(nil).PostPolicy), arg0, arg1)
}

// PutPolicy mocks base method.
func (m *MockPolicyClient) PutPolicy(arg0 context.Context, arg1 *storage.Policy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutPolicy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutPolicy indicates an expected call of PutPolicy.
func (mr *MockPolicyClientMockRecorder) PutPolicy(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPolicy", reflect.TypeOf((*MockPolicyClient)(nil).PutPolicy), arg0, arg1)
}

// TokenExchange mocks base method.
func (m *MockPolicyClient) TokenExchange(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenExchange", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// TokenExchange indicates an expected call of TokenExchange.
func (mr *MockPolicyClientMockRecorder) TokenExchange(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenExchange", reflect.TypeOf((*MockPolicyClient)(nil).TokenExchange), ctx)
}
