package test

import (
	"container/list"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strings"
)

// assertHelper 断言工具
type assertHelper struct {
	shouldChecks         *matchTree
	hasOrders            *matchTree
	validators           *matchTree
	expectTypeValidators map[string]func(string, any, any)
	actualTypeValidators map[string]func(string, any, any)
	numberTolerance      float64
	times                int
}

// NewAssertHelper 断言工具构造函数
func NewAssertHelper() *assertHelper {
	res := &assertHelper{
		shouldChecks: newMatchTree("root"),
		hasOrders:    newMatchTree("root"),
		validators:   newMatchTree("root"),
	}
	setNodeValue(res.shouldChecks, "**", true)
	setNodeValue(res.hasOrders, "**", true)
	setNodeValue(res.validators, "**", res.defaultAssertMethod)
	return res
}

// 字段匹配的结构
type matchTree struct {
	key      string
	value    any
	setted   bool
	children map[string]*matchTree
}

func newMatchTree(key string) *matchTree {
	res := new(matchTree)
	res.key = key
	res.children = map[string]*matchTree{}
	return res
}

// AssertError 断言错误
type AssertError struct {
	fieldName string
	expected  any
	actual    any
	message   string
}

func setNodeValue(node *matchTree, fieldName string, value any) {
	names := strings.Split(fieldName, ".")
	for _, name := range names {
		if _, ok := node.children[name]; !ok {
			node.children[name] = newMatchTree(name)
		}
		node = node.children[name]
	}
	node.setted = true
	node.value = value
}

func (eh *assertHelper) Check(fieldName string) *assertHelper {
	setNodeValue(eh.shouldChecks, fieldName, true)
	return eh
}

func (eh *assertHelper) UnCheck(fieldName string) *assertHelper {
	setNodeValue(eh.shouldChecks, fieldName, false)
	return eh
}

func (eh *assertHelper) Order(fieldName string) *assertHelper {
	setNodeValue(eh.hasOrders, fieldName, true)
	return eh
}

func (eh *assertHelper) UnOrder(fieldName string) *assertHelper {
	setNodeValue(eh.hasOrders, fieldName, false)
	return eh
}

func (eh *assertHelper) AddValidator(fieldName string, validateMethod func(string, any, any)) {
	setNodeValue(eh.validators, fieldName, validateMethod)
}

func (eh *assertHelper) AddExpectTypeValidator(typeName string, validateMethod func(string, any, any)) {
	eh.expectTypeValidators[typeName] = validateMethod
}

func (eh *assertHelper) AddActualTypeValidator(typeName string, validateMethod func(string, any, any)) {
	eh.actualTypeValidators[typeName] = validateMethod
}

func (eh *assertHelper) FloatPrecision(tolerance float64) {
	assertMethod := func(fieldName string, expect any, actual any) {
		ev := reflect.ValueOf(expect).Float()
		av := reflect.ValueOf(actual).Float()
		if math.Abs(ev-av) > tolerance {
			panic(AssertError{fieldName, ev, av, "(float equal)"})
		}
	}
	eh.expectTypeValidators["float64"] = assertMethod
	eh.expectTypeValidators["float32"] = assertMethod
	eh.actualTypeValidators["float64"] = assertMethod
	eh.actualTypeValidators["float32"] = assertMethod
}

// 断言方法
func (eh *assertHelper) Assert(expect any, actual any) *assertHelper {
	eh.times++
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(AssertError); ok {
				panic(fmt.Sprintf(`Error while validate (%dth)
    expect: %s,
    actual: %s,
    detail (%s):
        expect: %s, 
        actual: %s, 
        message: %s
`, eh.times, expect, actual, e.fieldName, e.expected, e.actual, e.message))
			} else {
				panic(r)
			}
		}
	}()

	eh.assert("", expect, actual)
	return eh
}

func (eh *assertHelper) assert(fieldName string, expect any, actual any) {
	if !getNodeValue(eh.shouldChecks, fieldName).(bool) {
		return
	}
	assertMethod := getNodeValue(eh.validators, fieldName).(func(string, any, any))
	assertMethod(fieldName, expect, actual)
}

func getNodeValue(root *matchTree, fieldName string) any {
	queue := list.New()
	queue.PushBack(root)
	names := strings.Split(fieldName, ".")
	for _, name := range names {
		for i := queue.Len(); i > 0; i-- {
			val := queue.Remove(queue.Front())
			if val != nil {
				node := val.(*matchTree)
				queue.PushBack(node.children[name])
				queue.PushBack(node.children["*"])
				queue.PushBack(node.children["**"])
				if node.key == "**" {
					queue.PushBack(node)
				}
			}
		}
	}
	for ele := queue.Front(); ele != nil; ele = ele.Next() {
		node := ele.Value.(*matchTree)
		if node != nil && node.setted {
			return node.value
		}
	}
	panic("Error in AssertHelper: no default value. Did you over write it?")
}

