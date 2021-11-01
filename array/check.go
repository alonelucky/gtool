package array

import (
	"reflect"
)

func Index(lst, v interface{}) int {
	if lst == nil {
		return -1
	}

	var inv = reflect.ValueOf(lst)
	var vv = indirect(reflect.ValueOf(v))
	var l = inv.Len()

	for i := 0; i < l; i++ {
		item := indirect(inv.Index(i))
		// if item.Kind() != vv.Kind() {
		// 	continue
		// }

		if item.Elem() == vv.Elem() {
			return i
		}

		// switch item.Kind() {
		// case reflect.Int, reflect.Int16, reflect.Int8, reflect.Int32, reflect.Int64:
		// 	if item.Int() == vv.Int() {
		// 		return i
		// 	}
		// case reflect.Uint, reflect.Uint16, reflect.Uint8, reflect.Uint32, reflect.Uint64:
		// 	if item.Uint() == vv.Uint() {
		// 		return i
		// 	}
		// case reflect.Bool:
		// 	if item.Bool() == vv.Bool() {
		// 		return i
		// 	}
		// case reflect.Float32, reflect.Float64:
		// 	if item.Float() == vv.Float() {
		// 		return i
		// 	}
		// default:
		// 	if item.Interface() == v {
		// 		return i
		// 	}
		// }
	}

	return -1
}

func Contains(lst, v interface{}) bool {
	return Index(lst, v) > -1
}

func isArray(v reflect.Value) bool {
	k := v.Kind()
	if k == reflect.Array || k == reflect.Slice {
		return true
	}
	return false
}

func indirect(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		v = reflect.Indirect(v)
	}
	return v
}

func IsInt(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int8, reflect.Int32, reflect.Int64:
		return true
	}
	return false
}

func IsUnit(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Uint, reflect.Uint16, reflect.Uint8, reflect.Uint32, reflect.Uint64:
		return true
	}
	return false
}

func IsFloat(v reflect.Value) bool {
	k := v.Kind()
	if k == reflect.Float32 || k == reflect.Float64 {
		return true
	}
	return false
}
