package arrays

import (
	"reflect"

	"github.com/alonelucky/gtool/reflects"
)

// IndexOf gets the position of the element in the slice.
//
// You can pass in the comparison function, and return true to prove that they are equal
func IndexOf(lst, v interface{}, fns ...func(a, b interface{}) bool) int {
	if lst == nil {
		return -1
	}

	var inv = reflect.ValueOf(lst)
	var vv = reflects.Indirect(reflect.ValueOf(v))
	var l = inv.Len()

	for i := 0; i < l; i++ {
		item := reflects.Indirect(inv.Index(i))
		if len(fns) > 0 && fns[0](item.Interface(), v) {
			return i
		}

		if !reflects.SameKind(item, vv) {
			continue
		}

		switch item.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if item.Int() == vv.Int() {
				return i
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if item.Uint() == vv.Uint() {
				return i
			}
		case reflect.Float32, reflect.Float64:
			if item.Float() == vv.Float() {
				return i
			}
		case reflect.String:
			if item.String() == vv.String() {
				return i
			}
		case reflect.Bool:
			if item.Bool() == vv.Bool() {
				return i
			}
		default:
			if item == vv {
				return i
			}
		}
	}

	return -1
}

// Contains Does the slice contain elements
//
// You can pass in the comparison function, and return true to prove that they are equal
func Contains(lst, v interface{}, fns ...func(a, b interface{}) bool) bool {
	return IndexOf(lst, v, fns...) > -1
}
