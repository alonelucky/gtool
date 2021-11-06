package array

import (
	"fmt"
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

func isInt(i reflect.Value) bool {
	ik := indirect(i).Kind()
	return ik > 1 && ik < 7
}

func isUint(i reflect.Value) bool {
	ik := indirect(i).Kind()
	return ik > 6 && ik < 12
}

func sameType(a, b reflect.Value) bool {
	var (
		ak = indirect(a).Kind()
		bk = indirect(b).Kind()
	)

	fmt.Println(uint(ak), uint(bk))

	if (ak > 1 && ak < 7) && (bk > 1 && bk < 7) {
		return true
	}

	if (ak > 7 && ak < 12) && (bk > 7 && bk < 12) {
		return true
	}

	if (ak == 13 || ak == 14) && (bk == 13 || bk == 14) {
		return true
	}

	return ak == bk
}
