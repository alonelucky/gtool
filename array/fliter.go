package array

import "reflect"

func Fliter(in, out interface{}, fn func(v interface{}, i int) bool) {
	if in == nil {
		return
	}

	var (
		inv  = reflect.ValueOf(in)
		outv = reflect.ValueOf(out)
	)

	if !isArray(inv) || outv.Type().Kind() != reflect.Ptr {
		return
	}

	l := inv.Len()
	outv = indirect(outv)

	if !isArray(outv) {
		return
	}

	for i := 0; i < l; i++ {
		v := fn(inv.Index(i), i)
		if v {
			reflect.Append(outv, reflect.ValueOf(v))
		}
	}
}
