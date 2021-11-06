package array

import (
	"reflect"
)

func Filter(in, out interface{}, fn func(v interface{}, i int) bool) {
	if in == nil {
		return
	}

	var (
		inv  = reflect.ValueOf(in)
		outv = reflect.ValueOf(out)
	)

	inv = indirect(inv)
	l := inv.Len()

	var r = reflect.MakeSlice(inv.Type(), 0, l)
	for i := 0; i < l; i++ {
		iv := inv.Index(i)
		v := fn(iv.Interface(), i)
		if v {
			r = reflect.Append(r, iv)
		}
	}

	outv = indirect(outv)
	if outv.Kind() == reflect.Slice {
		outv.Set(r)
	} else {
		inv.Set(r)
	}
}
