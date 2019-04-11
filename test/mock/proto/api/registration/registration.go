// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spiffe/spire/proto/api/registration (interfaces: RegistrationClient,RegistrationServer)

// Package mock_registration is a generated GoMock package.
package mock_registration

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	registration "github.com/spiffe/spire/proto/api/registration"
	common "github.com/spiffe/spire/proto/common"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockRegistrationClient is a mock of RegistrationClient interface
type MockRegistrationClient struct {
	ctrl     *gomock.Controller
	recorder *MockRegistrationClientMockRecorder
}

// MockRegistrationClientMockRecorder is the mock recorder for MockRegistrationClient
type MockRegistrationClientMockRecorder struct {
	mock *MockRegistrationClient
}

// NewMockRegistrationClient creates a new mock instance
func NewMockRegistrationClient(ctrl *gomock.Controller) *MockRegistrationClient {
	mock := &MockRegistrationClient{ctrl: ctrl}
	mock.recorder = &MockRegistrationClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegistrationClient) EXPECT() *MockRegistrationClientMockRecorder {
	return m.recorder
}

// CreateEntry mocks base method
func (m *MockRegistrationClient) CreateEntry(arg0 context.Context, arg1 *common.RegistrationEntry, arg2 ...grpc.CallOption) (*registration.RegistrationEntryID, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateEntry", varargs...)
	ret0, _ := ret[0].(*registration.RegistrationEntryID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEntry indicates an expected call of CreateEntry
func (mr *MockRegistrationClientMockRecorder) CreateEntry(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEntry", reflect.TypeOf((*MockRegistrationClient)(nil).CreateEntry), varargs...)
}

// CreateFederatedBundle mocks base method
func (m *MockRegistrationClient) CreateFederatedBundle(arg0 context.Context, arg1 *registration.FederatedBundle, arg2 ...grpc.CallOption) (*common.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateFederatedBundle", varargs...)
	ret0, _ := ret[0].(*common.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFederatedBundle indicates an expected call of CreateFederatedBundle
func (mr *MockRegistrationClientMockRecorder) CreateFederatedBundle(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFederatedBundle", reflect.TypeOf((*MockRegistrationClient)(nil).CreateFederatedBundle), varargs...)
}

// CreateJoinToken mocks base method
func (m *MockRegistrationClient) CreateJoinToken(arg0 context.Context, arg1 *registration.JoinToken, arg2 ...grpc.CallOption) (*registration.JoinToken, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateJoinToken", varargs...)
	ret0, _ := ret[0].(*registration.JoinToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJoinToken indicates an expected call of CreateJoinToken
func (mr *MockRegistrationClientMockRecorder) CreateJoinToken(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJoinToken", reflect.TypeOf((*MockRegistrationClient)(nil).CreateJoinToken), varargs...)
}

// DeleteEntry mocks base method
func (m *MockRegistrationClient) DeleteEntry(arg0 context.Context, arg1 *registration.RegistrationEntryID, arg2 ...grpc.CallOption) (*common.RegistrationEntry, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteEntry", varargs...)
	ret0, _ := ret[0].(*common.RegistrationEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteEntry indicates an expected call of DeleteEntry
func (mr *MockRegistrationClientMockRecorder) DeleteEntry(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEntry", reflect.TypeOf((*MockRegistrationClient)(nil).DeleteEntry), varargs...)
}

// DeleteFederatedBundle mocks base method
func (m *MockRegistrationClient) DeleteFederatedBundle(arg0 context.Context, arg1 *registration.DeleteFederatedBundleRequest, arg2 ...grpc.CallOption) (*common.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteFederatedBundle", varargs...)
	ret0, _ := ret[0].(*common.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFederatedBundle indicates an expected call of DeleteFederatedBundle
func (mr *MockRegistrationClientMockRecorder) DeleteFederatedBundle(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFederatedBundle", reflect.TypeOf((*MockRegistrationClient)(nil).DeleteFederatedBundle), varargs...)
}

// EvictAgent mocks base method
func (m *MockRegistrationClient) EvictAgent(arg0 context.Context, arg1 *registration.EvictAgentRequest, arg2 ...grpc.CallOption) (*registration.EvictAgentResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EvictAgent", varargs...)
	ret0, _ := ret[0].(*registration.EvictAgentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EvictAgent indicates an expected call of EvictAgent
func (mr *MockRegistrationClientMockRecorder) EvictAgent(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvictAgent", reflect.TypeOf((*MockRegistrationClient)(nil).EvictAgent), varargs...)
}

// FetchBundle mocks base method
func (m *MockRegistrationClient) FetchBundle(arg0 context.Context, arg1 *common.Empty, arg2 ...grpc.CallOption) (*registration.Bundle, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FetchBundle", varargs...)
	ret0, _ := ret[0].(*registration.Bundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBundle indicates an expected call of FetchBundle
func (mr *MockRegistrationClientMockRecorder) FetchBundle(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBundle", reflect.TypeOf((*MockRegistrationClient)(nil).FetchBundle), varargs...)
}

// FetchEntries mocks base method
func (m *MockRegistrationClient) FetchEntries(arg0 context.Context, arg1 *common.Empty, arg2 ...grpc.CallOption) (*common.RegistrationEntries, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FetchEntries", varargs...)
	ret0, _ := ret[0].(*common.RegistrationEntries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchEntries indicates an expected call of FetchEntries
func (mr *MockRegistrationClientMockRecorder) FetchEntries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchEntries", reflect.TypeOf((*MockRegistrationClient)(nil).FetchEntries), varargs...)
}

// FetchEntry mocks base method
func (m *MockRegistrationClient) FetchEntry(arg0 context.Context, arg1 *registration.RegistrationEntryID, arg2 ...grpc.CallOption) (*common.RegistrationEntry, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FetchEntry", varargs...)
	ret0, _ := ret[0].(*common.RegistrationEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchEntry indicates an expected call of FetchEntry
func (mr *MockRegistrationClientMockRecorder) FetchEntry(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchEntry", reflect.TypeOf((*MockRegistrationClient)(nil).FetchEntry), varargs...)
}

// FetchFederatedBundle mocks base method
func (m *MockRegistrationClient) FetchFederatedBundle(arg0 context.Context, arg1 *registration.FederatedBundleID, arg2 ...grpc.CallOption) (*registration.FederatedBundle, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FetchFederatedBundle", varargs...)
	ret0, _ := ret[0].(*registration.FederatedBundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchFederatedBundle indicates an expected call of FetchFederatedBundle
func (mr *MockRegistrationClientMockRecorder) FetchFederatedBundle(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchFederatedBundle", reflect.TypeOf((*MockRegistrationClient)(nil).FetchFederatedBundle), varargs...)
}

// ListAgents mocks base method
func (m *MockRegistrationClient) ListAgents(arg0 context.Context, arg1 *registration.ListAgentsRequest, arg2 ...grpc.CallOption) (*registration.ListAgentsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAgents", varargs...)
	ret0, _ := ret[0].(*registration.ListAgentsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAgents indicates an expected call of ListAgents
func (mr *MockRegistrationClientMockRecorder) ListAgents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAgents", reflect.TypeOf((*MockRegistrationClient)(nil).ListAgents), varargs...)
}

// ListByParentID mocks base method
func (m *MockRegistrationClient) ListByParentID(arg0 context.Context, arg1 *registration.ParentID, arg2 ...grpc.CallOption) (*common.RegistrationEntries, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListByParentID", varargs...)
	ret0, _ := ret[0].(*common.RegistrationEntries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByParentID indicates an expected call of ListByParentID
func (mr *MockRegistrationClientMockRecorder) ListByParentID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByParentID", reflect.TypeOf((*MockRegistrationClient)(nil).ListByParentID), varargs...)
}

// ListBySelector mocks base method
func (m *MockRegistrationClient) ListBySelector(arg0 context.Context, arg1 *common.Selector, arg2 ...grpc.CallOption) (*common.RegistrationEntries, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBySelector", varargs...)
	ret0, _ := ret[0].(*common.RegistrationEntries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBySelector indicates an expected call of ListBySelector
func (mr *MockRegistrationClientMockRecorder) ListBySelector(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySelector", reflect.TypeOf((*MockRegistrationClient)(nil).ListBySelector), varargs...)
}

// ListBySpiffeID mocks base method
func (m *MockRegistrationClient) ListBySpiffeID(arg0 context.Context, arg1 *registration.SpiffeID, arg2 ...grpc.CallOption) (*common.RegistrationEntries, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBySpiffeID", varargs...)
	ret0, _ := ret[0].(*common.RegistrationEntries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBySpiffeID indicates an expected call of ListBySpiffeID
func (mr *MockRegistrationClientMockRecorder) ListBySpiffeID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySpiffeID", reflect.TypeOf((*MockRegistrationClient)(nil).ListBySpiffeID), varargs...)
}

// ListFederatedBundles mocks base method
func (m *MockRegistrationClient) ListFederatedBundles(arg0 context.Context, arg1 *common.Empty, arg2 ...grpc.CallOption) (registration.Registration_ListFederatedBundlesClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFederatedBundles", varargs...)
	ret0, _ := ret[0].(registration.Registration_ListFederatedBundlesClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFederatedBundles indicates an expected call of ListFederatedBundles
func (mr *MockRegistrationClientMockRecorder) ListFederatedBundles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFederatedBundles", reflect.TypeOf((*MockRegistrationClient)(nil).ListFederatedBundles), varargs...)
}

// UpdateEntry mocks base method
func (m *MockRegistrationClient) UpdateEntry(arg0 context.Context, arg1 *registration.UpdateEntryRequest, arg2 ...grpc.CallOption) (*common.RegistrationEntry, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateEntry", varargs...)
	ret0, _ := ret[0].(*common.RegistrationEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEntry indicates an expected call of UpdateEntry
func (mr *MockRegistrationClientMockRecorder) UpdateEntry(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEntry", reflect.TypeOf((*MockRegistrationClient)(nil).UpdateEntry), varargs...)
}

// UpdateFederatedBundle mocks base method
func (m *MockRegistrationClient) UpdateFederatedBundle(arg0 context.Context, arg1 *registration.FederatedBundle, arg2 ...grpc.CallOption) (*common.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateFederatedBundle", varargs...)
	ret0, _ := ret[0].(*common.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFederatedBundle indicates an expected call of UpdateFederatedBundle
func (mr *MockRegistrationClientMockRecorder) UpdateFederatedBundle(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFederatedBundle", reflect.TypeOf((*MockRegistrationClient)(nil).UpdateFederatedBundle), varargs...)
}

// MockRegistrationServer is a mock of RegistrationServer interface
type MockRegistrationServer struct {
	ctrl     *gomock.Controller
	recorder *MockRegistrationServerMockRecorder
}

// MockRegistrationServerMockRecorder is the mock recorder for MockRegistrationServer
type MockRegistrationServerMockRecorder struct {
	mock *MockRegistrationServer
}

// NewMockRegistrationServer creates a new mock instance
func NewMockRegistrationServer(ctrl *gomock.Controller) *MockRegistrationServer {
	mock := &MockRegistrationServer{ctrl: ctrl}
	mock.recorder = &MockRegistrationServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegistrationServer) EXPECT() *MockRegistrationServerMockRecorder {
	return m.recorder
}

// CreateEntry mocks base method
func (m *MockRegistrationServer) CreateEntry(arg0 context.Context, arg1 *common.RegistrationEntry) (*registration.RegistrationEntryID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEntry", arg0, arg1)
	ret0, _ := ret[0].(*registration.RegistrationEntryID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEntry indicates an expected call of CreateEntry
func (mr *MockRegistrationServerMockRecorder) CreateEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEntry", reflect.TypeOf((*MockRegistrationServer)(nil).CreateEntry), arg0, arg1)
}

// CreateFederatedBundle mocks base method
func (m *MockRegistrationServer) CreateFederatedBundle(arg0 context.Context, arg1 *registration.FederatedBundle) (*common.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFederatedBundle", arg0, arg1)
	ret0, _ := ret[0].(*common.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFederatedBundle indicates an expected call of CreateFederatedBundle
func (mr *MockRegistrationServerMockRecorder) CreateFederatedBundle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFederatedBundle", reflect.TypeOf((*MockRegistrationServer)(nil).CreateFederatedBundle), arg0, arg1)
}

// CreateJoinToken mocks base method
func (m *MockRegistrationServer) CreateJoinToken(arg0 context.Context, arg1 *registration.JoinToken) (*registration.JoinToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJoinToken", arg0, arg1)
	ret0, _ := ret[0].(*registration.JoinToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJoinToken indicates an expected call of CreateJoinToken
func (mr *MockRegistrationServerMockRecorder) CreateJoinToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJoinToken", reflect.TypeOf((*MockRegistrationServer)(nil).CreateJoinToken), arg0, arg1)
}

// DeleteEntry mocks base method
func (m *MockRegistrationServer) DeleteEntry(arg0 context.Context, arg1 *registration.RegistrationEntryID) (*common.RegistrationEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEntry", arg0, arg1)
	ret0, _ := ret[0].(*common.RegistrationEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteEntry indicates an expected call of DeleteEntry
func (mr *MockRegistrationServerMockRecorder) DeleteEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEntry", reflect.TypeOf((*MockRegistrationServer)(nil).DeleteEntry), arg0, arg1)
}

// DeleteFederatedBundle mocks base method
func (m *MockRegistrationServer) DeleteFederatedBundle(arg0 context.Context, arg1 *registration.DeleteFederatedBundleRequest) (*common.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFederatedBundle", arg0, arg1)
	ret0, _ := ret[0].(*common.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFederatedBundle indicates an expected call of DeleteFederatedBundle
func (mr *MockRegistrationServerMockRecorder) DeleteFederatedBundle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFederatedBundle", reflect.TypeOf((*MockRegistrationServer)(nil).DeleteFederatedBundle), arg0, arg1)
}

// EvictAgent mocks base method
func (m *MockRegistrationServer) EvictAgent(arg0 context.Context, arg1 *registration.EvictAgentRequest) (*registration.EvictAgentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EvictAgent", arg0, arg1)
	ret0, _ := ret[0].(*registration.EvictAgentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EvictAgent indicates an expected call of EvictAgent
func (mr *MockRegistrationServerMockRecorder) EvictAgent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvictAgent", reflect.TypeOf((*MockRegistrationServer)(nil).EvictAgent), arg0, arg1)
}

// FetchBundle mocks base method
func (m *MockRegistrationServer) FetchBundle(arg0 context.Context, arg1 *common.Empty) (*registration.Bundle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBundle", arg0, arg1)
	ret0, _ := ret[0].(*registration.Bundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBundle indicates an expected call of FetchBundle
func (mr *MockRegistrationServerMockRecorder) FetchBundle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBundle", reflect.TypeOf((*MockRegistrationServer)(nil).FetchBundle), arg0, arg1)
}

// FetchEntries mocks base method
func (m *MockRegistrationServer) FetchEntries(arg0 context.Context, arg1 *common.Empty) (*common.RegistrationEntries, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchEntries", arg0, arg1)
	ret0, _ := ret[0].(*common.RegistrationEntries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchEntries indicates an expected call of FetchEntries
func (mr *MockRegistrationServerMockRecorder) FetchEntries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchEntries", reflect.TypeOf((*MockRegistrationServer)(nil).FetchEntries), arg0, arg1)
}

// FetchEntry mocks base method
func (m *MockRegistrationServer) FetchEntry(arg0 context.Context, arg1 *registration.RegistrationEntryID) (*common.RegistrationEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchEntry", arg0, arg1)
	ret0, _ := ret[0].(*common.RegistrationEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchEntry indicates an expected call of FetchEntry
func (mr *MockRegistrationServerMockRecorder) FetchEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchEntry", reflect.TypeOf((*MockRegistrationServer)(nil).FetchEntry), arg0, arg1)
}

// FetchFederatedBundle mocks base method
func (m *MockRegistrationServer) FetchFederatedBundle(arg0 context.Context, arg1 *registration.FederatedBundleID) (*registration.FederatedBundle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchFederatedBundle", arg0, arg1)
	ret0, _ := ret[0].(*registration.FederatedBundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchFederatedBundle indicates an expected call of FetchFederatedBundle
func (mr *MockRegistrationServerMockRecorder) FetchFederatedBundle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchFederatedBundle", reflect.TypeOf((*MockRegistrationServer)(nil).FetchFederatedBundle), arg0, arg1)
}

// ListAgents mocks base method
func (m *MockRegistrationServer) ListAgents(arg0 context.Context, arg1 *registration.ListAgentsRequest) (*registration.ListAgentsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAgents", arg0, arg1)
	ret0, _ := ret[0].(*registration.ListAgentsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAgents indicates an expected call of ListAgents
func (mr *MockRegistrationServerMockRecorder) ListAgents(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAgents", reflect.TypeOf((*MockRegistrationServer)(nil).ListAgents), arg0, arg1)
}

// ListByParentID mocks base method
func (m *MockRegistrationServer) ListByParentID(arg0 context.Context, arg1 *registration.ParentID) (*common.RegistrationEntries, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByParentID", arg0, arg1)
	ret0, _ := ret[0].(*common.RegistrationEntries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByParentID indicates an expected call of ListByParentID
func (mr *MockRegistrationServerMockRecorder) ListByParentID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByParentID", reflect.TypeOf((*MockRegistrationServer)(nil).ListByParentID), arg0, arg1)
}

// ListBySelector mocks base method
func (m *MockRegistrationServer) ListBySelector(arg0 context.Context, arg1 *common.Selector) (*common.RegistrationEntries, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySelector", arg0, arg1)
	ret0, _ := ret[0].(*common.RegistrationEntries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBySelector indicates an expected call of ListBySelector
func (mr *MockRegistrationServerMockRecorder) ListBySelector(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySelector", reflect.TypeOf((*MockRegistrationServer)(nil).ListBySelector), arg0, arg1)
}

// ListBySpiffeID mocks base method
func (m *MockRegistrationServer) ListBySpiffeID(arg0 context.Context, arg1 *registration.SpiffeID) (*common.RegistrationEntries, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySpiffeID", arg0, arg1)
	ret0, _ := ret[0].(*common.RegistrationEntries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBySpiffeID indicates an expected call of ListBySpiffeID
func (mr *MockRegistrationServerMockRecorder) ListBySpiffeID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySpiffeID", reflect.TypeOf((*MockRegistrationServer)(nil).ListBySpiffeID), arg0, arg1)
}

// ListFederatedBundles mocks base method
func (m *MockRegistrationServer) ListFederatedBundles(arg0 *common.Empty, arg1 registration.Registration_ListFederatedBundlesServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFederatedBundles", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListFederatedBundles indicates an expected call of ListFederatedBundles
func (mr *MockRegistrationServerMockRecorder) ListFederatedBundles(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFederatedBundles", reflect.TypeOf((*MockRegistrationServer)(nil).ListFederatedBundles), arg0, arg1)
}

// UpdateEntry mocks base method
func (m *MockRegistrationServer) UpdateEntry(arg0 context.Context, arg1 *registration.UpdateEntryRequest) (*common.RegistrationEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEntry", arg0, arg1)
	ret0, _ := ret[0].(*common.RegistrationEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEntry indicates an expected call of UpdateEntry
func (mr *MockRegistrationServerMockRecorder) UpdateEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEntry", reflect.TypeOf((*MockRegistrationServer)(nil).UpdateEntry), arg0, arg1)
}

// UpdateFederatedBundle mocks base method
func (m *MockRegistrationServer) UpdateFederatedBundle(arg0 context.Context, arg1 *registration.FederatedBundle) (*common.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFederatedBundle", arg0, arg1)
	ret0, _ := ret[0].(*common.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFederatedBundle indicates an expected call of UpdateFederatedBundle
func (mr *MockRegistrationServerMockRecorder) UpdateFederatedBundle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFederatedBundle", reflect.TypeOf((*MockRegistrationServer)(nil).UpdateFederatedBundle), arg0, arg1)
}
