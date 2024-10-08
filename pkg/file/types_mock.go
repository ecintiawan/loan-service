// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/file/types.go

// Package file is a generated GoMock package.
package file

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFile is a mock of File interface.
type MockFile struct {
	ctrl     *gomock.Controller
	recorder *MockFileMockRecorder
}

// MockFileMockRecorder is the mock recorder for MockFile.
type MockFileMockRecorder struct {
	mock *MockFile
}

// NewMockFile creates a new mock instance.
func NewMockFile(ctrl *gomock.Controller) *MockFile {
	mock := &MockFile{ctrl: ctrl}
	mock.recorder = &MockFileMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFile) EXPECT() *MockFileMockRecorder {
	return m.recorder
}

// Write mocks base method.
func (m *MockFile) Write(content []byte, filePath, fileName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", content, filePath, fileName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockFileMockRecorder) Write(content, filePath, fileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockFile)(nil).Write), content, filePath, fileName)
}

// MockPDFGenerator is a mock of PDFGenerator interface.
type MockPDFGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockPDFGeneratorMockRecorder
}

// MockPDFGeneratorMockRecorder is the mock recorder for MockPDFGenerator.
type MockPDFGeneratorMockRecorder struct {
	mock *MockPDFGenerator
}

// NewMockPDFGenerator creates a new mock instance.
func NewMockPDFGenerator(ctrl *gomock.Controller) *MockPDFGenerator {
	mock := &MockPDFGenerator{ctrl: ctrl}
	mock.recorder = &MockPDFGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPDFGenerator) EXPECT() *MockPDFGeneratorMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockPDFGenerator) Generate(content string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", content)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generate indicates an expected call of Generate.
func (mr *MockPDFGeneratorMockRecorder) Generate(content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockPDFGenerator)(nil).Generate), content)
}
