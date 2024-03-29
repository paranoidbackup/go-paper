// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// EncrypterInput is an autogenerated mock type for the EncrypterInput type
type EncrypterInput struct {
	mock.Mock
}

type EncrypterInput_Expecter struct {
	mock *mock.Mock
}

func (_m *EncrypterInput) EXPECT() *EncrypterInput_Expecter {
	return &EncrypterInput_Expecter{mock: &_m.Mock}
}

// Data provides a mock function with given fields:
func (_m *EncrypterInput) Data() []byte {
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

// EncrypterInput_Data_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Data'
type EncrypterInput_Data_Call struct {
	*mock.Call
}

// Data is a helper method to define mock.On call
func (_e *EncrypterInput_Expecter) Data() *EncrypterInput_Data_Call {
	return &EncrypterInput_Data_Call{Call: _e.mock.On("Data")}
}

func (_c *EncrypterInput_Data_Call) Run(run func()) *EncrypterInput_Data_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EncrypterInput_Data_Call) Return(_a0 []byte) *EncrypterInput_Data_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EncrypterInput_Data_Call) RunAndReturn(run func() []byte) *EncrypterInput_Data_Call {
	_c.Call.Return(run)
	return _c
}

// ProjectID provides a mock function with given fields:
func (_m *EncrypterInput) ProjectID() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// EncrypterInput_DocID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ID'
type EncrypterInput_DocID_Call struct {
	*mock.Call
}

// DocID is a helper method to define mock.On call
func (_e *EncrypterInput_Expecter) DocID() *EncrypterInput_DocID_Call {
	return &EncrypterInput_DocID_Call{Call: _e.mock.On("ID")}
}

func (_c *EncrypterInput_DocID_Call) Run(run func()) *EncrypterInput_DocID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EncrypterInput_DocID_Call) Return(_a0 string) *EncrypterInput_DocID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EncrypterInput_DocID_Call) RunAndReturn(run func() string) *EncrypterInput_DocID_Call {
	_c.Call.Return(run)
	return _c
}

// KeyCount provides a mock function with given fields:
func (_m *EncrypterInput) KeyCount() int {
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

// EncrypterInput_KeyCount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'KeyCount'
type EncrypterInput_KeyCount_Call struct {
	*mock.Call
}

// KeyCount is a helper method to define mock.On call
func (_e *EncrypterInput_Expecter) KeyCount() *EncrypterInput_KeyCount_Call {
	return &EncrypterInput_KeyCount_Call{Call: _e.mock.On("KeyCount")}
}

func (_c *EncrypterInput_KeyCount_Call) Run(run func()) *EncrypterInput_KeyCount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EncrypterInput_KeyCount_Call) Return(_a0 int) *EncrypterInput_KeyCount_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EncrypterInput_KeyCount_Call) RunAndReturn(run func() int) *EncrypterInput_KeyCount_Call {
	_c.Call.Return(run)
	return _c
}

// PublicKeys provides a mock function with given fields:
func (_m *EncrypterInput) PublicKeys() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Public")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// EncrypterInput_PublicKeys_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Public'
type EncrypterInput_PublicKeys_Call struct {
	*mock.Call
}

// PublicKeys is a helper method to define mock.On call
func (_e *EncrypterInput_Expecter) PublicKeys() *EncrypterInput_PublicKeys_Call {
	return &EncrypterInput_PublicKeys_Call{Call: _e.mock.On("Public")}
}

func (_c *EncrypterInput_PublicKeys_Call) Run(run func()) *EncrypterInput_PublicKeys_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EncrypterInput_PublicKeys_Call) Return(_a0 []string) *EncrypterInput_PublicKeys_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EncrypterInput_PublicKeys_Call) RunAndReturn(run func() []string) *EncrypterInput_PublicKeys_Call {
	_c.Call.Return(run)
	return _c
}

// NewEncrypterInput creates a new instance of EncrypterInput. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEncrypterInput(t interface {
	mock.TestingT
	Cleanup(func())
}) *EncrypterInput {
	mock := &EncrypterInput{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
