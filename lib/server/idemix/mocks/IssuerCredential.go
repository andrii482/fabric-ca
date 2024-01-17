/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
// Code generated by mockery v1.0.0

package mocks

import (
	idemix "github.com/IBM/idemix/bccsp/schemes/dlog/crypto"
	mock "github.com/stretchr/testify/mock"
)

// IssuerCredential is an autogenerated mock type for the IssuerCredential type
type IssuerCredential struct {
	mock.Mock
}

// GetIssuerKey provides a mock function with given fields:
func (_m *IssuerCredential) GetIssuerKey() (*idemix.IssuerKey, error) {
	ret := _m.Called()

	var r0 *idemix.IssuerKey
	if rf, ok := ret.Get(0).(func() *idemix.IssuerKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*idemix.IssuerKey)
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

// Load provides a mock function with given fields:
func (_m *IssuerCredential) Load() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIssuerKey provides a mock function with given fields:
func (_m *IssuerCredential) NewIssuerKey() (*idemix.IssuerKey, error) {
	ret := _m.Called()

	var r0 *idemix.IssuerKey
	if rf, ok := ret.Get(0).(func() *idemix.IssuerKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*idemix.IssuerKey)
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

// SetIssuerKey provides a mock function with given fields: key
func (_m *IssuerCredential) SetIssuerKey(key *idemix.IssuerKey) {
	_m.Called(key)
}

// Store provides a mock function with given fields:
func (_m *IssuerCredential) Store() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
