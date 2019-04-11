// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spiffe/spire/proto/agent/nodeattestor (interfaces: NodeAttestor,NodeAttestorServer,NodeAttestor_FetchAttestationDataClient)

// Package mock_nodeattestor is a generated GoMock package.
package mock_nodeattestor

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	nodeattestor "github.com/spiffe/spire/proto/agent/nodeattestor"
	plugin "github.com/spiffe/spire/proto/common/plugin"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockNodeAttestor is a mock of NodeAttestor interface
type MockNodeAttestor struct {
	ctrl     *gomock.Controller
	recorder *MockNodeAttestorMockRecorder
}

// MockNodeAttestorMockRecorder is the mock recorder for MockNodeAttestor
type MockNodeAttestorMockRecorder struct {
	mock *MockNodeAttestor
}

// NewMockNodeAttestor creates a new mock instance
func NewMockNodeAttestor(ctrl *gomock.Controller) *MockNodeAttestor {
	mock := &MockNodeAttestor{ctrl: ctrl}
	mock.recorder = &MockNodeAttestorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeAttestor) EXPECT() *MockNodeAttestorMockRecorder {
	return m.recorder
}

// FetchAttestationData mocks base method
func (m *MockNodeAttestor) FetchAttestationData(arg0 context.Context) (nodeattestor.NodeAttestor_FetchAttestationDataClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAttestationData", arg0)
	ret0, _ := ret[0].(nodeattestor.NodeAttestor_FetchAttestationDataClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAttestationData indicates an expected call of FetchAttestationData
func (mr *MockNodeAttestorMockRecorder) FetchAttestationData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAttestationData", reflect.TypeOf((*MockNodeAttestor)(nil).FetchAttestationData), arg0)
}

// MockNodeAttestorServer is a mock of NodeAttestorServer interface
type MockNodeAttestorServer struct {
	ctrl     *gomock.Controller
	recorder *MockNodeAttestorServerMockRecorder
}

// MockNodeAttestorServerMockRecorder is the mock recorder for MockNodeAttestorServer
type MockNodeAttestorServerMockRecorder struct {
	mock *MockNodeAttestorServer
}

// NewMockNodeAttestorServer creates a new mock instance
func NewMockNodeAttestorServer(ctrl *gomock.Controller) *MockNodeAttestorServer {
	mock := &MockNodeAttestorServer{ctrl: ctrl}
	mock.recorder = &MockNodeAttestorServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeAttestorServer) EXPECT() *MockNodeAttestorServerMockRecorder {
	return m.recorder
}

// Configure mocks base method
func (m *MockNodeAttestorServer) Configure(arg0 context.Context, arg1 *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Configure", arg0, arg1)
	ret0, _ := ret[0].(*plugin.ConfigureResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Configure indicates an expected call of Configure
func (mr *MockNodeAttestorServerMockRecorder) Configure(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configure", reflect.TypeOf((*MockNodeAttestorServer)(nil).Configure), arg0, arg1)
}

// FetchAttestationData mocks base method
func (m *MockNodeAttestorServer) FetchAttestationData(arg0 nodeattestor.NodeAttestor_FetchAttestationDataServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAttestationData", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// FetchAttestationData indicates an expected call of FetchAttestationData
func (mr *MockNodeAttestorServerMockRecorder) FetchAttestationData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAttestationData", reflect.TypeOf((*MockNodeAttestorServer)(nil).FetchAttestationData), arg0)
}

// GetPluginInfo mocks base method
func (m *MockNodeAttestorServer) GetPluginInfo(arg0 context.Context, arg1 *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPluginInfo", arg0, arg1)
	ret0, _ := ret[0].(*plugin.GetPluginInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPluginInfo indicates an expected call of GetPluginInfo
func (mr *MockNodeAttestorServerMockRecorder) GetPluginInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPluginInfo", reflect.TypeOf((*MockNodeAttestorServer)(nil).GetPluginInfo), arg0, arg1)
}

// MockNodeAttestor_FetchAttestationDataClient is a mock of NodeAttestor_FetchAttestationDataClient interface
type MockNodeAttestor_FetchAttestationDataClient struct {
	ctrl     *gomock.Controller
	recorder *MockNodeAttestor_FetchAttestationDataClientMockRecorder
}

// MockNodeAttestor_FetchAttestationDataClientMockRecorder is the mock recorder for MockNodeAttestor_FetchAttestationDataClient
type MockNodeAttestor_FetchAttestationDataClientMockRecorder struct {
	mock *MockNodeAttestor_FetchAttestationDataClient
}

// NewMockNodeAttestor_FetchAttestationDataClient creates a new mock instance
func NewMockNodeAttestor_FetchAttestationDataClient(ctrl *gomock.Controller) *MockNodeAttestor_FetchAttestationDataClient {
	mock := &MockNodeAttestor_FetchAttestationDataClient{ctrl: ctrl}
	mock.recorder = &MockNodeAttestor_FetchAttestationDataClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeAttestor_FetchAttestationDataClient) EXPECT() *MockNodeAttestor_FetchAttestationDataClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockNodeAttestor_FetchAttestationDataClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockNodeAttestor_FetchAttestationDataClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockNodeAttestor_FetchAttestationDataClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockNodeAttestor_FetchAttestationDataClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockNodeAttestor_FetchAttestationDataClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockNodeAttestor_FetchAttestationDataClient)(nil).Context))
}

// Header mocks base method
func (m *MockNodeAttestor_FetchAttestationDataClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockNodeAttestor_FetchAttestationDataClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockNodeAttestor_FetchAttestationDataClient)(nil).Header))
}

// Recv mocks base method
func (m *MockNodeAttestor_FetchAttestationDataClient) Recv() (*nodeattestor.FetchAttestationDataResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*nodeattestor.FetchAttestationDataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockNodeAttestor_FetchAttestationDataClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockNodeAttestor_FetchAttestationDataClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockNodeAttestor_FetchAttestationDataClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockNodeAttestor_FetchAttestationDataClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockNodeAttestor_FetchAttestationDataClient)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockNodeAttestor_FetchAttestationDataClient) Send(arg0 *nodeattestor.FetchAttestationDataRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockNodeAttestor_FetchAttestationDataClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockNodeAttestor_FetchAttestationDataClient)(nil).Send), arg0)
}

// SendMsg mocks base method
func (m *MockNodeAttestor_FetchAttestationDataClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockNodeAttestor_FetchAttestationDataClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockNodeAttestor_FetchAttestationDataClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockNodeAttestor_FetchAttestationDataClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockNodeAttestor_FetchAttestationDataClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockNodeAttestor_FetchAttestationDataClient)(nil).Trailer))
}
