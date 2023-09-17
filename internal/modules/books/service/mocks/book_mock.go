// Code generated by MockGen. DO NOT EDIT.
// Source: books_interface.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	service "github.com/Paul1k96/bookstorage/internal/modules/books/service"
	gomock "github.com/golang/mock/gomock"
)

// MockBookServicer is a mock of BookServicer interface.
type MockBookServicer struct {
	ctrl     *gomock.Controller
	recorder *MockBookServicerMockRecorder
}

// MockBookServicerMockRecorder is the mock recorder for MockBookServicer.
type MockBookServicerMockRecorder struct {
	mock *MockBookServicer
}

// NewMockBookServicer creates a new mock instance.
func NewMockBookServicer(ctrl *gomock.Controller) *MockBookServicer {
	mock := &MockBookServicer{ctrl: ctrl}
	mock.recorder = &MockBookServicerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookServicer) EXPECT() *MockBookServicerMockRecorder {
	return m.recorder
}

// GetAuthorsByBook mocks base method.
func (m *MockBookServicer) GetAuthorsByBook(in service.GetAuthorsByBookIn) service.GetAuthorsByBookOut {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorsByBook", in)
	ret0, _ := ret[0].(service.GetAuthorsByBookOut)
	return ret0
}

// GetAuthorsByBook indicates an expected call of GetAuthorsByBook.
func (mr *MockBookServicerMockRecorder) GetAuthorsByBook(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorsByBook", reflect.TypeOf((*MockBookServicer)(nil).GetAuthorsByBook), in)
}

// GetBooksByAuthor mocks base method.
func (m *MockBookServicer) GetBooksByAuthor(in service.GetBooksByAuthorIn) service.GetBooksByAuthorOut {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBooksByAuthor", in)
	ret0, _ := ret[0].(service.GetBooksByAuthorOut)
	return ret0
}

// GetBooksByAuthor indicates an expected call of GetBooksByAuthor.
func (mr *MockBookServicerMockRecorder) GetBooksByAuthor(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBooksByAuthor", reflect.TypeOf((*MockBookServicer)(nil).GetBooksByAuthor), in)
}
