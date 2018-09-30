// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

import v1 "github.com/stackrox/rox/generated/api/v1"

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// AddAlert provides a mock function with given fields: alert
func (_m *Store) AddAlert(alert *v1.Alert) error {
	ret := _m.Called(alert)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1.Alert) error); ok {
		r0 = rf(alert)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAlert provides a mock function with given fields: id
func (_m *Store) GetAlert(id string) (*v1.Alert, bool, error) {
	ret := _m.Called(id)

	var r0 *v1.Alert
	if rf, ok := ret.Get(0).(func(string) *v1.Alert); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Alert)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAlerts provides a mock function with given fields:
func (_m *Store) GetAlerts() ([]*v1.Alert, error) {
	ret := _m.Called()

	var r0 []*v1.Alert
	if rf, ok := ret.Get(0).(func() []*v1.Alert); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.Alert)
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

// ListAlert provides a mock function with given fields: id
func (_m *Store) ListAlert(id string) (*v1.ListAlert, bool, error) {
	ret := _m.Called(id)

	var r0 *v1.ListAlert
	if rf, ok := ret.Get(0).(func(string) *v1.ListAlert); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ListAlert)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListAlerts provides a mock function with given fields:
func (_m *Store) ListAlerts() ([]*v1.ListAlert, error) {
	ret := _m.Called()

	var r0 []*v1.ListAlert
	if rf, ok := ret.Get(0).(func() []*v1.ListAlert); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.ListAlert)
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

// UpdateAlert provides a mock function with given fields: alert
func (_m *Store) UpdateAlert(alert *v1.Alert) error {
	ret := _m.Called(alert)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1.Alert) error); ok {
		r0 = rf(alert)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
