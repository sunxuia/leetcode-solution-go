package test

import (
	"fmt"
	"reflect"
	"strconv"
)

type IEqual interface {
	Equal(other interface{}) bool
}

func DefaultAssertEqual(handler *Handler) {
	// IEqual (by reference)
	if expect, ok := handler.Expect.(IEqual); ok {
		if !expect.Equal(handler.Actual) {
			handler.Error("Not equal via IEqual")
		}
		return
	}

	ev, av := handler.Ev, handler.Av

	switch ev.Kind() {
	case reflect.Ptr, reflect.UnsafePointer, reflect.Uintptr:
		if handler.Expect == nil || ev.IsNil() {
			handler.ChangeValue(nil, handler.Actual)
		} else {
			handler.ChangeValue(av.Elem().Interface(), handler.Actual)
		}
		ev = handler.Ev
	}
	switch av.Kind() {
	case reflect.Ptr, reflect.UnsafePointer, reflect.Uintptr:
		if handler.Actual == nil || av.IsNil() {
			handler.ChangeValue(handler.Expect, nil)
		} else {
			handler.ChangeValue(handler.Expect, av.Elem().Interface())
		}
		av = handler.Av
	}
	if handler.Expect == nil || ev.Kind() == reflect.Invalid {
		if !(handler.Actual == nil || !av.IsValid() || av.CanInterface() && av.IsNil()) {
			handler.Error("Expect nil, actual not.")
		}
		return
	}
	if ev.CanAddr() && av.CanAddr() && ev.Pointer() == av.Pointer() {
		return
	}
	if ev.Type() != av.Type() {
		handler.Error(fmt.Sprint("Type not equal, expect ", ev.Type(), ", actual ", av.Type()))
	}

	switch ev.Kind() {
	case reflect.Float32, reflect.Float64:
		AssertFloat(handler)
	case reflect.Array, reflect.Slice:
		AssertArray(handler)
	case reflect.Map:
		AssertMap(handler)
	case reflect.Interface, reflect.Struct:
		AssertStruct(handler)
	case reflect.Chan, reflect.Func:
		panic(fmt.Sprint("Unsupported kind: ", ev.Kind(), " ."))
	default:
		if handler.Expect != handler.Actual {
			handler.Error("Value not equal.")
		}
	}
}

func AssertFloat(handler *Handler) {
	diff := handler.Ev.Float() - handler.Av.Float()
	if diff < 0 {
		diff = -diff
	}
	tolerance := handler.config.numberTolerance
	if diff > tolerance {
		if handler.config.numberTolerance == 0 {
			handler.Error("Value not equal.")
		} else {
			handler.Error(fmt.Sprintf("Difference %f larger than tolerance %f.", diff, tolerance))
		}
	}
}

func AssertArray(handler *Handler) {
	ev, av := handler.Ev, handler.Av
	length, lenA := ev.Len(), av.Len()
	if length != lenA {
		handler.Error(fmt.Sprint("Expect array length ", length, " actual length ", lenA))
	}

	if handler.IsOrdered() {
		for i := 0; i < length; i++ {
			cev, cav := ev.Index(i).Interface(), av.Index(i).Interface()
			child := handler.Child(strconv.Itoa(i), cev, cav)
			child.Assert()
		}
	} else {
		match := make([]bool, length, length)
		for i := 0; i < length; i++ {
			hasMatch := false
			for j := 0; j < length; j++ {
				if match[j] {
					continue
				}
				cev, cav := ev.Index(i).Interface(), av.Index(j).Interface()
				child := handler.Child(strconv.Itoa(i), cev, cav)
				err := child.AssertWithError()

				if err == nil {
					match[j] = true
					hasMatch = true
					break
				}
			}
			if !hasMatch {
				handler.Error(fmt.Sprint("Element at ", i, " has no match."))
			}
		}
	}
}

func AssertMap(handler *Handler) {
	ev, av := handler.Ev, handler.Av
	if ev.Len() != av.Len() {
		handler.Error(fmt.Sprint("map size not equal, expect ", ev.Len(), ", actual ", av.Len()))
	}
	for it := ev.MapRange(); it.Next(); {
		avi := av.MapIndex(it.Key())
		if !avi.IsValid() {
			handler.Error(fmt.Sprint("Expect key ", it.Key().Interface(), " not exist."))
		}
		child := handler.Child(fmt.Sprint(it.Key().Interface()), it.Value().Interface(), avi.Interface())
		child.Assert()
	}
}

func AssertStruct(handler *Handler) {
	// IEqual (by value)
	if expect, ok := handler.Expect.(IEqual); ok {
		if !expect.Equal(handler.Actual) {
			handler.Error("Not equal via IEqual")
		}
		return
	}

	ev, av := handler.Ev, handler.Av
	for i, size := 0, ev.NumField(); i < size; i++ {
		fieldName := ev.Type().Field(i).Name
		if 'A' <= fieldName[0] && fieldName[0] <= 'Z' {
			cev, aev := ev.FieldByName(fieldName), av.FieldByName(fieldName)
			if aev.IsZero() {
				handler.Error("Expect " + fieldName + " not exist in actual.")
			}
			child := handler.Child(fieldName, cev.Interface(), aev.Interface())
			child.Assert()
		}
	}
}
