package array

import "reflect"

func Each(in interface{}, fn func(v interface{}, i int)) {
	if in == nil {
		return
	}

	var inv = reflect.ValueOf(in)

	if !isArray(inv) {
		return
	}

	l := inv.Len()

	for i := 0; i < l; i++ {
		fn(inv.Index(i), i)
	}
}