func (eh *assertHelper) defaultAssertMethod(fieldName string, expect any, actual any) {
	// nil check
	if expect == nil || actual == nil {
		if expect != nil || actual != nil {
			panic(AssertError{fieldName, expect, actual, "(nil check)"})
		}
		return
	}

	// IEqual
	valE, okE := expect.(IEqual)
	valA, okA := actual.(IEqual)
	if okE && okA {
		if !valE.Equal(valA) {
			panic(AssertError{fieldName, expect, actual, "(validate via Equal method)"})
		}
		return
	}

	expectType := reflect.TypeOf(expect).String()
	actualType := reflect.TypeOf(actual).String()

	if validatorMethod, ok := eh.expectTypeValidators[expectType]; ok {
		validatorMethod(fieldName, expect, actual)
		return
	}
	if validatorMethod, ok := eh.actualTypeValidators[actualType]; ok {
		validatorMethod(fieldName, expect, actual)
		return
	}

	// array/ slice
	if em := arrayPattern.FindAllStringSubmatch(expectType, -1); em != nil {
		eh.assertArray(fieldName, expect, actual)
		return
	}

	// map
	if em := mapPattern.FindAllStringSubmatch(expectType, -1); em != nil {
		eh.assertMap(fieldName, expect, actual)
		return
	}

	// comparable
	if comp, fail := compare(expect, actual); !fail {
		if !comp {
			panic(AssertError{fieldName, expect, actual, ""})
		}
		return
	}

	// default
	panic(fmt.Sprint("Unknown type: expect ", expectType, ", actual ", actualType))
}

func compare(expect any, actual any) (res bool, fail bool) {
	defer func() {
		if r := recover(); r != nil {
			fail = true
		}
	}()
	res = expect == actual
	return
}

var arrayPattern = regexp.MustCompile("^\\[([^]+]*)](.+)$")

var mapPattern = regexp.MustCompile("^map\\[([^]+]+)](.+)$")

func (eh *assertHelper) assertArray(fieldName string, expect any, actual any) {
	actualType := reflect.TypeOf(actual).String()
	am := arrayPattern.FindAllStringSubmatch(actualType, -1)
	if am == nil {
		expectType := reflect.TypeOf(expect).String()
		panic(AssertError{fieldName, expect, actual,
			fmt.Sprintf("Expect array type %s while actual type is %s.", expectType, actualType)})
	}

	ve, va := reflect.ValueOf(expect), reflect.ValueOf(actual)
	length, lenA := ve.Len(), va.Len()
	if length != lenA {
		panic(AssertError{fieldName, expect, actual,
			fmt.Sprint("Expect array length ", length, " actual length ", lenA)})
	}

	if hasOrder := getNodeValue(eh.hasOrders, fieldName); hasOrder.(bool) {
		for i := 0; i < length; i++ {
			valE, valA := ve.Index(i).Interface(), va.Index(i).Interface()
			eh.assert(childFieldName(fieldName, i), valE, valA)
		}
	} else {
		match := make([]bool, length, length)
		for i := 0; i < length; i++ {
			hasMatch := false
			for j := 0; j < length; j++ {
				if match[j] {
					continue
				}
				valE, valA := ve.Index(i).Interface(), va.Index(j).Interface()
				err := eh.assertWithError(childFieldName(fieldName, i), valE, valA)
				if err == nil {
					match[j] = true
					hasMatch = true
					break
				}
			}
			if !hasMatch {
				panic(AssertError{fieldName, expect, actual,
					fmt.Sprint("Expect element at ", i, " cannot be found.")})
			}
		}
	}
}

func childFieldName(fieldName string, child any) string {
	str := fmt.Sprint(child)
	str = strings.ReplaceAll(str, ".", "_")
	if len(fieldName) > 0 {
		return fieldName + "." + str
	}
	return str
}

func (eh *assertHelper) assertWithError(fieldName string, expect any, actual any) (err *AssertError) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(AssertError); ok {
				err = &e
			} else {
				panic(r)
			}
		}
	}()

	eh.assert(fieldName, expect, actual)
	return
}

func (eh *assertHelper) assertMap(fieldName string, expect any, actual any) {
	expectType := reflect.TypeOf(expect).String()
	actualType := reflect.TypeOf(actual).String()
	em := mapPattern.FindAllStringSubmatch(expectType, -1)
	am := mapPattern.FindAllStringSubmatch(actualType, -1)
	if am == nil {
		panic(AssertError{fieldName, expect, actual,
			fmt.Sprintf("Expect map type %s while actual type is %s.", expectType, actualType)})
	}
	if em[0][1] != am[0][1] {
		panic(AssertError{fieldName, expect, actual,
			fmt.Sprintf("Map key not equal, expect type %s while actual type is %s.", expectType, actualType)})
	}

	ev, av := reflect.ValueOf(expect), reflect.ValueOf(actual)
	if ev.Len() != av.Len() {
		panic(AssertError{fieldName, expect, actual,
			fmt.Sprint("map size not equal, expect ", ev.Len(), ", actual ", av.Len())})
	}
	for it := ev.MapRange(); it.Next(); {
		avi := av.MapIndex(it.Key())
		if !avi.IsValid() {
			panic(AssertError{fieldName, expect, actual,
				fmt.Sprint("Expect key ", it.Key().Interface(), " not exist.")})
		}
		eh.assert(childFieldName(fieldName, it.Key().Interface()), it.Value().Interface(), avi.Interface())
	}
}

func (eh *assertHelper) AssertInRange(actual any, expects ...any) {
	eh.times++
	for _, expect := range expects {
		err := eh.assertWithError("", expect, actual)
		if err == nil {
			return
		}
	}

	panic(fmt.Sprintf(`Error while validate (%dth): none of expects match actual:
    expects: %s,
    actual : %s,
`, eh.times, expects, actual))
}
