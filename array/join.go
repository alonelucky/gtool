package array

import (
	"fmt"
	"reflect"
	"strings"
)

func Joins(lst interface{}, spec string) (s string) {
	if lst == nil {
		return
	}

	var bd strings.Builder
	var inv = reflect.ValueOf(lst)

	if !isArray(inv) {
		return
	}

	var l = inv.Len()
	for i := 0; i < l; i++ {
		bd.WriteString(fmt.Sprintf("%+v", inv.Index(i).Interface()))
		if i != l-1 {
			bd.WriteString(spec)
		}
	}

	return bd.String()
}
