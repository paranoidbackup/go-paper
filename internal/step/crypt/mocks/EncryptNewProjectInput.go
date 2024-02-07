// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// EncryptNewProjectInput is an autogenerated mock type for the EncryptNewProjectInput type
type EncryptNewProjectInput struct {
	mock.Mock
}

type EncryptNewProjectInput_Expecter struct {
	mock *mock.Mock
}

func (_m *EncryptNewProjectInput) EXPECT() *EncryptNewProjectInput_Expecter {
	return &EncryptNewProjectInput_Expecter{mock: &_m.Mock}
}

// Data provides a mock function with given fields:
func (_m *EncryptNewProjectInput) Data() []byte {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Data")
	}

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// EncryptNewProjectInput_Data_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Data'
type EncryptNewProjectInput_Data_Call struct {
	*mock.Call
}

// Data is a helper method to define mock.On call
func (_e *EncryptNewProjectInput_Expecter) Data() *EncryptNewProjectInput_Data_Call {
	return &EncryptNewProjectInput_Data_Call{Call: _e.mock.On("Data")}
}

func (_c *EncryptNewProjectInput_Data_Call) Run(run func()) *EncryptNewProjectInput_Data_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EncryptNewProjectInput_Data_Call) Return(_a0 []byte) *EncryptNewProjectInput_Data_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EncryptNewProjectInput_Data_Call) RunAndReturn(run func() []byte) *EncryptNewProjectInput_Data_Call {
	_c.Call.Return(run)
	return _c
}

// KeyCount provides a mock function with given fields:
func (_m *EncryptNewProjectInput) KeyCount() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for KeyCount")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// EncryptNewProjectInput_KeyCount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'KeyCount'
type EncryptNewProjectInput_KeyCount_Call struct {
	*mock.Call
}

// KeyCount is a helper method to define mock.On call
func (_e *EncryptNewProjectInput_Expecter) KeyCount() *EncryptNewProjectInput_KeyCount_Call {
	return &EncryptNewProjectInput_KeyCount_Call{Call: _e.mock.On("KeyCount")}
}

func (_c *EncryptNewProjectInput_KeyCount_Call) Run(run func()) *EncryptNewProjectInput_KeyCount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EncryptNewProjectInput_KeyCount_Call) Return(_a0 int) *EncryptNewProjectInput_KeyCount_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EncryptNewProjectInput_KeyCount_Call) RunAndReturn(run func() int) *EncryptNewProjectInput_KeyCount_Call {
	_c.Call.Return(run)
	return _c
}

// NewEncryptNewProjectInput creates a new instance of EncryptNewProjectInput. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEncryptNewProjectInput(t interface {
	mock.TestingT
	Cleanup(func())
}) *EncryptNewProjectInput {
	mock := &EncryptNewProjectInput{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}