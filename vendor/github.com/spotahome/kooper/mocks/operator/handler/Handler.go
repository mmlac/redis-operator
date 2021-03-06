// Code generated by mockery v1.0.0
package handler

import mock "github.com/stretchr/testify/mock"
import runtime "k8s.io/apimachinery/pkg/runtime"

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// Add provides a mock function with given fields: obj
func (_m *Handler) Add(obj runtime.Object) error {
	ret := _m.Called(obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(runtime.Object) error); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: _a0
func (_m *Handler) Delete(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
