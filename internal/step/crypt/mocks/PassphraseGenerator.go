// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PassphraseGenerator is an autogenerated mock type for the PassphraseGenerator type
type PassphraseGenerator struct {
	mock.Mock
}

type PassphraseGenerator_Expecter struct {
	mock *mock.Mock
}

func (_m *PassphraseGenerator) EXPECT() *PassphraseGenerator_Expecter {
	return &PassphraseGenerator_Expecter{mock: &_m.Mock}
}

// GeneratePassphrase provides a mock function with given fields:
func (_m *PassphraseGenerator) GeneratePassphrase() (string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GeneratePassphrase")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PassphraseGenerator_GeneratePassphrase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GeneratePassphrase'
type PassphraseGenerator_GeneratePassphrase_Call struct {
	*mock.Call
}

// GeneratePassphrase is a helper method to define mock.On call
func (_e *PassphraseGenerator_Expecter) GeneratePassphrase() *PassphraseGenerator_GeneratePassphrase_Call {
	return &PassphraseGenerator_GeneratePassphrase_Call{Call: _e.mock.On("GeneratePassphrase")}
}

func (_c *PassphraseGenerator_GeneratePassphrase_Call) Run(run func()) *PassphraseGenerator_GeneratePassphrase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PassphraseGenerator_GeneratePassphrase_Call) Return(_a0 string, _a1 error) *PassphraseGenerator_GeneratePassphrase_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PassphraseGenerator_GeneratePassphrase_Call) RunAndReturn(run func() (string, error)) *PassphraseGenerator_GeneratePassphrase_Call {
	_c.Call.Return(run)
	return _c
}

// NewPassphraseGenerator creates a new instance of PassphraseGenerator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPassphraseGenerator(t interface {
	mock.TestingT
	Cleanup(func())
}) *PassphraseGenerator {
	mock := &PassphraseGenerator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
