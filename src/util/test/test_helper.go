package test

// 空接口
type any interface{}

// 是否相等接口
type IEqual interface {
	Equal(other interface{}) bool
}
