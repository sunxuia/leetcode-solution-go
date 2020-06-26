package test

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

type MethodTestRunner struct {
	TestCases []func() *MethodTestCase
}

func NewMethodTestRunner() *MethodTestRunner {
	res := new(MethodTestRunner)
	res.TestCases = make([]func() *MethodTestCase, 0)
	return res
}

func (mtr *MethodTestRunner) AddTestCase(testCase *MethodTestCase) *MethodTestRunner {
	return mtr.AddTestCaseProvider(func() *MethodTestCase {
		return testCase
	})
}

func (mtr *MethodTestRunner) AddTestCaseProvider(provider func() *MethodTestCase) *MethodTestRunner {
	mtr.TestCases = append(mtr.TestCases, provider)
	return mtr
}

func (mtr *MethodTestRunner) Test(t *testing.T, method any) {
	if len(mtr.TestCases) == 0 {
		t.Skip()
		return
	}

	var testName string
	var idx int
	var expect, actual any
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(assertError); ok {
				t.Error(fmt.Sprintf(`Error while validate testCase %s
    expect: %s,
    actual: %s,
    detail (%s):
        expect: %s, 
        actual: %s, 
        message: %s`, testName, expect, actual, e.fieldName, e.expected, e.actual, e.message))
			} else {
				t.Error(r)
			}
		}
	}()
	for i, provider := range mtr.TestCases {
		testCase := provider()
		if len(testCase.name) > 0 {
			testName = testCase.name
		} else {
			testName = strconv.Itoa(i)
		}

		params := make([]reflect.Value, len(testCase.Arguments))
		for i := 0; i < len(params); i++ {
			params[i] = reflect.ValueOf(testCase.Arguments[i])
		}
		ret := reflect.ValueOf(method).Call(params)

		for idx, expect = range testCase.Expects {
			if idx >= 0 {
				actual = testCase.Arguments[idx]
			} else {
				actual = ret[-idx-1].Interface()
			}

			assertHelper := testCase.Asserts[idx]
			if assertHelper == nil {
				assertHelper = defaultAsserter
			}
			assertHelper.Assert(expect, actual)
		}
	}
}

var defaultAsserter = NewAssertHelper()

type MethodTestCase struct {
	name      string
	Asserts   map[int]*assertHelper
	Expects   map[int]any
	Arguments []any
}

func NewMethodTestCase(arguments ...any) *MethodTestCase {
	res := new(MethodTestCase)
	res.Asserts = map[int]*assertHelper{}
	res.Expects = map[int]any{}
	res.Arguments = arguments
	return res
}

func (mtc *MethodTestCase) ExpectReturn(vals ...any) *MethodTestCase {
	for i := range vals {
		mtc.Expects[-i-1] = vals[i]
	}
	return mtc
}

func (mtc *MethodTestCase) ExpectArgument(index int, val any) *MethodTestCase {
	mtc.Expects[index] = val
	return mtc
}

func (mtc *MethodTestCase) SetAssert(index int, assertHelper *assertHelper) *MethodTestCase {
	mtc.Asserts[index] = assertHelper
	return mtc
}
