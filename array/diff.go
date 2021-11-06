package array

import (
	"reflect"
)

// 存在A中，不存在B中,返回diff
func Diff(diff, a, b interface{}) {
	if a == nil {
		return
	}

	var (
		av = indirect(reflect.ValueOf(a))
		bv = indirect(reflect.ValueOf(b))
		dv = reflect.ValueOf(diff)
	)

	dv = indirect(dv)

	if b == nil {
		reflect.AppendSlice(dv, av)
		return
	}

	l := av.Len()
	var arr = reflect.MakeSlice(av.Type(), 0, l)
	for i := 0; i < l; i++ {
		v := av.Index(i)
		if !Contains(bv.Interface(), v.Interface()) {
			arr = reflect.Append(arr, v)
		}
	}

	dv.Set(arr)
}
