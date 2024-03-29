// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import ast "github.com/biodebox/yaastr/ast"
import mock "github.com/stretchr/testify/mock"

// ParentNode is an autogenerated mock type for the ParentNode type
type ParentNode struct {
	mock.Mock
}

// AppendNode provides a mock function with given fields: _a0
func (_m *ParentNode) AppendNode(_a0 ...ast.Node) {
	_va := make([]interface{}, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// DeleteNode provides a mock function with given fields: index
func (_m *ParentNode) DeleteNode(index int) {
	_m.Called(index)
}

// GetChildren provides a mock function with given fields:
func (_m *ParentNode) GetChildren() []ast.Node {
	ret := _m.Called()

	var r0 []ast.Node
	if rf, ok := ret.Get(0).(func() []ast.Node); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ast.Node)
		}
	}

	return r0
}

// GetParent provides a mock function with given fields:
func (_m *ParentNode) GetParent() ast.ParentNode {
	ret := _m.Called()

	var r0 ast.ParentNode
	if rf, ok := ret.Get(0).(func() ast.ParentNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ast.ParentNode)
		}
	}

	return r0
}

// InsertNode provides a mock function with given fields: _a0, _a1
func (_m *ParentNode) InsertNode(_a0 int, _a1 ...ast.Node) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// PrependNode provides a mock function with given fields: _a0
func (_m *ParentNode) PrependNode(_a0 ...ast.Node) {
	_va := make([]interface{}, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// SetParent provides a mock function with given fields: node
func (_m *ParentNode) SetParent(node ast.ParentNode) {
	_m.Called(node)
}
