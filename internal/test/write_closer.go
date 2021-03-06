package test

// Don't edit this file.
// This file is generated by gomic 0.5.2.
// https://github.com/suzuki-shunsuke/gomic

import (
	testing "testing"

	gomic "github.com/suzuki-shunsuke/gomic/gomic"
)

type (
	// WriteCloser is a mock.
	WriteCloser struct {
		t                      *testing.T
		name                   string
		callbackNotImplemented gomic.CallbackNotImplemented
		impl                   struct {
			Write func(p []byte) (n int, err error)
			Close func() error
		}
	}
)

// NewWriteCloser returns WriteCloser .
func NewWriteCloser(t *testing.T, cb gomic.CallbackNotImplemented) *WriteCloser {
	return &WriteCloser{
		t: t, name: "WriteCloser", callbackNotImplemented: cb}
}

// Write is a mock method.
func (mock WriteCloser) Write(p []byte) (n int, err error) {
	methodName := "Write" // nolint: goconst
	if mock.impl.Write != nil {
		return mock.impl.Write(p)
	}
	if mock.callbackNotImplemented != nil {
		mock.callbackNotImplemented(mock.t, mock.name, methodName)
	} else {
		gomic.DefaultCallbackNotImplemented(mock.t, mock.name, methodName)
	}
	return mock.fakeZeroWrite(p)
}

// SetFuncWrite sets a method and returns the mock.
func (mock *WriteCloser) SetFuncWrite(impl func(p []byte) (n int, err error)) *WriteCloser {
	mock.impl.Write = impl
	return mock
}

// SetReturnWrite sets a fake method.
func (mock *WriteCloser) SetReturnWrite(n int, err error) *WriteCloser {
	mock.impl.Write = func([]byte) (int, error) {
		return n, err
	}
	return mock
}

// fakeZeroWrite is a fake method which returns zero values.
func (mock WriteCloser) fakeZeroWrite(p []byte) (n int, err error) {
	return n, err
}

// Close is a mock method.
func (mock WriteCloser) Close() error {
	methodName := "Close" // nolint: goconst
	if mock.impl.Close != nil {
		return mock.impl.Close()
	}
	if mock.callbackNotImplemented != nil {
		mock.callbackNotImplemented(mock.t, mock.name, methodName)
	} else {
		gomic.DefaultCallbackNotImplemented(mock.t, mock.name, methodName)
	}
	return mock.fakeZeroClose()
}

// SetFuncClose sets a method and returns the mock.
func (mock *WriteCloser) SetFuncClose(impl func() error) *WriteCloser {
	mock.impl.Close = impl
	return mock
}

// SetReturnClose sets a fake method.
func (mock *WriteCloser) SetReturnClose(r0 error) *WriteCloser {
	mock.impl.Close = func() error {
		return r0
	}
	return mock
}

// fakeZeroClose is a fake method which returns zero values.
func (mock WriteCloser) fakeZeroClose() error {
	var (
		r0 error
	)
	return r0
}
