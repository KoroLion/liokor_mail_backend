// Code generated by MockGen. DO NOT EDIT.
// Source: liokor_mail/internal/pkg/mail (interfaces: MailUseCase)

// Package mocks is a generated GoMock package.
package mocks

import (
	mail "liokor_mail/internal/pkg/mail"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMailUseCase is a mock of MailUseCase interface.
type MockMailUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockMailUseCaseMockRecorder
}

// MockMailUseCaseMockRecorder is the mock recorder for MockMailUseCase.
type MockMailUseCaseMockRecorder struct {
	mock *MockMailUseCase
}

// NewMockMailUseCase creates a new mock instance.
func NewMockMailUseCase(ctrl *gomock.Controller) *MockMailUseCase {
	mock := &MockMailUseCase{ctrl: ctrl}
	mock.recorder = &MockMailUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMailUseCase) EXPECT() *MockMailUseCaseMockRecorder {
	return m.recorder
}

// CreateFolder mocks base method.
func (m *MockMailUseCase) CreateFolder(arg0 int, arg1 string) (mail.Folder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFolder", arg0, arg1)
	ret0, _ := ret[0].(mail.Folder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFolder indicates an expected call of CreateFolder.
func (mr *MockMailUseCaseMockRecorder) CreateFolder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFolder", reflect.TypeOf((*MockMailUseCase)(nil).CreateFolder), arg0, arg1)
}

// GetDialogues mocks base method.
func (m *MockMailUseCase) GetDialogues(arg0 string, arg1 int, arg2 string, arg3 int) ([]mail.Dialogue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDialogues", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]mail.Dialogue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDialogues indicates an expected call of GetDialogues.
func (mr *MockMailUseCaseMockRecorder) GetDialogues(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDialogues", reflect.TypeOf((*MockMailUseCase)(nil).GetDialogues), arg0, arg1, arg2, arg3)
}

// GetEmails mocks base method.
func (m *MockMailUseCase) GetEmails(arg0, arg1 string, arg2, arg3 int) ([]mail.DialogueEmail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmails", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]mail.DialogueEmail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmails indicates an expected call of GetEmails.
func (mr *MockMailUseCaseMockRecorder) GetEmails(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmails", reflect.TypeOf((*MockMailUseCase)(nil).GetEmails), arg0, arg1, arg2, arg3)
}

// GetFolders mocks base method.
func (m *MockMailUseCase) GetFolders(arg0 int) ([]mail.Folder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFolders", arg0)
	ret0, _ := ret[0].([]mail.Folder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFolders indicates an expected call of GetFolders.
func (mr *MockMailUseCaseMockRecorder) GetFolders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFolders", reflect.TypeOf((*MockMailUseCase)(nil).GetFolders), arg0)
}

// SendEmail mocks base method.
func (m *MockMailUseCase) SendEmail(arg0 mail.Mail) (mail.Mail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendEmail", arg0)
	ret0, _ := ret[0].(mail.Mail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendEmail indicates an expected call of SendEmail.
func (mr *MockMailUseCaseMockRecorder) SendEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEmail", reflect.TypeOf((*MockMailUseCase)(nil).SendEmail), arg0)
}

// UpdateFolder mocks base method.
func (m *MockMailUseCase) UpdateFolder(arg0 string, arg1, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFolder", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFolder indicates an expected call of UpdateFolder.
func (mr *MockMailUseCaseMockRecorder) UpdateFolder(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFolder", reflect.TypeOf((*MockMailUseCase)(nil).UpdateFolder), arg0, arg1, arg2)
}
