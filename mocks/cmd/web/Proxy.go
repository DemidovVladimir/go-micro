// Code generated by mockery 2.9.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
)

// Proxy is an autogenerated mock type for the Proxy type
type Proxy struct {
	mock.Mock
}

// Passthrough provides a mock function with given fields: s
func (_m *Proxy) Passthrough(s string) string {
	ret := _m.Called(s)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
