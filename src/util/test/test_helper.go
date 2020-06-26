package test

import "testing"

// 空接口
type any interface{}

// TestRunner 接口
type TestRunner interface {
	Test(t testing.T)
}

// 是否相等接口
type IEqual interface {
	Equal(other interface{}) bool
}
