// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import soap "github.com/nqmt/go-service/domain/soap"

// ISoapRepo is an autogenerated mock type for the ISoapRepo type
type ISoapRepo struct {
	mock.Mock
}

// GetSoap provides a mock function with given fields:
func (_m *ISoapRepo) GetSoap() ([]*soap.Soap, error) {
	ret := _m.Called()

	var r0 []*soap.Soap
	if rf, ok := ret.Get(0).(func() []*soap.Soap); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*soap.Soap)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
