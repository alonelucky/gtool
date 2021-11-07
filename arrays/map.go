package arrays

import (
	"reflect"

	"github.com/alonelucky/gtool/reflects"
)

// Map 遍历 input slice type, output slice type
//
// 通过fn函数定义修改后的返回信息
func Map(in, out interface{}, fn func(v interface{}, i int) interface{}) {
	if in == nil {
		return
	}

	var (
		inv  = reflect.ValueOf(in)
		outv = reflect.ValueOf(out)
	)

	if !reflects.IsArray(inv) || outv.Kind() != reflect.Ptr {
		return
	}

	l := inv.Len()

	outv = reflects.Indirect(outv)

	if !reflects.IsArray(outv) {
		return
	}

	for i := 0; i < l; i++ {
		v := fn(inv.Index(i).Interface(), i)
		outv.Set(reflect.Append(outv, reflect.ValueOf(v)))
	}
}

// Map 遍历 input slice type, change input
//
// 通过fn函数定义修改后的返回信息
func MapSelf(in interface{}, fn func(v interface{}, i int) interface{}) {
	if in == nil {
		return
	}

	var (
		inv = reflect.ValueOf(in)
		arr = reflects.Indirect(inv)
	)

	if inv.Kind() != reflect.Ptr {
		return
	}

	l := arr.Len()
	for i := 0; i < l; i++ {
		v := fn(arr.Index(i).Interface(), i)
		arr.Index(i).Set(reflect.ValueOf(v))
	}
}
