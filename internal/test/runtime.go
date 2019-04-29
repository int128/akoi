package test

// Don't edit this file.
// This file is generated by gomic 0.5.5-0.
// https://github.com/suzuki-shunsuke/gomic

import (
	testing "testing"

	gomic "github.com/suzuki-shunsuke/gomic/gomic"
)

type (
	// Runtime is a mock.
	Runtime struct {
		t                      *testing.T
		name                   string
		callbackNotImplemented gomic.CallbackNotImplemented
		impl                   struct {
			OS     func() (r0 string)
			Arch   func() (r0 string)
			NumCPU func() (r0 int)
		}
	}
)

// NewRuntime returns Runtime .
func NewRuntime(t *testing.T, cb gomic.CallbackNotImplemented) *Runtime {
	return &Runtime{
		t: t, name: "Runtime", callbackNotImplemented: cb}
}

// OS is a mock method.
func (mock Runtime) OS() (r0 string) {
	methodName := "OS" // nolint: goconst
	if mock.impl.OS != nil {
		return mock.impl.OS()
	}
	if mock.callbackNotImplemented != nil {
		mock.callbackNotImplemented(mock.t, mock.name, methodName)
	} else {
		gomic.DefaultCallbackNotImplemented(mock.t, mock.name, methodName)
	}
	return mock.fakeZeroOS()
}

// SetFuncOS sets a method and returns the mock.
func (mock *Runtime) SetFuncOS(impl func() (r0 string)) *Runtime {
	mock.impl.OS = impl
	return mock
}

// SetReturnOS sets a fake method.
func (mock *Runtime) SetReturnOS(r0 string) *Runtime {
	mock.impl.OS = func() string {
		return r0
	}
	return mock
}

// fakeZeroOS is a fake method which returns zero values.
func (mock Runtime) fakeZeroOS() (r0 string) {
	return r0
}

// Arch is a mock method.
func (mock Runtime) Arch() (r0 string) {
	methodName := "Arch" // nolint: goconst
	if mock.impl.Arch != nil {
		return mock.impl.Arch()
	}
	if mock.callbackNotImplemented != nil {
		mock.callbackNotImplemented(mock.t, mock.name, methodName)
	} else {
		gomic.DefaultCallbackNotImplemented(mock.t, mock.name, methodName)
	}
	return mock.fakeZeroArch()
}

// SetFuncArch sets a method and returns the mock.
func (mock *Runtime) SetFuncArch(impl func() (r0 string)) *Runtime {
	mock.impl.Arch = impl
	return mock
}

// SetReturnArch sets a fake method.
func (mock *Runtime) SetReturnArch(r0 string) *Runtime {
	mock.impl.Arch = func() string {
		return r0
	}
	return mock
}

// fakeZeroArch is a fake method which returns zero values.
func (mock Runtime) fakeZeroArch() (r0 string) {
	return r0
}

// NumCPU is a mock method.
func (mock Runtime) NumCPU() (r0 int) {
	methodName := "NumCPU" // nolint: goconst
	if mock.impl.NumCPU != nil {
		return mock.impl.NumCPU()
	}
	if mock.callbackNotImplemented != nil {
		mock.callbackNotImplemented(mock.t, mock.name, methodName)
	} else {
		gomic.DefaultCallbackNotImplemented(mock.t, mock.name, methodName)
	}
	return mock.fakeZeroNumCPU()
}

// SetFuncNumCPU sets a method and returns the mock.
func (mock *Runtime) SetFuncNumCPU(impl func() (r0 int)) *Runtime {
	mock.impl.NumCPU = impl
	return mock
}

// SetReturnNumCPU sets a fake method.
func (mock *Runtime) SetReturnNumCPU(r0 int) *Runtime {
	mock.impl.NumCPU = func() int {
		return r0
	}
	return mock
}

// fakeZeroNumCPU is a fake method which returns zero values.
func (mock Runtime) fakeZeroNumCPU() (r0 int) {
	return r0
}