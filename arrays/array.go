package arrays

import (
	"reflect"
)

func copyslice(dst, src reflect.Value) {
	l := src.Len()
	var arr = reflect.MakeSlice(src.Type(), 0, l)
	for i := 0; i < l; i++ {
		arr = reflect.Append(arr, src.Index(i))
	}

	dst.Set(arr)
}
