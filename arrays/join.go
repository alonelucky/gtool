package arrays

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/alonelucky/gtool/reflects"
)

// Joins 合并数组为字符串,支持全类型, 使用 fmt %+v拼接
func Joins(lst interface{}, spec string) (s string) {
	if lst == nil {
		return
	}

	var bd strings.Builder
	var inv = reflect.ValueOf(lst)

	if !reflects.IsArray(inv) {
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
