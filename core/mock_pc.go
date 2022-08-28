// Code generated by MockGen. DO NOT EDIT.
// Source: core/pc.go

// Package core is a generated GoMock package.
package core

import (
	os "os"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPC is a mock of PC interface.
type MockPC struct {
	ctrl     *gomock.Controller
	recorder *MockPCMockRecorder
}

// MockPCMockRecorder is the mock recorder for MockPC.
type MockPCMockRecorder struct {
	mock *MockPC
}

// NewMockPC creates a new mock instance.
func NewMockPC(ctrl *gomock.Controller) *MockPC {
	mock := &MockPC{ctrl: ctrl}
	mock.recorder = &MockPCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPC) EXPECT() *MockPCMockRecorder {
	return m.recorder
}

// Args mocks base method.
func (m *MockPC) Args() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Args")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Args indicates an expected call of Args.
func (mr *MockPCMockRecorder) Args() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Args", reflect.TypeOf((*MockPC)(nil).Args))
}

// ExecInteractive mocks base method.
func (m *MockPC) ExecInteractive(command, env []string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecInteractive", command, env)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecInteractive indicates an expected call of ExecInteractive.
func (mr *MockPCMockRecorder) ExecInteractive(command, env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecInteractive", reflect.TypeOf((*MockPC)(nil).ExecInteractive), command, env)
}

// ExecToString mocks base method.
func (m *MockPC) ExecToString(command, env []string) (int, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecToString", command, env)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ExecToString indicates an expected call of ExecToString.
func (mr *MockPCMockRecorder) ExecToString(command, env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecToString", reflect.TypeOf((*MockPC)(nil).ExecToString), command, env)
}

// Exit mocks base method.
func (m *MockPC) Exit(code int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Exit", code)
}

// Exit indicates an expected call of Exit.
func (mr *MockPCMockRecorder) Exit(code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exit", reflect.TypeOf((*MockPC)(nil).Exit), code)
}

// FileExists mocks base method.
func (m *MockPC) FileExists(filepath string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileExists", filepath)
	ret0, _ := ret[0].(bool)
	return ret0
}

// FileExists indicates an expected call of FileExists.
func (mr *MockPCMockRecorder) FileExists(filepath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileExists", reflect.TypeOf((*MockPC)(nil).FileExists), filepath)
}

// Getuid mocks base method.
func (m *MockPC) Getuid() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Getuid")
	ret0, _ := ret[0].(int)
	return ret0
}

// Getuid indicates an expected call of Getuid.
func (mr *MockPCMockRecorder) Getuid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Getuid", reflect.TypeOf((*MockPC)(nil).Getuid))
}

// Getwd mocks base method.
func (m *MockPC) Getwd() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Getwd")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Getwd indicates an expected call of Getwd.
func (mr *MockPCMockRecorder) Getwd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Getwd", reflect.TypeOf((*MockPC)(nil).Getwd))
}

// HomeDir mocks base method.
func (m *MockPC) HomeDir() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HomeDir")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HomeDir indicates an expected call of HomeDir.
func (mr *MockPCMockRecorder) HomeDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HomeDir", reflect.TypeOf((*MockPC)(nil).HomeDir))
}

// IsTerminal mocks base method.
func (m *MockPC) IsTerminal() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTerminal")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTerminal indicates an expected call of IsTerminal.
func (mr *MockPCMockRecorder) IsTerminal() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTerminal", reflect.TypeOf((*MockPC)(nil).IsTerminal))
}

// Printf mocks base method.
func (m *MockPC) Printf(format string, a ...interface{}) (int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a_2 := range a {
		varargs = append(varargs, a_2)
	}
	ret := m.ctrl.Call(m, "Printf", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Printf indicates an expected call of Printf.
func (mr *MockPCMockRecorder) Printf(format interface{}, a ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, a...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Printf", reflect.TypeOf((*MockPC)(nil).Printf), varargs...)
}

// Println mocks base method.
func (m *MockPC) Println(a ...interface{}) (int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a_2 := range a {
		varargs = append(varargs, a_2)
	}
	ret := m.ctrl.Call(m, "Println", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Println indicates an expected call of Println.
func (mr *MockPCMockRecorder) Println(a ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Println", reflect.TypeOf((*MockPC)(nil).Println), a...)
}

// ReadDir mocks base method.
func (m *MockPC) ReadDir(dirname string) ([]os.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadDir", dirname)
	ret0, _ := ret[0].([]os.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadDir indicates an expected call of ReadDir.
func (mr *MockPCMockRecorder) ReadDir(dirname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDir", reflect.TypeOf((*MockPC)(nil).ReadDir), dirname)
}

// ReadFile mocks base method.
func (m *MockPC) ReadFile(filename string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFile", filename)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFile indicates an expected call of ReadFile.
func (mr *MockPCMockRecorder) ReadFile(filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFile", reflect.TypeOf((*MockPC)(nil).ReadFile), filename)
}

// WriteFile mocks base method.
func (m *MockPC) WriteFile(filename string, data []byte, perm os.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteFile", filename, data, perm)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteFile indicates an expected call of WriteFile.
func (mr *MockPCMockRecorder) WriteFile(filename, data, perm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteFile", reflect.TypeOf((*MockPC)(nil).WriteFile), filename, data, perm)
}