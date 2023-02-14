// Code generated by mockery v2.19.0. DO NOT EDIT.

package mocks

import (
	compare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	mock "github.com/stretchr/testify/mock"

	types "github.com/aws-controllers-k8s/runtime/pkg/types"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AWSResourceDescriptor is an autogenerated mock type for the AWSResourceDescriptor type
type AWSResourceDescriptor struct {
	mock.Mock
}

// Delta provides a mock function with given fields: a, b
func (_m *AWSResourceDescriptor) Delta(a types.AWSResource, b types.AWSResource) *compare.Delta {
	ret := _m.Called(a, b)

	var r0 *compare.Delta
	if rf, ok := ret.Get(0).(func(types.AWSResource, types.AWSResource) *compare.Delta); ok {
		r0 = rf(a, b)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compare.Delta)
		}
	}

	return r0
}

// EmptyRuntimeObject provides a mock function with given fields:
func (_m *AWSResourceDescriptor) EmptyRuntimeObject() client.Object {
	ret := _m.Called()

	var r0 client.Object
	if rf, ok := ret.Get(0).(func() client.Object); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.Object)
		}
	}

	return r0
}

// GroupKind provides a mock function with given fields:
func (_m *AWSResourceDescriptor) GroupKind() *v1.GroupKind {
	ret := _m.Called()

	var r0 *v1.GroupKind
	if rf, ok := ret.Get(0).(func() *v1.GroupKind); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.GroupKind)
		}
	}

	return r0
}

// IsManaged provides a mock function with given fields: _a0
func (_m *AWSResourceDescriptor) IsManaged(_a0 types.AWSResource) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.AWSResource) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MarkAdopted provides a mock function with given fields: _a0
func (_m *AWSResourceDescriptor) MarkAdopted(_a0 types.AWSResource) {
	_m.Called(_a0)
}

// MarkManaged provides a mock function with given fields: _a0
func (_m *AWSResourceDescriptor) MarkManaged(_a0 types.AWSResource) {
	_m.Called(_a0)
}

// MarkUnmanaged provides a mock function with given fields: _a0
func (_m *AWSResourceDescriptor) MarkUnmanaged(_a0 types.AWSResource) {
	_m.Called(_a0)
}

// ResourceFromRuntimeObject provides a mock function with given fields: _a0
func (_m *AWSResourceDescriptor) ResourceFromRuntimeObject(_a0 client.Object) types.AWSResource {
	ret := _m.Called(_a0)

	var r0 types.AWSResource
	if rf, ok := ret.Get(0).(func(client.Object) types.AWSResource); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.AWSResource)
		}
	}

	return r0
}

type mockConstructorTestingTNewAWSResourceDescriptor interface {
	mock.TestingT
	Cleanup(func())
}

// NewAWSResourceDescriptor creates a new instance of AWSResourceDescriptor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAWSResourceDescriptor(t mockConstructorTestingTNewAWSResourceDescriptor) *AWSResourceDescriptor {
	mock := &AWSResourceDescriptor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
