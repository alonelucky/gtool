package array

import "reflect"

func Sort(in, out interface{}, fn func(a, v interface{}) int8) {
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

	l := inv.Len() - 1

	outv = indirect(outv)
	reflect.AppendSlice(outv, inv)

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			v := fn(inv.Index(i), inv.Index(j))
			if v > 0 {
				outv.Index(i).Set(inv.Index(j))
				outv.Index(j).Set(inv.Index(i))
			}
		}
	}
}
