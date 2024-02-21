// Code generated by MockGen. DO NOT EDIT.
// Source: ./stagemanager.go
//
// Generated by this command:
//
//	mockgen -destination=../mocks/ui/mock_stagemanager.go -source=./stagemanager.go
//

// Package mock_ui is a generated GoMock package.
package mock_ui

import (
	ui "keepassui/internal/ui"
	reflect "reflect"

	fyne "fyne.io/fyne/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockStagerController is a mock of StagerController interface.
type MockStagerController struct {
	ctrl     *gomock.Controller
	recorder *MockStagerControllerMockRecorder
}

// MockStagerControllerMockRecorder is the mock recorder for MockStagerController.
type MockStagerControllerMockRecorder struct {
	mock *MockStagerController
}

// NewMockStagerController creates a new mock instance.
func NewMockStagerController(ctrl *gomock.Controller) *MockStagerController {
	mock := &MockStagerController{ctrl: ctrl}
	mock.recorder = &MockStagerControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStagerController) EXPECT() *MockStagerControllerMockRecorder {
	return m.recorder
}

// GetContainer mocks base method.
func (m *MockStagerController) GetContainer() *fyne.Container {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainer")
	ret0, _ := ret[0].(*fyne.Container)
	return ret0
}

// GetContainer indicates an expected call of GetContainer.
func (mr *MockStagerControllerMockRecorder) GetContainer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainer", reflect.TypeOf((*MockStagerController)(nil).GetContainer))
}

// RegisterStager mocks base method.
func (m *MockStagerController) RegisterStager(stager ui.Stager) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterStager", stager)
}

// RegisterStager indicates an expected call of RegisterStager.
func (mr *MockStagerControllerMockRecorder) RegisterStager(stager any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterStager", reflect.TypeOf((*MockStagerController)(nil).RegisterStager), stager)
}

// TakeOver mocks base method.
func (m *MockStagerController) TakeOver(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TakeOver", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// TakeOver indicates an expected call of TakeOver.
func (mr *MockStagerControllerMockRecorder) TakeOver(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TakeOver", reflect.TypeOf((*MockStagerController)(nil).TakeOver), name)
}

// MockStager is a mock of Stager interface.
type MockStager struct {
	ctrl     *gomock.Controller
	recorder *MockStagerMockRecorder
}

// MockStagerMockRecorder is the mock recorder for MockStager.
type MockStagerMockRecorder struct {
	mock *MockStager
}

// NewMockStager creates a new mock instance.
func NewMockStager(ctrl *gomock.Controller) *MockStager {
	mock := &MockStager{ctrl: ctrl}
	mock.recorder = &MockStagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStager) EXPECT() *MockStagerMockRecorder {
	return m.recorder
}

// ExecuteOnTakeOver mocks base method.
func (m *MockStager) ExecuteOnTakeOver() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ExecuteOnTakeOver")
}

// ExecuteOnTakeOver indicates an expected call of ExecuteOnTakeOver.
func (mr *MockStagerMockRecorder) ExecuteOnTakeOver() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteOnTakeOver", reflect.TypeOf((*MockStager)(nil).ExecuteOnTakeOver))
}

// GetPaintedContainer mocks base method.
func (m *MockStager) GetPaintedContainer() *fyne.Container {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPaintedContainer")
	ret0, _ := ret[0].(*fyne.Container)
	return ret0
}

// GetPaintedContainer indicates an expected call of GetPaintedContainer.
func (mr *MockStagerMockRecorder) GetPaintedContainer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaintedContainer", reflect.TypeOf((*MockStager)(nil).GetPaintedContainer))
}

// GetStageName mocks base method.
func (m *MockStager) GetStageName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStageName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetStageName indicates an expected call of GetStageName.
func (mr *MockStagerMockRecorder) GetStageName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStageName", reflect.TypeOf((*MockStager)(nil).GetStageName))
}
