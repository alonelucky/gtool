package array

import "reflect"

func Map(in, out interface{}, fn func(v interface{}, i int) interface{}) {
	if in == nil {
		return
	}

	var (
		inv  = reflect.ValueOf(in)
		outv = reflect.ValueOf(out)
	)

	if !isArray(inv) || outv.Kind() != reflect.Ptr {
		return
	}

	l := inv.Len()

	outv = indirect(outv)

	if !isArray(outv) {
		return
	}

	for i := 0; i < l; i++ {
		v := fn(inv.Index(i).Interface(), i)
		outv.Set(reflect.Append(outv, reflect.ValueOf(v)))
	}
}

func MapSelf(in interface{}, fn func(v interface{}, i int) interface{}) {
	if in == nil {
		return
	}

	var (
		inv = reflect.ValueOf(in)
		arr = indirect(inv)
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
