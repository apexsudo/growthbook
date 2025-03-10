// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/apexsudo/analytics (interfaces: Client)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_analytics.go -package=mocks -typed github.com/apexsudo/analytics Client
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
	isgomock struct{}
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockClientMockRecorder) Close() *MockClientCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
	return &MockClientCloseCall{Call: call}
}

// MockClientCloseCall wrap *gomock.Call
type MockClientCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientCloseCall) Return(arg0 error) *MockClientCloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientCloseCall) Do(f func() error) *MockClientCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientCloseCall) DoAndReturn(f func() error) *MockClientCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Identify mocks base method.
func (m *MockClient) Identify(userID string, traits map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Identify", userID, traits)
	ret0, _ := ret[0].(error)
	return ret0
}

// Identify indicates an expected call of Identify.
func (mr *MockClientMockRecorder) Identify(userID, traits any) *MockClientIdentifyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identify", reflect.TypeOf((*MockClient)(nil).Identify), userID, traits)
	return &MockClientIdentifyCall{Call: call}
}

// MockClientIdentifyCall wrap *gomock.Call
type MockClientIdentifyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientIdentifyCall) Return(arg0 error) *MockClientIdentifyCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientIdentifyCall) Do(f func(string, map[string]any) error) *MockClientIdentifyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientIdentifyCall) DoAndReturn(f func(string, map[string]any) error) *MockClientIdentifyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Track mocks base method.
func (m *MockClient) Track(userID, event string, properties map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Track", userID, event, properties)
	ret0, _ := ret[0].(error)
	return ret0
}

// Track indicates an expected call of Track.
func (mr *MockClientMockRecorder) Track(userID, event, properties any) *MockClientTrackCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Track", reflect.TypeOf((*MockClient)(nil).Track), userID, event, properties)
	return &MockClientTrackCall{Call: call}
}

// MockClientTrackCall wrap *gomock.Call
type MockClientTrackCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientTrackCall) Return(arg0 error) *MockClientTrackCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientTrackCall) Do(f func(string, string, map[string]any) error) *MockClientTrackCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientTrackCall) DoAndReturn(f func(string, string, map[string]any) error) *MockClientTrackCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
