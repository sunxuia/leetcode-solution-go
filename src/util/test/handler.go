package test

import (
	"reflect"
	"strings"
)

// 测试过程中用来表示一个测试对象
type Handler struct {
	caseName string
	config   *assertConfig
	visited  map[visit]bool

	FullExpect interface{}
	FullActual interface{}

	FieldName string
	Expect    interface{}
	Actual    interface{}
	Ev        reflect.Value
	Av        reflect.Value
}

type visit struct {
	a1 uintptr
	a2 uintptr
}

func (h *Handler) Child(childName string, expect, actual interface{}) *Handler {
	res := &Handler{
		caseName:   h.caseName,
		config:     h.config,
		FullExpect: h.FullExpect,
		FullActual: h.FullActual,
	}
	res.ChangeValue(expect, actual)
	if strings.Contains(childName, ".") {
		childName = strings.ReplaceAll(childName, ".", "_")
	}
	if len(h.FieldName) > 0 {
		res.FieldName = h.FieldName + "." + childName
	} else {
		res.FieldName = childName
	}
	return res
}

func (h *Handler) ChangeValue(expect, actual interface{}) {
	h.Expect = expect
	h.Actual = actual
	h.Ev = reflect.ValueOf(expect)
	h.Av = reflect.ValueOf(actual)
}

func (h *Handler) Error(mesage string) {
	err := &AssertError{
		handler: h,
		message: mesage,
	}
	panic(err)
}

func (h *Handler) Assert() {
	if !h.config.shouldChecks.getNodeValue(h.FieldName).(bool) || h.visit() {
		return
	}

	if v := h.config.validators.getNodeValue(h.FieldName); v != nil {
		v.(func(*Handler))(h)
		return
	}

	eTypeName := "nil"
	if h.Ev.IsValid() {
		eTypeName = h.Ev.Type().String()
	}
	if v, ok := h.config.expectTypeValidators[eTypeName]; ok {
		v(h)
		return
	}

	aTypeName := "nil"
	if h.Av.IsValid() {
		aTypeName = h.Av.Type().String()
	}
	if v, ok := h.config.actualTypeValidators[aTypeName]; ok {
		v(h)
		return
	}

	DefaultAssertEqual(h)
}

func (h *Handler) visit() bool {
	if h.Expect == nil || h.Actual == nil ||
		!h.Ev.IsValid() || !h.Av.IsValid() ||
		!h.Ev.CanAddr() || !h.Av.CanAddr() {
		return false
	}
	v := visit{h.Ev.UnsafeAddr(), h.Av.UnsafeAddr()}
	if h.visited[v] {
		return true
	}
	h.visited[v] = true
	return false
}

func (h *Handler) AssertWithError() (err interface{}) {
	defer func() {
		err = recover()
	}()
	h.Assert()
	return
}

func (h *Handler) IsOrdered() bool {
	return h.config.hasOrders.getNodeValue(h.FieldName).(bool)
}
